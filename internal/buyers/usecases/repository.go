package usecases

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/domain"
)

type BuyerRepository interface {
	Create(firstName string, lastName string, address string, document string) (domain.Buyer, error)
	GetAll() (domain.Buyer, error)
	GetBuyerById(id int) (domain.Buyer, error)
	UpdateBuyerById(id int, firstName string, lastName string, address string, document string) (domain.Buyer, error)
	DeleteBuyerById(id int) error
}