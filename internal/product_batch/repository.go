package product_batch

import (
	"context"
)

type Repository interface {
	GetById(ctx context.Context, id int) ([]ProductsReport, error)
	Add(ctx context.Context, id int, batchNumber int, currentQuantity int, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minimumTemperature int, productID int, sectionID int) (ProductBatch, error)
	LastID(ctx context.Context) (int, error)
	HasBatchNumber(ctx context.Context, number int) (bool, error)
}
