-- name: CreateCustomer :one
INSERT INTO customers (
  name,
  email
) VALUES ($1, $2) RETURNING *;

-- name: CreateCloudResource :one
INSERT INTO cloud_resources (
  customer_id,
  name,
  type,
  region
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateCloudResource :exec
UPDATE cloud_resources
  set name = $2,
  type = $3,
  region = $4
WHERE id = $1;

-- name: DeleteCloudResource :exec
DELETE FROM cloud_resources WHERE id = $1;

-- name: FindCloudResourceByID :one
SELECT * FROM cloud_resources WHERE id = $1 LIMIT 1;

-- name: FindCloudResourceByCustomer :many
SELECT * FROM cloud_resources WHERE customer_id = $1;


