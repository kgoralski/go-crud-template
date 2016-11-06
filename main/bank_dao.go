package main

import (
	"errors"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Bank struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

const (
	DB_QUERY_FAIL    = "DB_QUERY_FAIL"
	DB_NOT_SUPPORTED = "DB_NOT_SUPPORTED"
	ENTITY_NOT_EXIST = "ENTITY_NOT_EXIST"
	sqlConnection    = "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8"
)

func getBanks() ([]Bank, error) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	var banks = []Bank{}
	err := db.Select(&banks, "SELECT * FROM banks")
	if err != nil {
		return nil, &httpError{err, DB_QUERY_FAIL, http.StatusConflict}
	}
	return banks, nil
}

func getBankById(id int) (*Bank, error) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	var bank = Bank{}
	err := db.Get(&bank, "SELECT * FROM banks WHERE id=?", id)
	if err != nil {
		return nil, &httpError{err, DB_QUERY_FAIL, http.StatusConflict}
	}
	return &bank, nil
}

func createBank(bank Bank) (int, error) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	result, err := db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return -1, &httpError{err, DB_QUERY_FAIL, http.StatusConflict}
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return -1, &httpError{err, DB_NOT_SUPPORTED, http.StatusConflict}
	}
	return int(lastId), nil
}

func deleteAllBanks() error {
	db := sqlx.MustConnect("mysql", sqlConnection)
	_, err := db.Exec("TRUNCATE table banks")
	if err != nil {
		return &httpError{err, DB_QUERY_FAIL, http.StatusConflict}
	}
	return nil
}

func deleteBankById(id int) error {
	db := sqlx.MustConnect("mysql", sqlConnection)
	res, err := db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return &httpError{err, DB_QUERY_FAIL, http.StatusConflict}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return &httpError{err, DB_QUERY_FAIL, http.StatusBadRequest}
	}
	if affect == 0 {
		return &httpError{errors.New(ENTITY_NOT_EXIST), ENTITY_NOT_EXIST, http.StatusNotFound}
	}
	return nil
}

func updateBank(bank Bank) (*Bank, error) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	res, err := db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.Id)
	if err != nil {
		return nil, &httpError{err, DB_QUERY_FAIL, http.StatusConflict}
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return nil, &httpError{errors.New(ENTITY_NOT_EXIST), ENTITY_NOT_EXIST, http.StatusNotFound}
	}
	return &bank, nil
}
