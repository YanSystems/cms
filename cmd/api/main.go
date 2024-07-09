package main

import (
	"github.com/YanSystems/cms/pkg/server"
)

func main() {
	api := server.Server{}
	server := api.NewServer()
	api.Run(server)
}
