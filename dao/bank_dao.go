package dao

import (
	_ "github.com/go-sql-driver/mysql" // db driver registration
	"github.com/jmoiron/sqlx"
	e "github.com/kgoralski/go-crud-template/handleErr"
	"github.com/pkg/errors"
)

const (
	sqlConnection = "admin:Admin.123@tcp(go-app-db:3306)/bank_db?charset=utf8"
	mysql         = "mysql"
)

// DBAccess initialised on startup
var DBAccess *BankAPI

// Bank db entity shared in app for simplicity
type Bank struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// BankAPI has db *sqlx.DB inside
type BankAPI struct {
	db *sqlx.DB
}

// NewBankAPI establishing db connection
func NewBankAPI() (*BankAPI, error) {
	db, err := sqlx.Connect(mysql, sqlConnection)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}
	return &BankAPI{db: db}, nil
}

// GetBanks returns all banks from database
func GetBanks() ([]Bank, error) {
	var banks = []Bank{}
	if err := DBAccess.db.Select(&banks, "SELECT * FROM banks"); err != nil {
		return nil, errors.Wrap(err, e.DbQueryFail)
	}
	return banks, nil
}

// GetBankByID returns single bank by ID
func GetBankByID(id int) (*Bank, error) {
	var bank = Bank{}
	if err := DBAccess.db.Get(&bank, "SELECT * FROM banks WHERE id=?", id); err != nil {
		return nil, errors.Wrap(err, e.DbQueryFail)
	}
	return &bank, nil
}

// CreateBank is creating single bank inside database
func CreateBank(bank Bank) (int, error) {
	result, err := DBAccess.db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, errors.Wrap(err, e.DbQueryFail)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, e.DbNotSupported)
	}
	return int(lastID), nil
}

// DeleteAllBanks deletes all banks inside database
func DeleteAllBanks() error {
	if _, err := DBAccess.db.Exec("TRUNCATE table banks"); err != nil {
		return errors.Wrap(err, e.DbQueryFail)
	}
	return nil
}

// DeleteBankByID deletes single bank by ID
func DeleteBankByID(id int) error {
	res, err := DBAccess.db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return errors.Wrap(err, e.DbQueryFail)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, e.DbQueryFail)
	}
	if affect == 0 {
		return errors.New(e.EntityNotExist)
	}
	return nil
}

// UpdateBank updates single bank in database
func UpdateBank(bank Bank) (*Bank, error) {
	res, err := DBAccess.db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, errors.Wrap(err, e.DbQueryFail)
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return nil, errors.New(e.EntityNotExist)
	}
	return &bank, nil
}
