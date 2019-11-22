package createAddress

import (
	"fmt"
	"os/exec"

	"github.com/iotaledger/iota.go/address"
	"github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/trinary"
)

func New(endpoint string) (addr, seed string, err error) {
	rawSeed, err := createSeed()
	if err != nil {
		return "", "", err
	}
	seed = trinary.Trytes(rawSeed)

	// compose a new API instance
	//api, err := iotaAPI.ComposeAPI(iotaAPI.HTTPClientSettings{URI: endpoint})
	//must(err)

	// GetNewAddress retrieves the first unspent from address through IRI
	//addresses, err := api.GetNewAddress(seed, iotaAPI.GetNewAddressOptions{})
	address, err := address.GenerateAddress(seed, 1, consts.SecurityLevelMedium, true)
	if err != nil {
		return "", "", err
	}
	return address, seed, nil
}

func createSeed() (string, error) {
	out, err := exec.Command("sh", "-c", "cat /dev/urandom |tr -dc A-Z9|head -c${1:-81}").Output()
	if err != nil {
		return "", err
	}
	seed := string(out)
	fmt.Println("Seed: ", seed)
	return seed, nil
}
