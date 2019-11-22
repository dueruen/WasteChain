package pkg

import (
	"fmt"
	"net"

	pb "github.com/dueruen/WasteChain/service/authentication/gen/proto"
	"github.com/dueruen/WasteChain/service/authentication/pkg/auth"
	"github.com/dueruen/WasteChain/service/authentication/pkg/storage/postgres"
	"github.com/dueruen/WasteChain/service/authentication/pkg/transport"
	"google.golang.org/grpc"

	grpctransport "github.com/dueruen/WasteChain/service/authentication/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

const port = ":50054"

func Run() {
	storage, err := postgres.NewStorage("localhost", "5432", "root", "root", "root")
	defer postgres.Close(storage)
	if err != nil {
		fmt.Printf("Storage err: %v\n", err)
	}

	authService := auth.NewService(storage)

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(authService)
	}

	var (
		ocTracing             = kitoc.GRPCServerTrace()
		serverOptions         = []kitgrpc.ServerOption{ocTracing}
		authenticationService = grpctransport.NewGrpcServer(endpoints, serverOptions)
		grpcListener, _       = net.Listen("tcp", port)
		grpcServer            = grpc.NewServer()
	)

	pb.RegisterAuthenticationServiceServer(grpcServer, authenticationService)

	fmt.Printf("Autehntication service is listening on port %s...\n", port)

	err = grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}
