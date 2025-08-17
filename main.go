package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	apiCfg = APIConfig{}
)

func init() {
	godotenv.Load()

	port := os.Getenv("PORT")
	validateEnv(port, "PORT")
	apiCfg.port = port

	dbURL := os.Getenv("DB_URL")
	validateEnv(dbURL, "DB_URL")
	apiCfg.dbURL = dbURL
}

func main() {
}

func validateEnv(variable, varName string) {
	if variable == "" {
		log.Fatal("failed to load env variable:", varName)
	}
}
