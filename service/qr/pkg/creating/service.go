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
	var code []byte
	code, err := qrcode.Encode(*dataString, qrCode.Medium, 256)
	return code, err
}
