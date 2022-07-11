package products_rec

type ProductRepository interface {
	GetById(id int) (Product, error)
}
