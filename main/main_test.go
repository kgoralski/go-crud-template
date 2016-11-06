package main

import "log"

func init() {
	go startServer()
	deleteAllBanks()
}

func errTestPanic(err error) {
	if err != nil {
		switch e := err.(type) {
		case *httpError:
			log.Print(e)
			panic(e)
		default:
			log.Print(err)
			panic(err)
		}
	}
}

