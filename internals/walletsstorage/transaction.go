package walletsstorage

import "time"

type Transaction struct {
	Timestamp time.Time `json:"timestamp"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Amount    float32   `json:"amount"`
}
