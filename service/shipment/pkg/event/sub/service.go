package sub

import (
	"fmt"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	"github.com/dueruen/WasteChain/service/shipment/pkg/eventvalidating"
	"github.com/nats-io/go-nats"
)

type eventHandler struct {
	natsConn      *nats.EncodedConn
	validationSrv eventvalidating.Service
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

func StartListening(url string, validationSrv eventvalidating.Service) error {
	conn, err := connectToNats(url)
	if err != nil {
		return err
	}

	handler := eventHandler{conn, validationSrv}

	handler.listenToBlockchain()

	return nil
}

func (handler *eventHandler) listenToBlockchain() {
	handler.natsConn.QueueSubscribe(pb.BlockchainSubjectTypes_Published.String(),
		"queue", func(e *pb.PublishedEvent) {
			handler.validationSrv.ValidateLatestHistoryEvent(e.ShipmentID)
		})
}
