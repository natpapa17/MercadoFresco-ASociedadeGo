package usecases

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
)

type CarrierRepository interface {
	Create(cid string, companyName string, address string, telephone string, localityId int) (domain.Carrier, error)
	GetNumberOfCarriersPerLocality(localityId int) (int, error)
	GetByCid(cid string) (domain.Carrier, error)
}
