package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"

type LocalityRepository interface {
	GetById(id int) (domain.Locality, error)
	GetAll() (domain.Localities, error)
}
