package impl

import (
	"log"

	"github.com/cristhianRamirez3010/market-place-go/src/configuration/constants"
	"github.com/cristhianRamirez3010/market-place-go/src/configuration/dbConnector"
	"github.com/cristhianRamirez3010/market-place-go/src/v1/models"
	"github.com/cristhianRamirez3010/market-place-go/src/v1/queries/domDocuments"
)

type DocumentsRepository struct {
	constants constants.Constants
}

func BuildDocumentsRepository() DocumentsRepository {
	return DocumentsRepository{
		constants: constants.BuildConstants(),
	}
}

func (d *DocumentsRepository) loadConnection() dbConnector.DBConnector {
	return dbConnector.BuildDBConnector(
		d.constants.GetConnectionString(),
		d.constants.GetMaxOpenDbConn(),
		d.constants.GetMasIdleDbConn(),
		d.constants.GetMaxDbLifetime(),
	)
}

func (d DocumentsRepository) FindAllDocuments() []models.DomDocumentModel {
	var documents []models.DomDocumentModel

	db, err := d.loadConnection().ConnectDBMysql()
	if err != nil {
		log.Fatal("problems with db connection", err)
		return nil
	}
	defer db.Close()
	rows, err := db.Query(domDocuments.GET_ALL_DOCUMENTS)
	if err != nil {
		log.Fatal("problems with the query or Data base", err)
		return nil
	}

	for rows.Next() {
		document := models.DomDocumentModel{}
		document.ScanModel(rows)
		documents = append(documents, document)
	}

	return documents
}
