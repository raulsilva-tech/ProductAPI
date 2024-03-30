package database

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *ProductRepository) Update(ctx context.Context, product *entity.Product) error {

	_, err := r.Queries.GetProduct(ctx, product.ID.String())
	if err != nil {
		return err
	}
	return r.Queries.UpdateProduct(ctx, sqlc.UpdateProductParams{
		ID:            product.ID.String(),
		Name:          product.Name,
		Description:   product.Description,
		ProductTypeID: sql.NullString{String: product.ProductType.ID.String(), Valid: false},
	})

}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {

	_, err := r.Queries.GetProduct(ctx, id)
	if err != nil {
		return err
	}
	err = r.Queries.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) GetById(ctx context.Context, id string) (entity.Product, error) {

	sqlcProduct, err := r.Queries.GetProduct(ctx, id)
	if err != nil {
		return entity.Product{}, err
	}

	ptUUID := uuid.Nil
	if sqlcProduct.ProductTypeID.String != "" {
		ptUUID, err = uuid.Parse(sqlcProduct.ProductTypeID.String)
		if err != nil {
			return entity.Product{}, fmt.Errorf("invalid product type id. Error: %v", err.Error())
		}
	}
	return entity.Product{
		ID:          uuid.MustParse(sqlcProduct.ID),
		Name:        sqlcProduct.Name,
		Description: sqlcProduct.Description,
		CreatedAt:   sqlcProduct.CreatedAt,
		ProductType: entity.ProductType{
			ID: ptUUID,
		},
	}, nil
}
