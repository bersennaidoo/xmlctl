// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package repository

import (
	"context"
	//"database/sql"
)

const create = `-- name: Create :one
INSERT INTO xmld (
  name, config
) VALUES (
  $1, $2
)
RETURNING name, config
`

type CreateParams struct {
	Name   string         `json:"name"`
	Config []byte    `json:"config"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Xmld, error) {
	row := q.db.QueryRowContext(ctx, create, arg.Name, arg.Config)
	var i Xmld
	err := row.Scan(&i.Name, &i.Config)
	return i, err
}

const get = `-- name: Get :one
SELECT name, config FROM xmld
WHERE name = $1 LIMIT 1
`

func (q *Queries) Get(ctx context.Context, name string) (Xmld, error) {
	row := q.db.QueryRowContext(ctx, get, name)
	var i Xmld
	err := row.Scan(&i.Name, &i.Config)
	return i, err
}
