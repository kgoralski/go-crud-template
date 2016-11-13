package rest

import (
	"log"
	"net/http"

	"fmt"
	"github.com/gorilla/mux"
	"github.com/kgoralski/go-crud-template/dao"
)
var db, err = dao.NewBankAPI()

// StartServer starts server with REST handlers and initialise db connection pool
func StartServer() {
	if err != nil {
		log.Fatal(fmt.Errorf("FATAL: %+v\n", err))
	}
	dao.DBAccess = db

	r := mux.NewRouter()
	r.HandleFunc("/rest/banks/", commonHeaders(getBanksHandler)).Methods("GET")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", commonHeaders(getBankByIDHandler)).Methods("GET")
	r.HandleFunc("/rest/banks/", commonHeaders(createBankHanlder)).Methods("POST")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", commonHeaders(deleteBankByIDHandler)).Methods("DELETE")
	r.HandleFunc("/rest/banks/{id:[0-9]+}", commonHeaders(updateBankHanlder)).Methods("PUT")
	r.HandleFunc("/rest/banks/", commonHeaders(deleteAllBanksHandler)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
