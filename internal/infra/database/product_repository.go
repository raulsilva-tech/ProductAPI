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

	err := r.Queries.CreateProduct(ctx, sqlc.CreateProductParams{
		ID:            product.ID.String(),
		Name:          product.Name,
		Description:   product.Description,
		CreatedAt:     product.CreatedAt,
		ProductTypeID: sql.NullString{String: product.ProductType.ID.String(), Valid: false},
	})

	if err != nil {
		return entity.Product{}, err
	}

	sqlcProduct, err := r.Queries.GetProduct(ctx, product.ID.String())
	if err != nil {
		return entity.Product{}, err
	}

	id, err := uuid.Parse(sqlcProduct.ID)
	if err != nil {
		return entity.Product{}, err
	}

	var typeUUID uuid.UUID
	if sqlcProduct.ProductTypeID.String != "" {
		typeUUID, err = uuid.Parse(sqlcProduct.ProductTypeID.String)
		if err != nil {
			return entity.Product{}, err
		}
	}

	return entity.Product{
		ID:          id,
		Name:        sqlcProduct.Name,
		Description: sqlcProduct.Description,
		CreatedAt:   sqlcProduct.CreatedAt,
		ProductType: entity.ProductType{
			ID: typeUUID,
		},
	}, nil
}

func (r *ProductRepository) List(ctx context.Context) ([]entity.Product, error) {
	sqlcProducts, err := r.Queries.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]entity.Product, len(sqlcProducts))

	for i, sqlcProduct := range sqlcProducts {
		pt := entity.ProductType{ID: uuid.Nil}
		if sqlcProduct.ProductTypeID.Valid {
			pt = entity.ProductType{ID: uuid.MustParse(sqlcProduct.ProductTypeID.String)}
		}
		products[i] = entity.Product{
			ID:          uuid.MustParse(sqlcProduct.ID),
			Name:        sqlcProduct.Name,
			Description: sqlcProduct.Description,
			CreatedAt:   sqlcProduct.CreatedAt,
			ProductType: pt,
		}
	}

	return products, nil
}
