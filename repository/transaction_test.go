package repository_test

import (
	"database/sql/driver"
	"pismo/entity"
	"pismo/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestInsertTransaction(t *testing.T) {
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

	sqlInsert := `INSERT INTO "transactions" (.+) RETURNING`
	mock.ExpectBegin()
	mock.ExpectQuery(sqlInsert).WithArgs(1, 1, 123.45, AnyTime{}, AnyTime{}).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repository := repository.NewTransactionRepository(gdb)

	transaction := entity.Transaction{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          123.45,
	}

	_, err = repository.Insert(transaction)
	if err != nil {
		t.Errorf("error was not expected while inserting: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertTransactionWithError(t *testing.T) {
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

	sqlInsert := `INSERT INTO "transactions" (.+) RETURNING`
	mock.ExpectBegin()
	mock.ExpectQuery(sqlInsert).WithArgs(12345).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repository := repository.NewTransactionRepository(gdb)

	transaction := entity.Transaction{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          123.45,
	}

	_, err = repository.Insert(transaction)
	if err == nil {
		t.Errorf("an error was  expected while inserting: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err == nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
