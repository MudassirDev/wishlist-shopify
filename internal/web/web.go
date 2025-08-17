package web

import (
	"database/sql"
	"net/http"

	"github.com/MudassirDev/shopify-wishlist/db/database"
	"github.com/rs/cors"
)

var (
	cfg = config{}
)

func CreateMuxServe(conn *sql.DB) http.Handler {
	cfg.DB = database.New(conn)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/cartentry", cfg.handleCartEntry)
	mux.HandleFunc("GET /api/cartentry", cfg.handleGetEntry)

	return configureCors(mux)
}

func configureCors(mux *http.ServeMux) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	return c.Handler(mux)
}
