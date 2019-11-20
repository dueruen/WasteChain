package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"
	"github.com/nats-io/go-nats"
	"google.golang.org/grpc"
)

const (
	//Correct data
	constData = "ALL VERIFICATIONS SHOULD -- PASS"

//Check that it fails
//fmt.Println("ALL VERIFICATIONS SHOULD -- FAIL")
//constData = "ALL VERIFICATIONS SHOULD -- FAIL"
)

func main() {
	fmt.Println("SUB: Signature client")
	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer cc.Close()

	fmt.Println("DATA: ", constData)

	client := pb.NewSignatureServiceClient(cc)

	err = Listen("localhost:4222", client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Infinite loop started...")
	for {
	}
}

func Listen(url string, signClient pb.SignatureServiceClient) (err error) {
	conn, err := connectToNats(url)
	if err != nil {
		return
	}

	conn.QueueSubscribe(pb.SignSubjectTypes_SIGN_DONE.String(), "done-queue", func(e *pb.DoneEvent) {
		switch e.EventType {
		case pb.DoneEventType_DOUBLE_SIGN_DONE:
			{
				fmt.Printf("Double sign done event\n")
				res, err := signClient.DoubleVerify(context.Background(), &pb.DoubleVerifyRequest{
					Data:                   []byte(constData),
					CurrentHolderID:        "user1",
					CurrentHolderSignature: e.CurrentHolderSignature,
					NewHolderID:            "user2",
					NewHolderSignature:     e.NewHolderSignature,
				})
				if err != nil {
					fmt.Println("Double sign Error: ", err)
				} else {
					fmt.Println("Double sign verified, ok: ", res.Ok)
				}
			}
		case pb.DoneEventType_SINGLE_SIGN_DONE:
			{
				fmt.Printf("Single sign done\n")
				res, err := signClient.SingleVerify(context.Background(), &pb.SingleVerifyRequest{
					UserID:    "user1",
					Data:      []byte(constData),
					Signature: e.CurrentHolderSignature,
				})
				if err != nil {
					fmt.Println("Single sign Error: ", err)
				} else {
					fmt.Println("Single sign verified, ok: ", res.Ok)
				}
			}
		default:

		}
	})

	conn.QueueSubscribe(pb.SignSubjectTypes_DOUBLE_SIGN_NEEDED.String(), "sign-queue", func(e *pb.DoubleSignNeededEvent) {
		fmt.Printf("Double sign event\n")
		fmt.Println("QR code is: ", e.QRCode)
		_, err := signClient.ContinueDoubleSign(context.Background(), &pb.ContinueDoubleSignRequest{
			ContinueID:        e.ContinueID,
			NewHolderID:       "user2",
			NewHolderPassword: "user2",
		})
		if err != nil {
			fmt.Println("ContinueDoubleSign ERROR: ", err)
		} else {
			fmt.Println("ContinueDoubleSign done")
		}
	})
	return
}

func connectToNats(url string) (encodedConn *nats.EncodedConn, err error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return
	}
	return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
}
