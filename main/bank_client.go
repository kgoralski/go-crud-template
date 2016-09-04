package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getAllBanks() []Bank {
	resp, err := http.Get("http://localhost:8080/rest/banks/")
	if err != nil {
		panic(err)
	}
	var banks []Bank
	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &banks)
	defer resp.Body.Close()
	return banks
}

func getOneBank(id int) Bank {
	idStr := strconv.Itoa(id)
	resp, err := http.Get("http://localhost:8080/rest/banks/" + idStr)

	if err != nil {
		panic(err)
	}
	var bank Bank
	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &bank)
	defer resp.Body.Close()
	return bank
}

func postBank(bank Bank) int {
	buf, _ := json.Marshal(bank)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("http://localhost:8080/rest/banks/", "text/plain", body)
	response, _ := ioutil.ReadAll(r.Body)
	var id int
	json.Unmarshal(response, &id)

	return id
}
