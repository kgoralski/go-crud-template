package main

func init() {
	go startServer()
	deleteAllBanks()
}
