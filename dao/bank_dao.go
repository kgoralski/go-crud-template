package dao

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	e "github.com/kgoralski/go-crud-template/handleErr"
)

const sqlConnection = "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8"

type DBAccess interface {
	GetBanks()
}

type Bank struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type BankAPI struct {
	db *sqlx.DB
}

func NewBankAPI() (*BankAPI, error) {
	db, err := sqlx.Connect("mysql", sqlConnection)
	if err != nil {
		return nil, &e.DbError{err, e.DbConnectionFail}
	}
	return &BankAPI{db: db}, nil
}

func GetBanks(b *BankAPI) ([]Bank, error) {
	var banks = []Bank{}
	if err := b.db.Select(&banks, "SELECT * FROM banks"); err != nil {
		return nil, &e.DbError{err, e.DbQueryFail}
	}
	return banks, nil
}

func GetBankByID(id int, b *BankAPI) (*Bank, error) {
	var bank = Bank{}
	if err := b.db.Get(&bank, "SELECT * FROM banks WHERE id=?", id); err != nil {
		return nil, &e.DbError{err, e.DbQueryFail}
	}
	return &bank, nil
}

func CreateBank(bank Bank, b *BankAPI) (int, error) {
	result, err := b.db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, &e.DbError{err, e.DbQueryFail}
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, &e.DbError{err, e.DbNotSupported}
	}
	return int(lastID), nil
}

func DeleteAllBanks(b *BankAPI) error {
	if _, err := b.db.Exec("TRUNCATE table banks"); err != nil {
		return &e.DbError{err, e.DbQueryFail}
	}
	return nil
}

func DeleteBankByID(id int, b *BankAPI) error {
	res, err := b.db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return &e.DbError{err, e.DbQueryFail}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return &e.DbError{err, e.DbQueryFail}
	}
	if affect == 0 {
		return &e.DbError{errors.New(e.EntityNotExist), e.EntityNotExist}
	}
	return nil
}

func UpdateBank(bank Bank, b *BankAPI) (*Bank, error) {
	res, err := b.db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, &e.DbError{err, e.DbQueryFail}
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return nil, &e.DbError{errors.New(e.EntityNotExist), e.EntityNotExist}
	}
	return &bank, nil
}
