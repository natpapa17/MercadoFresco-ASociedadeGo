package sections

import (
	"context"
	"errors"
)

var ss []Section = []Section{}

type Service interface {
	GetAll(ctx context.Context) ([]Section, error)
	GetById(ctx context.Context, id int) (Section, error)
	LastID(ctx context.Context) (int, error)
	HasSectionNumber(ctx context.Context, number int) (bool, error)
	Add(ctx context.Context, sectionNumber int, currentTemperature float32, minimumTemprarature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) (Section, error)
	UpdateById(ctx context.Context, id int, section Section) (Section, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll(ctx context.Context) ([]Section, error) {
	ss, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (s service) GetById(ctx context.Context, id int) (Section, error) {
	ss, err := s.repository.GetById(ctx, id)
	if err != nil {
		return Section{}, err
	}
	return ss, nil
}

func (s *service) LastID(ctx context.Context) (int, error) {
	id, err := s.repository.LastID(ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) HasSectionNumber(ctx context.Context, number int) (bool, error) {
	has, err := s.repository.HasSectionNumber(ctx, number)
	if err != nil {
		return true, err
	}
	return has, nil
}

func (s *service) Add(ctx context.Context, sectionNumber int, currentTemperature float32, minimumTemprarature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) (Section, error) {
	has, err := s.HasSectionNumber(ctx, sectionNumber)
	if err != nil {
		return Section{}, errors.New("unable to add")
	}

	if has {
		return Section{}, errors.New("section already exists")
	}

	id, err := s.LastID(ctx)
	if err != nil {
		return Section{}, errors.New("unable to add")
	}

	id++

	ss, err := s.repository.Add(ctx, id, sectionNumber, currentTemperature, minimumTemprarature, currentCapacity, minimumCapacity, maximumCapacity, warehouseID, productTypeID)
	if err != nil {
		return Section{}, err
	}
	return ss, nil
}

func (s *service) UpdateById(ctx context.Context, id int, section Section) (Section, error) {
	has, err := s.HasSectionNumber(ctx, id)

	if err != nil {
		return Section{}, errors.New("unable to update")
	}

	if !has {
		return Section{}, errors.New("inexistent section")
	}

	ss, err := s.repository.UpdateById(ctx, id, section)
	if err != nil {
		return Section{}, err
	}
	return ss, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
