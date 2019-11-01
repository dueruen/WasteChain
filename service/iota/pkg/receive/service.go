package receive

import (
	"bytes"
	"sort"
	"strings"

	compress "github.com/dueruen/WasteChain/service/iota/pkg/compress"
	iotaAPI "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
)

func Receive(address, endpoint string) ([]string, error) {
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

	messages := make([]string, 0)
	for _, buf := range buffers {
		suf := strings.Replace(buf.String(), "9", "", -1)
		msg, err := converter.TrytesToASCII(suf)
		if err != nil {
			return nil, err
		}
		decompressData, err := compress.Decompress(msg)
		if err != nil {
			return nil, err
		}

		messages = append(messages, string(decompressData))
	}
	return messages, nil
}
