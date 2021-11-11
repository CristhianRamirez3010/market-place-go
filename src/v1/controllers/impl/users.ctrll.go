package impl

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func BuildUsersController() UsersController {
	return UsersController{}
}

type test struct {
	IdLineaBaseDpto int
	IdMunicipio     int
	IdLineaBaseMpo  int
	LineaBase       int
	Other           int
}

func (u UsersController) CreateUser(c *gin.Context) {
	var json test
	err := c.BindJSON(&json)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, json)
}
