package app

import (
	"ewallet/cmd/appconfig"
	"ewallet/internals/walletsstorage"
	"net/http"
)

type App struct {
	cfg            *appconfig.AppConfig
	walletsStorage *walletsstorage.WalletsStorage
}

func New(cfg *appconfig.AppConfig) (*App, error) {
	ws, err := walletsstorage.New(cfg)
	if err != nil {
		return nil, err
	}
	return &App{
		cfg:            cfg,
		walletsStorage: ws,
	}, nil
}

func (h *App) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200 OK"))
}

func (h *App) Start(ch chan error) {
	http.HandleFunc("/health", h.Health)
	http.HandleFunc("/wallet/list", h.ListWallets)
	http.HandleFunc("/transactions", h.ListTxs)
	http.HandleFunc("/wallet/", h.WalletHalndler)
	http.HandleFunc("/send", h.Transfer)

	ch <- http.ListenAndServe(h.cfg.Hostname, nil)
}
