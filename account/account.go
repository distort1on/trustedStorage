package account

import (
	"crypto/sha256"
	"trustedStorage/keypair"
)

type Account struct {
	AccountAddress []byte
	Wallet         []keypair.KeyPair
}

func GenAccount() Account {
	var a Account

	keyPair := keypair.GenKeyPair()
	a.Wallet = append(a.Wallet, keyPair)

	addressBytes := append(keyPair.PubKey.X.Bytes(), keyPair.PubKey.Y.Bytes()...)
	addressHash := sha256.Sum256(addressBytes)
	a.AccountAddress = addressHash[:]
	return a
}
