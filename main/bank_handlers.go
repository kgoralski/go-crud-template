package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJSON)
		fn(w, r)
	}
}

func getBanksHandler(w http.ResponseWriter, r *http.Request) {
	banks, err := getBanks()
	if err != nil {
		handleHTTPError(w, err)
		return
	}
	json.NewEncoder(w).Encode(banks)
}

func getBankbyIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	b, err := getBankByID(id)
	if err != nil {
		handleHTTPError(w, err)
		return
	}
	json.NewEncoder(w).Encode(b)
}

func createBankHanlder(w http.ResponseWriter, r *http.Request) {
	var bank Bank
	json.NewDecoder(r.Body).Decode(&bank)
	id, err := createBank(bank)
	if err != nil {
		handleHTTPError(w, err)
		return
	}
	json.NewEncoder(w).Encode(id)
}

func deleteBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = deleteBankByID(id)
	if err != nil {
		handleHTTPError(w, err)
		return
	}
}

func updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var bank Bank
	json.NewDecoder(r.Body).Decode(&bank)
	updatedBank, err := updateBank(Bank{id, bank.Name})
	if err != nil {
		handleHTTPError(w, err)
		return
	}
	json.NewEncoder(w).Encode(updatedBank)
}

func deleteAllBanksHandler(w http.ResponseWriter, r *http.Request) {
	err := deleteAllBanks()
	if err != nil {
		handleHTTPError(w, err)
		return
	}
}
