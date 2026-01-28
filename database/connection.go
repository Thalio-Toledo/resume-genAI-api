package database

import (
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
)

func Open() (*sql.DB, error) {
	connString := "sqlserver://DESKTOP-TH88BP7?database=resume-gen-db&trusted_connection=yes"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	// Testa a conexão
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
