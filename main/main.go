package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	startServer()
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/rest/banks/", commonHeaders(getBanksHandler)).Methods("GET")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", commonHeaders(getBankbyIdHandler)).Methods("GET")
	r.HandleFunc("/rest/banks/", commonHeaders(createBankHanlder)).Methods("POST")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", commonHeaders(deleteBankByIdHandler)).Methods("DELETE")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", commonHeaders(updateBankHanlder)).Methods("PUT")
	r.HandleFunc("/rest/banks/", commonHeaders(deleteAllBanksHandler)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
