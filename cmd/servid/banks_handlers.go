package servid

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kgoralski/go-crud-template/internal/banks"

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

func (s *Server) getBanksHandler(w http.ResponseWriter, _ *http.Request) {
	banksDAO, err := s.db.GetBanks()
	if err != nil {
		handleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(banksDAO); err != nil {
		handleErrors(w, err)
		return
	}
}

func (s *Server) getBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		handleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}
	b, err := s.db.GetBankByID(id)
	if err != nil {
		handleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(b); err != nil {
		handleErrors(w, err)
		return
	}
}

func (s *Server) createBankHanlder(w http.ResponseWriter, r *http.Request) {
	var bank banks.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		handleErrors(w, err)
		return
	}
	id, err := s.db.CreateBank(bank)
	if err != nil {
		handleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(id); err != nil {
		handleErrors(w, err)
		return
	}
}

func (s *Server) deleteBankByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		handleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}

	if err = s.db.DeleteBankByID(id); err != nil {
		handleErrors(w, err)
		return
	}
}

func (s *Server) updateBankHanlder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		handleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}
	var bank banks.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		handleErrors(w, err)
		return
	}
	updatedBank, err := s.db.UpdateBank(banks.Bank{ID: id, Name: bank.Name})
	if err != nil {
		handleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(updatedBank); err != nil {
		handleErrors(w, err)
		return
	}
}

func (s *Server) deleteAllBanksHandler(w http.ResponseWriter, _ *http.Request) {
	if err := s.db.DeleteAllBanks(); err != nil {
		handleErrors(w, err)
		return
	}
}
