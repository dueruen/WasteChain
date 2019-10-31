package pkg

import "testing"

func TestDoubleSign(t *testing.T) {
	producerPrivateKey, err := CreateKeyPair()
	if err != nil {
		t.Error(err)
	}
	transporterPrivateKey, err := CreateKeyPair()
	if err != nil {
		t.Error(err)
	}

	message := []byte("Hello")

	proHash, proSign, err := Sign(message, producerPrivateKey)
	if err != nil {
		t.Error(err)
	}

	traHash, traSign, err := Sign(proHash, transporterPrivateKey)
	if err != nil {
		t.Error(err)
	}

	if Verify(&transporterPrivateKey.PublicKey, traHash, traSign) != nil {
		t.Error(err)
	}

	if Verify(&producerPrivateKey.PublicKey, proHash, proSign) != nil {
		t.Error(err)
	}
}
