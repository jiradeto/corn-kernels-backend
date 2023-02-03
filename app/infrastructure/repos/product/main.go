package productrepo

import (
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"gorm.io/gorm"
)

type any interface{}

// Repo ...
type Repo interface {
	CreateOneProduct(tx *gorm.DB, input CreateOneProductInput) (*entities.Product, error)
	FindOneProduct(tx *gorm.DB, input FindOneProductInput) (*entities.Product, error)
	GetProductList(tx *gorm.DB, input GetProductListInput) ([]*entities.Product, error)
	GetDailyStockReport(tx *gorm.DB) (any, error)
}

type repo struct {
	DB *gorm.DB
}

// New is a constructor method of Repo
func New(db *gorm.DB) Repo {
	return &repo{
		DB: db,
	}
}

func (repo *repo) selectDB(injectedDB *gorm.DB) *gorm.DB {
	if injectedDB == nil {
		return repo.DB
	}
	return injectedDB
}

type DailyStockMovement struct {
	ProductID     string `json:"productID"`
	ProductName   string `json:"productName"`
	TotalQuantity int    `json:"totalQuantity"`
	MovementType  string `json:"movementType"`
	MovementDate  string `json:"movementDate"`
}

func dailyStockMovements(query *gorm.DB, movementType string) []DailyStockMovement {
	var stockInQuery string = `
		SELECT 
		products.id AS product_id,
		products.name AS product_name,
		SUM(stock_movements.quantity) AS total_quantity,
		stock_movements.movement_type,
		DATE(stock_movements.created_at) AS movement_date
		FROM 
		products 
		JOIN stock_movements ON products.id = stock_movements.product_id
		WHERE 
		stock_movements.movement_type  = ?
		GROUP BY 
		products.name, stock_movements.movement_type, movement_date
		ORDER BY product_id, movement_date ASC;
	`
	var dailyStockMovementsIn []DailyStockMovement
	query.Raw(stockInQuery, movementType).Scan(&dailyStockMovementsIn)
	return dailyStockMovementsIn
}

type getCurrentAvailableStocksItems struct {
	ProductName  string `json:"productName"`
	ProductID    int    `json:"productID"`
	CurrentStock int    `json:"currentStock"`
}

func getCurrentAvailableStocks(query *gorm.DB) []getCurrentAvailableStocksItems {
	var stockInQuery string = `
		SELECT products.name AS product_name, product_stocks.product_id, SUM(product_stocks.quantity) AS current_stock
		FROM product_stocks
		JOIN products ON product_stocks.product_id = products.id
		GROUP BY product_stocks.product_id;
	`
	var getCurrentAvailableStocksResponse []getCurrentAvailableStocksItems
	query.Raw(stockInQuery).Scan(&getCurrentAvailableStocksResponse)
	return getCurrentAvailableStocksResponse
}

func (repo *repo) GetDailyStockReport(tx *gorm.DB) (any, error) {
	query := repo.selectDB(tx)
	dailyStockMovementsIn := dailyStockMovements(query, "in")
	dailyStockMovementsOut := dailyStockMovements(query, "out")

	dailyMovement := map[string]interface{}{
		"in":  dailyStockMovementsIn,
		"out": dailyStockMovementsOut,
	}
	currentAvailableStocks := getCurrentAvailableStocks(query)

	dashboardData := map[string]interface{}{
		"dailyMovement":   dailyMovement,
		"availableStocks": currentAvailableStocks,
	}

	return dashboardData, nil
}
