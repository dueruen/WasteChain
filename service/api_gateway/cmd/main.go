package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	pb "github.com/dueruen/WasteChain/service/api_gateway/gen/proto"
	"github.com/dueruen/WasteChain/service/api_gateway/graphql"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
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
		SignatureClient:      signatureService,
		AuthenticationClient: authService,
		ShipmentClient:       shipmentService,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Check against your desired domains here
			return r.Host == "example.org"
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	router.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	router.Handle("/query",
		handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver}), handler.WebsocketUpgrader(upgrader)),
	)

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		panic(err)
	}
}
