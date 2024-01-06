package repository

import "github.com/jmoiron/sqlx"

// Repository is a struct to hold the database connection
type Repository struct {
	Db *sqlx.DB
}

// Config is a struct to hold the database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
	Driver   string
}

// GetConnection returns a connection to the database
func (r *Repository) GetConnection() *sqlx.DB {
	var dsn string
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
