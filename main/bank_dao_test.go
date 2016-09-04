package main

import (
	"fmt"
	"testing"
)

func TestGetBanks(t *testing.T) {
	fmt.Println("TestGetBanks")
	deleteAllBanks()
	createBank(Bank{Name: "BZWBK"})
	createBank(Bank{Name: "MBANK"})

	banks := getBanks()
	if len(banks) != 2 {
		t.Fail()
	}
}

func TestCreateBank(t *testing.T) {
	fmt.Println("TestCreateBank")
	deleteAllBanks()
	id := createBank(Bank{Name: "MBANK"})

	if id == 0 {
		t.Fail()
	}
}

func TestDeleteAllBanks(t *testing.T) {
	fmt.Println("TestDeleteAllBanks")
	createBank(Bank{Name: "BZWBK"})
	createBank(Bank{Name: "MBANK"})
	deleteAllBanks()
	banks := getBanks()
	if len(banks) != 0 {
		t.Fail()
	}
}

func TestGetBankById(t *testing.T) {
	fmt.Println("TestGetBankById")
	deleteAllBanks()
	id := createBank(Bank{Name: "Santander"})
	bank := getBankById(int(id))
	if bank.Name != "Santander" {
		t.Fail()
	}
}

func TestDeleteBankById(t *testing.T) {
	fmt.Println("TestDeleteBankById")
	deleteAllBanks()
	id := createBank(Bank{Name: "Santander"})

	deleteBankById(int(id))
	banks := getBanks()
	if len(banks) != 0 {
		t.Fail()
	}
}

func TestUpdateBank(t *testing.T) {
	fmt.Println("TestUpdateBank")
	deleteAllBanks()
	id := createBank(Bank{Name: "MBANK"})

	bank := updateBank(Bank{Id: int(id), Name: "Santander"})
	if bank.Name != "Santander" {
		t.Fail()
	}
}
