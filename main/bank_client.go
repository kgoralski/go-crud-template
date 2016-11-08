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
	if err := json.NewDecoder(resp.Body).Decode(&banks); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return banks, nil
}

func getOneBank(id int) (*Bank, error) {
	resp, err := http.Get(fmt.Sprintf(baseURL+"%d", id))
	if err != nil {
		return nil, err
	}
	var bank Bank
	if err := json.NewDecoder(resp.Body).Decode(&bank); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &bank, nil
}

func postBank(bank Bank) (int, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(bank); err != nil {
		return 0, err
	}
	r, err := http.Post(baseURL, "text/plain", buf)
	if err != nil {
		return 0, err
	}
	var id int
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		return 0, err
	}
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
	if _, err := http.DefaultClient.Do(req); err != nil {
		return err
	}
	return nil
}
