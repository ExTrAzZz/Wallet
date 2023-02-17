package walletsstorage

import (
	"errors"
	"ewallet/cmd/appconfig"
	"os"
	"sort"
	"sync"
	"time"
)

const (
	DEFAULT_CASH_AMOUNT    = 100.0
	DEFAULT_ADDRESS_LENGTH = 10
	DEFAULT_COUNT_WALLET   = 10
)

type WalletsStorage struct {
	cfg *appconfig.AppConfig

	addrWallets  map[string]*Wallet
	wallets      []*Wallet
	transactions []*Transaction

	ioBlocker sync.Mutex

	syncChan chan int
}

func New(cfg *appconfig.AppConfig) (*WalletsStorage, error) {
	h := &WalletsStorage{
		cfg:          cfg,
		wallets:      []*Wallet{},
		transactions: []*Transaction{},
		ioBlocker:    sync.Mutex{},
		addrWallets:  make(map[string]*Wallet),
		syncChan:     make(chan int),
	}
	err := h.Init()
	if err != nil {
		return nil, err
	}
	go h.Sync(h.syncChan)
	return h, nil
}

func (h *WalletsStorage) Init() error {
	err := h.Load()
	if err != nil {
		if err == os.ErrNotExist {
			// Initialize new file
			h.wallets = make([]*Wallet, DEFAULT_COUNT_WALLET)
			for i := range h.wallets {
				h.wallets[i], err = CreateWallet()
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}
	for _, w := range h.wallets {
		h.addrWallets[w.Address] = w
	}
	return nil
}

func (h *WalletsStorage) ListWallets() []string {
	h.ioBlocker.Lock()
	defer h.ioBlocker.Unlock()
	buff := []string{}
	for _, w := range h.wallets {
		buff = append(buff, w.Address)
	}
	return buff
}

func (h *WalletsStorage) ListLastTxs(amount int) []*Transaction {
	h.ioBlocker.Lock()
	defer h.ioBlocker.Unlock()
	if amount > len(h.transactions) {
		amount = len(h.transactions)
	}

	buff := make([]*Transaction, len(h.transactions))
	copy(buff, h.transactions)

	sort.Slice(buff, func(i, j int) bool {
		return buff[i].Timestamp.After(buff[j].Timestamp)
	})
	return buff[:amount]
}

func (h *WalletsStorage) GetWallet(addr string) (*Wallet, error) {
	wallet, ok := h.addrWallets[addr]
	if !ok {
		return nil, errors.New("Wallet with such address not found")
	}
	return wallet, nil
}

func (h *WalletsStorage) Transfer(from, to string, amount float32) error {
	if amount < 0 {
		return errors.New("Not valid amount value")
	}
	h.ioBlocker.Lock()
	defer h.ioBlocker.Unlock()
	fromW, ok := h.addrWallets[from]
	if !ok {
		return errors.New("Origin wallet not found")
	}
	toW, ok := h.addrWallets[to]
	if !ok {
		return errors.New("Target wallet not found")
	}
	if fromW.Amount < amount {
		return errors.New("Origin wallet not has enough balance")
	}
	toW.Amount += amount
	fromW.Amount -= amount
	h.transactions = append(h.transactions, &Transaction{
		Timestamp: time.Now(),
		From:      fromW.Address,
		To:        toW.Address,
		Amount:    amount,
	})
	//Tell thread to save db to disk
	h.syncChan <- 1
	return nil
}
