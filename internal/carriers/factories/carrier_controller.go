package factories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases"
)

func MakeCarrierController() *adapters.CarrierController {
	cr := adapters.CreateCarrierMySQLRepository(db.GetInstance())
	lr := adapters.CreateLocalityMySQLRepository(db.GetInstance())
	cs := usecases.CreateCarrierService(cr, lr)
	cc := adapters.CreateCarryController(cs)

	return cc
}
