-- +goose Up
CREATE TABLE wishtlist_entries (
  product_id INTEGER NOT NULL,
  customer_id INTEGER NOT NULL,
  UNIQUE(product_id, customer_id)
);

-- +goose Down
DROP TABLE wishtlist_entries;
