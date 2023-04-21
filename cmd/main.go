package main

import (
	"exchange-api/cmd/server"
	"exchange-api/pkg/config"
)

func main() {

	config.Migrate()

	server.StartServer()
}
