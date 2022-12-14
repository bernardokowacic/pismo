package repository_test

import (
	"pismo/entity"
	"pismo/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub gorm connection", err)
	}

	sqlInsert := `INSERT INTO "accounts" (.+) RETURNING`
	mock.ExpectBegin()
	mock.ExpectQuery(sqlInsert).WithArgs("12345678900", 5000.00).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repository := repository.NewAccountRepository(gdb)

	account := entity.Account{
		DocumentNumber:       "12345678900",
		AvailableCreditLimit: 5000.00,
	}

	_, err = repository.Insert(account)
	if err != nil {
		t.Errorf("error was not expected while inserting: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertWithError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub gorm connection", err)
	}

	sqlInsert := `INSERT INTO "accounts" (.+) RETURNING`
	mock.ExpectBegin()
	mock.ExpectQuery(sqlInsert).WithArgs(12345).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repository := repository.NewAccountRepository(gdb)

	account := entity.Account{
		DocumentNumber:       "12345678900",
		AvailableCreditLimit: 5000.00,
	}

	_, err = repository.Insert(account)
	if err == nil {
		t.Errorf("an error was  expected while inserting: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err == nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFind(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub gorm connection", err)
	}

	sql := `SELECT (.+) FROM "accounts" WHERE (.+) ORDER BY "accounts"."id" LIMIT 1`
	mock.ExpectQuery(sql).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "document_number", "available_creadit_limit"}).AddRow(1, "12345678900", 5000.00))

	repository := repository.NewAccountRepository(gdb)

	_, err = repository.Find(uint64(1))
	if err != nil {
		t.Errorf("error was not expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindWithError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub gorm connection", err)
	}

	sql := `SELECT (.+) FROM "accounts" WHERE (.+) ORDER BY "accounts"."id" LIMIT 1`
	mock.ExpectQuery(sql).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"id", "document_number"}).AddRow(1, "12345678900"))

	repository := repository.NewAccountRepository(gdb)

	_, err = repository.Find(uint64(1))
	if err == nil {
		t.Errorf("an error was expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err == nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub gorm connection", err)
	}

	sqlUpdate := `Update "accounts" set available_creadit_limit = $1 where id = $2`
	mock.ExpectBegin()
	mock.ExpectQuery(sqlUpdate).WithArgs(5000.00, 1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repository := repository.NewAccountRepository(gdb)

	err = repository.UpdateBalance(uint64(1), 5000.00)
	if err != nil {
		t.Errorf("an unexpected error was expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateBalanceWithError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub gorm connection", err)
	}

	sqlUpdate := `Update "accounts" set available_creadit_limit = $1`
	mock.ExpectBegin()
	mock.ExpectQuery(sqlUpdate).WithArgs("5000.00").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repository := repository.NewAccountRepository(gdb)

	err = repository.UpdateBalance(uint64(1), 5000.00)
	if err == nil {
		t.Errorf("an error was expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err == nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
