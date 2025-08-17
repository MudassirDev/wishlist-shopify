-- name: CreateWishlistEntry :one
INSERT INTO wishtlist_entries (
  product_id, customer_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetWishlistEntries :many
SELECT * FROM wishtlist_entries WHERE customer_id = ?;
