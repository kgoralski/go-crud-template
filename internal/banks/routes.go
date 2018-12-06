package banks

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/kgoralski/go-crud-template/cmd/middleware"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// Router structs represents Banks Handlers
type Router struct {
	r   *chi.Mux
	dao bankRepository
}

// NewRouter is creating New Bank Router Handlers
func NewRouter(r *chi.Mux, db *sqlx.DB) *Router {
	return &Router{r, &bankDAO{db: db}}
}

// Routes , all banks routes
func (h *Router) Routes() {
	h.r.Get("/rest/banks/", middleware.CommonHeaders(h.getBanks))
	h.r.Get("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.getBankByID))
	h.r.Post("/rest/banks/", middleware.CommonHeaders(h.createBank))
	h.r.Delete("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.deleteBankByID))
	h.r.Put("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.updateBank))
	h.r.Delete("/rest/banks/", middleware.CommonHeaders(h.deleteAllBanks))
}

func (h *Router) getBanks(w http.ResponseWriter, _ *http.Request) {
	banksDAO, err := h.dao.getBanks()
	if err != nil {
		middleware.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(banksDAO); err != nil {
		middleware.HandleErrors(w, err)
		return
	}
}

func (h *Router) getBankByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		middleware.HandleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}
	b, err := h.dao.getBankByID(id)
	if err != nil {
		middleware.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(b); err != nil {
		middleware.HandleErrors(w, err)
		return
	}
}

func (h *Router) createBank(w http.ResponseWriter, r *http.Request) {
	var bank Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		middleware.HandleErrors(w, err)
		return
	}
	id, err := h.dao.createBank(bank)
	if err != nil {
		middleware.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(id); err != nil {
		middleware.HandleErrors(w, err)
		return
	}
}

func (h *Router) deleteBankByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		middleware.HandleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}

	if err = h.dao.deleteBankByID(id); err != nil {
		middleware.HandleErrors(w, err)
		return
	}
}

func (h *Router) updateBank(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		middleware.HandleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
		return
	}
	var bank Bank
	if errDecode := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		middleware.HandleErrors(w, errDecode)
		return
	}
	updatedBank, err := h.dao.updateBank(Bank{ID: id, Name: bank.Name})
	if err != nil {
		middleware.HandleErrors(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(updatedBank); err != nil {
		middleware.HandleErrors(w, err)
		return
	}
}

func (h *Router) deleteAllBanks(w http.ResponseWriter, _ *http.Request) {
	if err := h.dao.deleteAllBanks(); err != nil {
		middleware.HandleErrors(w, err)
		return
	}
}
