package dbConnector

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBConnector struct {
	connectionsString string
	maxOpenDbConn     int
	masIdleDbConn     time.Duration
	maxDbLifetime     time.Duration
}

//Method to build DBConnector struct
func BuildDBConnector(connectionsString string, maxOpenDbConn int, masIdleDbConn time.Duration, maxDbLifetime time.Duration) DBConnector {
	return DBConnector{
		connectionsString: connectionsString,
		masIdleDbConn:     masIdleDbConn,
		maxDbLifetime:     maxDbLifetime * time.Minute,
		maxOpenDbConn:     maxOpenDbConn,
	}
}

func (connector DBConnector) ConnectDBMysql() (*sql.DB, error) {
	if connector.connectionsString == "" {
		log.Fatal("not connection string")
		return nil, nil
	}
	db, err := sql.Open("mysql", connector.connectionsString)
	if err != nil {
		log.Fatal("error with the connection", err.Error())
		return nil, err
	}
	db.SetMaxOpenConns(connector.maxOpenDbConn)
	db.SetConnMaxIdleTime(connector.masIdleDbConn)
	db.SetConnMaxLifetime(connector.maxDbLifetime)
	return db, nil
}
