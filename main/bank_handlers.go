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
	db, err := NewBankAPI()
	if err != nil {
		handleErrors(w, &dbError{err, dbConnectionFail})
		return
	}
	banks, err := getBanks(db)
	if err != nil {
		handleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(banks); err != nil {
		handleErrors(w, err)
		return
	}
}

func getBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleErrors(w, &httpError{err, http.StatusText(http.StatusBadRequest), http.StatusBadRequest})
		return
	}
	db, err := NewBankAPI()
	if err != nil {
		handleErrors(w, &dbError{err, dbConnectionFail})
		return
	}
	b, err := getBankByID(id, db)
	if err != nil {
		handleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(b); err != nil {
		handleErrors(w, err)
		return
	}
}

func createBankHanlder(w http.ResponseWriter, r *http.Request) {
	var bank Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		handleErrors(w, err)
		return
	}
	db, err := NewBankAPI()
	if err != nil {
		handleErrors(w, &dbError{err, dbConnectionFail})
		return
	}
	id, err := createBank(bank, db)
	if err != nil {
		handleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(id); err != nil {
		handleErrors(w, err)
		return
	}
}

func deleteBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleErrors(w, &httpError{err, http.StatusText(http.StatusBadRequest), http.StatusBadRequest})
		return
	}
	db, err := NewBankAPI()
	if err != nil {
		handleErrors(w, &dbError{err, dbConnectionFail})
		return
	}
	if err = deleteBankByID(id, db); err != nil {
		handleErrors(w, err)
		return
	}
}

func updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleErrors(w, &httpError{err, http.StatusText(http.StatusBadRequest), http.StatusBadRequest})
		return
	}
	var bank Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		handleErrors(w, err)
		return
	}
	db, err := NewBankAPI()
	if err != nil {
		handleErrors(w, &dbError{err, dbConnectionFail})
		return
	}
	updatedBank, err := updateBank(Bank{id, bank.Name}, db)
	if err != nil {
		handleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(updatedBank); err != nil {
		handleErrors(w, err)
		return
	}
}

func deleteAllBanksHandler(w http.ResponseWriter, r *http.Request) {
	db, err := NewBankAPI()
	if err != nil {
		handleErrors(w, &dbError{err, dbConnectionFail})
		return
	}
	if err := deleteAllBanks(db); err != nil {
		handleErrors(w, err)
		return
	}
}
