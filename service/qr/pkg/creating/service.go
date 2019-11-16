package creating

import (
    qrcode "github.com/skip2/go-qrcode"
    pb "github.com/dueruen/WasteChain/service/qr/gen/proto"
)

type Service interface {
	CreateQRCode(qrCode *pb.CreateQRCode) (*pb., error)
}


func (srv *service) CreateQRCode(string *dataString) (*[]byte, error) {
	var code []byte
	code, err := qrcode.Encode(dataString, qrCode.Medium, 256)
	return code, err
}
