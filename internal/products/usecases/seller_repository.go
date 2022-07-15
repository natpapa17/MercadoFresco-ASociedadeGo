package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"

type Repository interface {
	GetById(id int) (domain.Seller, error)
}
