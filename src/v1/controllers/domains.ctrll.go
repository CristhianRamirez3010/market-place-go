package controllers

import (
	"github.com/cristhianRamirez3010/market-place-go/src/v1/controllers/impl"
	"github.com/gin-gonic/gin"
)

type IDomainsController interface {
	GetDocuments(c *gin.Context)
}

func BuildIDomainsController() IDomainsController {
	return impl.BuildDomainsController()
}
