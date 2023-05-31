package postgres

import (
  "context"
  "database/sql"

  //"github.com/jackc/pgx/v5"

  "github.com/bersennaidoo/xmlctl/internal/domain/model"
)

type PostgresRepo struct {
  DB       *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
  return &PostgresRepo{
    DB: db,
  }
}

func (pr *PostgresRepo) Create(ctx context.Context, app model.AppConfig) *model.AppConfig {
  query := `
  INSERT INTO xmld (config)
  VALUES ($1)
  RETURNING config`

  row := pr.DB.QueryRowContext(ctx, query, app.Config,)

  appconfig := &model.AppConfig{}
  _ = row.Scan(&appconfig.Config)

  return appconfig
}

func (pr *PostgresRepo) Get(ctx context.Context, id string) *model.AppConfig {
  return &model.AppConfig{}
}
