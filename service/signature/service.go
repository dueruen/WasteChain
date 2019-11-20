package signature

import (
	"fmt"
	"log"
	"net"

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

const port = ":50053"

func Run() {
	//Create storage
	storage, err := postgres.NewStorage("localhost", "5432", "root", "root", "root")
	defer postgres.Close(storage)
	if err != nil {
		fmt.Printf("Storage err: %v\n", err)
	}

	//Connect Pub to NATS
	pubEventHandler, errPub := pub.NewEventHandler("localhost:4222")
	if errPub != nil {
		log.Fatalf("Could not connect to NATS %v", errPub)
	}
	fmt.Printf("Pub Connection to NATS service made\n")

	//Connect to QR Service
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to QR service %v", err)
	} else {
		fmt.Printf("Connection to QR service made\n")
	}
	defer cc.Close()

	qrClient := pb.NewQRServiceClient(cc)

	//Create Key Service
	keySrv := key.NewService(storage)

	//Create Sign Service
	signSrv := sign.NewService(storage, keySrv, pubEventHandler, qrClient)

	//Connect Sub to NATS
	errSub := sub.StartListening("localhost:4222", signSrv)
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

	fmt.Printf("Shipping service is listening on port %s...\n", port)

	err = grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}
