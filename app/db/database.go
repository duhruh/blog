package db

import (
	"context"
	"errors"
	"regexp"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"

	errors2 "github.com/duhruh/blog/app/errors"
	"github.com/duhruh/blog/config"
)

func NewDatabaseConnection(config config.ApplicationConfig, logger log.Logger) DatabaseConnection {
	settings := config.DatabaseConnection()
	conn := mysql.ConnectionURL{
		Host:     settings.Get("host").(string),
		Database: settings.Get("database").(string),
		User:     settings.Get("user").(string),
		Password: settings.Get("password").(string),
	}

	return &databaseConnection{
		connectionUrl: conn,
		logger:        logger,
	}
}

type DatabaseConnection interface {
	ConnectionURL() mysql.ConnectionURL
	Open() sqlbuilder.Database
	Connection() sqlbuilder.Database
	ConnectionWithContext(cxt context.Context) sqlbuilder.Database
	Close() error
}

type databaseConnection struct {
	connectionUrl mysql.ConnectionURL
	session       sqlbuilder.Database
	logger        log.Logger
}

func (db *databaseConnection) ConnectionURL() mysql.ConnectionURL {
	return db.connectionUrl
}

func (db *databaseConnection) Open() sqlbuilder.Database {
	sess, err := mysql.Open(db.ConnectionURL())
	if err != nil {
		panic(err)
	}

	sess.SetLogging(true)
	sess.SetLogger(db)

	db.session = sess

	return db.session
}

func (db *databaseConnection) Close() error {
	return db.session.Close()
}

func (db *databaseConnection) Connection() sqlbuilder.Database {
	return db.ConnectionWithContext(context.Background())
}

func (db *databaseConnection) Log(q *db.QueryStatus) {
	var re = regexp.MustCompile(`[\n\t\s]+`)
	s := re.ReplaceAllString(q.Query, " ")
	err := q.Err
	if err != nil {
		err = errors2.New(err)
	}

	level.Debug(db.logger).Log(
		"query", s,
		"took", q.End.Sub(q.Start),
		"error", err,
		"trace", errors2.StackTrace(err),
		"count", q.Context.Value("count"),
	)
}

func (db *databaseConnection) ConnectionWithContext(cxt context.Context) sqlbuilder.Database {
	if db.session == nil {
		panic(errors.New("nope"))
	}
	return db.session.WithContext(cxt)
}
