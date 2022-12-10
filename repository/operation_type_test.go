package repository_test

import (
	"pismo/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFindOperationType(t *testing.T) {
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

	sql := `SELECT (.+) FROM "operation_types" WHERE (.+) ORDER BY "operation_types"."id" LIMIT 1`
	mock.ExpectQuery(sql).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "content"}).AddRow(1, "content"))

	repository := repository.NewOperationTypeRepository(gdb)

	_, err = repository.Find(uint64(1))
	if err != nil {
		t.Errorf("error was not expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindOperationTypeWithError(t *testing.T) {
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

	sql := `SELECT "id", "content" FROM "operation_types" WHERE (.+) ORDER BY "operation_types"."id" LIMIT 1`
	mock.ExpectQuery(sql).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"id", "content"}).AddRow(1, "content"))

	repository := repository.NewOperationTypeRepository(gdb)

	_, err = repository.Find(uint64(1))
	if err == nil {
		t.Errorf("an error was expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err == nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
