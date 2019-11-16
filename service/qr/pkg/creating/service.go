package creating

import (
	qrcode "github.com/skip2/go-qrcode"
)

type Service interface {
	CreateQRCode(dataString *string) (*[]byte, error)
}

type service struct {
}

func (srv *service) CreateQRCode(dataString *string) (*[]byte, error) {
	code, err := qrcode.Encode(*dataString, qrcode.Medium, 256)
	return &code, err
}

func NewService() Service {
	return &service{}
}
