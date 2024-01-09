package main

import (
	"os"

	"github.com/joho/godotenv"

	listNode "github.com/joaofilippe/americanas-desafio/internal/list_node"
	"github.com/joaofilippe/americanas-desafio/internal/repository"
	"github.com/joaofilippe/americanas-desafio/internal/webserver"
)

func main() {
	if os.Args[1] == "debug-mode" {
		// Load environment variables
		err := godotenv.Load("../config/.env")
		if err != nil {
			panic(err)
		}
	} else if os.Args[1] == "run-mode" {
		// Load environment variables
		err := godotenv.Load("config/.env")
		if err != nil {
			panic(err)
		}
	}

	repository := getRepository()

	service := listNode.Service{
		Repository: repository,
	}

	webApp := webserver.WebApp{
		Application: &service,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	webApp.Server().Run(":" + port)
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
	if err := repository.CreateListNodeTable(); err != nil {
		panic(err)
	}

	return repository
}
