package client

import (
	"testing"

	"github.com/kgoralski/go-crud-template/handleErr"
	"github.com/kgoralski/go-crud-template/rest"
	"github.com/stretchr/testify/assert"
)

func init() {
	go rest.StartServer()
	deleteBanks()
}

const (
	bzwbk     = "BZWBK"
	mbank     = "MBANK"
	santander = "SANTANDER"
	alior     = "ALIOR"
	ing       = "ING"
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

func TestGetBanksClient(t *testing.T) {
	deleteBanks()
	postBank(bank{Name: bzwbk})
	postBank(bank{Name: mbank})
	banks, err := getAllBanks()
	logFatalOnTest(t, err)
	assert.Len(t, banks, 2, "Expected size is 2")

}

func TestGetOneBankClient(t *testing.T) {
	id, err := postBank(bank{Name: santander})
	logFatalOnTest(t, err)
	bank, errQuery := getOneBank(id)
	logFatalOnTest(t, errQuery)

	assert.Equal(t, santander, bank.Name, "Expected that both names are equal")
}

func TestCreateBankClient(t *testing.T) {
	bank := bank{Name: alior}
	id, err := postBank(bank)
	logFatalOnTest(t, err)

	createdBank, errQuery := getOneBank(id)
	logFatalOnTest(t, errQuery)

	assert.Equal(t, alior, createdBank.Name, "Expected that both names are equal")
}

func TestDeleteSingleBankClient(t *testing.T) {
	id, err := postBank(bank{Name: ing})
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
	postBank(bank{Name: bzwbk})
	postBank(bank{Name: mbank})
	postBank(bank{Name: alior})
	deleteBanks()
	banks, err := getAllBanks()
	logFatalOnTest(t, err)

	assert.Empty(t, banks)
}
