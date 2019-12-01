package pkg

import (
	"fmt"
	"log"
	"net"
	"os"

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

func Run() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":50051"
	}
	dbString := os.Getenv("DB_STRING")
	if dbString == "" {
		dbString = "host=db port=5432 user=root dbname=root password=root sslmode=disable"
	}
	sign := os.Getenv("SIGN")
	if len(sign) == 0 {
		sign = "signature:50053"
	}
	auth := os.Getenv("AUTH")
	if len(auth) == 0 {
		auth = "auth:50054"
	}

	storage, err := postgres.NewStorage(dbString)
	defer postgres.Close(storage)
	if err != nil {
		fmt.Printf("Storage err: %v\n", err)
	}

	//Connect to Signature Service
	signConn, err := grpc.Dial(sign, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Signature service %v", err)
	} else {
		fmt.Printf("Connection to Signature service made\n")
	}
	defer signConn.Close()

	signClient := pb.NewSignatureServiceClient(signConn)

	//Connect to Authentication Service
	authConn, err := grpc.Dial(auth, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Authentication service %v", err)
	} else {
		fmt.Printf("Connection to Authentication service made\n")
	}
	defer authConn.Close()

	authClient := pb.NewAuthenticationServiceClient(authConn)

	creatingService := creating.NewService(storage, authClient, signClient)
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
