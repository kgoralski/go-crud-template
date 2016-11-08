package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL string = "http://localhost:8080/rest/banks/"

func getAllBanks() ([]Bank, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	var banks []Bank
	json.NewDecoder(resp.Body).Decode(&banks)
	defer resp.Body.Close()
	return banks, nil
}

func getOneBank(id int) (*Bank, error) {
	resp, err := http.Get(fmt.Sprintf(baseURL+"%d", id))
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
	r, err := http.Post(baseURL, "text/plain", buf)
	if err != nil {
		return -1, err
	}
	var id int
	json.NewDecoder(r.Body).Decode(&id)

	return id, nil
}

func deleteBank(id int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(baseURL+"%d", id), nil)
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
	req, err := http.NewRequest(http.MethodDelete, baseURL, nil)
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
