package main

import (
	"testing"
)

func init() {
	go startServer()
	deleteAllBanks()
}

func logFatalOnTest(t *testing.T, err error) {
	if err != nil {
		switch e := err.(type) {
		case *httpError:
			t.Fatal(e)
		default:
			t.Fatal(err)
		}
	}
}
