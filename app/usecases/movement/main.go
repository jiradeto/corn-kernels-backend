package movementusecase

import (
	"context"

	"github.com/jiradeto/corn-kernels-backend/app/entities"
	movementrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/movement"
	productrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/product"
	stockrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/stock"
)

// UseCase is a declarative interface for movement usecase
type UseCase interface {
	CreateOneStockMovement(ctx context.Context, input CreateOneStockMovementInput) (*entities.StockMovement, error)
	GetMovementList(ctx context.Context, input GetMovementListInput) ([]*entities.StockMovement, error)
}

type useCase struct {
	StockMovementRepo movementrepo.Repo
	ProductRepo       productrepo.Repo
	StockRepo         stockrepo.Repo
}

// New is a constructor method of UseCase
func New(
	stockmovementRepo movementrepo.Repo,
	productRepo productrepo.Repo,
	stockRepo stockrepo.Repo,
) UseCase {
	return &useCase{
		StockMovementRepo: stockmovementRepo,
		ProductRepo:       productRepo,
		StockRepo:         stockRepo,
	}
}
