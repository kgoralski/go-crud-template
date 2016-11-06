package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	BZWBK     = "BZWBK"
	MBANK     = "MBANK"
	SANTANDER = "SANTANDER"
	ALIOR     = "ALIOR"
	ING       = "ING"
)

func TestGetBanksClient(t *testing.T) {
	deleteAllBanks()
	postBank(Bank{Name: BZWBK})
	postBank(Bank{Name: MBANK})
	banks, err := getAllBanks()
	panicOnErrInTest(err)
	assert.Len(t, banks, 2, "Expected size is 2")

}

func TestGetOneBankClient(t *testing.T) {
	id, err := postBank(Bank{Name: SANTANDER})
	panicOnErrInTest(err)
	bank, errQuery := getOneBank(id)
	panicOnErrInTest(errQuery)

	assert.Equal(t, SANTANDER, bank.Name, "Expected that both names are equal")
}

func TestCreateBankClient(t *testing.T) {
	bank := Bank{Name: ALIOR}
	id, err := postBank(bank)
	panicOnErrInTest(err)

	createdBank, errQuery := getOneBank(id)
	panicOnErrInTest(errQuery)

	assert.Equal(t, ALIOR, createdBank.Name, "Expected that both names are equal")
}

func TestDeleteSingleBankClient(t *testing.T) {
	id, err := postBank(Bank{Name: ING})
	panicOnErrInTest(err)
	deleteBank(id)
	banks, errQuery := getAllBanks()
	panicOnErrInTest(errQuery)

	for _, bank := range banks {
		assert.NotEqual(t, ING, bank)
	}
}

func TestDeleteAllBankClient(t *testing.T) {
	deleteAllBanks()
	postBank(Bank{Name: BZWBK})
	postBank(Bank{Name: MBANK})
	postBank(Bank{Name: ALIOR})
	deleteBanks()
	banks, err := getAllBanks()
	panicOnErrInTest(err)

	assert.Empty(t, banks)
}
