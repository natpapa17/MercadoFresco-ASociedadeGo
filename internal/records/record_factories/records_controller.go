package record_factories

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/adapters"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/usecases"
)

func MakeRecordsController() *adapters.RecordsController {
	rr := adapters.NewRecordMysqlRepository(db.GetInstance())
	rrp := adapters.NewMysqlProductRepository(db.GetInstance())
	rs := usecases.NewRecordsService(rr, rrp)
	rc := adapters.NewRecordController(rs)

	return rc
}
