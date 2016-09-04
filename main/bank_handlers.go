package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getBanksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m := getBanks()
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func getBankbyIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	m := getBankById(id)
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func createBankHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bank Bank
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &bank)
	id := createBank(bank)

	j, _ := json.Marshal(id)
	w.Write(j)
}

func deleteBankByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	deleteBankById(id)
}

func updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bank Bank
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &bank)
	updatedBank := updateBank(bank)

	j, _ := json.Marshal(updatedBank)
	w.Write(j)
}
