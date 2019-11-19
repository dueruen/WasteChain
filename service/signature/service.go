package signature

import (
	"fmt"
	"log"
	"net"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"
	"github.com/dueruen/WasteChain/service/signature/pkg/event"
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
	storage, err := postgres.NewStorage("localhost", "5432", "root", "root", "root")
	defer postgres.Close(storage)
	if err != nil {
		fmt.Printf("Storage err: %v\n", err)
	}

	eventHandler, errSub := event.NewEventHandler("localhost:4222")
	if errSub != nil {
		log.Fatalf("Could not connect to NATS %v", errSub)
	}
	fmt.Printf("Connection to NATS service made\n")

	keySrv := key.NewService(storage)
	signSrv := sign.NewService(storage, keySrv, eventHandler)
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
