package main

import (
	"fmt"
	"os"

	"github.com/joaofilippe/americanas-desafio/api"
	"github.com/joaofilippe/americanas-desafio/repository"
)

func main() {
	repository := getRepository()

	err := repository.Db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	api.Server().Run(":8080")
}

func getRepository() *repository.Repository {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	configRepo := repository.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Driver:   "postgres",
		DbName:   dbName,
		SSLMode:  sslMode,
	}

	repository := repository.NewRepository(configRepo)
	repository.GetConnection()

	return repository
}
