// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tokens.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteTokens = `-- name: DeleteTokens :exec
DELETE FROM tokens
WHERE admin_id = $1
`

func (q *Queries) DeleteTokens(ctx context.Context, adminID int32) error {
	_, err := q.db.Exec(ctx, deleteTokens, adminID)
	return err
}

const getToken = `-- name: GetToken :one
SELECT hash, admin_id, expires_at FROM tokens
WHERE hash = $1
`

func (q *Queries) GetToken(ctx context.Context, hash []byte) (Token, error) {
	row := q.db.QueryRow(ctx, getToken, hash)
	var i Token
	err := row.Scan(&i.Hash, &i.AdminID, &i.ExpiresAt)
	return i, err
}

const saveToken = `-- name: SaveToken :one
INSERT INTO tokens (hash, admin_id, expires_at)
VALUES ($1, $2, $3)
RETURNING hash, admin_id, expires_at
`

type SaveTokenParams struct {
	Hash      []byte             `json:"hash"`
	AdminID   int32              `json:"admin_id"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
}

func (q *Queries) SaveToken(ctx context.Context, arg SaveTokenParams) (Token, error) {
	row := q.db.QueryRow(ctx, saveToken, arg.Hash, arg.AdminID, arg.ExpiresAt)
	var i Token
	err := row.Scan(&i.Hash, &i.AdminID, &i.ExpiresAt)
	return i, err
}
