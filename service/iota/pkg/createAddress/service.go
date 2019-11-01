package createAddress

import (
	"fmt"
	"os/exec"

	srv "github.com/dueruen/WasteChain/service/iota/pkg"
	"github.com/iotaledger/iota.go/address"
	"github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/trinary"
)

func CreateAddress(endpoint string) (res *srv.TransportInfo, err error) {
	rawSeed, err := createSeed()
	if err != nil {
		return &srv.TransportInfo{}, err
	}
	seed := trinary.Trytes(rawSeed)

	// compose a new API instance
	//api, err := iotaAPI.ComposeAPI(iotaAPI.HTTPClientSettings{URI: endpoint})
	//must(err)

	// GetNewAddress retrieves the first unspent from address through IRI
	//addresses, err := api.GetNewAddress(seed, iotaAPI.GetNewAddressOptions{})
	address, err := address.GenerateAddress(seed, 1, consts.SecurityLevelMedium, true)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nYour new address: ", address)
	return &srv.TransportInfo{Address: address, Seed: seed}, nil
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
