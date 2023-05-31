package service

import (
  "context"

  "github.com/bersennaidoo/xmlctl/internal/repository/postgres"
  "github.com/bersennaidoo/xmlctl/internal/domain/model"
)

type AppConfigService struct {
  DB         *postgres.PostgresRepo
}

func NewAppConfigService(db *postgres.PostgresRepo) *AppConfigService {
  return &AppConfigService{
    DB: db,
  }
}

func (acs *AppConfigService) Create(ctx context.Context, app model.AppConfig) *model.AppConfig {
  appconfig := acs.DB.Create(ctx, app)

  return appconfig
}

func (acs *AppConfigService) Get(ctx context.Context, id string) *model.AppConfig {
  return &model.AppConfig{}
}
