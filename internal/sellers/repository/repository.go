package repository

import (
	"errors"
	"fmt"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)


var sl []domain.Seller = []domain.Seller{}

type Repository interface{
	GetAll() ([]domain.Seller, error)
	GetById(id int) (domain.Seller, error)
	Store(id int, cid int, companyName string, address string , telephone string , localityId int) (domain.Seller , error)
	LastID() (int, error)
	Update(id , cid int, companyName, address, telephone string, localityId int) (domain.Seller, error)
	Delete(id int) error
	

}

type repository struct{

	db store.Store
}

func (r *repository) LastID() (int, error) {
	var sl[]domain.Seller
	if err := r.db.Read(&sl); err != nil {
		return 0, err
	}

	if len(sl) == 0 {
		return 0, nil
	}

	return sl[len(sl)-1].Id, nil
}

func (r *repository) GetAll() ([]domain.Seller, error) {
	var sl []domain.Seller 
	if err := r.db.Read(&sl); err != nil {
		return []domain.Seller{}, nil
	}
	return sl, nil
}

func (r *repository) GetById(id int) (domain.Seller, error){
	var sl []domain.Seller 
	if err := r.db.Read(&sl); err != nil {
		return domain.Seller{}, nil
	}
	

	for _, s := range sl{
		if s.Id == id{
			return s, nil
		}
	}

	
	return domain.Seller{},  errors.New("nao encontrado")
}

func (r *repository) Store(id int, cid int, companyName string, address string , telephone string, localityId int) (domain.Seller, error) {
	var sl []domain.Seller 
	if err := r.db.Read(&sl); err != nil {
		return domain.Seller{}, err
	}
	s := domain.Seller{id, cid, companyName, address, telephone, localityId}
	sl = append(sl, s)
	if err := r.db.Write(sl); err != nil {
		return domain.Seller{}, err
	}
	return s, nil
}

func (r repository) Update(id , cid int, companyName, address, telephone string, localityId int) (domain.Seller, error) {
	if err := r.db.Read(&sl); err != nil {
		return domain.Seller{}, nil
	}
	s := domain.Seller{Id: id, Cid: cid, CompanyName: companyName, Address: address, Telephone: telephone, LocalityId : localityId}
	updated := false
	for i := range sl {
		if sl[i].Id == id {
			s.Id = id
			sl[i] =s
			updated = true
		}
	}


	if err := r.db.Write(&sl); err != nil {
		fmt.Println("Write Error")
		return domain.Seller{}, err
	}


	if !updated {
		return domain.Seller{}, fmt.Errorf("vendedor %d não encontrado", id)
	}
	return s, nil
}

func (r repository) Delete(id int) error {
	var sl []domain.Seller 
	if err := r.db.Read(&sl); err != nil {
		return  nil
	}

	deleted := false
	var index int
	for i := range sl {
		if sl[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("vendedor %d nao encontrado", id)
	}

	sl= append(sl[:index], sl[index+1:]...)

	if err := r.db.Write(&sl); err != nil {
		fmt.Println("Write Error")
		return err
	}
	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}