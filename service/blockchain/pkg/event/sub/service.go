package sub

import (
	"encoding/json"

	pb "github.com/dueruen/WasteChain/service/blockchain/gen/proto"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/publish"
	"github.com/nats-io/go-nats"
)

type eventHandler struct {
	natsConn   *nats.EncodedConn
	publishSrv publish.Service
}

func StartListening(url string, publishSrv publish.Service) error {
	conn, err := connectToNats(url)
	if err != nil {
		return err
	}
	handler := eventHandler{conn, publishSrv}

	handler.listenToSignature()

	return nil
}

type PublishData struct {
	CurrentHolderSignature []byte `json:"CHS"`
	NewHolderSignature     []byte `json:"NHS"`
}

func (handler *eventHandler) listenToSignature() {
	handler.natsConn.QueueSubscribe(pb.SignSubjectTypes_SIGN_DONE.String(), "queue", func(e *pb.DoneEvent) {
		data, _ := json.Marshal(&PublishData{
			CurrentHolderSignature: e.CurrentHolderSignature,
			NewHolderSignature:     e.NewHolderSignature,
		})
		handler.publishSrv.Publish(e.ShipmentID, data)
	})
}

func connectToNats(url string) (encodedConn *nats.EncodedConn, err error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return
	}
	return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
}
