package dao

import (
	"errors"

	_ "github.com/go-sql-driver/mysql" // db driver registration
	"github.com/jmoiron/sqlx"
	e "github.com/kgoralski/go-crud-template/handleErr"
)

const (
	sqlConnection = "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8"
	mysql         = "mysql"
)

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
		return nil, &e.DbError{Err: err, Message: e.DbConnectionFail}
	}
	return &BankAPI{db: db}, nil
}

// GetBanks returns all banks in database
func GetBanks(b *BankAPI) ([]Bank, error) {
	var banks = []Bank{}
	if err := b.db.Select(&banks, "SELECT * FROM banks"); err != nil {
		return nil, &e.DbError{Err: err, Message: e.DbQueryFail}
	}
	return banks, nil
}

// GetBankByID returns single bank by ID
func GetBankByID(id int, b *BankAPI) (*Bank, error) {
	var bank = Bank{}
	if err := b.db.Get(&bank, "SELECT * FROM banks WHERE id=?", id); err != nil {
		return nil, &e.DbError{Err: err, Message: e.DbQueryFail}
	}
	return &bank, nil
}

// CreateBank is creating single bank inside database
func CreateBank(bank Bank, b *BankAPI) (int, error) {
	result, err := b.db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, &e.DbError{Err: err, Message: e.DbQueryFail}
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, &e.DbError{Err: err, Message: e.DbNotSupported}
	}
	return int(lastID), nil
}

// DeleteAllBanks deletes all banks inside database
func DeleteAllBanks(b *BankAPI) error {
	if _, err := b.db.Exec("TRUNCATE table banks"); err != nil {
		return &e.DbError{Err: err, Message: e.DbQueryFail}
	}
	return nil
}

// DeleteBankByID deletes single bank by ID
func DeleteBankByID(id int, b *BankAPI) error {
	res, err := b.db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return &e.DbError{Err: err, Message: e.DbQueryFail}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return &e.DbError{Err: err, Message: e.DbQueryFail}
	}
	if affect == 0 {
		return &e.DbError{Err: errors.New(e.EntityNotExist), Message: e.EntityNotExist}
	}
	return nil
}

// UpdateBank updates single bank in database
func UpdateBank(bank Bank, b *BankAPI) (*Bank, error) {
	res, err := b.db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, &e.DbError{Err: err, Message: e.DbQueryFail}
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return nil, &e.DbError{Err: errors.New(e.EntityNotExist), Message: e.EntityNotExist}
	}
	return &bank, nil
}
