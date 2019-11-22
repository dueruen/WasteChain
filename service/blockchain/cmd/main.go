package main

import "github.com/dueruen/WasteChain/service/blockchain/pkg"

func main() {
	pkg.Run()

	// info, err := createAddress.CreateAddress(endpoint)
	// fmt.Println("info: ", info, " err: ", err)
	//addr, _, err := service.CreateAddress(endpoint)

	// if err != nil {
	// 	panic("err")
	// }
	//fmt.Println("addr: ", addr, " seed: ", seed)
	//service.Send(addr, seed, endpoint)
	//sender.Send(service.TransportInfo{Address: addrSender, Seed: seedSender}, service.TransportInfo{Address: addrReceiver, Seed: seedReceiver}, endpoint)

	// messages, err := service.Receive(addrReceiver, endpoint)
	// if err != nil {
	// 	panic("err")
	// }
	// fmt.Println("Messages: ", messages)

	// remainderAddressNoCheck, err := address.GenerateAddress(seed, 1, consts.SecurityLevelMedium, false)
	// if err != nil {
	// 	// handle error
	// 	return
	// }
	// fmt.Println("NOaddress: ", remainderAddressNoCheck)
	// remainderAddress, err := address.GenerateAddress(seed, 1, consts.SecurityLevelMedium, true)
	// if err != nil {
	// 	// handle error
	// 	return
	// }
	// fmt.Println("WIaddress: ", remainderAddress)
	// fmt.Println("TRaddress: ", remainderAddress[0:81])
	// fmt.Println("CHaddress: ", remainderAddress[81:90])
}
