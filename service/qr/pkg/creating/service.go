package creating

import (
	pb "github.com/dueruen/WasteChain/service/qr/gen/proto"
	qrcode "github.com/skip2/go-qrcode"
)

type EventHandler interface {
	QRCreated(event *pb.QRCreatedEvent)
}

type Service interface {
	CreateQRCode(id, dataString string) (*[]byte, error)
}

type service struct {
	eventHandler EventHandler
}

func NewService(eventHandler EventHandler) Service {
	return &service{eventHandler}
}

func (srv *service) CreateQRCode(id, dataString string) (*[]byte, error) {
	code, err := qrcode.Encode(dataString, qrcode.Medium, 256)
	//Public event
	srv.eventHandler.QRCreated(&pb.QRCreatedEvent{
		ID:     id,
		QRCode: code,
	})

	return &code, err
}
