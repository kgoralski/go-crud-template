package dao

import (
	"testing"

	"github.com/kgoralski/go-crud-template/handleErr"
	"github.com/stretchr/testify/assert"
)

const (
	bzwbk     = "BZWBK"
	mbank     = "MBANK"
	santander = "SANTANDER"
)

func logFatalOnTest(t *testing.T, err error) {
	if err != nil {
		switch e := err.(type) {
		case *handleErr.HTTPError:
			t.Fatal(e)
		case *handleErr.DbError:
			t.Fatal(e)
		default:
			t.Fatal(err)
		}
	}
}

func TestGetBanks(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	DeleteAllBanks(db)
	CreateBank(Bank{Name: bzwbk}, db)
	CreateBank(Bank{Name: mbank}, db)

	banks, err := GetBanks(db)
	logFatalOnTest(t, err)
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	id, err := CreateBank(Bank{Name: mbank}, db)
	logFatalOnTest(t, err)
	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	CreateBank(Bank{Name: bzwbk}, db)
	CreateBank(Bank{Name: mbank}, db)
	DeleteAllBanks(db)
	banks, err := GetBanks(db)
	logFatalOnTest(t, err)
	assert.Empty(t, banks)
}

func TestGetBankById(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	id, err := CreateBank(Bank{Name: santander}, db)
	logFatalOnTest(t, err)
	bank, err := GetBankByID(int(id), db)
	logFatalOnTest(t, err)

	assert.Equal(t, santander, bank.Name)
}

func TestDeleteBankById(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)

	id, err := CreateBank(Bank{Name: mbank}, db)
	logFatalOnTest(t, err)

	DeleteBankByID(int(id), db)
	banks, err := GetBanks(db)
	logFatalOnTest(t, err)

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	db, err := NewBankAPI()
	logFatalOnTest(t, err)
	id, err := CreateBank(Bank{Name: mbank}, db)
	logFatalOnTest(t, err)
	bank, err := UpdateBank(Bank{ID: int(id), Name: bzwbk}, db)
	logFatalOnTest(t, err)

	assert.Equal(t, bzwbk, bank.Name)
}
