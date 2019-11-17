package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
)

func main() {
	key, _ := CreateKeyPair()
	fmt.Println("KEY: ", key)
	res, _ := json.Marshal(&key)
	fmt.Println(string(res))
	un := rsa.PrivateKey{}
	json.Unmarshal(res, &un)
	fmt.Println(un)
}

func CreateKeyPair() (privateKey *rsa.PrivateKey, err error) {
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
