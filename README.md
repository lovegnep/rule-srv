# RuleSrv Service

This is the RuleSrv service

Generated with

```
micro new --namespace=go.micro --type=service rule-srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.rule-srv
- Type: service
- Alias: rule-srv

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./rule-srv-service
```

Build a docker image
```
make docker
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