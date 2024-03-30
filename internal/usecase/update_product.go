package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/ProductAPI/internal/entity"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
)

type UpdateProductInputDTO struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ProductTypeId string `json:"product_type_id"`
}

type UpdateProductUseCase struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewUpdateProductUseCase(repo database.ProductRepositoryInterface) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		ProductRepository: repo,
	}
}

func (u *UpdateProductUseCase) Execute(ctx context.Context, input UpdateProductInputDTO) error {

	pUUID, err := uuid.Parse(input.Id)

	ptUUID := uuid.Nil
	if input.ProductTypeId != "" {
		ptUUID, err = uuid.Parse(input.ProductTypeId)
	}

	if err != nil {
		return err
	}
	p := entity.Product{
		ID:          pUUID,
		Name:        input.Name,
		Description: input.Description,
		ProductType: entity.ProductType{
			ID: ptUUID,
		},
	}

	return u.ProductRepository.Update(ctx, &p)
}
