package sellers

import (
	"errors"
	"fmt"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)


var sl []Seller = []Seller{}

type Repository interface{
	GetAll() ([]Seller, error)
	GetById(id int) (Seller, error)
	Store(id int, cid int, companyName string, address string , telephone string ) (Seller , error)
	LastID() (int, error)
	Update(id , cid int, companyName, address, telephone string) (Seller, error)
	Delete(id int) error
	

}

type repository struct{

	db store.Store
}

func (r *repository) LastID() (int, error) {
	var sl[]Seller
	if err := r.db.Read(&sl); err != nil {
		return 0, err
	}

	if len(sl) == 0 {
		return 0, nil
	}

	return sl[len(sl)-1].Id, nil
}

func (r *repository) GetAll() ([]Seller, error) {
	var sl []Seller 
	if err := r.db.Read(&sl); err != nil {
		return []Seller{}, nil
	}
	return sl, nil
}

func (r *repository) GetById(id int) (Seller, error){
	var sl []Seller 
	if err := r.db.Read(&sl); err != nil {
		return Seller{}, nil
	}
	

	for _, s := range sl{
		if s.Id == id{
			return s, nil
		}
	}

	
	return Seller{},  errors.New("Nao encontrado")
}

func (r *repository) Store(id int, cid int, companyName string, address string , telephone string) (Seller, error) {
	var sl []Seller 
	if err := r.db.Read(&sl); err != nil {
		return Seller{}, err
	}
	s := Seller{id, cid, companyName, address, telephone}
	sl = append(sl, s)
	if err := r.db.Write(sl); err != nil {
		return Seller{}, err
	}
	return s, nil
}

func (r repository) Update(id , cid int, companyName, address, telephone string) (Seller, error) {
	if err := r.db.Read(&sl); err != nil {
		return Seller{}, nil
	}
	s := Seller{Id: id, Cid: cid, CompanyName: companyName, Address: address, Telephone: telephone}
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
		return Seller{}, err
	}


	if !updated {
		return Seller{}, fmt.Errorf("vendedor %d não encontrado", id)
	}
	return s, nil
}

func (r repository) Delete(id int) error {
	var sl []Seller 
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