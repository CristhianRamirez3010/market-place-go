package handlers

import (
	"github.com/cristhianRamirez3010/market-place-go/src/v1/handlers/impl"
	"github.com/cristhianRamirez3010/market-place-go/src/v1/models"
)

type IDomainsHandler interface {
	GetAllDocuments() []models.DomDocumentModel
}

func BuildIDomainsHandler() IDomainsHandler {
	return impl.BuildDomainsHandler()
}
