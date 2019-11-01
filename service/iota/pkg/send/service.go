package send

import (
	"fmt"

	srv "github.com/dueruen/WasteChain/service/iota/pkg"
	compress "github.com/dueruen/WasteChain/service/iota/pkg/compress"
	iotaAPI "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
)

var data = "{\"TEST\":\"A lot of test data\",\"sign1\":\"148e6b80ac74340bca25aa126bf010542c72d2bf7483216a63d1742f30e1948d174a08df12ad5e3a90291b20b1349edcd3272c79726b554ecf1df682eaeba2949ce4b158b2383c0bb25383242b25d2e821ceacc834811d8fb13cf61457df81dcfb563c4b0f141efd09fd5d11d4efb349ee7edb5e7f4bdf6026026072361695d85016eb01e08e1985f87dcbcc0b74e16306b599d1976134fb30515c461066aa53de5a3fd7277a54ad3c08609b96d78423f90ebf79cbbcf738f9eafd9009e0c69edc430f55246d1ad21596cb7af172cbc6e22d88f6680ef218ce2f2962aaa9f74d4131c09b0aaabdc2483639cb028145fe3df6e5d22242ac0bacebb248dcadbafb\",\"sign2\":\"148e6b80ac74340bca25aa126bf010542c72d2bf7483216a63d1742f30e1948d174a08df12ad5e3a90291b20b1349edcd3272c79726b554ecf1df682eaeba2949ce4b158b2383c0bb25383242b25d2e821ceacc834811d8fb13cf61457df81dcfb563c4b0f141efd09fd5d11d4efb349ee7edb5e7f4bdf6026026072361695d85016eb01e08e1985f87dcbcc0b74e16306b599d1976134fb30515c461066aa53de5a3fd7277a54ad3c08609b96d78423f90ebf79cbbcf738f9eafd9009e0c69edc430f55246d1ad21596cb7af172cbc6e22d88f6680ef218ce2f2962aaa9f74d4131c09b0aaabdc2483639cb028145fe3df6e5d22242ac0bacebb248dcadbafb\"}"

//var data = "H4sIAAAAAAAA/+2SPW4kMAiF7zJ1CsAY45xlGn6jNGlSrvbuYXavEckFtqwH7+P9eT6+Pz++8Pl4fz6QtcQVLA4vBg+jbYYk3oCwmeJQkvdhXYRishIPUy8ovKyvi4FmI1nuWnaBLjqB4+JbGbnojMY9o7g3VzRmi1JZudHlG8WOW52WrgB32lMQ0xRJpYRRFqGLFTG1RzdakPfJVsxo37KCHRoZqxNu5875ytX+GqFOpe86zT6Ngf6dQ0tQ7k7dgFIOWKBjSHfryfCYSQ4XyoIZ+97Ee2QcjSRs3MGCIGK2V9a21XnoHNtsOR5U4PqVPMq0+kINvRvu0Wdp37LOC3ALQl6EeEHvTSyJloT7SvixxsHmIUWUqi2iUE2oUdR0hczszlKScWFMP5gHz6BZk6zpBqQDqWuN6RqSNEhtXFnUMJ7FhaVb+/Px9j8O9BuH3zhMHP7+AO2BQTkfBAAA"

//var data = "{\"senderID\":\"TESTID\",\"test\":23213}"
var tag = "WASTECHAIN"

const mwm = 9
const depth = 3

func Send(senderInfo srv.TransportInfo, receiverInfo srv.TransportInfo, endpoint string) error {
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
	fmt.Println("Compress: ", compressData)

	message, err := converter.ASCIIToTrytes(compressData)
	if err != nil {
		return err
	}

	transfers := bundle.Transfers{
		{
			// must be 90 trytes long (include the checksum)
			Address: receiverInfo.Address,
			Value:   0,
			Message: message,
			Tag:     trinary.Trytes(tag),
		},
	}

	// We need to pass an options object, since we want to use the defaults it stays empty
	prepTransferOpts := iotaAPI.PrepareTransfersOptions{}

	trytes, err := httpAPI.PrepareTransfers(senderInfo.Seed, transfers, prepTransferOpts)
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
