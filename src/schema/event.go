package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Event struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
	UserID primitive.ObjectID `bson:"_userId" json:"_userId"`
	EventType int32 `bson:"eventType" json:"eventType"`
	RefID primitive.ObjectID `bson:"_refId" json:"_refId"`
	Status int32 `bson:"status" json:"status"`
	Created time.Time `bson:"created" json:"created"`
	Updated time.Time `bson:"updated" json:"updated"`
}
