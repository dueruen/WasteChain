package event

import (
	"log"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"
	"github.com/nats-io/go-nats"
)

type eventHandler struct {
	natsConn *nats.EncodedConn
}

func NewEventHandler(url string) (*eventHandler, error) {
	conn, err := connectToNats(url)
	if err != nil {
		return nil, err
	}
	return &eventHandler{conn}, nil
}

func (handler *eventHandler) DoubleSignNeeded(event *pb.DoubleSignNeededEvent) {
	err := handler.natsConn.Publish(pb.SubjectTypes_DOUBLE_SIGN_NEEDED.String(), event)
	if err != nil {
		log.Fatal(err)
	}
}

func (handler *eventHandler) DoubleSignDone(event *pb.DoneEvent) {
	err := handler.natsConn.Publish(pb.SubjectTypes_SIGN_DONE.String(), event)
	if err != nil {
		log.Fatal(err)
	}
}

func (handler *eventHandler) SingleSignDone(event *pb.DoneEvent) {
	err := handler.natsConn.Publish(pb.SubjectTypes_SIGN_DONE.String(), event)
	if err != nil {
		log.Fatal(err)
	}
}

func connectToNats(url string) (encodedConn *nats.EncodedConn, err error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return
	}
	return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
}
