package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Bank struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

const DB_QUERY_FAIL = "DB_QUERY_FAIL"
const DB_NOT_SUPPORTED = "DB_NOT_SUPPORTED"
const ENTITY_NOT_EXIST = "ENTITY_NOT_EXIST"

const sqlConnection string = "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8"

func getBanks() ([]Bank, *HttpError) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	var banks = []Bank{}
	err := db.Select(&banks, "SELECT * FROM banks")
	if err != nil {
		return nil, &HttpError{err, DB_QUERY_FAIL, 409}
	}
	return banks, nil
}

func getBankById(id int) (*Bank, *HttpError) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	var bank = Bank{}
	err := db.Get(&bank, "SELECT * FROM banks WHERE id=?", id)
	if err != nil {
		return nil, &HttpError{err, DB_QUERY_FAIL, 409}
	}
	return &bank, nil
}

func createBank(bank Bank) (int, *HttpError) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	result, err := db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return -1, &HttpError{err, DB_QUERY_FAIL, 409}
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return -1, &HttpError{err, DB_NOT_SUPPORTED, 409}
	}
	return int(lastId), nil
}

func deleteAllBanks() *HttpError {
	db := sqlx.MustConnect("mysql", sqlConnection)
	_, err := db.Exec("TRUNCATE table banks")
	if err != nil {
		return &HttpError{err, DB_QUERY_FAIL, 409}
	}
	return nil
}

func deleteBankById(id int) *HttpError {
	db := sqlx.MustConnect("mysql", sqlConnection)
	res, err := db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return &HttpError{err, DB_QUERY_FAIL, 409}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return &HttpError{err, DB_QUERY_FAIL, 400}
	}
	if affect == 0 {
		return &HttpError{errors.New(ENTITY_NOT_EXIST), ENTITY_NOT_EXIST, 404}
	}
	return nil
}

func updateBank(bank Bank) (*Bank, *HttpError) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	res, err := db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.Id)
	if err != nil {
		return nil, &HttpError{err, DB_QUERY_FAIL, 409}
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return nil, &HttpError{errors.New(ENTITY_NOT_EXIST), ENTITY_NOT_EXIST, 404}
	}
	return &bank, nil
}
