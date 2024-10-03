// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (name, description, price)
VALUES ($1, $2, $3)
RETURNING id, name, description, price
`

type CreateProductParams struct {
	Name        pgtype.Text    `json:"name"`
	Description pgtype.Text    `json:"description"`
	Price       pgtype.Numeric `json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct, arg.Name, arg.Description, arg.Price)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :one
DELETE FROM products
WHERE id = $1
RETURNING id, name, description, price
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRow(ctx, deleteProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, description, price FROM products
WHERE id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
SELECT id, name, description, price FROM products
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetProducts(ctx context.Context, arg GetProductsParams) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET name = $1, description = $2, price = $3
WHERE id = $4
RETURNING id, name, description, price
`

type UpdateProductParams struct {
	Name        pgtype.Text    `json:"name"`
	Description pgtype.Text    `json:"description"`
	Price       pgtype.Numeric `json:"price"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}
