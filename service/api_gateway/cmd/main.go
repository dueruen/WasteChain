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
	accountConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to account service %v", err)
	}
	defer accountConn.Close()
	accountService := pb.NewAccountServiceClient(accountConn)
	fmt.Printf("Connection to account service made\n")

	resolver := graphql.Resolver{
		AccountClient: accountService,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	// http.Handle("/query", handler.GraphQL(
	// 	graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver}),
	// 	handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
	// 		// send this panic somewhere
	// 		log.Print(err)
	// 		debug.PrintStack()
	// 		return errors.New("Something went wrong!!!!")
	// 	}),
	// ))
	// log.Fatal(http.ListenAndServe(":8081", nil))

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
