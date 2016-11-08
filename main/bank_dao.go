package main

import (
	"errors"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
Bank struct declared here only for example reasons
*/
type Bank struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

const (
	dbQueryFail    = "DB_QUERY_FAIL"
	dbNotSupported = "DB_NOT_SUPPORTED"
	entityNotExist = "ENTITY_NOT_EXIST"
	sqlConnection  = "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8"
)

func getBanks() ([]Bank, error) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	var banks = []Bank{}
	err := db.Select(&banks, "SELECT * FROM banks")
	if err != nil {
		return nil, &httpError{err, dbQueryFail, http.StatusConflict}
	}
	return banks, nil
}

func getBankByID(id int) (*Bank, error) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	var bank = Bank{}
	err := db.Get(&bank, "SELECT * FROM banks WHERE id=?", id)
	if err != nil {
		return nil, &httpError{err, dbQueryFail, http.StatusConflict}
	}
	return &bank, nil
}

func createBank(bank Bank) (int, error) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	result, err := db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, &httpError{err, dbQueryFail, http.StatusConflict}
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, &httpError{err, dbNotSupported, http.StatusConflict}
	}
	return int(lastID), nil
}

func deleteAllBanks() error {
	db := sqlx.MustConnect("mysql", sqlConnection)
	_, err := db.Exec("TRUNCATE table banks")
	if err != nil {
		return &httpError{err, dbQueryFail, http.StatusConflict}
	}
	return nil
}

func deleteBankByID(id int) error {
	db := sqlx.MustConnect("mysql", sqlConnection)
	res, err := db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return &httpError{err, dbQueryFail, http.StatusConflict}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return &httpError{err, dbQueryFail, http.StatusBadRequest}
	}
	if affect == 0 {
		return &httpError{errors.New(entityNotExist), entityNotExist, http.StatusNotFound}
	}
	return nil
}

func updateBank(bank Bank) (*Bank, error) {
	db := sqlx.MustConnect("mysql", sqlConnection)
	res, err := db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, &httpError{err, dbQueryFail, http.StatusConflict}
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		return nil, &httpError{errors.New(entityNotExist), entityNotExist, http.StatusNotFound}
	}
	return &bank, nil
}
