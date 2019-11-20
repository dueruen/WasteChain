package pkg

import (
	"fmt"
	"log"
	"net"

	"github.com/dueruen/WasteChain/service/qr/pkg/creating"
	"github.com/dueruen/WasteChain/service/qr/pkg/event"
	"github.com/dueruen/WasteChain/service/qr/pkg/transport"
	grpctransport "github.com/dueruen/WasteChain/service/qr/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	pb "github.com/dueruen/WasteChain/service/qr/gen/proto"
)

const port = ":50052"

func Run() {

	eventHandler, errSub := event.NewEventHandler("localhost:4222")
	if errSub != nil {
		log.Fatalf("Could not connect to NATS %v", errSub)
	}
	fmt.Printf("Connection to NATS service made\n")

	creatingService := creating.NewService(eventHandler)
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
