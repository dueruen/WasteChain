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
