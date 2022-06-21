package sections

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

type Repository interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	LastID() (int, error)
	HasSectionNumber(number int) (bool, error)
	Add(id int, sectionNumber int, currentTemperature float32, minimumTemprarature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) (Section, error)
	UpdateById(id int, section Section) (Section, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Section, error) {
	var ss []Section
	if err := r.db.Read(&ss); err != nil {
		return []Section{}, nil
	}
	return ss, nil
}

func (r *repository) GetById(id int) (Section, error) {
	var ss []Section
	if err := r.db.Read(&ss); err != nil {
		return Section{}, errors.New("Unable to read database.")
	}

	for _, s := range ss {
		if s.ID == id {
			return s, nil
		}
	}

	return Section{}, errors.New("Id not found.")
}

func (r *repository) LastID() (int, error) {
	var ss []Section
	if err := r.db.Read(&ss); err != nil {
		return 0, err
	}

	if len(ss) == 0 {
		return 0, nil
	}

	return ss[len(ss)-1].ID, nil
}

func (r *repository) HasSectionNumber(number int) (bool, error) {
	var ss []Section
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

func (r *repository) Add(id int, sectionNumber int, currentTemperature float32, minimumTemprarature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) (Section, error) {
	var ss []Section
	if err := r.db.Read(&ss); err != nil {
		return Section{}, err
	}

	section := Section{
		ID:                  id,
		SectionNumber:       sectionNumber,
		CurrentTemperature:  currentTemperature,
		MinimumTemprarature: minimumTemprarature,
		CurrentCapacity:     currentCapacity,
		MinimumCapacity:     minimumCapacity,
		MaximumCapacity:     maximumCapacity,
		WarehouseID:         warehouseID,
		ProductTypeID:       productTypeID,
	}

	ss = append(ss, section)
	if err := r.db.Write(ss); err != nil {
		return Section{}, err
	}
	return section, nil
}

func (r *repository) UpdateById(id int, section Section) (Section, error) {
	var ss []Section
	if err := r.db.Read(&ss); err != nil {
		return Section{}, err
	}

	nss := func(old *[]Section, new Section) *[]Section {
		for i, s := range ss {
			if s.ID == id {
				if section.CurrentCapacity != 0 {
					s.CurrentCapacity = section.CurrentCapacity
				}
				if section.CurrentTemperature != 0.0 {
					s.CurrentTemperature = section.CurrentTemperature
				}
				if section.MinimumTemprarature != 0.0 {
					s.MinimumTemprarature = section.MinimumTemprarature
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
		return Section{}, err
	}

	ns, err := r.GetById(id)
	if err != nil {
		return Section{}, err
	}

	return ns, nil
}

func (r *repository) Delete(id int) error {
	var ss []Section
	if err := r.db.Read(&ss); err != nil {
		return err
	}

	found := false
	nss := func(old *[]Section, found *bool) *[]Section {

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
