package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBanks(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	deleteAllBanks(db)
	createBank(Bank{Name: bzwbk}, db)
	createBank(Bank{Name: mbank}, db)

	banks, err := getBanks(db)
	logFatalOnTest(t, err)
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	id, err := createBank(Bank{Name: mbank}, db)
	logFatalOnTest(t, err)
	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	createBank(Bank{Name: bzwbk}, db)
	createBank(Bank{Name: mbank}, db)
	deleteAllBanks(db)
	banks, err := getBanks(db)
	logFatalOnTest(t, err)
	assert.Empty(t, banks)
}

func TestGetBankById(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	id, err := createBank(Bank{Name: santander}, db)
	logFatalOnTest(t, err)
	bank, err := getBankByID(int(id), db)
	logFatalOnTest(t, err)

	assert.Equal(t, santander, bank.Name)
}

func TestDeleteBankById(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	id, err := createBank(Bank{Name: mbank}, db)
	logFatalOnTest(t, err)

	deleteBankByID(int(id), db)
	banks, err := getBanks(db)
	logFatalOnTest(t, err)

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)
	id, err := createBank(Bank{Name: mbank}, db)
	logFatalOnTest(t, err)
	bank, err := updateBank(Bank{ID: int(id), Name: bzwbk}, db)
	logFatalOnTest(t, err)

	assert.Equal(t, bzwbk, bank.Name)
}
