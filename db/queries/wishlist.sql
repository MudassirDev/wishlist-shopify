-- name: CreateWishlistEntry :one
INSERT INTO wishlist_entries (
  product_id, customer_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetWishlistEntries :many
SELECT * FROM wishlist_entries WHERE customer_id = ?;

-- name: DeleteWishlistEntry :exec
DELETE FROM wishlist_entries
WHERE customer_id = ?
AND product_id = ?;
