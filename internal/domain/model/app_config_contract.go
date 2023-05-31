package model

import (
  "context"
)

type AppConfigContract interface {
  Create(ctx context.Context, app AppConfig) *AppConfig
  Get(ctx context.Context, id string) *AppConfig
}
