package main

import (
	"log"

	"github.com/AjinkyaSalunke22/SocioGram---Go-Lang-APIs-PostgreSQL.git/initializers"
	"github.com/AjinkyaSalunke22/SocioGram---Go-Lang-APIs-PostgreSQL.git/models"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.SocioGram{})
	log.Println("Model creation successfull")
}