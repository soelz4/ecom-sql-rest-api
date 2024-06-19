package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func InitMySQLStorage(db *sql.DB) {
	// Start DataBase Connection
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DataBase Succecfully Connected")
}
