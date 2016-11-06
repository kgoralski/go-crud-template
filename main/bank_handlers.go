package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	Content_Type     = "Content-Type"
	APPLICATION_JSON = "application/json"
)

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(Content_Type, APPLICATION_JSON)
		fn(w, r)
	}
}

func getBanksHandler(w http.ResponseWriter, r *http.Request) {
	banks, err := getBanks()
	if err != nil {
		log.Print(err.Err)
		http.Error(w, err.Message, err.Code)
		return
	}
	json.NewEncoder(w).Encode(banks)
}

func getBankbyIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, errParse := strconv.Atoi(vars["id"])
	if errParse != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	b, err := getBankById(id)
	if err != nil {
		log.Print(err.Err)
		http.Error(w, err.Message, err.Code)
		return
	}
	json.NewEncoder(w).Encode(b)
}

func createBankHanlder(w http.ResponseWriter, r *http.Request) {
	var bank Bank
	json.NewDecoder(r.Body).Decode(&bank)
	id, err := createBank(bank)
	if err != nil {
		log.Print(err.Err)
		http.Error(w, err.Message, err.Code)
		return
	}
	json.NewEncoder(w).Encode(id)
}

func deleteBankByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, errParse := strconv.Atoi(vars["id"])
	if errParse != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err := deleteBankById(id)
	if err != nil {
		log.Print(err.Err)
		http.Error(w, err.Message, err.Code)
		return
	}
}

func updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, errParse := strconv.Atoi(vars["id"])
	if errParse != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var bank Bank
	json.NewDecoder(r.Body).Decode(&bank)
	updatedBank, err := updateBank(Bank{id, bank.Name})
	if err != nil {
		log.Print(err.Err)
		http.Error(w, err.Message, err.Code)
		return
	}
	json.NewEncoder(w).Encode(updatedBank)
}

func deleteAllBanksHandler(w http.ResponseWriter, r *http.Request) {
	err := deleteAllBanks()
	if err != nil {
		log.Print(err.Err)
		http.Error(w, err.Message, err.Code)
		return
	}
}
