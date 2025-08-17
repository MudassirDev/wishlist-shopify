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

	mux.HandleFunc("POST /api/wishlistentry", cfg.handleWishlistEntry)
	mux.HandleFunc("POST /api/deleteentry", cfg.handleDeleteEntry)
	mux.HandleFunc("GET /api/wishlistentries/{customerID}", cfg.handleGetEntry)

	return configureCors(mux)
}

func configureCors(mux *http.ServeMux) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://theuvpen.myshopify.com/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	return c.Handler(mux)
}
