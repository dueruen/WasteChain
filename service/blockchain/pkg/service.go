package pkg

import (
	"fmt"
	"log"
	"net"

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

const port = ":50056"

var endpoint = "https://nodes.devnet.thetangle.org"

func Run() {
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

	receiveService := receive.NewService(storage, endpoint, pubEventHandler)
	publishService := publish.NewService(storage, endpoint, pubEventHandler)

	//Connect Sub to NATS
	errSub := sub.StartListening("localhost:4222", publishService)
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
