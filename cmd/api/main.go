package main

func main() {
	api := Server{}
	server := api.NewServer()
	api.Run(server)
}
