package keypair

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
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
	x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Println(err)
	}
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	return pemEncoded
}

func EncodePublicKey(publicKey *ecdsa.PublicKey) []byte {
	x509EncodedPub, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Println(err)
	}
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	return pemEncodedPub
}

func DecodePublicKey(pemEncodedPub []byte) (*ecdsa.PublicKey, error) {
	var err error

	blockPub, _ := pem.Decode(pemEncodedPub)
	if blockPub == nil {
		log.Println("Error while decoding public key")
		err = errors.New("Error while decoding public key")
		return nil, err
	}

	x509EncodedPub := blockPub.Bytes
	genericPublicKey, err := x509.ParsePKIXPublicKey(x509EncodedPub)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return publicKey, nil
}

func DecodePrivateKey(pemEncoded []byte) *ecdsa.PrivateKey {
	block, _ := pem.Decode(pemEncoded)
	if block == nil {
		log.Println("Error while decoding private key")
	}

	x509Encoded := block.Bytes
	privateKey, err := x509.ParseECPrivateKey(x509Encoded)
	if err != nil {
		log.Println(err)
	}

	return privateKey
}
