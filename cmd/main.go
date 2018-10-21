package main

import (
	"github.com/kgoralski/go-crud-template/cmd/servid"
)

func main() {
	server := servid.NewServer()
	server.Start()
}
