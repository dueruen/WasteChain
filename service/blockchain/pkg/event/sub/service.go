package sub

import (
	"fmt"
	"time"

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
		handler.publishSrv.Publish(e.ShipmentID, e.Data)
	})
}

func connectToNats(url string) (encodedConn *nats.EncodedConn, err error) {
	i := 5
	for i > 0 {
		conn, err := nats.Connect(url)
		if err != nil {
			fmt.Println("Can't connect to nats, sleeping for 2 sec, err: ", err)
			time.Sleep(2 * time.Second)
			i--
			continue
		} else {
			fmt.Println("Connected to storanatsge")
			return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
		}
	}
	panic("Not connected to storage")
}
