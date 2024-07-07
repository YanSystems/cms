package main

import (
	api "github.com/YanSystems/cms/pkg/api"
)

func main() {
	api := api.Server{}
	server := api.NewServer()
	api.Run(server)
}
