CREATE TABLE customers (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  email text UNIQUE NOT NULL
);

CREATE TABLE cloud_resources(
  id BIGSERIAL PRIMARY KEY,
  customer_id BIGINT NOT NULL,
  name text NOT NULL,
  type text NOT NULL,
  region text NOT NULL,
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);