package usecases

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
)

type CarrierService interface {
	Create(cid string, companyName string, address string, telephone string, localityId int) (domain.Carrier, error)
	GetNumberOfCarriersPerLocality(localityId int) (domain.NumberOfCarriersPerLocality, error)
}

type carrierService struct {
	carrierRepository  CarrierRepository
	localityRepository LocalityRepository
}

func CreateCarrierService(cr CarrierRepository, lr LocalityRepository) CarrierService {
	return &carrierService{
		carrierRepository:  cr,
		localityRepository: lr,
	}
}

func (s *carrierService) Create(cid string, companyName string, address string, telephone string, localityId int) (domain.Carrier, error) {
	c, err := s.carrierRepository.GetByCid(cid)

	if c.Id != 0 {
		return domain.Carrier{}, ErrCidInUse
	}

	if !errors.Is(err, ErrNoElementFound) {
		return domain.Carrier{}, err
	}

	_, err = s.localityRepository.GetById(localityId)

	if errors.Is(err, ErrNoElementFound) {
		return domain.Carrier{}, ErrInvalidLocalityId
	}

	if err != nil {
		return domain.Carrier{}, nil
	}

	carrier, err := s.carrierRepository.Create(cid, companyName, address, telephone, localityId)

	if err != nil {
		return domain.Carrier{}, err
	}

	return carrier, nil
}

func (s *carrierService) GetNumberOfCarriersPerLocality(localityId int) (domain.NumberOfCarriersPerLocality, error) {
	locality, err := s.localityRepository.GetById(localityId)

	if errors.Is(err, ErrNoElementFound) {
		return domain.NumberOfCarriersPerLocality{}, ErrInvalidLocalityId
	}

	if err != nil {
		return domain.NumberOfCarriersPerLocality{}, err
	}

	carriersPerLocality, err := s.carrierRepository.GetNumberOfCarriersPerLocality(localityId)

	if err != nil {
		return domain.NumberOfCarriersPerLocality{}, err
	}

	result := domain.NumberOfCarriersPerLocality{
		LocalityId:    locality.Id,
		LocalityName:  locality.Name,
		CarriersCount: carriersPerLocality,
	}

	return result, nil
}
