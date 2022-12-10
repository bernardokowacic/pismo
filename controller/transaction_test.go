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

func TestInsertTransaction(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	args := struct {
		AccountID     uint64  `json:"account_id"`
		OperationType uint64  `json:"operation_type_id"`
		Amount        float64 `json:"amount"`
	}{
		AccountID:     1,
		OperationType: 4,
		Amount:        123.56,
	}
	body, _ := json.Marshal(args)

	request := entity.Transaction{
		AccountID:       args.AccountID,
		OperationTypeID: args.OperationType,
		Amount:          args.Amount,
	}

	response := entity.Transaction{
		ID:              1,
		AccountID:       args.AccountID,
		OperationTypeID: args.AccountID,
		Amount:          args.Amount,
	}

	transactionService.On("Insert", request).Return(response, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInsertTransactionWithBindingError(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	args := struct {
		AccountID     uint64  `json:"account_id"`
		OperationType uint64  `json:"operation_type"`
		Amount        float64 `json:"amount"`
	}{
		AccountID:     1,
		OperationType: 4,
		Amount:        123.56,
	}
	body, _ := json.Marshal(args)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestInsertTransactionWithServiceError(t *testing.T) {
	accountService := &serviceMock.AccountServiceInterface{}
	transactionService := &serviceMock.TransactionServiceInterface{}
	router := api.Start(accountService, transactionService)

	args := struct {
		AccountID     uint64  `json:"account_id"`
		OperationType uint64  `json:"operation_type_id"`
		Amount        float64 `json:"amount"`
	}{
		AccountID:     1,
		OperationType: 4,
		Amount:        123.56,
	}
	body, _ := json.Marshal(args)

	request := entity.Transaction{
		AccountID:       args.AccountID,
		OperationTypeID: args.OperationType,
		Amount:          args.Amount,
	}

	transactionService.On("Insert", request).Return(entity.Transaction{}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
