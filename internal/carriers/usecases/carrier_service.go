package usecases

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
)

type CarrierService interface {
	Create(cid string, companyName string, address string, telephone string, localityId int) (domain.Carrier, error)
	GetNumberOfCarriersPerLocalities(localitiesIds []int) (domain.ReportsNumberOfCarriersPerLocality, error)
	GetAllNumberOfCarriersPerLocality() (domain.ReportsNumberOfCarriersPerLocality, error)
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

	if err != nil && !errors.Is(err, ErrNoElementFound) {
		return domain.Carrier{}, err
	}

	_, err = s.localityRepository.GetById(localityId)

	if err != nil && errors.Is(err, ErrNoElementFound) {
		return domain.Carrier{}, ErrInvalidLocalityId
	}

	if err != nil {
		return domain.Carrier{}, err
	}

	carrier, err := s.carrierRepository.Create(cid, companyName, address, telephone, localityId)

	if err != nil {
		return domain.Carrier{}, err
	}

	return carrier, nil
}

func (s *carrierService) GetNumberOfCarriersPerLocalities(localitiesIds []int) (domain.ReportsNumberOfCarriersPerLocality, error) {
	reports := domain.ReportsNumberOfCarriersPerLocality{}

	for _, localityId := range localitiesIds {
		locality, err := s.localityRepository.GetById(localityId)

		if errors.Is(err, ErrNoElementFound) {
			continue
		}

		if err != nil {
			return domain.ReportsNumberOfCarriersPerLocality{}, err
		}

		carriersPerLocality, err := s.carrierRepository.GetNumberOfCarriersPerLocality(localityId)

		if err != nil {
			return domain.ReportsNumberOfCarriersPerLocality{}, err
		}

		report := domain.ReportNumberOfCarriersPerLocality{
			LocalityId:    locality.Id,
			LocalityName:  locality.Name,
			CarriersCount: carriersPerLocality,
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func (s *carrierService) GetAllNumberOfCarriersPerLocality() (domain.ReportsNumberOfCarriersPerLocality, error) {
	localities, err := s.localityRepository.GetAll()

	if err != nil {
		return domain.ReportsNumberOfCarriersPerLocality{}, err
	}

	reports := domain.ReportsNumberOfCarriersPerLocality{}

	for _, locality := range localities {
		carriersPerLocality, err := s.carrierRepository.GetNumberOfCarriersPerLocality(locality.Id)

		if err != nil {
			return domain.ReportsNumberOfCarriersPerLocality{}, err
		}

		report := domain.ReportNumberOfCarriersPerLocality{
			LocalityId:    locality.Id,
			LocalityName:  locality.Name,
			CarriersCount: carriersPerLocality,
		}

		reports = append(reports, report)
	}

	return reports, nil
}
