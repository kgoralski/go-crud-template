package banks

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // DB driver registration
	"github.com/jmoiron/sqlx"
	dbErrors "github.com/kgoralski/go-crud-template/cmd/middleware"
	"github.com/pkg/errors"
)

// Bank DB entity shared in app for simplicity
type Bank struct {
	ID   int    `json:"id" DB:"id"`
	Name string `json:"name" DB:"name"`
}

type repository interface {
	getAll() ([]Bank, error)
	get(id int) (*Bank, error)
	create(bank Bank) (int, error)
	deleteAll() error
	update(bank Bank) (*Bank, error)
	delete(id int) error
}

type banksRepository struct {
	db *sqlx.DB
}

func (r *banksRepository) getAll() ([]Bank, error) {
	var banks = []Bank{}
	if err := r.db.Select(&banks, "SELECT * FROM banks"); err != nil {
		return nil, dbErrors.ErrDbQuery{Err: errors.Wrap(err, dbErrors.DbQueryFail)}
	}
	return banks, nil
}

func (r *banksRepository) get(id int) (*Bank, error) {
	var bank = Bank{}
	if err := r.db.Get(&bank, "SELECT * FROM banks WHERE id=?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrEntityNotFound{Err: errors.Wrap(err, dbErrors.EntityNotExist)}
		}
		return nil, dbErrors.ErrDbQuery{Err: errors.Wrap(err, dbErrors.DbQueryFail)}
	}
	return &bank, nil
}

func (r *banksRepository) create(bank Bank) (int, error) {
	result, err := r.db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, dbErrors.ErrDbQuery{Err: errors.Wrap(err, dbErrors.DbQueryFail)}
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, dbErrors.ErrDbNotSupported{Err: errors.Wrap(err, dbErrors.DbNotSupported)}
	}
	return int(lastID), nil
}

func (r *banksRepository) deleteAll() error {
	if _, err := r.db.Exec("TRUNCATE table banks"); err != nil {
		return dbErrors.ErrDbQuery{Err: errors.Wrap(err, dbErrors.DbQueryFail)}
	}
	return nil
}

func (r *banksRepository) delete(id int) error {
	res, err := r.db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return dbErrors.ErrDbQuery{Err: errors.Wrap(err, dbErrors.DbQueryFail)}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return dbErrors.ErrDbQuery{Err: errors.Wrap(err, dbErrors.DbQueryFail)}
	}
	if affect == 0 {
		return dbErrors.ErrEntityNotFound{Err: errors.Wrap(err, dbErrors.EntityNotExist)}
	}
	return nil
}

func (r *banksRepository) update(bank Bank) (*Bank, error) {
	res, err := r.db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, dbErrors.ErrDbQuery{Err: errors.Wrap(err, dbErrors.DbQueryFail)}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, dbErrors.ErrDbQuery{Err: errors.Wrap(err, dbErrors.DbQueryFail)}
	}
	if affect == 0 {
		return nil, dbErrors.ErrEntityNotFound{Err: errors.Wrap(err, dbErrors.EntityNotExist)}
	}
	return &bank, nil
}
