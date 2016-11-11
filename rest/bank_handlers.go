package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kgoralski/go-crud-template/dao"
	e "github.com/kgoralski/go-crud-template/handleErr"

	"github.com/gorilla/mux"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJSON)
		fn(w, r)
	}
}

func getBanksHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dao.NewBankAPI()
	if err != nil {
		e.HandleErrors(w, &e.DbError{Err: err, Message: e.DbConnectionFail})
		return
	}
	banks, err := dao.GetBanks(db)
	if err != nil {
		e.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(banks); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func getBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		e.HandleErrors(w, &e.HTTPError{Err: err, Message: http.StatusText(http.StatusBadRequest), Code: 400})
		return
	}
	db, err := dao.NewBankAPI()
	if err != nil {
		e.HandleErrors(w, &e.DbError{Err: err, Message: e.DbConnectionFail})
		return
	}
	b, err := dao.GetBankByID(id, db)
	if err != nil {
		e.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(b); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func createBankHanlder(w http.ResponseWriter, r *http.Request) {
	var bank dao.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		e.HandleErrors(w, err)
		return
	}
	db, err := dao.NewBankAPI()
	if err != nil {
		e.HandleErrors(w, &e.DbError{Err: err, Message: e.DbConnectionFail})
		return
	}
	id, err := dao.CreateBank(bank, db)
	if err != nil {
		e.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(id); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func deleteBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		e.HandleErrors(w, &e.HTTPError{Err: err, Message: http.StatusText(http.StatusBadRequest), Code: 400})
		return
	}

	db, err := dao.NewBankAPI()
	if err != nil {
		e.HandleErrors(w, &e.DbError{Err: err, Message: e.DbConnectionFail})
		return
	}
	if err = dao.DeleteBankByID(id, db); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		e.HandleErrors(w, &e.HTTPError{Err: err, Message: http.StatusText(http.StatusBadRequest), Code: 400})
		return
	}
	var bank dao.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		e.HandleErrors(w, err)
		return
	}
	db, err := dao.NewBankAPI()
	if err != nil {
		e.HandleErrors(w, &e.DbError{Err: err, Message: e.DbConnectionFail})
		return
	}
	updatedBank, err := dao.UpdateBank(dao.Bank{ID: id, Name: bank.Name}, db)
	if err != nil {
		e.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(updatedBank); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func deleteAllBanksHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dao.NewBankAPI()
	if err != nil {
		e.HandleErrors(w, &e.DbError{Err: err, Message: e.DbConnectionFail})
		return
	}
	if err := dao.DeleteAllBanks(db); err != nil {
		e.HandleErrors(w, err)
		return
	}
}
