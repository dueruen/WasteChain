package pkg

import (
	"fmt"
	"net"

	"github.com/dueruen/WasteChain/service/qr/pkg/creating"
	"github.com/dueruen/WasteChain/service/qr/pkg/transport"
	grpctransport "github.com/dueruen/WasteChain/service/qr/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	pb "github.com/dueruen/WasteChain/service/qr/gen/proto"
)

const port = ":50051"

func Run() {

	creatingService := creating.NewService()
	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(creatingService)
	}

	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		QRService       = grpctransport.NewGrpcServer(endpoints, serverOptions)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	pb.RegisterQRServiceServer(grpcServer, QRService)

	fmt.Printf("QR service is listening on port %s...\n", port)

	err := grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)

}
