package repositories

import (
	"github.com/cristhianRamirez3010/market-place-go/src/v1/models"
	"github.com/cristhianRamirez3010/market-place-go/src/v1/repositories/impl"
)

type IDocumentsRepository interface {
	FindAllDocuments() []models.DomDocumentModel
}

func BuildIDocumentsRepository() IDocumentsRepository {
	return impl.BuildDocumentsRepository()
}
