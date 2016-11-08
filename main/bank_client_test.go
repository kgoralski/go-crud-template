package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	bzwbk     = "BZWBK"
	mbank     = "MBANK"
	santander = "SANTANDER"
	alior     = "ALIOR"
	ing       = "ING"
)

func TestGetBanksClient(t *testing.T) {
	deleteBanks()
	postBank(Bank{Name: bzwbk})
	postBank(Bank{Name: mbank})
	banks, err := getAllBanks()
	logFatalOnTest(t, err)
	assert.Len(t, banks, 2, "Expected size is 2")

}

func TestGetOneBankClient(t *testing.T) {
	id, err := postBank(Bank{Name: santander})
	logFatalOnTest(t, err)
	bank, errQuery := getOneBank(id)
	logFatalOnTest(t, errQuery)

	assert.Equal(t, santander, bank.Name, "Expected that both names are equal")
}

func TestCreateBankClient(t *testing.T) {
	bank := Bank{Name: alior}
	id, err := postBank(bank)
	logFatalOnTest(t, err)

	createdBank, errQuery := getOneBank(id)
	logFatalOnTest(t, errQuery)

	assert.Equal(t, alior, createdBank.Name, "Expected that both names are equal")
}

func TestDeleteSingleBankClient(t *testing.T) {
	id, err := postBank(Bank{Name: ing})
	logFatalOnTest(t, err)
	deleteBank(id)
	banks, errQuery := getAllBanks()
	logFatalOnTest(t, errQuery)

	for _, bank := range banks {
		assert.NotEqual(t, ing, bank)
	}
}

func TestDeleteAllBankClient(t *testing.T) {
	deleteBanks()
	postBank(Bank{Name: bzwbk})
	postBank(Bank{Name: mbank})
	postBank(Bank{Name: alior})
	deleteBanks()
	banks, err := getAllBanks()
	logFatalOnTest(t, err)

	assert.Empty(t, banks)
}
