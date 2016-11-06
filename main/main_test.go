package main

func init() {
	go startServer()
	deleteAllBanks()
}

func panicInTest(err error) {
	if err != nil {
		panic(err)
	}
}
func checkHttpErr(err *HttpError) {
	if err != nil {
		panic(err)
	}
}
