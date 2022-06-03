package main

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
