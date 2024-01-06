package main

import (
	"fmt"

	"github.com/joaofilippe/americanas-desafio/api"
	"github.com/joaofilippe/americanas-desafio/repository"
)

func main() {
	configRepo := repository.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "12345678",
		DbName:   "list_node",
		Driver:   "postgres",
		SSLMode:  "disable",
	}

	repository := repository.NewRepository(configRepo)
	repository.GetConnection()

	err := repository.Db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	api.Server().Run(":8080")
}
