package web

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MudassirDev/shopify-wishlist/db/database"
)

func (c *config) handleCartEntry(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Items      string `json:"items"`
		CustomerID int    `json:"customer_id"`
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&req); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "failed to decode request", err)
		return
	}

	entry, err := c.DB.CreateCartEntry(context.Background(), database.CreateCartEntryParams{
		Items:      req.Items,
		CustomerID: int64(req.CustomerID),
	})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "failed to created entry", err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, entry)
}

func (c *config) handleGetEntry(w http.ResponseWriter, r *http.Request) {
}
