package walletsstorage

import (
	"crypto/rand"
	"encoding/hex"
)

type Wallet struct {
	Address string  `json:"address"`
	Amount  float32 `json:"amount"`
}

func CreateWallet() (*Wallet, error) {
	b := make([]byte, DEFAULT_ADDRESS_LENGTH)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	wal := &Wallet{
		Address: hex.EncodeToString(b),
		Amount:  DEFAULT_CASH_AMOUNT,
	}
	return wal, nil
}
