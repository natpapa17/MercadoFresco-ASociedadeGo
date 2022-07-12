package newController

import (
	controllers "github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/seller"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	repository "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/repository/mySql"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/services"
)


func NewSellerController() *controllers.SellerController {

	sellerRepo := repository.CreateMySQLRepository(db.GetInstance())
	sellerService := services.NewService(sellerRepo)
	sellercontroller := controllers.NewSeller(sellerService)

	return sellercontroller
}