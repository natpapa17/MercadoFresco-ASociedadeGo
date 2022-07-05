package factories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
)

func MakeWarehouseController() *adapters.WarehouseController {
	wr := adapters.CreateMySQLRepository(db.GetInstance())
	ws := usecases.CreateService(wr)
	wc := adapters.CreateWarehouseController(ws)

	return wc
}
