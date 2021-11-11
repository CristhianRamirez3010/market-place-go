package impl

import (
	"net/http"

	"github.com/cristhianRamirez3010/market-place-go/src/v1/controllers"
	"github.com/cristhianRamirez3010/market-place-go/src/v1/endpoints/domainsEndpoints"
	"github.com/cristhianRamirez3010/market-place-go/src/v1/endpoints/usersEndpoints"
	"github.com/gin-gonic/gin"
)

type Routes struct{}

func BuildRoutes() Routes {
	return Routes{}
}

func (r Routes) domainsRoutes(gin *gin.Engine) {
	contrl := controllers.BuildIUsersController()
	endpoints := usersEndpoints.BuildUserEndpoints()
	gin.POST(endpoints.Create, contrl.CreateUser)
}

func (r Routes) userRoutes(gin *gin.Engine) {
	contrl := controllers.BuildIDomainsController()
	endpoints := domainsEndpoints.BuildDomainsEndpoints()
	gin.GET(endpoints.GetDocuments, contrl.GetDocuments)
}

func (r Routes) LoadRouteHandler() http.Handler {
	gin := gin.Default()
	r.domainsRoutes(gin)
	r.userRoutes(gin)
	return gin
}
