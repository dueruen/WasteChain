package sub

import (
	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	"github.com/dueruen/WasteChain/service/shipment/pkg/eventvalidating"
	"github.com/nats-io/go-nats"
)

type eventHandler struct {
	natsConn      *nats.EncodedConn
	validationSrv eventvalidating.Service
}

func connectToNats(url string) (encodedConn *nats.EncodedConn, err error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return
	}
	return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
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
