package fsys

import (
  "os"
  "context"

  "github.com/bersennaidoo/xmlctl/internal/repository"
)

type Handler struct {
  Qsrv            *repository.Queries
  AppName         string
  Fname           string
}

func NewHandler(qsrv *repository.Queries, fname string, appname string) *Handler {

  return &Handler{
    Qsrv:     qsrv,
    AppName:  appname,
    Fname:    fname,
  }
}

func (h *Handler) UploadFile() repository.Xmld {
  data, _ := os.ReadFile(h.Fname)

  xmld := repository.CreateParams{
    Name: h.AppName,
    Config: data,
  }

  ctx, cancel := context.WithCancel(context.Background())

  defer cancel()

  app, _ := h.Qsrv.Create(ctx, xmld)

  return app
}

