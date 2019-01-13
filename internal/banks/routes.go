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
	r     *chi.Mux
	store *banksStore
}

// NewRouter is creating New Bank Router Handlers
func NewRouter(r *chi.Mux, db *sqlx.DB) *Router {
	return &Router{r, &banksStore{db: db}}
}

// Routes , all banks routes
func (h *Router) Routes() {
	h.r.Get("/rest/banks/", middleware.CommonHeaders(h.getBanks()))
	h.r.Get("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.getBankByID()))
	h.r.Post("/rest/banks/", middleware.CommonHeaders(h.createBank()))
	h.r.Delete("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.deleteBankByID()))
	h.r.Put("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.updateBank()))
	h.r.Delete("/rest/banks/", middleware.CommonHeaders(h.deleteAllBanks()))
}

func (h *Router) getBanks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repository, err := h.store.getAll()
		if err != nil {
			middleware.HandleErrors(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(repository); err != nil {
			middleware.HandleErrors(w, err)
			return
		}
	}

}

func (h *Router) getBankByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			middleware.HandleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
			return
		}
		b, err := h.store.get(id)
		if err != nil {
			middleware.HandleErrors(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(b); err != nil {
			middleware.HandleErrors(w, err)
			return
		}
	}
}

func (h *Router) createBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bank Bank
		if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
			middleware.HandleErrors(w, err)
			return
		}
		id, err := h.store.create(bank)
		if err != nil {
			middleware.HandleErrors(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(id); err != nil {
			middleware.HandleErrors(w, err)
			return
		}
	}
}

func (h *Router) deleteBankByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			middleware.HandleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
			return
		}

		if err = h.store.delete(id); err != nil {
			middleware.HandleErrors(w, err)
			return
		}
	}
}

func (h *Router) updateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		updatedBank, err := h.store.update(Bank{ID: id, Name: bank.Name})
		if err != nil {
			middleware.HandleErrors(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(updatedBank); err != nil {
			middleware.HandleErrors(w, err)
			return
		}
	}
}

func (h *Router) deleteAllBanks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.store.deleteAll(); err != nil {
			middleware.HandleErrors(w, err)
			return
		}
	}
}
