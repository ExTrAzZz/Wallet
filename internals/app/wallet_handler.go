package app

import (
	"encoding/json"
	"net/http"
	"strings"
)

type walletBalance struct {
	Balance float32 `json:"balance"`
}

func (h *App) WalletHalndler(w http.ResponseWriter, r *http.Request) {
	requestParts := strings.Split(r.RequestURI, "/")[1:]
	if len(requestParts) < 3 {
		http.Error(w, "query is not valid", http.StatusBadRequest)
		return
	}
	walletAddr := requestParts[1]
	action := requestParts[2]

	switch action {
	case "balance":
		wallet, err := h.walletsStorage.GetWallet(walletAddr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(&walletBalance{
			Balance: wallet.Amount,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		break
	default:
		http.Error(w, "action is not valid", http.StatusBadRequest)
		return
	}

}
