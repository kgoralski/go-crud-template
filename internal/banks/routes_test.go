package banks

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/kgoralski/go-crud-template/internal/banks/domain"
	"github.com/kgoralski/go-crud-template/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterGetAllBanksWithGoMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// given
	mockService := mock.NewMockService(mockCtrl)
	mockResponse := []domain.Bank{{
		ID:   1,
		Name: "Santander",
	}, {
		ID:   2,
		Name: "MBANK",
	}}
	mockService.EXPECT().GetBanks().Return(mockResponse, nil).Times(1)

	h := &Router{chi.NewRouter(), mockService}

	// when
	r, err := http.NewRequest(http.MethodGet, "/rest/banks/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	h.r.Get("/rest/banks/", h.getBanks())
	h.r.ServeHTTP(w, r)

	// then
	assert.Equal(t, http.StatusOK, w.Code)
	var b []domain.Bank
	err = json.NewDecoder(w.Body).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, mockResponse, b)
}

func TestRouterGetBankWithTestify(t *testing.T) {
	// given
	mockService := new(mock.Service)
	mockResponse := &domain.Bank{
		ID:   1,
		Name: "Santander",
	}
	mockService.On("GetBank", 1).Return(mockResponse, nil).Once()

	h := &Router{chi.NewRouter(), mockService}
	h.Routes()

	// when
	r, err := http.NewRequest(http.MethodGet, "/rest/banks/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	h.r.ServeHTTP(w, r)

	// then
	assert.Equal(t, http.StatusOK, w.Code)
	var b domain.Bank
	err = json.NewDecoder(w.Body).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *mockResponse, b)
}
