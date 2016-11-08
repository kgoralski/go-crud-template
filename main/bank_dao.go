package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	dbQueryFail      = "DB_QUERY_FAIL"
	dbNotSupported   = "DB_NOT_SUPPORTED"
	entityNotExist   = "ENTITY_NOT_EXIST"
	dbConnectionFail = "DB_CONNECTION_FAIL"
	sqlConnection    = "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8"
)

/*
Bank struct Public only for example reasons
*/
type Bank struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

/*
BankAPI struct Public only for example reasons
*/
type BankAPI struct {
	db *sqlx.DB
}

/*
NewBankAPI struct Public only for example reasons
*/
func NewBankAPI() (*BankAPI, error) {
	db, err := sqlx.Connect("mysql", sqlConnection)
	if err != nil {
		return nil, err
	}
	return &BankAPI{db: db}, nil
}

func getBanks(b *BankAPI) ([]Bank, error) {
	var banks = []Bank{}
	if err := b.db.Select(&banks, "SELECT * FROM banks"); err != nil {
		return nil, &dbError{err, dbQueryFail}
	}
	return banks, nil
}

func getBankByID(id int, b *BankAPI) (*Bank, error) {
	var bank = Bank{}
	if err := b.db.Get(&bank, "SELECT * FROM banks WHERE id=?", id); err != nil {
		return nil, &dbError{err, dbQueryFail}
	}
	return &bank, nil
}

func createBank(bank Bank, b *BankAPI) (int, error) {
	result, err := b.db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, &dbError{err, dbQueryFail}
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, &dbError{err, dbNotSupported}
	}
	return int(lastID), nil
}

func deleteAllBanks(b *BankAPI) error {
	if _, err := b.db.Exec("TRUNCATE table banks"); err != nil {
		return &dbError{err, dbQueryFail}
	}
	return nil
}

func deleteBankByID(id int, b *BankAPI) error {
	res, err := b.db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return &dbError{err, dbQueryFail}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return &dbError{err, dbQueryFail}
	}
	if affect == 0 {
		return &dbError{errors.New(entityNotExist), entityNotExist}
	}
	return nil
}

func updateBank(bank Bank, b *BankAPI) (*Bank, error) {
	res, err := b.db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, &dbError{err, dbQueryFail}
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return nil, &dbError{errors.New(entityNotExist), entityNotExist}
	}
	return &bank, nil
}
