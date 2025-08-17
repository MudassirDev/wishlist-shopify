package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MudassirDev/shopify-wishlist/internal/web"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var (
	apiCfg = APIConfig{}
	//go:embed db/schema/*.sql
	embedMigrations embed.FS
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

func init() {

	conn, err := sql.Open("libsql", apiCfg.dbURL)
	if err != nil {
		log.Fatal("DB connection failed")
	}
	defer conn.Close()

	goose.SetDialect("turso")
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(conn, "db/schema"); err != nil {
		log.Fatal("failed to run migrations ", err)
	}
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
