package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBanks(t *testing.T) {
	deleteAllBanks()
	createBank(Bank{Name: BZWBK})
	createBank(Bank{Name: MBANK})

	banks, err := getBanks()
	panicOnHttpErrInTest(err)
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	id, err := createBank(Bank{Name: MBANK})
	panicOnHttpErrInTest(err)
	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	createBank(Bank{Name: BZWBK})
	createBank(Bank{Name: MBANK})
	deleteAllBanks()
	banks, err := getBanks()
	panicOnHttpErrInTest(err)
	assert.Empty(t, banks)
}

func TestGetBankById(t *testing.T) {
	id, err := createBank(Bank{Name: SANTANDER})
	panicOnHttpErrInTest(err)
	bank, errQuery := getBankById(int(id))
	panicOnHttpErrInTest(errQuery)

	assert.Equal(t, SANTANDER, bank.Name)
}

func TestDeleteBankById(t *testing.T) {
	id, err := createBank(Bank{Name: MBANK})
	panicOnHttpErrInTest(err)

	deleteBankById(int(id))
	banks, errQuery := getBanks()
	panicOnHttpErrInTest(errQuery)

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	id, err := createBank(Bank{Name: MBANK})
	panicOnHttpErrInTest(err)
	bank, errQuery := updateBank(Bank{Id: int(id), Name: BZWBK})
	panicOnHttpErrInTest(errQuery)

	assert.Equal(t, BZWBK, bank.Name)
}
