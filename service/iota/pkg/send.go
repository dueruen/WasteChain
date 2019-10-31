package pkg

import (
	"fmt"

	iotaAPI "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
)

var data = "DATA"

//var data = "{\"senderID\":\"TESTID\",\"test\":23213}"
var tag = "WASTECHAIN"

const mwm = 9
const depth = 3

func Send(senderInfo TransportInfo, receiverInfo TransportInfo, endpoint string) {
	// compose a new API instance, we provide no PoW function so this uses remote PoW
	httpAPI, err := iotaAPI.ComposeAPI(iotaAPI.HTTPClientSettings{URI: endpoint})
	must(err)

	// convert a ascii message for the transaction to trytes,if possible
	message, err := converter.ASCIIToTrytes(data)
	must(err)

	transfers := bundle.Transfers{
		{
			// must be 90 trytes long (include the checksum)
			Address: receiverInfo.Address,
			Value:   0,
			Message: message,
			Tag:     trinary.Trytes(tag),
		},
	}

	// create inputs for the transfer
	// inputs := []iotaAPI.Input{
	// 	{
	// 		// must be 90 trytes long (inlcude the checksum)
	// 		Address:  senderInfo.Address,
	// 		Security: consts.SecurityLevelMedium,
	// 		KeyIndex: 0,
	// 		Balance:  0,
	// 	},
	// }

	// We need to pass an options object, since we want to use the defaults it stays empty
	prepTransferOpts := iotaAPI.PrepareTransfersOptions{}
	//prepTransferOpts := iotaAPI.PrepareTransfersOptions{Inputs: inputs, RemainderAddress: &senderInfo.Address}

	trytes, err := httpAPI.PrepareTransfers(senderInfo.Seed, transfers, prepTransferOpts)
	must(err)

	// Send the transaction to the tangle using given depth and minimum weight magnitude
	bndl, err := httpAPI.SendTrytes(trytes, depth, mwm)
	must(err)

	fmt.Println("\nbroadcasted bundle with tail tx hash: ", bundle.TailTransactionHash(bndl))
	// Check what the bundle looks like on thetangle!
	fmt.Printf("https://devnet.thetangle.org/bundle/%s\n\n", bndl[0].Bundle)
}
