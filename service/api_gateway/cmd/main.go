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

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8081"
	}
	acco := os.Getenv("ACCO")
	if len(acco) == 0 {
		acco = "account:50051"
	}
	sign := os.Getenv("SIGN")
	if len(sign) == 0 {
		sign = "signature:50053"
	}
	auth := os.Getenv("AUTH")
	if len(auth) == 0 {
		auth = "auth:50054"
	}
	ship := os.Getenv("SHIP")
	if len(ship) == 0 {
		ship = "shipment:50055"
	}

	accountConn, err := grpc.Dial(acco, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to account service %v", err)
	}
	defer accountConn.Close()
	accountService := pb.NewAccountServiceClient(accountConn)
	fmt.Printf("Connection to account service made\n")

	//Connect to Signature service
	signatureConn, err := grpc.Dial(sign, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Signature service %v", err)
	}
	defer signatureConn.Close()
	signatureService := pb.NewSignatureServiceClient(signatureConn)
	fmt.Printf("Connection to Signature service made\n")

	//Connect to Authentication service
	authConn, err := grpc.Dial(auth, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Authentication service %v", err)
	}
	defer authConn.Close()
	authService := pb.NewAuthenticationServiceClient(authConn)
	fmt.Printf("Connection to Authentication service made\n")

	//Connect to Shipment service
	shipmentConn, err := grpc.Dial(ship, grpc.WithInsecure())
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

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-Requested-With", "Accept", "Authorization", "Accept-Language", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Check against your desired domains here
			return r.Host == "http://localhost:8081" || r.Host == "http://localhost:3000"
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
