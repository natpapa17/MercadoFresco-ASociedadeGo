package products_rec

type Repository interface {
	GetById(id int) (Product, error)
}
