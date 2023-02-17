package app

import (
	"encoding/json"
	"net/http"
)

func (h *App) ListWallets(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(h.walletsStorage.ListWallets())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
