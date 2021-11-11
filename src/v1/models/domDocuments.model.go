package models

import (
	"database/sql"
	"log"
)

type DomDocumentModel struct {
	Id           int64          `db:"id"`
	Name         string         `db:"name"`
	Type         string         `db:"type"`
	CreationDate sql.NullTime   `db:"creation_date"`
	UpdateDate   sql.NullTime   `db:"update_date"`
	CreationUser sql.NullString `db:"creation_user"`
	UpdateUser   sql.NullString `db:"update_user"`
}

func (m *DomDocumentModel) ScanModel(rows *sql.Rows) {

	err := rows.Scan(
		&m.Id,
		&m.Name,
		&m.Type,
		&m.CreationDate,
		&m.UpdateDate,
		&m.CreationUser,
		&m.UpdateUser,
	)
	if err != nil {
		log.Fatal("err", err.Error())
	}
}
