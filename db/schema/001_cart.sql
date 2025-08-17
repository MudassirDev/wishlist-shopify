-- +goose Up
CREATE TABLE cart_entries (
  items TEXT NOT NULL,
  customer_id INTEGER NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE cart_entries;
