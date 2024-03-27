package usecase

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/ProductAPI/internal/entity"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
)

type CreateProductInputDTO struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	ProductTypeId string `json:"product_type_id"`
}

type CreateProductOutputDTO struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   sql.NullTime       `json:"created_at"`
	ProductType entity.ProductType `json:"product_type"`
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

	typeId := uuid.Nil
	var err error

	if input.ProductTypeId != "" {
		typeId, err = uuid.Parse(input.ProductTypeId)
		if err != nil {

			return CreateProductOutputDTO{}, err
		}
	}

	product, err := entity.NewProduct(input.Name, input.Description, entity.ProductType{ID: typeId})
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
		ProductType: record.ProductType,
	}, nil
}
