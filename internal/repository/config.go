package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

// Repository is a struct to hold the database connection
type Repository struct {
	Db     *sqlx.DB
	Config Config
}

// Config is a struct to hold the database configuration
type Config struct {
	Host       string
	Port       string
	User       string
	Password   string
	DbName     string
	Driver     string
	StringConn string
}

// NewRepository returns a new Repository
func NewRepository(config Config) *Repository {
	stringConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	config.StringConn = stringConn

	return &Repository{
		Config: config,
	}
}

// GetConnection returns a connection to the database
func (r *Repository) GetConnection() {
	db, err := sqlx.Open(r.Config.Driver, r.Config.StringConn)
	if err != nil {
		panic(err)
	}

	r.Db = db
}
