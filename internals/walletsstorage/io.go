package walletsstorage

import (
	"encoding/json"
	"fmt"
	"os"
)

type saveFile struct {
	Wallets      []*Wallet      `json:"wallets"`
	Transactions []*Transaction `json:"transactions"`
}

func (h *WalletsStorage) Load() error {
	h.ioBlocker.Lock()
	defer h.ioBlocker.Unlock()
	file, err := os.Open(h.cfg.FileName)
	if err != nil {
		return err
	}
	defer file.Close()
	dump := &saveFile{}
	err = json.NewDecoder(file).Decode(dump)
	if err != nil {
		return err
	}
	h.transactions = dump.Transactions
	h.wallets = dump.Wallets
	return nil
}

func (h *WalletsStorage) Save() error {

	dump := &saveFile{
		Wallets:      h.wallets,
		Transactions: h.transactions,
	}
	h.ioBlocker.Lock()
	defer h.ioBlocker.Unlock()
	file, err := os.Create(h.cfg.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(dump)
	if err != nil {
		return err
	}
	return nil
}

func (h *WalletsStorage) Sync(save chan int) {
	for {
		<-save
		err := h.Save()
		if err != nil {
			fmt.Printf("Failed to save to disk:%s\n", err.Error())
		} else {
			fmt.Println("Saved to disk")
		}
	}
}
