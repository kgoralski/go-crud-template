package main

import (
	"github.com/kgoralski/go-crud-template/rest"
)

func main() {
	server := rest.NewServer()
	server.Start()
}
