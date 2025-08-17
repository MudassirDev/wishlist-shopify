package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MudassirDev/shopify-wishlist/db/database"
)

func (c *config) handleCreateCart(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Items      string `json:"items"`
		CustomerID int    `json:"customer_id"`
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{err: 'failed to decode'}"))
		return
	}

	entry, err := c.DB.CreateCartEntry(context.Background(), database.CreateCartEntryParams{
		Items:      req.Items,
		CustomerID: int64(req.CustomerID),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{err: 'failed to create'}"))
		fmt.Println(err)
		return
	}
	data, err := json.Marshal(entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{err: 'failed to marshal'}"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
