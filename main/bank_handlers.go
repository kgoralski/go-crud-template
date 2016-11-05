package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBanksHandler(w http.ResponseWriter, r *http.Request) {
	// You can set wrap your handler in another handler - functions are types in go (function can implement interface)
	w.Header().Set("Content-Type", "application/json")
	m := getBanks()

	// instead: json.NewEncoder(w).Encode(m)
	b, err := json.Marshal(m)

	checkErr(err)
	w.Write(b)
}

func getBankbyIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	checkErr(err)
	m := getBankById(id)
	b, err := json.Marshal(m)

	checkErr(err)
	w.Write(b)
}

func createBankHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bank Bank
	b, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	json.Unmarshal(b, &bank)
	id := createBank(bank)

	j, err := json.Marshal(id)
	checkErr(err)
	w.Write(j)
}

func deleteBankByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	checkErr(err)
	deleteBankById(id)
}

func updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bank Bank
	b, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	json.Unmarshal(b, &bank)
	updatedBank := updateBank(bank)

	j, err := json.Marshal(updatedBank)
	checkErr(err)
	w.Write(j)
}

func deleteAllBanksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deleteAllBanks()
}
