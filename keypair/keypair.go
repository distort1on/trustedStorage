package keypair

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
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

func (k *KeyPair) PrintKeyPair() {
	fmt.Println(k.privKey, k.PubKey)
}

func (k *KeyPair) GetPrivateKey() *ecdsa.PrivateKey {
	return k.privKey
}

func EncodePrivateKey(privateKey *ecdsa.PrivateKey) []byte {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	return pemEncoded
}

func EncodePublicKey(publicKey *ecdsa.PublicKey) []byte {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	return pemEncodedPub
}

func DecodePublicKey(pemEncodedPub []byte) *ecdsa.PublicKey {
	blockPub, _ := pem.Decode(pemEncodedPub)

	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return publicKey
}

func DecodePrivateKey(pemEncoded []byte) *ecdsa.PrivateKey {
	block, _ := pem.Decode(pemEncoded)
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	return privateKey
}
