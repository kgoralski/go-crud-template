package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/rest/banks/", getBanksHandler).Methods("GET")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", getBankbyIdHandler).Methods("GET")
	r.HandleFunc("/rest/banks/", createBankHanlder).Methods("POST")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", deleteBankByIdHandler).Methods("DELETE")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", updateBankHanlder).Methods("PUT")
	r.HandleFunc("/rest/banks/all", deleteAllBanksHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
