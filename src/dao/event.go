package dao

import(
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"rule-srv/src/schema"
	"rule-srv/src/services/mongodb"
)

var EventDao *Event

type Event struct {
	db *mongo.Collection
}

func init() {
	EventDao = &Event{db:mongodb.Client.Collection("event")}
}

func (e *Event) Create (ctx context.Context, event schema.Event) (interface{}, error) {
	rsp, err := e.db.InsertOne(ctx, event)
	if err != nil {
		return nil, err
	}
	return rsp.InsertedID, nil
}

func (e *Event) Find (ctx context.Context, query bson.M) ([]schema.Event, error) {
	var events []schema.Event
	cursor, err := e.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}

func (e *Event) FindOne (ctx context.Context, query bson.D) (*schema.Event, error) {
	var event schema.Event
	if err := e.db.FindOne(ctx, query).Decode(&event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *Event) UpdateOne (ctx context.Context, query, update bson.D) (*mongo.UpdateResult, error) {
	rsp, err := e.db.UpdateOne(ctx, query, update)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
