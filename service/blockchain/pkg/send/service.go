package send

import (
	"fmt"

	compress "github.com/dueruen/WasteChain/service/blockchain/pkg/compress"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/createAddress"
	iotaAPI "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
)

const tag = "WASTECHAIN"
const mwm = 9
const depth = 3

type Repository interface {
	GetShipmentAddress(shipmentID string) (addr string, err error)
	SaveShipmentInfo(shpmentID, addr, seed string)
}

type Service interface {
	Publish(shipmentID string, data []byte) error
}

type service struct {
	repo     Repository
	endpoint string
}

func NewService(repo Repository, endpoint string) Service {
	return &service{repo, endpoint}
}

func (srv *service) Publish(shipmentID string, data []byte) error {
	shipmentAddr, err := srv.repo.GetShipmentAddress(shipmentID)
	if err != nil {
		return err
	}
	if shipmentAddr == "" {
		shipmentAddr, shipmentSeed, err := createAddress.New(srv.endpoint)
		if err != nil {
			return err
		}
		srv.repo.SaveShipmentInfo(shipmentID, shipmentAddr, shipmentSeed)
	}
	err = sendToIOTA(shipmentAddr, data, srv.endpoint)
	if err != nil {
		return err
	}
	return nil
}

func sendToIOTA(shipmentAddr string, data []byte, endpoint string) error {
	// compose a new API instance, we provide no PoW function so this uses remote PoW
	httpAPI, err := iotaAPI.ComposeAPI(iotaAPI.HTTPClientSettings{URI: endpoint})
	if err != nil {
		return err
	}

	// convert a ascii message for the transaction to trytes,if possible
	compressData, err := compress.Compress(data)
	if err != nil {
		return err
	}

	message, err := converter.ASCIIToTrytes(compressData)
	if err != nil {
		return err
	}

	transfers := bundle.Transfers{
		{
			// must be 90 trytes long (include the checksum)
			Address: shipmentAddr,
			Value:   0,
			Message: message,
			Tag:     trinary.Trytes(tag),
		},
	}

	// We need to pass an options object, since we want to use the defaults it stays empty
	prepTransferOpts := iotaAPI.PrepareTransfersOptions{}

	trytes, err := httpAPI.PrepareTransfers(DEV_NET_SEED, transfers, prepTransferOpts)
	if err != nil {
		return err
	}

	// Send the transaction to the tangle using given depth and minimum weight magnitude
	bndl, err := httpAPI.SendTrytes(trytes, depth, mwm)
	if err != nil {
		return err
	}

	fmt.Println("\nbroadcasted bundle with tail tx hash: ", bundle.TailTransactionHash(bndl))
	// Check what the bundle looks like on thetangle!
	fmt.Printf("https://devnet.thetangle.org/bundle/%s\n\n", bndl[0].Bundle)
	return nil
}
