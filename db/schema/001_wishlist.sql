-- +goose Up
CREATE TABLE wishlist_entries (
  product_id INTEGER NOT NULL,
  customer_id INTEGER NOT NULL,
  UNIQUE(product_id, customer_id)
);

-- +goose Down
DROP TABLE wishlist_entries;
