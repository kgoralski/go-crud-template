package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBanksClient(t *testing.T) {
	// You don't need to print test name, run tests with -v: go test -v (will print test name and stdout)
	fmt.Println("TestGetBanksClient")
	deleteAllBanks()
	createBank(Bank{Name: "BZWBK"})
	createBank(Bank{Name: "MBANK"})
	banks := getAllBanks()

	// Wrong order - assert.Len(t, expected, actual, "msg")
	assert.Len(t, banks, 2, "Expected size is 2")

}

func TestGetOneBankClient(t *testing.T) {
	fmt.Println("TestGetOneBankClient")
	id := createBank(Bank{Name: "Santander"})
	bank := getOneBank(int(id))

	// Wrong order - assert.Len(t, expected, actual, "msg")
	assert.Equal(t, bank.Name, "Santander", "Expected that both names are equal")
}

func TestCreateBankClient(t *testing.T) {
	fmt.Println("TestCreateBankClient")
	bank := Bank{Name: "Alior"}
	id := postBank(bank)
	createdBank := getOneBank(id)

	// Wrong order - assert.Equal(t, expected, actual, "msg")
	assert.Equal(t, createdBank.Name, "Alior", "Expected that both names are equal")
}

func TestDeleteSingleBankClient(t *testing.T) {
	fmt.Println("TestDeleteSingleBankClient")
	id := createBank(Bank{Name: "Santander"})
	deleteBank(id)
	banks := getAllBanks()

	// Wrong order - assert.NotEqual(t, expected, actual)
	for _, bank := range banks {
		assert.NotEqual(t, bank, "Santander")
	}
}

func TestDeleteAllBankClient(t *testing.T) {
	fmt.Println("TestDeleteAllBankClient")
	createBank(Bank{Name: "BZWBK"})
	createBank(Bank{Name: "MBANK"})
	createBank(Bank{Name: "ING"})
	deleteBanks()
	banks := getAllBanks()

	assert.Empty(t, banks)
}
