package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
	UserID primitive.ObjectID `bson:"_userId" json:"_userId"`
	EventType int32 `bson:"eventType" json:"eventType"`
}
