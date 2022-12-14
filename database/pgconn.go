package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreatePGConn() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PASSWORD"))
	// db, _ = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	// db = db.Exec("CREATE DATABASE pismo;")
	// if db.Error != nil {
	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=pismo port=5432 sslmode=disable TimeZone=America/Sao_Paulo", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PASSWORD"))
	db, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		return nil, err
	}
	// }

	return db, nil
}
