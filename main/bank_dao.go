package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Bank struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getBanks() []Bank {
	db := sqlx.MustConnect("mysql", "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8")
	var banks = []Bank{}
	err := db.Select(&banks, "SELECT * FROM banks")
	checkErr(err)
	return banks
}

func getBankById(id int) Bank {
	db := sqlx.MustConnect("mysql", "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8")
	var bank = Bank{}
	err := db.Get(&bank, "SELECT * FROM banks WHERE id=?", id)
	checkErr(err)
	return bank
}

func createBank(bank Bank) int {
	db := sqlx.MustConnect("mysql", "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8")
	result, err := db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	checkErr(err)
	lastId, lastIdErr := result.LastInsertId()
	checkErr(lastIdErr)
	return int(lastId)
}

func deleteAllBanks() {
	db := sqlx.MustConnect("mysql", "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8")
	db.Exec("TRUNCATE table banks")
}

func deleteBankById(id int) {
	db := sqlx.MustConnect("mysql", "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8")
	db.Exec("DELETE from banks where id=?", id)
}

func updateBank(bank Bank) Bank {
	db := sqlx.MustConnect("mysql", "admin:Admin.123@tcp(localhost:3306)/bank_db?charset=utf8")
	db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.Id)
	return bank
}
