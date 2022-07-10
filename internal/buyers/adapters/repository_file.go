package adapters

import (
	"errors"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases"
)

type BuyerFileRepository struct{
	file store.Store
}

func CreateBuyerRepository(file store.Store) usecases.BuyerRepository {
	return &BuyerFileRepository{
		file: file,
	}
}

func (r *BuyerFileRepository) Create(firstName string, lastName string, address string, document string) (Buyer, error) {
	var bs []Buyer
	if err := r.file.Read(&bs); err != nil {
		return Buyer{}, err
	}

	lastId, err := r.lastId()

	if err != nil {
		return Buyer{}, err
	}
	b := Buyer{
		ID:                 lastId + 1,
		FirstName:      firstName,
		LastName:          lastName,
		Address:            address,
		DocumentNumber:    document,
	}

	bs = append(bs, b)

	if err := r.file.Write(bs); err != nil {
		return Buyer{}, err
	}
	return b, nil
}

func (r *repository) GetAll() ([]Buyer, error) {
	var bs []Buyer
	if err := r.file.Read(&bs); err != nil {
		return []Buyer{}, nil
	}
	return bs, nil
}

func (r *repository) GetBuyerById(id int) (Buyer, error) {
	var bs []Buyer
	if err := r.file.Read(&bs); err != nil {
		return Buyer{}, nil
	}

	for _, b := range bs {
		if b.ID == id {
			return b, nil
		}
	}

	return Buyer{}, &NoElementInFileError{errors.New("can't find element with this id")}

}

func (r *repository) UpdateBuyerById(id int, firstName string, lastName string, address string, document string) (Buyer, error) {
	var bs []Buyer
	if err := r.file.Read(&bs); err != nil {
		return Buyer{}, nil
	}
	result, found := Buyer{}, false
	for i, b := range bs {
		if b.ID == id {
			bs[i], found = Buyer{
				ID:                 id,
				FirstName:      firstName,
				LastName:          lastName,
				Address:            address,
				DocumentNumber:    document,
			}, true
			result = bs[i]
			break
		}
	}

	if !found {
		return Buyer{}, &NoElementInFileError{errors.New("can't find element with this id")}
	}
	
	if err := r.file.Write(bs); err != nil {
		return Buyer{}, err
	}

	return result, nil
}

func (r *repository) DeleteBuyerById(id int) error {
	var bs []Buyer
	if err := r.file.Read(&bs); err != nil {
		return nil
	}
	found := false
	for i, b := range bs {
		if b.ID == id {
			newBs := []Buyer{}
			newBs = append(newBs, bs[:i]...)
			newBs = append(newBs, bs[i+1:]...)
			bs = newBs
			found = true
			break
		}
	}

	if !found {
		return &NoElementInFileError{errors.New("can't find element with this id")}
	}

	if err := r.file.Write(bs); err != nil {
		return err
	}

	return nil
}

func (r *repository) lastId() (int, error) {
	var bs []Buyer
	if err := r.file.Read(&bs); err != nil {
		return 0, err
	}

	if len(bs) == 0 {
		return 0, nil
	}

	return bs[len(bs)-1].ID, nil
}