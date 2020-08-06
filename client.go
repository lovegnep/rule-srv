
package main
/*
import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	rulesrv "rule-srv/proto/rule-srv"
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
		UserId: "111",
		Event:  0,
		RefId:  "2222",
	})
	if err != nil {
		fmt.Println("rsp error:", err.Error())
		return
	}
	fmt.Println(rsp)
}
*/