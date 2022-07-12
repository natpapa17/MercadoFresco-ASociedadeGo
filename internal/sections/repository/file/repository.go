package sections

import (
	"context"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) sections.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]sections.Section, error) {
	var ss []sections.Section
	if err := r.db.Read(&ss); err != nil {
		return []sections.Section{}, nil
	}
	return ss, nil
}

func (r *repository) GetById(ctx context.Context, id int) (sections.Section, error) {
	var ss []sections.Section
	if err := r.db.Read(&ss); err != nil {
		return sections.Section{}, errors.New("Unable to read database.")
	}

	for _, s := range ss {
		if s.ID == id {
			return s, nil
		}
	}

	return sections.Section{}, errors.New("Id not found.")
}

func (r *repository) LastID(ctx context.Context) (int, error) {
	var ss []sections.Section
	if err := r.db.Read(&ss); err != nil {
		return 0, err
	}

	if len(ss) == 0 {
		return 0, nil
	}

	return ss[len(ss)-1].ID, nil
}

func (r *repository) HasSectionNumber(ctx context.Context, number int) (bool, error) {
	var ss []sections.Section
	if err := r.db.Read(&ss); err != nil {
		return true, err
	}

	if len(ss) == 0 {
		return false, nil
	}

	for _, s := range ss {
		if s.SectionNumber == number {
			return true, nil
		}
	}

	return false, nil
}

func (r *repository) Add(ctx context.Context, id int, sectionNumber int, currentTemperature float32, minimumTemperature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) (sections.Section, error) {
	var ss []sections.Section
	if err := r.db.Read(&ss); err != nil {
		return sections.Section{}, err
	}

	section := sections.Section{
		ID:                 id,
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}

	ss = append(ss, section)
	if err := r.db.Write(ss); err != nil {
		return sections.Section{}, err
	}
	return section, nil
}

func (r *repository) UpdateById(ctx context.Context, id int, section sections.Section) (sections.Section, error) {
	var ss []sections.Section
	if err := r.db.Read(&ss); err != nil {
		return sections.Section{}, err
	}

	nss := func(old *[]sections.Section, new sections.Section) *[]sections.Section {
		for i, s := range ss {
			if s.ID == id {
				if section.CurrentCapacity != 0 {
					s.CurrentCapacity = section.CurrentCapacity
				}
				if section.CurrentTemperature != 0.0 {
					s.CurrentTemperature = section.CurrentTemperature
				}
				if section.MinimumTemperature != 0.0 {
					s.MinimumTemperature = section.MinimumTemperature
				}
				if section.CurrentCapacity != 0 {
					s.CurrentCapacity = section.CurrentCapacity
				}
				if section.MinimumCapacity != 0 {
					s.MinimumCapacity = section.MinimumCapacity
				}
				if section.MaximumCapacity != 0 {
					s.MaximumCapacity = section.MaximumCapacity
				}
				if section.WarehouseID != 0 {
					s.WarehouseID = section.WarehouseID
				}
				if section.ProductTypeID != 0 {
					s.ProductTypeID = section.ProductTypeID
				}

				ss[i] = s
			}
		}
		return &ss
	}(&ss, section)

	if err := r.db.Write(nss); err != nil {
		return sections.Section{}, err
	}

	ns, err := r.GetById(ctx, id)
	if err != nil {
		return sections.Section{}, err
	}

	return ns, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	var ss []sections.Section
	if err := r.db.Read(&ss); err != nil {
		return err
	}

	found := false
	nss := func(old *[]sections.Section, found *bool) *[]sections.Section {

		for i, s := range ss {
			if s.ID == id {
				ss = append(ss[:i], ss[i+1:]...)
				*found = true
			}
		}

		return &ss

	}(&ss, &found)

	if !found {
		return errors.New("Id not found")
	}

	if err := r.db.Write(nss); err != nil {
		return err
	}

	return nil
}
