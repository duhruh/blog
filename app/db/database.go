package db

import (
	"errors"
	"github.com/duhruh/blog/config"
	"upper.io/db.v3/lib/sqlbuilder"
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

	return &databaseConnection{
		connectionUrl: conn,
	}
}

type DatabaseConnection interface {
	ConnectionURL() mysql.ConnectionURL
	Open() sqlbuilder.Database
	Connection() sqlbuilder.Database
}

type databaseConnection struct {
	connectionUrl mysql.ConnectionURL
	session       sqlbuilder.Database
}

func (db *databaseConnection) ConnectionURL() mysql.ConnectionURL {
	return db.connectionUrl
}

func (db *databaseConnection) Open() sqlbuilder.Database {
	sess, err := mysql.Open(db.ConnectionURL())
	if err != nil {
		panic(err)
	}

	db.session = sess

	return db.session
}

func (db *databaseConnection) Connection() sqlbuilder.Database {
	if db.session == nil {
		panic(errors.New("nope"))
	}
	return db.session
}
