package log

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"gopkg.in/olivere/elastic.v5"
)

var (
	ErrNoElasticClientProvided = errors.New("no elastic client provided")
	ErrCannotCreateIndex = errors.New("cannot create index")
)

type IndexNameFunc func() string

func NewElasticSearchLogger(client *elastic.Client, host string, index string, logger log.Logger) (log.Logger, error) {
	return generateElasticLogger(client, host, func() string { return index }, logger)
}

func generateElasticLogger(client *elastic.Client, host string, indexFunc IndexNameFunc, logger log.Logger) (log.Logger, error) {
	ctx, cancel := context.WithCancel(context.TODO())

	if client == nil {
		return logger, ErrNoElasticClientProvided
	}
	exists, err := client.IndexExists(indexFunc()).Do(ctx)
	if err != nil {
		return logger, err
	}

	if !exists {
		createIndex, err := client.CreateIndex(indexFunc()).Do(ctx)
		if err != nil {
			return logger, err
		}
		if !createIndex.Acknowledged {
			return logger, ErrCannotCreateIndex
		}
	}

	return elasticLogger{
		next:      logger,
		client:    client,
		ctx:       ctx,
		ctxCancel: cancel,
		host:      host,
		index:     indexFunc,
	}, nil
}

type elasticLogger struct {
	next log.Logger

	client *elastic.Client

	index IndexNameFunc

	ctx context.Context

	ctxCancel context.CancelFunc

	host string
}

func (el elasticLogger) Log(keyvals ...interface{}) error {
	var msg map[string]interface{}
	msg = make(map[string]interface{})
	msg["host"] = el.host

	for i := 0; i < len(keyvals); i += 2 {
		msg[keyvals[i].(string)] = keyvals[i+1]
	}

	el.client.
		Index().
		Index(el.index()).
		Type("log").
		BodyJson(msg).
		Do(el.ctx)

	return el.next.Log(keyvals...)
}
