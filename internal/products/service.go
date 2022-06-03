package products

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (*Product, error)
	Create(ProductCode int, Description string, Width float64, Height float64, Length float64, NetWeight float64, ExpirationRate string, RecommendedFreezingTemperature float64, FreezingRate float64, ProductTypeId string, SellerId string) (Product, error)
	Update(Id int, ProductCode int, Description string, Width float64, Height float64, Length float64, NetWeight float64, ExpirationRate string, RecommendedFreezingTemperature float64, FreezingRate float64, ProductTypeId string, SellerId string) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s service) GetById(id int) (*Product, error) {
	ps, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s service) Create(ProductCode int, Description string, Width float64, Height float64, Length float64, NetWeight float64, ExpirationRate string, RecommendedFreezingTemperature float64, FreezingRate float64, ProductTypeId string, SellerId string) (Product, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Create(lastID, ProductCode, Description, Width, Height, Length, NetWeight, ExpirationRate, RecommendedFreezingTemperature, FreezingRate, ProductTypeId, SellerId)

	if err != nil {
		return Product{}, err
	}

	return product, nil

}

func (s service) Update(lastId, Id int, ProductCode int, Description string, Width float64, Height float64, Length float64, NetWeight float64, ExpirationRate string, RecommendedFreezingTemperature float64, FreezingRate float64, ProductTypeId string, SellerId string) (Product, error) {
	product, err := s.repository.Update(Id, ProductCode, Description, Width, Height, Length, NetWeight, ExpirationRate, RecommendedFreezingTemperature, FreezingRate, ProductTypeId, SellerId)
	if err != nil {
		return Product{}, err
	}
	return product, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
