-- +goose Up
CREATE TABLE wishlist_entries (
  product_handle TEXT NOT NULL,
  customer_id INTEGER NOT NULL,
  UNIQUE(product_handle, customer_id)
);

-- +goose Down
DROP TABLE wishlist_entries;
