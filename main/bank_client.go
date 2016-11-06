package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const BASE_URL string = "http://localhost:8080/rest/banks/"

func getAllBanks() ([]Bank, error) {
	resp, err := http.Get(BASE_URL)
	if err != nil {
		return nil, err
	}
	var banks []Bank
	json.NewDecoder(resp.Body).Decode(&banks)
	defer resp.Body.Close()
	return banks, nil
}

func getOneBank(id int) (*Bank, error) {
	resp, err := http.Get(fmt.Sprintf(BASE_URL+"%d", id))
	if err != nil {
		return nil, err
	}
	var bank Bank
	json.NewDecoder(resp.Body).Decode(&bank)
	defer resp.Body.Close()
	return &bank, nil
}

func postBank(bank Bank) (int, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(bank)
	if err != nil {
		return -1, err
	}
	r, err := http.Post(BASE_URL, "text/plain", buf)
	if err != nil {
		return -1, err
	}
	var id int
	json.NewDecoder(r.Body).Decode(&id)

	return int(id), nil
}

func deleteBank(id int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(BASE_URL+"%d", id), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func deleteBanks() error {
	req, err := http.NewRequest(http.MethodDelete, BASE_URL, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
