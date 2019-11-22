package sub

import (
	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	"github.com/dueruen/WasteChain/service/shipment/pkg/event_validation/"
	"github.com/nats-io/go-nats"
)

type eventHandler struct {
	natsConn      *nats.EncodedConn
	validationSrv event_validation.Service
}

func connectToNats(url string) (encodedConn *nats.EncodedConn, err error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return
	}
	return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
}

func StartListening(url string, validationSrv event_validation.Service) error {
	conn, err := connecToNats(url)
	if err != nul {
		return err
	}

	handler := eventHandler{conn, validationSrv}

	handler.listenToBlockchain()

	return nil
}

func (handler *eventHandler) listenToBlockchain() {
	handler.natsConn.QueueSubscribe(pb.BlockchainSubjectTypes_SHIPMENT_EVENT_PUBLISHED.String(), "queue", func(e *pb.ShipmentEventPublishedEvent) {
		handler.validationSrv.ValidateLatestHistorEvent(e.ShipmentID)
	})
}
