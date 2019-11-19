package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"
	"google.golang.org/grpc"
)

const (
	constData = "Some data"
)

func main() {
	fmt.Println("PUB: Signature client")
	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer cc.Close()

	client := pb.NewSignatureServiceClient(cc)

	_, err = client.CreateKeys(context.Background(), &pb.CreateKeysRequest{
		UserID:     "user1",
		Passphrase: "user1",
	})
	if err != nil {
		fmt.Println("ERR CreateKey user1: ", err)
	}

	_, err = client.CreateKeys(context.Background(), &pb.CreateKeysRequest{
		UserID:     "user2",
		Passphrase: "user2",
	})
	if err != nil {
		fmt.Println("ERR CreateKey user2: ", err)
	}

	signRes, err := client.SingleSign(context.Background(), &pb.SingleSignRequest{
		UserID:   "user1",
		Password: "user1",
		Data:     []byte(constData),
	})
	if err != nil {
		fmt.Println("ERR SingleSign: ", err)
	} else {
		fmt.Println(signRes)
	}

	res, err := client.StartDoubleSign(context.Background(), &pb.StartDoubleSignRequest{
		Data:                  []byte(constData),
		CurrentHolderID:       "user1",
		CurrentHolderPassword: "user1",
	})
	if err != nil {
		fmt.Println("ERR StartDoubleSign: ", err)
	} else {
		fmt.Println(res)
	}
}
