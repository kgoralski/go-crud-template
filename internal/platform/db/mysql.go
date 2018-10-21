package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const mysql = "mysql"

func New(sqlConnection string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(mysql, sqlConnection)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}
	return db, nil
}
