package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server/routes"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

func NewServer() Server {
	return Server{"8080", gin.Default()}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.Server)
	log.Println("server is running at port: 8080")
	log.Fatal(router.Run(":" + s.Port))
}