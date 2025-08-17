package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MudassirDev/shopify-wishlist/internal/web"
	"github.com/joho/godotenv"
)

var (
	apiCfg = APIConfig{}
)

func init() {
	godotenv.Load()

	port := os.Getenv("PORT")
	validateEnv(port, "PORT")
	apiCfg.port = ":" + port

	dbURL := os.Getenv("DB_URL")
	validateEnv(dbURL, "DB_URL")
	apiCfg.dbURL = dbURL

	handler := web.CreateMuxServe(dbURL)
	apiCfg.handler = handler
}

func main() {
	srv := http.Server{
		Addr:    apiCfg.port,
		Handler: apiCfg.handler,
	}

	fmt.Printf("Server is listening on http://localhost%s\n", apiCfg.port)
	log.Fatal(srv.ListenAndServe())
}

func validateEnv(variable, varName string) {
	if variable == "" {
		log.Fatal("failed to load env variable:", varName)
	}
}
