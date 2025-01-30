package main

import (
	"github.com/AjinkyaSalunke22/SocioGram---Go-Lang-APIs-PostgreSQL.git/controllers"
	"github.com/AjinkyaSalunke22/SocioGram---Go-Lang-APIs-PostgreSQL.git/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/socio", controllers.CreateSocio)
	r.GET("/socios", controllers.GetSocios)
	r.GET("/socio/:id", controllers.GetSocio)
	r.PUT("/socio/:id", controllers.UpdateSocio)
	r.DELETE("/socio/:id", controllers.DeleteSocio)
	r.Run() 
}