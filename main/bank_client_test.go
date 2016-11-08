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
	deleteAllBanks()
	postBank(Bank{Name: bzwbk})
	postBank(Bank{Name: mbank})
	banks, err := getAllBanks()
	logFatalOnTest(err)
	assert.Len(t, banks, 2, "Expected size is 2")

}

func TestGetOneBankClient(t *testing.T) {
	id, err := postBank(Bank{Name: santander})
	logFatalOnTest(err)
	bank, errQuery := getOneBank(id)
	logFatalOnTest(errQuery)

	assert.Equal(t, santander, bank.Name, "Expected that both names are equal")
}

func TestCreateBankClient(t *testing.T) {
	bank := Bank{Name: alior}
	id, err := postBank(bank)
	logFatalOnTest(err)

	createdBank, errQuery := getOneBank(id)
	logFatalOnTest(errQuery)

	assert.Equal(t, alior, createdBank.Name, "Expected that both names are equal")
}

func TestDeleteSingleBankClient(t *testing.T) {
	id, err := postBank(Bank{Name: ing})
	logFatalOnTest(err)
	deleteBank(id)
	banks, errQuery := getAllBanks()
	logFatalOnTest(errQuery)

	for _, bank := range banks {
		assert.NotEqual(t, ing, bank)
	}
}

func TestDeleteAllBankClient(t *testing.T) {
	deleteAllBanks()
	postBank(Bank{Name: bzwbk})
	postBank(Bank{Name: mbank})
	postBank(Bank{Name: alior})
	deleteBanks()
	banks, err := getAllBanks()
	logFatalOnTest(err)

	assert.Empty(t, banks)
}
