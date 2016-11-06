package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getAllBanks() []Bank {
	// base url could be a const
	resp, err := http.Get("http://localhost:8080/rest/banks/")

	// https://go-proverbs.github.io:
	// - Don't panic.
	// Don't just check errors, handle them gracefully.
	// This (and all other methods) should return errors (or handle them instead of panic())
	checkErr(err)
	var banks []Bank
	// There's no need to use ioutil.ReadAll()
	// Instead you an do: err := json.NewDecoder(resp.Body).Decode(&banks)
	b, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	// json.Unmarshal also returns error
	json.Unmarshal(b, &banks)
	// This is only good here because panic() will be called on error - otherwise - null pointer dereference
	defer resp.Body.Close()
	return banks
}

func getOneBank(id int) Bank {
	idStr := strconv.Itoa(id)
	// fmt.Sprintf("http://localhost:8080/rest/banks/%s", id)
	resp, err := http.Get("http://localhost:8080/rest/banks/" + idStr)

	// As aboce
	checkErr(err)
	var bank Bank
	b, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	json.Unmarshal(b, &bank)
	defer resp.Body.Close()
	return bank
}

func postBank(bank Bank) int {
	// Handle errors :)
	buf, _ := json.Marshal(bank)
	// Instead:
	// var buf bytes.Buffer
	// err := json.NewEncoder(buf).Encode(bank)
	body := bytes.NewBuffer(buf)
	r, err := http.Post("http://localhost:8080/rest/banks/", "text/plain", body)
	checkErr(err)
	response, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	var id int
	json.Unmarshal(response, &id)

	return id
}

func deleteBank(id int) {
	// fmt.Sprintf
	idStr := strconv.Itoa(id)
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/rest/banks/"+idStr, nil)
	// handle errors
	resp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()
}

func deleteBanks() {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/rest/banks/", nil)
	// handle errors
	resp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()
}
