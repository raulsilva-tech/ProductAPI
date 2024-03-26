package usecase

import (
	"context"
	"database/sql"

	"github.com/raulsilva-tech/ProductAPI/internal/entity"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
)

type CreateProductInputDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProductOutputDTO struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   sql.NullTime `json:"created_at"`
}

type CreateProductUseCase struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewCreateProductUseCase(repo database.ProductRepositoryInterface) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: repo,
	}
}

func (u *CreateProductUseCase) Execute(ctx context.Context, input CreateProductInputDTO) (CreateProductOutputDTO, error) {

	product, err := entity.NewProduct(input.Name, input.Description)
	if err != nil {
		return CreateProductOutputDTO{}, err
	}

	record, err := u.ProductRepository.Save(ctx, product)
	if err != nil {
		return CreateProductOutputDTO{}, err
	}

	return CreateProductOutputDTO{
		ID:          record.ID.String(),
		Name:        record.Name,
		Description: record.Description,
		CreatedAt:   record.CreatedAt,
	}, nil
}
