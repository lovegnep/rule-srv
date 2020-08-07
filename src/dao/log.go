package dao

import(
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"rule-srv/src/schema"
	"rule-srv/src/services/mongodb"
)

var LogDao *Log

type Log struct {
	db *mongo.Collection
}

func init() {
	LogDao = &Log{db:mongodb.Client.Collection("log")}
}

func (e *Log) Create (ctx context.Context, event schema.Log) (interface{}, error) {
	rsp, err := e.db.InsertOne(ctx, event)
	if err != nil {
		return nil, err
	}
	return rsp.InsertedID, nil
}

func (e *Log) Find (ctx context.Context, query bson.M) ([]schema.Log, error) {
	var logs []schema.Log
	cursor, err := e.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}

func (e *Log) FindOne (ctx context.Context, query bson.M) (*schema.Log, error) {
	var log schema.Log
	if err := e.db.FindOne(ctx, query).Decode(&log); err != nil {
		return nil, err
	}
	return &log, nil
}
