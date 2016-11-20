package dao

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/spf13/viper"
)

const (
	bzwbk     = "BZWBK"
	mbank     = "MBANK"
	santander = "SANTANDER"
)

func logFatalOnTest(t *testing.T, err error) {
	if err != nil {
		log.Fatal(fmt.Errorf("FATAL: %+v\n", err))
		t.Fatal(err)
	}
}

func setupConf() {
	viper.SetConfigName("conf_test")
	viper.AddConfigPath("../_conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("FATAL: %+v\n", err))
	}
}

func init() {
	setupConf()
	db, err := NewBankAPI(viper.GetString("database.URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("FATAL: %+v\n", err))
	}
	DBAccess = db
	DeleteAllBanks()
}

func TestGetBanks(t *testing.T) {
	DeleteAllBanks()
	CreateBank(Bank{Name: bzwbk})
	CreateBank(Bank{Name: mbank})

	banks, err := GetBanks()
	logFatalOnTest(t, err)
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	id, err := CreateBank(Bank{Name: mbank})
	logFatalOnTest(t, err)
	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	CreateBank(Bank{Name: bzwbk})
	CreateBank(Bank{Name: mbank})
	DeleteAllBanks()
	banks, err := GetBanks()
	logFatalOnTest(t, err)
	assert.Empty(t, banks)
}

func TestGetBankById(t *testing.T) {
	id, err := CreateBank(Bank{Name: santander})
	logFatalOnTest(t, err)
	bank, err := GetBankByID(int(id))
	logFatalOnTest(t, err)

	assert.Equal(t, santander, bank.Name)
}

func TestDeleteBankById(t *testing.T) {
	id, err := CreateBank(Bank{Name: mbank})
	logFatalOnTest(t, err)

	DeleteBankByID(int(id))
	banks, err := GetBanks()
	logFatalOnTest(t, err)

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	id, err := CreateBank(Bank{Name: mbank})
	logFatalOnTest(t, err)
	bank, err := UpdateBank(Bank{ID: int(id), Name: bzwbk})
	logFatalOnTest(t, err)

	assert.Equal(t, bzwbk, bank.Name)
}
