package usecase

import (
	"context"

	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
)

type GetProductUseCase struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewGetProductUseCase(repo database.ProductRepositoryInterface) *GetProductUseCase {
	return &GetProductUseCase{
		ProductRepository: repo,
	}
}

func (u *GetProductUseCase) Execute(ctx context.Context, id string) (ProductOutputDTO, error) {

	p, err := u.ProductRepository.GetById(ctx, id)
	if err != nil {
		return ProductOutputDTO{}, err
	}

	return ProductOutputDTO{
		ID:            p.ID.String(),
		Name:          p.Name,
		Description:   p.Description,
		ProductTypeId: p.ProductType.ID.String(),
		CreatedAt:     p.CreatedAt,
	}, nil

}
