package app

import (
	"github.com/duhruh/blog/config"
	"upper.io/db.v3/mysql"
)

func NewDatabaseConnection(config config.ApplicationConfig) DatabaseConnection {
	settings := config.DatabaseConnection()
	conn := mysql.ConnectionURL{
		Host:     settings.Get("host").(string),
		Database: settings.Get("database").(string),
		User:     settings.Get("user").(string),
		Password: settings.Get("password").(string),
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
