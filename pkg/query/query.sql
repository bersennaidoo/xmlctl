-- name: Create :one
INSERT INTO xmld (
  name, config
) VALUES (
  $1, $2
)
RETURNING *;

-- name: Get :one
SELECT * FROM xmld
WHERE name = $1 LIMIT 1;
