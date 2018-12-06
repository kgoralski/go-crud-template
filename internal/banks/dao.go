package banks

import (
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

type bankRepository interface {
	getBanks() ([]Bank, error)
	getBankByID(id int) (*Bank, error)
	createBank(bank Bank) (int, error)
	deleteAllBanks() error
	updateBank(bank Bank) (*Bank, error)
	deleteBankByID(id int) error
}

type bankDAO struct {
	db *sqlx.DB
}

func (bankAPI *bankDAO) getBanks() ([]Bank, error) {
	var banks = []Bank{}
	if err := bankAPI.db.Select(&banks, "SELECT * FROM banks"); err != nil {
		return nil, errors.Wrap(err, dbErrors.DbQueryFail)
	}
	return banks, nil
}

func (bankAPI *bankDAO) getBankByID(id int) (*Bank, error) {
	var bank = Bank{}
	if err := bankAPI.db.Get(&bank, "SELECT * FROM banks WHERE id=?", id); err != nil {
		return nil, errors.Wrap(err, dbErrors.DbQueryFail)
	}
	return &bank, nil
}

func (bankAPI *bankDAO) createBank(bank Bank) (int, error) {
	result, err := bankAPI.db.Exec("INSERT into banks (name) VALUES (?)", bank.Name)
	if err != nil {
		return 0, errors.Wrap(err, dbErrors.DbQueryFail)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, dbErrors.DbNotSupported)
	}
	return int(lastID), nil
}

func (bankAPI *bankDAO) deleteAllBanks() error {
	if _, err := bankAPI.db.Exec("TRUNCATE table banks"); err != nil {
		return errors.Wrap(err, dbErrors.DbQueryFail)
	}
	return nil
}

func (bankAPI *bankDAO) deleteBankByID(id int) error {
	res, err := bankAPI.db.Exec("DELETE from banks where id=?", id)
	if err != nil {
		return errors.Wrap(err, dbErrors.DbQueryFail)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, dbErrors.DbQueryFail)
	}
	if affect == 0 {
		return errors.New(dbErrors.EntityNotExist)
	}
	return nil
}

func (bankAPI *bankDAO) updateBank(bank Bank) (*Bank, error) {
	res, err := bankAPI.db.Exec("UPDATE banks SET name=? WHERE id=?", bank.Name, bank.ID)
	if err != nil {
		return nil, errors.Wrap(err, dbErrors.DbQueryFail)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, dbErrors.DbQueryFail)
	}
	if affect == 0 {
		return nil, errors.New(dbErrors.EntityNotExist)
	}
	return &bank, nil
}
