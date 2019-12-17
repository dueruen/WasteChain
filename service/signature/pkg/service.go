package pkg

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"
	"github.com/dueruen/WasteChain/service/signature/pkg/event/pub"
	"github.com/dueruen/WasteChain/service/signature/pkg/event/sub"
	"github.com/dueruen/WasteChain/service/signature/pkg/key"
	"github.com/dueruen/WasteChain/service/signature/pkg/sign"
	"github.com/dueruen/WasteChain/service/signature/pkg/storage/postgres"
	"github.com/dueruen/WasteChain/service/signature/pkg/transport"
	grpctransport "github.com/dueruen/WasteChain/service/signature/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func Run() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":50053"
	}
	dbString := os.Getenv("DB_STRING")
	if dbString == "" {
		dbString = "host=db port=5432 user=root dbname=root password=root sslmode=disable"
	}
	block := os.Getenv("BLOC")
	if block == "" {
		block = "blockchain:50056"
	}
	qr := os.Getenv("QR")
	if qr == "" {
		qr = "qr:50052"
	}
	nats := os.Getenv("NATS")
	if nats == "" {
		nats = "nats:4222"
	}

	//Create storage
	storage, err := postgres.NewStorage(dbString)
	defer postgres.Close(storage)
	if err != nil {
		fmt.Printf("Storage err: %v\n", err)
	}

	//Connect Pub to NATS
	pubEventHandler, errPub := pub.NewEventHandler(nats)
	if errPub != nil {
		log.Fatalf("Could not connect to NATS %v", errPub)
	}
	fmt.Printf("Pub Connection to NATS service made\n")

	//Connect to QR Service
	cc, err := grpc.Dial(qr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to QR service %v", err)
	} else {
		fmt.Printf("Connection to QR service made\n")
	}
	defer cc.Close()

	qrClient := pb.NewQRServiceClient(cc)

	//Connect to Blockchain Service
	ccBlock, err := grpc.Dial(block, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Blockchain service %v", err)
	} else {
		fmt.Printf("Connection to Blockchain service made\n")
	}
	defer ccBlock.Close()

	blockClient := pb.NewBlockchainServiceClient(ccBlock)

	//Create Key Service
	keySrv := key.NewService(storage)

	//Create Sign Service
	signSrv := sign.NewService(storage, keySrv, pubEventHandler, qrClient, blockClient)

	//Connect Sub to NATS
	errSub := sub.StartListening(nats, signSrv)
	if errSub != nil {
		log.Fatalf("Could not connect to NATS %v", errSub)
	}
	fmt.Printf("Sub Connection to NATS service made\n")

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(keySrv, signSrv)
	}

	var (
		ocTracing        = kitoc.GRPCServerTrace()
		serverOptions    = []kitgrpc.ServerOption{ocTracing}
		signatureService = grpctransport.NewGRPCServer(endpoints, serverOptions)
		grpcListener, _  = net.Listen("tcp", port)
		grpcServer       = grpc.NewServer()
	)

	pb.RegisterSignatureServiceServer(grpcServer, signatureService)

	fmt.Printf("Signing service is listening on port %s...\n", port)

	err = grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}
