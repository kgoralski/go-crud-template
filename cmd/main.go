package main

import (
	"github.com/kgoralski/go-crud-template/cmd/banks-api"
)

func main() {
	server := banks_api.NewServer()
	server.Start()
}
