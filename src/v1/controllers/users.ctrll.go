package controllers

import (
	"github.com/cristhianRamirez3010/market-place-go/src/v1/controllers/impl"
	"github.com/gin-gonic/gin"
)

type IUsersController interface {
	CreateUser(c *gin.Context)
}

func BuildIUsersController() IUsersController {
	return impl.BuildUsersController()
}
