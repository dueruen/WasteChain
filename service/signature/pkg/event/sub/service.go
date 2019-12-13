package sub

import (
	"fmt"
	"time"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"
	"github.com/dueruen/WasteChain/service/signature/pkg/sign"
	"github.com/nats-io/go-nats"
)

type eventHandler struct {
	natsConn *nats.EncodedConn
	signSrv  sign.Service
}

func StartListening(url string, signSrv sign.Service) error {
	conn, err := connectToNats(url)
	if err != nil {
		return err
	}
	handler := eventHandler{conn, signSrv}

	handler.listenToQR()

	return nil
}

func (handler *eventHandler) listenToQR() {
	handler.natsConn.QueueSubscribe(pb.QRSubjectTypes_QR_CREATED.String(), "queue", func(e *pb.QRCreatedEvent) {
		handler.signSrv.FinishStartDoubleSign(e.ID, e.QRCode)
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
