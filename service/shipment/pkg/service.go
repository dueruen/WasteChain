package pkg

import (
	"fmt"
	"log"
	"net"

	"github.com/dueruen/WasteChain/service/shipment/pkg/creating"
	"github.com/dueruen/WasteChain/service/shipment/pkg/event/sub"
	"github.com/dueruen/WasteChain/service/shipment/pkg/eventvalidating"
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
type validationRepository = eventvalidating.Repository

type creatingService = creating.Service
type listingService = listing.Service
type transferingService = transfering.Service
type processingService = processing.Service
type validationService = eventvalidating.Service

const port = ":50055"

func Run() {
	storage, err := postgres.NewStorage("localhost", "5432", "root", "root", "root")
	defer postgres.Close(storage)
	if err != nil {
		fmt.Printf("Storage err: %v\n", err)
	}

	//Creating validation service
	validationService := eventvalidating.NewService(storage)

	//Connect to Signature Service
	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Signature service %v", err)
	} else {
		fmt.Printf("Connection to Signature service made\n")
	}
	defer cc.Close()

	signClient := pb.NewSignatureServiceClient(cc)

	//Connect to Account Service
	ac, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Account service %v", err)
	} else {
		fmt.Printf("Connection to Account service made\n")
	}
	defer ac.Close()

	accountClient := pb.NewAccountServiceClient(ac)

	//Connect Sub to NATS
	errSub := sub.StartListening("localhost:4222", validationService)
	if errSub != nil {
		log.Fatalf("Could not connect to NATS %v", errSub)
	}
	fmt.Printf("Sub connection to NATS service made\n")

	creatingService := creating.NewService(storage, signClient, accountClient)
	listingService := listing.NewService(storage)
	transferingService := transfering.NewService(storage, signClient)
	processingService := processing.NewService(storage, signClient)

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
