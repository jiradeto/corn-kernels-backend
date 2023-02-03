package productusecase

import (
	"context"

	"github.com/jiradeto/corn-kernels-backend/app/entities"
	productrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/product"
	stockrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/stock"
)

type any interface{}

type UseCase interface {
	CreateOneProduct(ctx context.Context, input CreateOneProductInput) (*entities.Product, error)
	GetProductList(ctx context.Context, input GetProductListInput) ([]*entities.Product, error)
	GetDashboardData(ctx context.Context) (any, error)
}

type useCase struct {
	ProductRepo productrepo.Repo
	StockRepo   stockrepo.Repo
}

// New is a constructor method of UseCase
func New(
	productRepo productrepo.Repo,
	stockRepo stockrepo.Repo,
) UseCase {
	return &useCase{
		ProductRepo: productRepo,
		StockRepo:   stockRepo,
	}
}
