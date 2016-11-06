package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpError struct {
	Error   error
	Message string
	Code    int
}

func getBanksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	banks, err := getBanks()
	if err != nil {
		log.Print(err.Error)
		http.Error(w, err.Message, err.Code)
		return
	}
	json.NewEncoder(w).Encode(banks)
}

func getBankbyIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, errParse := strconv.Atoi(vars["id"])
	if errParse != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	b, err := getBankById(id)
	if err != nil {
		log.Print(err.Error)
		http.Error(w, err.Message, err.Code)
		return
	}
	json.NewEncoder(w).Encode(b)
}

func createBankHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bank Bank
	json.NewDecoder(r.Body).Decode(&bank)
	id, err := createBank(bank)
	if err != nil {
		log.Print(err.Error)
		http.Error(w, err.Message, err.Code)
		return
	}
	json.NewEncoder(w).Encode(id)
}

func deleteBankByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, errParse := strconv.Atoi(vars["id"])
	if errParse != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	err := deleteBankById(id)
	if err != nil {
		log.Print(err.Error)
		http.Error(w, err.Message, err.Code)
		return
	}
}

func updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bank Bank
	json.NewDecoder(r.Body).Decode(&bank)
	updatedBank, err := updateBank(bank)
	if err != nil {
		log.Print(err.Error)
		http.Error(w, err.Message, err.Code)
		return
	}
	json.NewEncoder(w).Encode(updatedBank)
}

func deleteAllBanksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := deleteAllBanks()
	if err != nil {
		log.Print(err.Error)
		http.Error(w, err.Message, err.Code)
		return
	}
}
