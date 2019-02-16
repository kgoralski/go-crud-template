package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // DB driver registration
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//Store interface for Bank persistence layer
type Store interface {
	getAll() ([]Bank, error)
	get(id int) (*Bank, error)
	create(bank Bank) (int, error)
	deleteAll() error
	update(bank Bank) (*Bank, error)
	delete(id int) error
}

//BankStore for persistence
type BankStore struct {
	db *sqlx.DB
}

//NewStore creates new BankStore for Banks
func NewStore(db *sqlx.DB) *BankStore {
	return &BankStore{db: db}
}

func (s *BankStore) getAll() ([]Bank, error) {
	var banks []Bank
	if err := s.db.Select(&banks, "SELECT * FROM banks"); err != nil {
		return nil, ErrDbQuery{Err: errors.WithStack(err)}
	}
	return banks, nil
}

func (s *BankStore) get(id int) (*Bank, error) {
	var bank Bank
	if err := s.db.Get(&bank, "SELECT * FROM banks WHERE id=?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrEntityNotFound{Err: errors.WithStack(err)}
		}
		return nil, ErrDbQuery{Err: errors.WithStack(err)}
	}
	return &bank, nil
}

func (s *BankStore) create(bank Bank) (int, error) {
	result, err := s.db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, ErrDbQuery{Err: errors.WithStack(err)}
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, ErrDbNotSupported{Err: errors.WithStack(err)}
	}
	return int(lastID), nil
}

func (s *BankStore) deleteAll() error {
	if _, err := s.db.Exec("TRUNCATE table banks"); err != nil {
		return ErrDbQuery{Err: errors.WithStack(err)}
	}
	return nil
}

func (s *BankStore) delete(id int) error {
	res, err := s.db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return ErrDbQuery{Err: errors.WithStack(err)}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return ErrDbQuery{Err: errors.WithStack(err)}
	}
	if affect == 0 {
		return ErrEntityNotFound{Err: errors.WithStack(err)}
	}
	return nil
}

func (s *BankStore) update(bank Bank) (*Bank, error) {
	res, err := s.db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, ErrDbQuery{Err: errors.WithStack(err)}
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, ErrDbQuery{Err: errors.WithStack(err)}
	}
	if affect == 0 {
		return nil, ErrEntityNotFound{Err: errors.WithStack(err)}
	}
	return &bank, nil
}
