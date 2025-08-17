-- name: CreateCartEntry :one
INSERT INTO cart_entries (
  items, customer_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetCartEntry :one
SELECT * FROM cart_entries WHERE customer_id = ?;
