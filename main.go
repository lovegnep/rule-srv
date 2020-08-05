package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"rule-srv/src/handler"
	"rule-srv/subscriber"

	rulesrv "rule-srv/proto/rule-srv"
	_ "rule-srv/src/services/mongodb"
	_ "rule-srv/src/services/redis"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.rule-srv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	rulesrv.RegisterRuleSrvHandler(service.Server(), new(handler.RuleSrv))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.rule-srv", service.Server(), new(subscriber.RuleSrv))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
