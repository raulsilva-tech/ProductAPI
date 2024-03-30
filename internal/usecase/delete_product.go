package usecase

import (
	"context"

	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
)

type DeleteProductUseCase struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewDeleteProductUseCase(repo database.ProductRepositoryInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		ProductRepository: repo,
	}
}

func (u *DeleteProductUseCase) Execute(ctx context.Context, id string) error {

	return u.ProductRepository.Delete(ctx, id)
}
