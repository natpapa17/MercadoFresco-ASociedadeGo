package main

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/cmd/server"
)

func main() {
	envFilePath, err := filepath.Abs("" + ".env")
	if err != nil {
		log.Fatal("failed to create .env path")
	}
	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("failed to load .env")
	}
	s := server.NewServer()
	s.Run()
}
