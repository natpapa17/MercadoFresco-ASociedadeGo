package services
import(
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/repository"

) 

type Service interface {
	GetAll() ([]domain.Seller, error)
	GetById(id int) (domain.Seller, error)
	Store( cid int, companyName string, address string , telephone string, localityId int ) (domain.Seller, error)
	Update(id , cid int, companyName, address, telephone string, localityId int) (domain.Seller, error)
	Delete(id int) error
}

type service struct {
	repository repository.Repository

}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]domain.Seller, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return sl, nil

}

func (s service) GetById (id int) (domain.Seller, error){
	seller , err := s.repository.GetById(id)
	if err != nil{
		return domain.Seller{}, err
	}
	return seller, nil
}


func (s service) Store(cid int, companyName string, address string , telephone string, localityId int) (domain.Seller, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return domain.Seller{}, err
	}

	lastID++

	seller, err := s.repository.Store(lastID, cid, companyName, address, telephone, localityId)

	if err != nil {
		return domain.Seller{}, err
	}

	

	return seller, nil

}

func (s service) Update(id , cid int, companyName, address, telephone string, localityId int) (domain.Seller, error) {
	seller, err := s.repository.Update(id, cid, companyName, address, telephone, localityId)
	if err != nil {
		return domain.Seller{}, err
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
