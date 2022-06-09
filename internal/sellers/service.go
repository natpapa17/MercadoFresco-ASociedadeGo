package sellers

type Service interface {
	GetAll() ([]Seller, error)
	GetById(id int) (Seller, error)
	Store( cid int, companyName string, address string , telephone string ) (Seller, error)
	Update(id , cid int, companyName, address, telephone string) (Seller, error)
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

func (s service) GetAll() ([]Seller, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return sl, nil

}

func (s service) GetById (id int) (Seller, error){
	seller , err := s.repository.GetById(id)
	if err != nil{
		return Seller{}, err
	}
	return seller, nil
}


func (s service) Store(cid int, companyName string, address string , telephone string) (Seller, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Seller{}, err
	}

	lastID++

	seller, err := s.repository.Store(lastID, cid, companyName, address, telephone)

	if err != nil {
		return Seller{}, err
	}

	

	return seller, nil

}

func (s service) Update(id , cid int, companyName, address, telephone string) (Seller, error) {
	seller, err := s.repository.Update(id, cid, companyName, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, err
}



func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
