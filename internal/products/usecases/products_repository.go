package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"

type RepositoryProduct interface {
	GetAll() (domain.Products, error)
	GetById(id int) (domain.Product, error)
	GetByCode(product_code string) (domain.Product, error)
	Create(product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (domain.Product, error)
	Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (domain.Product, error)
	Delete(id int) error
}
