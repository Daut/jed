package testutils

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
)

var user = "jed"
var password = "jed"
var databaseName = "jed_shop"

func NewDBContainer(t *testing.T, queries []string) *gnomock.Container {
	t.Helper()
	queries = append([]string{"CREATE EXTENSION IF NOT EXISTS pgcrypto;"}, queries...)
	p := postgres.Preset(
		postgres.WithUser(user, password),
		postgres.WithDatabase(databaseName),
		postgres.WithQueries(queries...),
		postgres.WithQueriesFile("../../../db/schema.sql"),
		postgres.WithTimezone("Europe/Belgrade"),
	)
	container, err := gnomock.Start(p)
	if err != nil {
		t.Fatal(err)
	}
	return container
}

func NewDBConn(t *testing.T, container *gnomock.Container) *pgx.Conn {
	t.Helper()
	ctx := context.Background()
	port := container.DefaultPort()
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, container.Host, port, databaseName)
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		t.Fatal(err)
	}
	return conn
}
