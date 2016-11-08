package main

import "log"

func init() {
	go startServer()
	deleteAllBanks()
}

func logFatalOnTest(err error) {
	if err != nil {
		switch e := err.(type) {
		case *httpError:
			log.Fatal(e)
		default:
			log.Fatal(err)
		}
	}
}
