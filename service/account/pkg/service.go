package pkg

import (
	"fmt"
	"net"

	"github.com/dueruen/WasteChain/service/account/pkg/creating"
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
	"github.com/dueruen/WasteChain/service/account/pkg/storage/postgres"
	"github.com/dueruen/WasteChain/service/account/pkg/transport"
	grpctransport "github.com/dueruen/WasteChain/service/account/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
)

type listingRepository = listing.Repository
type creatingRepository = creating.Repository

type listingService = listing.Service
type creatingService = creating.Service

const port = ":50051"

func Run() {
	storage, err := postgres.NewStorage("localhost", "5432", "root", "root", "root")
	defer postgres.Close(storage)
	if err != nil {
		fmt.Printf("Storage err: %v\n", err)
	}
	creatingService := creating.NewService(storage)
	listingService := listing.NewService(storage)

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(creatingService, listingService)
	}

	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		accountService  = grpctransport.NewGrpcServer(endpoints, serverOptions)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	pb.RegisterAccountServiceServer(grpcServer, accountService)

	fmt.Printf("Account service is listening on port %s...\n", port)

	err = grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}
