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
	checkErr(err)
	var banks []Bank
	b, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	json.Unmarshal(b, &banks)
	defer resp.Body.Close()
	return banks
}

func getOneBank(id int) Bank {
	idStr := strconv.Itoa(id)
	resp, err := http.Get("http://localhost:8080/rest/banks/" + idStr)
	checkErr(err)
	var bank Bank
	b, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	json.Unmarshal(b, &bank)
	defer resp.Body.Close()
	return bank
}

func postBank(bank Bank) int {
	buf, _ := json.Marshal(bank)
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
	idStr := strconv.Itoa(id)
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/rest/banks/"+idStr, nil)
	resp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()
}

func deleteBanks() {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/rest/banks/", nil)
	resp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()
}
