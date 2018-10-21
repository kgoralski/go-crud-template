package banks

import (
	"fmt"
	"github.com/kgoralski/go-crud-template/internal/platform/db"
	log "github.com/sirupsen/logrus"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	bzwbk     = "BZWBK"
	mbank     = "MBANK"
	santander = "SANTANDER"
)

func logFatalOnTest(t *testing.T, err error) {
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
		t.Fatal(err)
	}
}

func setupConf() {
	viper.SetConfigName("conf_test")
	viper.AddConfigPath("./../../configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
}

var dbAccess *BankAPI

func init() {
	setupConf()
	mysql, err := db.New(viper.GetString("database.URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
	bankAPI, err := NewBankAPI(mysql)
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
	dbAccess = bankAPI
}

func TestGetBanks(t *testing.T) {
	dbAccess.DeleteAllBanks()
	dbAccess.CreateBank(Bank{Name: bzwbk})
	dbAccess.CreateBank(Bank{Name: mbank})

	banks, err := dbAccess.GetBanks()
	logFatalOnTest(t, err)
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	id, err := dbAccess.CreateBank(Bank{Name: mbank})
	logFatalOnTest(t, err)
	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	dbAccess.CreateBank(Bank{Name: bzwbk})
	dbAccess.CreateBank(Bank{Name: mbank})
	dbAccess.DeleteAllBanks()
	banks, err := dbAccess.GetBanks()
	logFatalOnTest(t, err)
	assert.Empty(t, banks)
}

func TestGetBankById(t *testing.T) {
	id, err := dbAccess.CreateBank(Bank{Name: santander})
	logFatalOnTest(t, err)
	bank, err := dbAccess.GetBankByID(int(id))
	logFatalOnTest(t, err)

	assert.Equal(t, santander, bank.Name)
}

func TestDeleteBankById(t *testing.T) {
	id, err := dbAccess.CreateBank(Bank{Name: mbank})
	logFatalOnTest(t, err)

	dbAccess.DeleteBankByID(int(id))
	banks, err := dbAccess.GetBanks()
	logFatalOnTest(t, err)

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	id, err := dbAccess.CreateBank(Bank{Name: mbank})
	logFatalOnTest(t, err)
	bank, err := dbAccess.UpdateBank(Bank{ID: int(id), Name: bzwbk})
	logFatalOnTest(t, err)

	assert.Equal(t, bzwbk, bank.Name)
}
