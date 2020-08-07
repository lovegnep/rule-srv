
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	rulesrv "rule-srv/proto/rule-srv"
	"rule-srv/src/constants"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.rule-srv.client"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	client := rulesrv.NewRuleSrvService("go.micro.service.rule-srv", service.Client())
	rsp, err := client.Event(context.TODO(), &rulesrv.EventRequest{
		UserId: primitive.NewObjectID().Hex(),
		Event:  constants.SickLeave,
		//RefId:  "2222",
	})
	if err != nil {
		fmt.Println("rsp error:", err.Error())
		return
	}
	fmt.Println(rsp)
}
