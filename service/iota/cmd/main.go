package main

import (
	"fmt"

	service "github.com/dueruen/WasteChain/service/iota/pkg"
)

var endpoint = "https://nodes.devnet.thetangle.org"

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
