package product_factories

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/usecases"
)

func MakeProductController() *adapters.ProductController {
	pr := adapters.NewProductMysqlRepository(db.GetInstance())
	ps := usecases.NewProductService(pr)
	pc := adapters.NewProductController(ps)

	return pc
}
