package usecase

import (
	"context"
	"database/sql"

	"github.com/raulsilva-tech/ProductAPI/internal/entity"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
)

type ProductOutputDTO struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   sql.NullTime       `json:"created_at"`
	ProductType entity.ProductType `json:"product_type"`
}

type ListProductUseCase struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewListProductUseCase(repo database.ProductRepositoryInterface) *ListProductUseCase {
	return &ListProductUseCase{
		ProductRepository: repo,
	}
}

func (l *ListProductUseCase) Execute(ctx context.Context) ([]ProductOutputDTO, error) {

	productList, err := l.ProductRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	outputList := make([]ProductOutputDTO, len(productList))
	for i, product := range productList {
		outputList[i] = ProductOutputDTO{
			ID:          product.ID.String(),
			Name:        product.Name,
			Description: product.Description,
			CreatedAt:   product.CreatedAt,
			ProductType: product.ProductType,
		}
	}

	return outputList, nil
}
