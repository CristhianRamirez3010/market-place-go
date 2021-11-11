package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cristhianRamirez3010/market-place-go/src/configuration/constants"
	"github.com/cristhianRamirez3010/market-place-go/src/v1/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	constanst := constants.BuildConstants()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", constanst.GetPort()),
		Handler: routes.BuildIRoutes().LoadRouteHandler(),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}
