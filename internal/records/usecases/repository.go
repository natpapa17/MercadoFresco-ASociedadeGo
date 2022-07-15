package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/record_domain"

type RecordsRepository interface {
	GetRecordsPerProduct(product_id int) (int, error)
	Create(last_update_date string, purchase_price int, sale_price int, product_id int) (record_domain.Records, error)
}

type ProductRepository interface {
	GetById(id int) (record_domain.Product, error)
}
