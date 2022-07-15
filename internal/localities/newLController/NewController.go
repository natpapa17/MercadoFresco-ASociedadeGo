package newLController

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/controllers/locality"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/db"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/repository"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/services"
)

func NewLocalityController() *locality.LocalityController{
	lr := repository.CreateMySQLRepository(db.GetInstance())
	pr := repository.CreateProvincyRepository(db.GetInstance())
	ar := repository.CreateCountryRepository(db.GetInstance())
	ls := services.NewService(lr, pr, ar)
	lc := locality.NewLocality(ls)
	return lc
}