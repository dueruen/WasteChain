package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"
	pb "github.com/dueruen/WasteChain/service/api_gateway/gen/proto"
	"github.com/dueruen/WasteChain/service/api_gateway/graphql"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

func main() {
	//Connect to Account service
	accountConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to account service %v", err)
	}
	defer accountConn.Close()
	accountService := pb.NewAccountServiceClient(accountConn)
	fmt.Printf("Connection to account service made\n")

	//Connect to QR service
	qrConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to QR service %v", err)
	}
	defer qrConn.Close()
	qrService := pb.NewQRServiceClient(qrConn)
	fmt.Printf("Connection to QR service made\n")

	//Connect to Signature service
	signatureConn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Signature service %v", err)
	}
	defer signatureConn.Close()
	signatureService := pb.NewSignatureServiceClient(signatureConn)
	fmt.Printf("Connection to Signature service made\n")

	//Connect to Authentication service
	authConn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Authentication service %v", err)
	}
	defer authConn.Close()
	authService := pb.NewAuthenticationServiceClient(authConn)
	fmt.Printf("Connection to Authentication service made\n")

	//Connect to Shipment service
	shipmentConn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to shipment service %v", err)
	}
	defer shipmentConn.Close()
	shipmentService := pb.NewShipmentServiceClient(shipmentConn)
	fmt.Printf("Connection to shipment service made\n")

	resolver := graphql.Resolver{
		AccountClient:        accountService,
		QRClient:             qrService,
		SignatureClient:      signatureService,
		AuthenticationClient: authService,
		ShipmentClient:       shipmentService,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver}),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			// send this panic somewhere
			log.Print(err)
			debug.PrintStack()
			return errors.New("Something went wrong!!!!")
		}),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
