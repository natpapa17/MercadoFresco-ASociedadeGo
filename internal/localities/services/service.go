package services

import (


	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/repository"
)

type Service interface {
	Create(name string, province_id, country_id int) (domain.FullLocality, error)
	ReportById(id int) (repository.LocalityReport, error)
	ReportAll() ([]repository.LocalityReport, error)
}

type service struct {
	localityrepository repository.LocalityRepository
	provinceRepository repository.ProvincyRepository
	countryRepository repository.CountryRepository
	
}



func (s *service) Create(name string, province_id, country_id int) (domain.FullLocality,  error) {
	province, err := s.provinceRepository.GetById(province_id)
	

	if err != nil{
		return domain.FullLocality{}, err
	}

	country, a:= s.countryRepository.GetById(country_id)
	if a != nil{
		return domain.FullLocality{}, err
	}

	if err != nil{
		return domain.FullLocality{}, err
	}
	locality, err := s.localityrepository.Create(name, province_id)

	if err != nil{
		return domain.FullLocality{}, err
	}
	fullLocality := domain.FullLocality{
		Id: locality.Id,
		Name: name,
		ProvinceName: province.Name,
		CountryName: country.Name,


	}


	return fullLocality, nil
}

func (s *service) ReportAll() ([]repository.LocalityReport, error) {
	data, err := s.localityrepository.ReportAll()

	if err != nil{
		return []repository.LocalityReport{}, err
	}

	return data, nil
}


func (s *service) ReportById(id int) (repository.LocalityReport, error) {
	report, err := s.localityrepository.ReportById(id)
	if err != nil {
		return repository.LocalityReport{}, err
	}
	return report, nil
}

func NewService(l repository.LocalityRepository, p repository.ProvincyRepository, c repository.CountryRepository ) Service {
	return &service{
		localityrepository: l,
		provinceRepository: p,
		countryRepository: c,
	}
}
