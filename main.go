package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/joaofilippe/americanas-desafio/api"
	listNode "github.com/joaofilippe/americanas-desafio/list_node"
	"github.com/joaofilippe/americanas-desafio/repository"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	repository := getRepository()

	service := listNode.Service{
		Repository: repository,
	}

	webApp := api.WebApp{
		Application: &service,
	}

	webApp.Server().Run(":8080")
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
