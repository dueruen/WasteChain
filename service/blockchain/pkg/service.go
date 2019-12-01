package pkg

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/dueruen/WasteChain/service/blockchain/gen/proto"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/event/pub"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/event/sub"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/publish"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/receive"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/storage/postgres"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/transport"
	grpctransport "github.com/dueruen/WasteChain/service/blockchain/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func Run() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":50056"
	}
	endpoint := os.Getenv("ENDPOINT")
	if len(port) == 0 {
		port = "https://nodes.devnet.thetangle.org"
	}
	dbString := os.Getenv("DB_STRING")
	if dbString == "" {
		dbString = "host=db port=5432 user=root dbname=root password=root sslmode=disable"
	}
	nats := os.Getenv("NATS")
	if len(nats) == 0 {
		nats = "localhost:4222"
	}

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

	receiveService := receive.NewService(storage, endpoint, pubEventHandler)
	publishService := publish.NewService(storage, endpoint, pubEventHandler)

	//Connect Sub to NATS
	errSub := sub.StartListening(nats, publishService)
	if errSub != nil {
		log.Fatalf("Could not connect to NATS %v", errSub)
	}
	fmt.Printf("Sub Connection to NATS service made\n")

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(receiveService, publishService)
	}

	var (
		ocTracing         = kitoc.GRPCServerTrace()
		serverOptions     = []kitgrpc.ServerOption{ocTracing}
		blockchainService = grpctransport.NewGrpcServer(endpoints, serverOptions)
		grpcListener, _   = net.Listen("tcp", port)
		grpcServer        = grpc.NewServer()
	)

	pb.RegisterBlockchainServiceServer(grpcServer, blockchainService)

	fmt.Printf("Blockchain service is listening on port %s...\n", port)

	err = grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}
