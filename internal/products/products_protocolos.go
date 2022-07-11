package products

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	GetByCode(product_code string) (Product, error)
	Create(product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error)
	Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error)
	Delete(id int) error
}
