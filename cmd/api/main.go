package main

import (
	"log"
	"os"

	"github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/server"
	postgres_database "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres"
	"github.com/axel-andrade/secret-gift-api/internal/infra/bootstrap"
	"github.com/joho/godotenv"
)

/*
* The init function is called after all the variable declarations in the package have evaluated their initializers, and
* those are evaluated only after all the imported packages have been initialized.
 */
func init() {
	if os.Getenv("ENV") != "production" && os.Getenv("ENV") != "testing" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	postgres_database.ConnectDB()
}

func main() {
	dependecies := bootstrap.LoadDependencies()

	server := server.NewServer(os.Getenv("PORT"))
	server.AddRoutes(dependecies)
	server.Run()
}
