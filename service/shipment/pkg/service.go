package pkg

import (
	"fmt"
	"net"

	"github.com/dueruen/WasteChain/service/shipment/pkg/creating"
	"github.com/dueruen/WasteChain/service/shipment/pkg/listing"
	"github.com/dueruen/WasteChain/service/shipment/pkg/processing"
	"github.com/dueruen/WasteChain/service/shipment/pkg/storage/postgres"
	"github.com/dueruen/WasteChain/service/shipment/pkg/transfering"
	"github.com/dueruen/WasteChain/service/shipment/pkg/transport"
	grpctransport "github.com/dueruen/WasteChain/service/shipment/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type creatingRepository = creating.Repository
type listingRepository = listing.Repository
type transferingRepository = transfering.Repository
type processingRepository = processing.Repository

type creatingService = creating.Service
type listingService = listing.Service
type transferingService = transfering.Service
type processingService = processing.Service

const port = ":50051"

func Run() {
	storage, err := postgres.NewStorage("localhost", "5433", "root", "root", "root")
	defer postgres.Close(storage)
	if err != nil {
		fmt.Printf("Storage err: %v\n", err)
	}
	creatingService := creating.NewService(storage)
	listingService := listing.NewService(storage)
	transferingService := transfering.NewService(storage)
	processingService := processing.NewService(storage)

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(creatingService, listingService, transferingService, processingService)
	}

	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		shipmentService = grpctransport.NewGrpcServer(endpoints, serverOptions)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	pb.RegisterShipmentServiceServer(grpcServer, shipmentService)

	fmt.Printf("Shipment service is listening on port %s...\n", port)

	err = grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}
