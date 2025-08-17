-- name: CreateCartEntry :one
INSERT INTO cart_entries (
  items, customer_id
) VALUES (
  ?, ?
)
RETURNING *;
