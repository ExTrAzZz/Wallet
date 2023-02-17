package app

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *App) ListTxs(w http.ResponseWriter, r *http.Request) {

	amount, err := strconv.ParseInt(r.URL.Query().Get("count"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	txs := h.walletsStorage.ListLastTxs(int(amount))

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(txs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
