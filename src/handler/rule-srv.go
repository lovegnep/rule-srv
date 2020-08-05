package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	rulesrv "rule-srv/proto/rule-srv"
)

type RuleSrv struct{}

func (e *RuleSrv) Event(ctx context.Context, req *rulesrv.EventRequest, rsp *rulesrv.EventResponse) error {
	rsp.Status = 1
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *RuleSrv) Call(ctx context.Context, req *rulesrv.Request, rsp *rulesrv.Response) error {
	log.Info("Received RuleSrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *RuleSrv) Stream(ctx context.Context, req *rulesrv.StreamingRequest, stream rulesrv.RuleSrv_StreamStream) error {
	log.Infof("Received RuleSrv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&rulesrv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *RuleSrv) PingPong(ctx context.Context, stream rulesrv.RuleSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&rulesrv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
