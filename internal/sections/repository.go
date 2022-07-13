package sections

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Section, error)
	GetById(ctx context.Context, id int) (Section, error)
	LastID(ctx context.Context) (int, error)
	HasSectionNumber(ctx context.Context, number int) (bool, error)
	Add(ctx context.Context, id int, sectionNumber int, currentTemperature float32, minimumTemprarature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) (Section, error)
	UpdateById(ctx context.Context, id int, section Section) (Section, error)
	Delete(ctx context.Context, id int) error
}
