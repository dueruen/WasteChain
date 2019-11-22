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
	address, err := address.GenerateAddress(seed, 1, consts.SecurityLevelMedium, true)
	fmt.Println("NEW addr: ", address)
	fmt.Println("LEN: ", len(address))
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
