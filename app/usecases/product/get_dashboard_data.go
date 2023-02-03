package productusecase

import (
	"context"
)

func (uc *useCase) GetDashboardData(ctx context.Context) (any, error) {
	data, err := uc.ProductRepo.GetDailyStockReport(nil)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}
