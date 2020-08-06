package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"rule-srv/src/handler"

	rulesrv "rule-srv/proto/rule-srv"
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

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
