package main

import (
	"net/http"

	"github.com/daut/jed/cmd/api/router"
	"github.com/daut/jed/internal/utils"
	db "github.com/daut/jed/sqlc"
	"github.com/jackc/pgx/v5"
)

type Application struct {
	Queries *db.Queries
	Logger  *utils.Logger
	Router  http.Handler
}

func NewApp(conn *pgx.Conn) *Application {
	queries := db.New(conn)
	logger := utils.NewLogger()
	router := router.New(queries, logger)
	return &Application{
		Queries: queries,
		Logger:  logger,
		Router:  router,
	}
}
