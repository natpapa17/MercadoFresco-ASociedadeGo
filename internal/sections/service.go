package sections

import "errors"

var ss []Section = []Section{}

type Service interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	LastID() (int, error)
	HasSectionNumber(number int) (bool, error)
	Add(section Section) (Section, error)
	UpdateById(id int, section Section) (Section, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]Section, error) {
	ss, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (s service) GetById(id int) (Section, error) {
	ss, err := s.repository.GetById(id)
	if err != nil {
		return Section{}, err
	}
	return ss, nil
}

func (s *service) LastID() (int, error) {
	id, err := s.repository.LastID()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) HasSectionNumber(number int) (bool, error) {
	has, err := s.repository.HasSectionNumber(number)
	if err != nil {
		return true, err
	}
	return has, nil
}

func (s *service) Add(section Section) (Section, error) {
	has, err := s.HasSectionNumber(section.SectionNumber)
	if err != nil {
		return Section{}, errors.New("unable to add")
	}

	if has {
		return Section{}, errors.New("section already exists")
	}

	id, err := s.LastID()
	if err != nil {
		return Section{}, errors.New("unable to add")
	}

	id++

	section.ID = id

	ss, err := s.repository.Add(section)
	if err != nil {
		return Section{}, err
	}
	return ss, nil
}

func (s *service) UpdateById(id int, section Section) (Section, error) {
	has, err := s.HasSectionNumber(section.SectionNumber)
	if err != nil {
		return Section{}, errors.New("unable to update")
	}

	if !has {
		return Section{}, errors.New("inexistent section")
	}

	ss, err := s.repository.UpdateById(id, section)
	if err != nil {
		return Section{}, err
	}
	return ss, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
