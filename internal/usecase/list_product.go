package usecase

import (
	"context"
	"database/sql"

	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
)

type ProductOutputDTO struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	CreatedAt     sql.NullTime `json:"created_at"`
	ProductTypeId string       `json:"product_type_id"`
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
			ID:            product.ID.String(),
			Name:          product.Name,
			Description:   product.Description,
			CreatedAt:     product.CreatedAt,
			ProductTypeId: product.ProductType.ID.String(),
		}
	}

	return outputList, nil
}
