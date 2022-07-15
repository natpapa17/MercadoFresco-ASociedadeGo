package product_batch

import (
	"context"
	"errors"
)

type Service interface {
	GetById(ctx context.Context, id int) ([]ProductsReport, error)
	Add(ctx context.Context, batchNumber int, currentQuantity int, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minumumTemperature int, productID int, sectionID int) (ProductBatch, error)
	LastID(ctx context.Context) (int, error)
	HasBatchNumber(ctx context.Context, number int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetById(ctx context.Context, id int) ([]ProductsReport, error) {
	prl, err := s.repository.GetById(ctx, id)
	if err != nil {
		return []ProductsReport{}, err
	}

	if len(prl) == 0 {
		return []ProductsReport{}, errors.New("not found")
	}

	return prl, err
}

func (s *service) Add(ctx context.Context, batchNumber int, currentQuantity int, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minimumTemperature int, productID int, sectionID int) (ProductBatch, error) {
	has, err := s.HasBatchNumber(ctx, batchNumber)
	if err != nil {
		return ProductBatch{}, err
	}

	if has {
		return ProductBatch{}, errors.New("batch number already exists")
	}

	id, err := s.LastID(ctx)
	if err != nil {
		return ProductBatch{}, err
	}

	id++

	return s.repository.Add(ctx, id, batchNumber, currentQuantity, currentTemperature, dueDate, initialQuantity, manufacturingDate, manufacturingHour, minimumTemperature, productID, sectionID)
}

func (s *service) LastID(ctx context.Context) (int, error) {
	id, err := s.repository.LastID(ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) HasBatchNumber(ctx context.Context, number int) (bool, error) {
	return s.repository.HasBatchNumber(ctx, number)
}
