package impl

import (
	"github.com/cristhianRamirez3010/market-place-go/src/v1/models"
	"github.com/cristhianRamirez3010/market-place-go/src/v1/repositories"
)

type DomainsHandler struct {
	documentsRepository repositories.IDocumentsRepository
}

func BuildDomainsHandler() DomainsHandler {
	return DomainsHandler{
		documentsRepository: repositories.BuildIDocumentsRepository(),
	}
}

func (d DomainsHandler) GetAllDocuments() []models.DomDocumentModel {
	return d.documentsRepository.FindAllDocuments()
}
