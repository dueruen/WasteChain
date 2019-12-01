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

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}
	acco := os.Getenv("ACCO")
	if len(acco) == 0 {
		acco = "localhost:50051"
	}

	accountConn, err := grpc.Dial(acco, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to account service %v", err)
	}
	defer accountConn.Close()
	accountService := pb.NewAccountServiceClient(accountConn)
	fmt.Printf("Connection to account service made\n")

	resolver := graphql.Resolver{
		AccountClient: accountService,
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
	log.Fatal(http.ListenAndServe(port, nil))
}
