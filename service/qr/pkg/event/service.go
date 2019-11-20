package event

import (
	"log"

	pb "github.com/dueruen/WasteChain/service/qr/gen/proto"
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

func (handler *eventHandler) QRCreated(event *pb.QRCreatedEvent) {
	err := handler.natsConn.Publish(pb.QRSubjectTypes_QR_CREATED.String(), event)
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
