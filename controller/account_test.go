package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pismo/api"
	"pismo/entity"
	serviceMock "pismo/mocks/service"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInsertAccount(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	args := struct {
		DocumentNumber string `json:"document_number"`
	}{
		DocumentNumber: "12345678900",
	}
	body, _ := json.Marshal(args)

	request := entity.Account{
		DocumentNumber: args.DocumentNumber,
	}

	response := entity.Account{
		ID:             1,
		DocumentNumber: args.DocumentNumber,
	}

	accountService.On("Insert", request).Return(response, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInsertAccountWithBindingError(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	args := struct {
		DocumentNumber uint `json:"document_number"`
	}{
		DocumentNumber: 12345678900,
	}
	body, _ := json.Marshal(args)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestInsertAccountWithServiceError(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	args := struct {
		DocumentNumber string `json:"document_number"`
	}{
		DocumentNumber: "12345678900",
	}
	body, _ := json.Marshal(args)

	request := entity.Account{
		DocumentNumber: args.DocumentNumber,
	}

	accountService.On("Insert", request).Return(entity.Account{}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFindAccount(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	response := entity.Account{
		ID:             1,
		DocumentNumber: "12345678900",
	}

	accountService.On("Get", uint64(1)).Return(response, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/account/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFindAccountWithBindingError(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/account/'1'", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestFindAccountWithServiceError(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	accountService.On("Get", uint64(1)).Return(entity.Account{}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/account/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
