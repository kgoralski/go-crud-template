package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBanks(t *testing.T) {
	deleteAllBanks()
	fmt.Println("TestGetBanks")
	createBank(Bank{Name: "BZWBK"})
	createBank(Bank{Name: "MBANK"})

	banks := getBanks()
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	fmt.Println("TestCreateBank")
	id := createBank(Bank{Name: "MBANK"})

	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	fmt.Println("TestDeleteAllBanks")
	createBank(Bank{Name: "BZWBK"})
	createBank(Bank{Name: "MBANK"})
	deleteAllBanks()
	banks := getBanks()
	assert.Empty(t, banks)
}

func TestGetBankById(t *testing.T) {
	fmt.Println("TestGetBankById")
	id := createBank(Bank{Name: "Santander"})
	bank := getBankById(int(id))

	assert.Equal(t, bank.Name, "Santander")
}

func TestDeleteBankById(t *testing.T) {
	fmt.Println("TestDeleteBankById")
	id := createBank(Bank{Name: "MBANK"})

	deleteBankById(int(id))
	banks := getBanks()

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	fmt.Println("TestUpdateBank")
	id := createBank(Bank{Name: "MBANK"})
	bank := updateBank(Bank{Id: int(id), Name: "BZWBK"})

	assert.Equal(t, bank.Name, "BZWBK")
}
