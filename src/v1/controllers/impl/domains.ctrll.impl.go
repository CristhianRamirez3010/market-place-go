package impl

import (
	"net/http"

	"github.com/cristhianRamirez3010/market-place-go/src/v1/handlers"
	"github.com/gin-gonic/gin"
)

type DomainsController struct {
	userHandler handlers.IDomainsHandler
}

func BuildDomainsController() DomainsController {
	return DomainsController{
		userHandler: handlers.BuildIDomainsHandler(),
	}
}

func (d DomainsController) GetDocuments(c *gin.Context) {
	c.JSON(http.StatusOK, d.userHandler.GetAllDocuments())
}
