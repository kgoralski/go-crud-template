package rest

import (
	"encoding/json"
	"github.com/kgoralski/go-crud-template/dao"
	e "github.com/kgoralski/go-crud-template/handleErr"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)


func connectDB(w http.ResponseWriter) *dao.BankAPI {
	db, err := dao.NewBankAPI()
	if err != nil {
		e.HandleErrors(w, err)
		return nil
	}
	return db
}

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJSON)
		fn(w, r)
	}
}

func getBanksHandler(w http.ResponseWriter, r *http.Request) {
	banks, err := dao.GetBanks(connectDB(w))
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
		e.HandleErrors(w, &e.HTTPError{err, http.StatusText(http.StatusBadRequest), http.StatusBadRequest})
		return
	}
	b, err := dao.GetBankByID(id, connectDB(w))
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
	id, err := dao.CreateBank(bank, connectDB(w))
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
		e.HandleErrors(w, &e.HTTPError{err, http.StatusText(http.StatusBadRequest), http.StatusBadRequest})
		return
	}

	if err = dao.DeleteBankByID(id, connectDB(w)); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		e.HandleErrors(w, &e.HTTPError{err, http.StatusText(http.StatusBadRequest), http.StatusBadRequest})
		return
	}
	var bank dao.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		e.HandleErrors(w, err)
		return
	}
	updatedBank, err := dao.UpdateBank(dao.Bank{id, bank.Name}, connectDB(w))
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
	if err := dao.DeleteAllBanks(connectDB(w)); err != nil {
		e.HandleErrors(w, err)
		return
	}
}
