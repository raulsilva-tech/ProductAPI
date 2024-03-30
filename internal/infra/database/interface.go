package database

import (
	"context"

	"github.com/raulsilva-tech/ProductAPI/internal/entity"
)

type ProductRepositoryInterface interface {
	Save(ctx context.Context, product *entity.Product) (entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]entity.Product, error)
	GetById(ctx context.Context, id string) (entity.Product, error)
}
