package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"rule-srv/src/config"
	"rule-srv/src/util"
	"time"
)

var Client *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.Mongodb.Url))
	if err != nil {
		util.Sugar.Infow("mongodb connect fail")
		panic(err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		util.Sugar.Infow("mongodb ping fail")
		panic(err)
	}
	util.Sugar.Infow("mongodb connected.")
}
