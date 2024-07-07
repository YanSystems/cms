package main

import (
	"context"
	"log"

	utils "github.com/YanSystems/cms/pkg/utils/db"
)

func main() {
	client, err := utils.ConnectToDB("./../../.env")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	api := Server{
		Port: "8000",
		DB:   client.Database("content"),
	}

	server := api.NewServer()
	api.Run(server)
}
