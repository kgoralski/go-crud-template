package main

import (
	"fmt"
	"testing"
)

func TestGetBanksClient(t *testing.T) {
	fmt.Println("TestGetBanksClient")
	deleteAllBanks()
	createBank(Bank{Name: "BZWBK"})
	createBank(Bank{Name: "MBANK"})
	banks := getAllBanks()

	if len(banks) != 2 {
		t.Fail()
	}
}

func TestGetOneBankClient(t *testing.T) {
	fmt.Println("TestGetOneBankClient")
	deleteAllBanks()
	id := createBank(Bank{Name: "Santander"})
	bank := getOneBank(int(id))

	if bank.Name != "Santander" {
		t.Fail()
	}
}

func TestCreateBankClient(t *testing.T) {
	fmt.Println("TestCreateBankClient")
	deleteAllBanks()
	bank := Bank{Name: "Alior"}
	id := postBank(bank)
	cretedBank := getOneBank(id)

	if cretedBank.Name != "Alior" {
		t.Fail()
	}
}
