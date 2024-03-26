package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/ProductAPI/internal/entity"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/database/sqlc"
)

type ProductRepository struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

func NewProductRepository(dbConn *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB:      dbConn,
		Queries: sqlc.New(dbConn),
	}
}

func (r *ProductRepository) Save(ctx context.Context, product *entity.Product) (entity.Product, error) {

	sqlcProduct, err := r.Queries.CreateProduct(ctx, sqlc.CreateProductParams{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
	})

	if err != nil {
		return entity.Product{}, err
	}

	id, err := uuid.Parse(sqlcProduct.ID)
	if err != nil {
		return entity.Product{}, err
	}

	return entity.Product{
		ID:          id,
		Name:        sqlcProduct.Name,
		Description: sqlcProduct.Description,
		CreatedAt:   sqlcProduct.CreatedAt,
	}, nil
}
