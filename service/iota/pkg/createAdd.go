package pkg

import (
	"fmt"
	"os/exec"

	iotaAPI "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/trinary"
)

// var endpoint = "https://nodes.devnet.thetangle.org"

func CreateAddress(endpoint string) (addr, seed string, err error) {
	rawSeed, err := createSeed()
	if err != nil {
		return "", "", err
	}
	seed = trinary.Trytes(rawSeed)

	// compose a new API instance
	api, err := iotaAPI.ComposeAPI(iotaAPI.HTTPClientSettings{URI: endpoint})
	must(err)

	// GetNewAddress retrieves the first unspent from address through IRI
	addresses, err := api.GetNewAddress(seed, iotaAPI.GetNewAddressOptions{})
	must(err)

	fmt.Println("\nYour new address: ", addresses[0])
	return addresses[0], seed, nil
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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
