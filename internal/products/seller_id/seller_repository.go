package seller_id

type Repository interface {
	GetById(id int) (Seller, error)
}
