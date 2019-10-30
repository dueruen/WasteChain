package pkg

import (
	"fmt"

	. "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
)

// var endpoint = "https://nodes.devnet.thetangle.org"

// We need a dummy seed even though we don't sign, because the API requires a seed to send
//var seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")

//var address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

// Our data is very long here, it needs to be split over several transactions, 3 in this case
var data = "{\"senderID\":\"KZZDTHNQWCNQ9RKHNI7TM2LB40242WHZ5D5130J7ZQHJNS9HT1ELFJIV86D2ENHEXO54ES4WBEOXBX2JD\",\"test\":23213}"
var tag = "WASTECHAIN"

const mwm = 9
const depth = 3

func Send(address, seed, endpoint string) {
	// compose a new API instance, we provide no PoW function so this uses remote PoW
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	must(err)

	// convert a ascii message for the transaction to trytes,if possible
	message, err := converter.ASCIIToTrytes(data)
	must(err)

	transfers := bundle.Transfers{
		{
			// must be 90 trytes long (include the checksum)
			Address: address,
			Value:   0,
			Message: message,
			Tag:     trinary.Trytes(tag),
		},
	}
	// We need to pass an options object, since we want to use the defaults it stays empty
	prepTransferOpts := PrepareTransfersOptions{}

	trytes, err := api.PrepareTransfers(seed, transfers, prepTransferOpts)
	must(err)

	// Send the transaction to the tangle using given depth and minimum weight magnitude
	bndl, err := api.SendTrytes(trytes, depth, mwm)
	must(err)

	fmt.Println("\nbroadcasted bundle with tail tx hash: ", bundle.TailTransactionHash(bndl))
	// Check what the bundle looks like on thetangle!
	fmt.Printf("https://devnet.thetangle.org/bundle/%s\n\n", bndl[0].Bundle)
}

// func must(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
