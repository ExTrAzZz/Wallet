package main

import (
	"ewallet/cmd/appconfig"
	"ewallet/internals/app"
)

func main() {
	config, err := appconfig.LoadAppConfig()
	if err != nil {
		panic(err)
	}
	application, err := app.New(config)
	if err != nil {
		panic(err)
	}

	waiter := make(chan error, 1)
	application.Start(waiter)
	panic(<-waiter)
}
