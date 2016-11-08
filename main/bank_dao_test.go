package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBanks(t *testing.T) {
	deleteAllBanks()
	createBank(Bank{Name: bzwbk})
	createBank(Bank{Name: mbank})

	banks, err := getBanks()
	logFatalOnTest(err)
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	id, err := createBank(Bank{Name: mbank})
	logFatalOnTest(err)
	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	createBank(Bank{Name: bzwbk})
	createBank(Bank{Name: mbank})
	deleteAllBanks()
	banks, err := getBanks()
	logFatalOnTest(err)
	assert.Empty(t, banks)
}

func TestGetBankById(t *testing.T) {
	id, err := createBank(Bank{Name: santander})
	logFatalOnTest(err)
	bank, errQuery := getBankByID(int(id))
	logFatalOnTest(errQuery)

	assert.Equal(t, santander, bank.Name)
}

func TestDeleteBankById(t *testing.T) {
	id, err := createBank(Bank{Name: mbank})
	logFatalOnTest(err)

	deleteBankByID(int(id))
	banks, errQuery := getBanks()
	logFatalOnTest(errQuery)

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	id, err := createBank(Bank{Name: mbank})
	logFatalOnTest(err)
	bank, errQuery := updateBank(Bank{ID: int(id), Name: bzwbk})
	logFatalOnTest(errQuery)

	assert.Equal(t, bzwbk, bank.Name)
}
