package routes

import (
	"net/http"

	"github.com/cristhianRamirez3010/market-place-go/src/v1/routes/impl"
)

type IRoutes interface {
	LoadRouteHandler() http.Handler
}

func BuildIRoutes() IRoutes {
	return impl.BuildRoutes()
}
