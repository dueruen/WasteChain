package receive

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"sort"

	compress "github.com/dueruen/WasteChain/service/blockchain/pkg/compress"
	iotaAPI "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
)

type Repository interface {
	GetShipmentAddress(shipmentID string) (addr string, err error)
}

type Service interface {
	GetShipmentData(shipmentID string) (error, [][]byte)
}

type service struct {
	repo     Repository
	endpoint string
}

func NewService(repo Repository, endpoint string) Service {
	return &service{repo, endpoint}
}

func (srv *service) GetShipmentData(shipmentID string) (error, [][]byte) {
	shipmentAddr, err := srv.repo.GetShipmentAddress(shipmentID)
	if err != nil {
		fmt.Println("BLOCK GetShipmentAddress err: ", err)
		return err, nil
	}
	if shipmentAddr == "" {
		fmt.Println("BLOCK Shipment don't exists!!")
		return errors.New("Shipment don't exists!!"), nil
	}
	history, err := receive(shipmentAddr, srv.endpoint)
	if err != nil {
		fmt.Println("BLOCK receive err: ", err)
		return err, nil
	}

	return nil, history
}

func receive(address, endpoint string) ([][]byte, error) {
	query := iotaAPI.FindTransactionsQuery{Addresses: trinary.Hashes{address}}

	api, err := iotaAPI.ComposeAPI(iotaAPI.HTTPClientSettings{URI: endpoint})
	if err != nil {
		return nil, err
	}

	// Find Transaction Objects uses the connected node to find transactions based on our query
	transactions, err := api.FindTransactionObjects(query)
	if err != nil {
		return nil, err
	}

	// We need to sort all transactions by index first so we can concatenate them
	sort.Slice(transactions[:], func(i, j int) bool {
		return transactions[i].CurrentIndex < transactions[j].CurrentIndex
	})

	// We define a buffer to concatenate all sorted transactions
	buffers := make([]bytes.Buffer, 0)
	//var buffer

	for i, tx := range transactions {
		// We add the sorted Transaction Signature Message Fragment to the buffer
		buffers = append(buffers, bytes.Buffer{})
		buffers[i].WriteString(tx.SignatureMessageFragment)
	}
	// fmt.Println("BLOCK BUFFER SIZE: ", len(buffers))
	messages := make([][]byte, 0)
	for _, buf := range buffers {
		r := regexp.MustCompile("\\d{2,}")
		suf := r.ReplaceAllString(buf.String(), "")
		msg, err := converter.TrytesToASCII(suf)
		if err != nil {
			fmt.Println("BLOCK TrytesToASCII err: ", err)
			return nil, errors.New("TrytesToASCII: " + err.Error())
		}
		decompressData, err := compress.Decompress(msg)
		if err != nil {
			fmt.Println("BLOCK Decompress err: ", err)
			return nil, errors.New("Decompress: " + err.Error())
		}
		messages = append(messages, decompressData)
	}
	return messages, nil
}
