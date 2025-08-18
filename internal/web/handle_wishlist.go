package web

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MudassirDev/shopify-wishlist/db/database"
)

func (c *config) handleWishlistEntry(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ProductHandle string `json:"product_handle"`
		CustomerID    int64  `json:"customer_id"`
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&req); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "failed to decode request", err)
		return
	}

	entry, err := c.DB.CreateWishlistEntry(context.Background(), database.CreateWishlistEntryParams{
		ProductHandle: req.ProductHandle,
		CustomerID:    req.CustomerID,
	})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "failed to created entry", err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, entry)
}

func (c *config) handleGetEntry(w http.ResponseWriter, r *http.Request) {
	rawCustomerID := r.PathValue("customerID")
	customerID, _ := strconv.Atoi(rawCustomerID)
	wishlistEntries, err := c.DB.GetWishlistEntries(context.Background(), int64(customerID))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "no wishlist entries", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, wishlistEntries)
}

func (c *config) handleDeleteEntry(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ProductHandle string `json:"product_handle"`
		CustomerID    int64  `json:"customer_id"`
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&req); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "failed to decode request", err)
		return
	}

	err := c.DB.DeleteWishlistEntry(context.Background(), database.DeleteWishlistEntryParams{
		CustomerID:    req.CustomerID,
		ProductHandle: req.ProductHandle,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "failed to delete entry", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, struct {
		Msg string `json:"message"`
	}{
		Msg: "Deleted successfully",
	})
}
