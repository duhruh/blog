package app

import (
	"github.com/duhruh/tackle"
	"upper.io/db.v3/mysql"
)

func NewDatabaseConnection(config tackle.Config) DatabaseConnection {
	settings := config.DatabaseConnection()
	conn := mysql.ConnectionURL{
		Host:     settings["host"],
		Database: settings["database"],
		User:     settings["user"],
		Password: settings["password"],
	}
	return databaseConnection{
		connection: conn,
	}
}

type DatabaseConnection interface {
	ConnectionURL() mysql.ConnectionURL
}

type databaseConnection struct {
	connection mysql.ConnectionURL
}

func (db databaseConnection) ConnectionURL() mysql.ConnectionURL {
	return db.connection
}
