package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func CreateKeyPair() (privateKey *rsa.PrivateKey, err error) {
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func Sign(message []byte, privateKey *rsa.PrivateKey) (hashed, signature []byte, err error) {
	hash := sha256.New()
	hash.Write(message)

	opts := rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed = pssh.Sum(nil)
	signature, err = rsa.SignPSS(
		rand.Reader,
		privateKey,
		newhash,
		hashed,
		&opts,
	)

	if err != nil {
		return nil, nil, err
	}
	return hashed, signature, nil
}

func Verify(publicKey *rsa.PublicKey, hashed, signature []byte) (err error) {
	opts := rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}
	err = rsa.VerifyPSS(
		publicKey,
		crypto.SHA256,
		hashed,
		signature,
		&opts,
	)
	if err != nil {
		return err
	}
	return nil
}
