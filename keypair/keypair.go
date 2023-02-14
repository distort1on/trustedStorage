package keypair

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type KeyPair struct {
	privKey *ecdsa.PrivateKey
	PubKey  *ecdsa.PublicKey
}

func GenKeyPair() KeyPair {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return KeyPair{privateKey, &privateKey.PublicKey}
}

func (k KeyPair) PrintKeyPair() {
	fmt.Println(k.privKey, k.PubKey)
}

func (k KeyPair) GetPrivateKey() *ecdsa.PrivateKey {
	return k.privKey
}
