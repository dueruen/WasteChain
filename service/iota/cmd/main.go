package main

import (
	"fmt"

	service "github.com/dueruen/WasteChain/service/iota/pkg"
)

var endpoint = "https://nodes.devnet.thetangle.org"

// The address we want to fetch all transactions for
// This should contain a set of transactions as provided by the previous example
// It won't work if you ran the previous example multiple times for the same address
//var address string // = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")
//var seed string

// We need a query object containing the address we want to look for

func main() {
	//address, seed, err := service.CreateAddress(endpoint)
	//address, _, err := service.CreateAddress(endpoint)

	// if err != nil {
	// 	panic("err")
	// }
	//service.Send(address, seed, endpoint)

	messages, err := service.Recive(address, endpoint)
	if err != nil {
		panic("err")
	}
	fmt.Println("Messages: ", messages)
}
