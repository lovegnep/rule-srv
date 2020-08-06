package handler

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rule-srv/src/config"
	"rule-srv/src/constants"
	"rule-srv/src/schema"
	"rule-srv/src/util"
	"time"

	log "github.com/micro/go-micro/v2/logger"

	rulesrv "rule-srv/proto/rule-srv"
	"rule-srv/src/dao"
)

type RuleSrv struct{}

func (e *RuleSrv) Event(ctx context.Context, req *rulesrv.EventRequest, rsp *rulesrv.EventResponse) error {
	var notify = true
	var query bson.D
	var update bson.D
	var targetUserID primitive.ObjectID
	dur, _ := time.ParseDuration(fmt.Sprintf("-%dh", config.Cfg.Timeout))
	switch req.Event {
	case constants.SickLeave:
		notify = false
	case constants.Approve:
		if req.RefId == "" {
			return errors.New("refId is not valid.")
		}
		_id, _ := primitive.ObjectIDFromHex(req.RefId)
		event, err := dao.EventDao.FindOne(ctx, bson.D{
			{"_id", _id},
		})
		if err != nil {
			return err
		}
		targetUserID = event.UserID
		query = bson.D{
			{"_id", _id},
			{"created", bson.D{{"$gt", time.Now().Add(dur)}}},
			{"status", constants.EventStatusInit},
		}
		update = bson.D{
			{"$set", bson.D{
				{"status", constants.EventStatusApproved},
				{"updated", time.Now()},
			}},
		}

	case constants.CriticallyIll:
		_userID, _ := primitive.ObjectIDFromHex(req.UserId)
		targetUserID = _userID
		query = bson.D{
			{"_userId", _userID},
			{"created", bson.D{{"$gt", time.Now().Add(dur)}}},
			{"status", constants.EventStatusInit},
		}
		update = bson.D{
			{"$set", bson.D{
				{"status", constants.EventStatusIll},
				{"updated", time.Now()},
			}},
		}
	}
	if notify {
		// 更新之前病假的事件状态
		result, err := dao.EventDao.UpdateOne(ctx, query, update)
		if err != nil {
			return err
		}
		if result.ModifiedCount != 1 {
			return errors.New("refId not exist or timeout or has been approved.")
		}

		// 记录日志
		_, err = dao.LogDao.Create(ctx, schema.Log{
			UserID:    targetUserID,
			EventType: constants.GoToDoctor,
			Created:   time.Now(),
			Updated:   time.Now(),
		})
		if err != nil {
			util.Sugar.Error("create log fail", err.Error(), targetUserID.String(), constants.GoToDoctor)
		}
		// 发送通知到消息队列
	}
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return err
	}
	event := schema.Event{
		UserID:    userID,
		EventType: 0,
		Status:    constants.EventStatusInit,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	if req.RefId != "" {
		_id, _ := primitive.ObjectIDFromHex(req.RefId)
		event.RefID = _id
	}
	_, err = dao.EventDao.Create(ctx, event)
	if err != nil {
		return err
	}
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
