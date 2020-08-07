# RuleSrv Service

该服务由以下命令产生

```
micro new --namespace=go.micro --type=service rule-srv
```

## 用法

该demo集成了makefile

构建可执行文件

```
make build
```

运行服务
```
./rule-srv-service
```

构建镜像
```
make docker
```

## 依赖

请先安装mongodb和redis，并修改对应的配置

另外，以下依赖也需要安装
```
brew install protobuf
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/protoc-gen-micro/v2
```

## 数据模型

共有两种数据，事件和日志。请病假、主管审批，病重均为事件，另外主管审批和病重会产生小明去医院的日志。
主管审批、病重均依赖于事件请病假，当这两种事件发生时要进行合法性校验，即判断之前是否有请病假的事件。
主管审批、病重这两个事件如果合法，则会修改对应的请病假事件的状态。

- 事件

谁做了什么事情，以及事情的状态
```
type Event struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
	UserID primitive.ObjectID `bson:"_userId" json:"_userId"`
	EventType int32 `bson:"eventType" json:"eventType"`
	RefID primitive.ObjectID `bson:"_refId" json:"_refId"`
	Status int32 `bson:"status" json:"status"`
	Created time.Time `bson:"created" json:"created"`
	Updated time.Time `bson:"updated" json:"updated"`
}
```
- 日志

只有当发生合法的审批或者病重时才会产生日志
```
type Log struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	UserID    primitive.ObjectID `bson:"_userId" json:"_userId"`
	EventType int32              `bson:"eventType" json:"eventType"`
	Created   time.Time          `bson:"created" json:"created"`
	Updated   time.Time          `bson:"updated" json:"updated"`
}
```
## 特别说明

本demo中的消息队列用redis的列表实现