package signature

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func SignData(privKey *ecdsa.PrivateKey, message []byte) []byte {
	messageHash := sha256.Sum256(message)
	sig, err := ecdsa.SignASN1(rand.Reader, privKey, messageHash[:])
	if err != nil {
		panic(err)
	}
	return sig
}

func VerifySignature(sig []byte, message []byte, pubKey *ecdsa.PublicKey) bool {
	messageHash := sha256.Sum256(message)
	return ecdsa.VerifyASN1(pubKey, messageHash[:], sig)
}

func PrintSignature(sig []byte) {
	fmt.Printf("%x\n", sig)
}
