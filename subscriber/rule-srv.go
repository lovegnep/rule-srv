package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	rulesrv "rule-srv/proto/rule-srv"
)

type RuleSrv struct{}

func (e *RuleSrv) Handle(ctx context.Context, msg *rulesrv.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *rulesrv.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
