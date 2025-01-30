package controllers

import (
	"github.com/AjinkyaSalunke22/SocioGram---Go-Lang-APIs-PostgreSQL.git/initializers"
	"github.com/AjinkyaSalunke22/SocioGram---Go-Lang-APIs-PostgreSQL.git/models"
	"github.com/gin-gonic/gin"
)

func CreateSocio(c *gin.Context) {
	var eventBody struct{
		Post string
		Image string
	}

	c.Bind(&eventBody)

	socio := models.SocioGram{Post: eventBody.Post, Image: eventBody.Image}

	result := initializers.DB.Create(&socio)

	if result.Error != nil{
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "Server error in creating event",
		})
		return
	}

	if eventBody.Post == "" || eventBody.Image == "" {
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "Sociogram fields cannot be empty",
		})
		return
	}

	c.JSON(200, gin.H{
		"event" : socio,
		"message": "SocioGram creation successfull",
	})
}

func GetSocios(c *gin.Context) {
	
	var socios []models.SocioGram
	initializers.DB.Find(&socios)

	if socios == nil{
		c.Status(404)
		c.JSON(404, gin.H{
			"message": "No event is registered at this time",
		})
		return
	}

	c.JSON(200, gin.H{
		"event" : socios,
		"message": "Events fetched successfully",
	})

}

func GetSocio(c *gin.Context) {
	
	id := c.Param("id")

	var socio []models.SocioGram
	initializers.DB.Find(&socio, id)

	if len(socio) == 0{
		c.Status(404)
		c.JSON(404, gin.H{
			"message": "No socio is registered with this ID",
		})
		return
	}

	c.JSON(200, gin.H{
		"event" : socio,
		"message": "SocioGram fetched successfully",
	})
}

func UpdateSocio(c *gin.Context){

	id := c.Param("id")

	var socio struct{
		Post string
		Image string
	}

	c.Bind(&socio)

	var db_socio []models.SocioGram
	initializers.DB.First(&db_socio, id)

	if len(db_socio) == 0{
		c.Status(404)
		c.JSON(404, gin.H{
			"message": "No socio is registered with this ID",
		})
		return
	}

	initializers.DB.Model(&db_socio).Updates(models.SocioGram{Post: socio.Post, Image: socio.Image})

	c.JSON(200, gin.H{
		"event" : db_socio,
		"message": "SocioGram updated successfully",
	})

}

func DeleteSocio(c *gin.Context){
	id := c.Param("id")

	var db_socio []models.SocioGram
	initializers.DB.First(&db_socio, id)

	if len(db_socio) == 0{
		c.Status(404)
		c.JSON(404, gin.H{
			"message": "No socio is registered with this ID",
		})
		return
	}

	initializers.DB.Delete(&models.SocioGram{}, id)

	c.JSON(200, gin.H{
		"event" : "",
		"message": "SocioGram deleted successfully",
	})

}