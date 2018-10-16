package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kgoralski/go-crud-template/dao"
	e "github.com/kgoralski/go-crud-template/handleErr"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
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

func (s *server) getBanksHandler(w http.ResponseWriter, _ *http.Request) {
	banks, err := s.db.GetBanks()
	if err != nil {
		e.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(banks); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func (s *server) getBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		e.HandleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}
	b, err := s.db.GetBankByID(id)
	if err != nil {
		e.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(b); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func (s *server) createBankHanlder(w http.ResponseWriter, r *http.Request) {
	var bank dao.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		e.HandleErrors(w, err)
		return
	}
	id, err := s.db.CreateBank(bank)
	if err != nil {
		e.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(id); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func (s *server) deleteBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		e.HandleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}

	if err = s.db.DeleteBankByID(id); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func (s *server) updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		e.HandleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}
	var bank dao.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		e.HandleErrors(w, err)
		return
	}
	updatedBank, err := s.db.UpdateBank(dao.Bank{ID: id, Name: bank.Name})
	if err != nil {
		e.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(updatedBank); err != nil {
		e.HandleErrors(w, err)
		return
	}
}

func (s *server) deleteAllBanksHandler(w http.ResponseWriter, _ *http.Request) {
	if err := s.db.DeleteAllBanks(); err != nil {
		e.HandleErrors(w, err)
		return
	}
}
