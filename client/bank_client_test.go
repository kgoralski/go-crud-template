package client

import (
	"testing"

	"fmt"
	"log"

	"time"

	"github.com/kgoralski/go-crud-template/rest"
	"github.com/stretchr/testify/assert"
	"sync"
)

const (
	bzwbk     = "BZWBK"
	mbank     = "MBANK"
	santander = "SANTANDER"
	alior     = "ALIOR"
	ing       = "ING"
)

func init() {
	go rest.StartServer()
	time.Sleep(time.Millisecond * 1)
}

func logFatalOnTest(t *testing.T, err error) {
	if err != nil {
		log.Fatal(fmt.Errorf("FATAL: %+v\n", err))
		t.Fatal(err)
	}
}

func TestDeleteAllBankClient(t *testing.T) {
	postBank(bank{Name: bzwbk})
	postBank(bank{Name: mbank})
	postBank(bank{Name: alior})
	deleteBanks()
	banks, err := getAllBanks()
	logFatalOnTest(t, err)

	assert.Empty(t, banks)
}

func TestGetBanksClient(t *testing.T) {
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
