package app

import (
	"encoding/json"
	"net/http"
)

type transferRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float32 `json:"amount"`
}

func (h *App) Transfer(w http.ResponseWriter, r *http.Request) {
	tr := &transferRequest{}
	err := json.NewDecoder(r.Body).Decode(tr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.walletsStorage.Transfer(tr.From, tr.To, tr.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
