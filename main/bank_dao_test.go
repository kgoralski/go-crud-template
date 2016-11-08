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
	logFatalOnTest(t, err)
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	id, err := createBank(Bank{Name: mbank})
	logFatalOnTest(t, err)
	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	createBank(Bank{Name: bzwbk})
	createBank(Bank{Name: mbank})
	deleteAllBanks()
	banks, err := getBanks()
	logFatalOnTest(t, err)
	assert.Empty(t, banks)
}

func TestGetBankById(t *testing.T) {
	id, err := createBank(Bank{Name: santander})
	logFatalOnTest(t, err)
	bank, err := getBankByID(int(id))
	logFatalOnTest(t, err)

	assert.Equal(t, santander, bank.Name)
}

func TestDeleteBankById(t *testing.T) {
	id, err := createBank(Bank{Name: mbank})
	logFatalOnTest(t, err)

	deleteBankByID(int(id))
	banks, err := getBanks()
	logFatalOnTest(t, err)

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	id, err := createBank(Bank{Name: mbank})
	logFatalOnTest(t, err)
	bank, err := updateBank(Bank{ID: int(id), Name: bzwbk})
	logFatalOnTest(t, err)

	assert.Equal(t, bzwbk, bank.Name)
}
