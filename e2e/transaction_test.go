package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pismo/api"
	"pismo/database"
	"pismo/entity"
	"pismo/repository"
	"pismo/service"
	"testing"

	"github.com/go-playground/assert"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func TestEndToEnd(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	dbConn, err := database.CreatePGConn()
	if err != nil {
		t.Fatal(err.Error())
	}
	database.Migrate(dbConn)
	database.Seed(dbConn)

	accountRepository := repository.NewAccountRepository(dbConn)
	transactionRepository := repository.NewTransactionRepository(dbConn)
	operationTypeRepository := repository.NewOperationTypeRepository(dbConn)
	accountService := service.NewAccountService(accountRepository)
	transactionService := service.NewTransactionService(transactionRepository, operationTypeRepository, accountRepository)
	router := api.Start(accountService, transactionService)

	accountArgs := entity.Account{
		DocumentNumber: "12345678900",
	}
	accountBody, _ := json.Marshal(accountArgs)

	req, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(accountBody))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	var accountResponse entity.Account
	json.Unmarshal(res.Body.Bytes(), &accountResponse)

	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/account/%d", accountResponse.ID), nil)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)

	var accountFound entity.Account
	json.Unmarshal(res.Body.Bytes(), &accountFound)

	transactionArgs := entity.Transaction{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          123.45,
	}
	transactionBody, _ := json.Marshal(transactionArgs)

	req, _ = http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(transactionBody))
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
