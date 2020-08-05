package config

import (
	microConfig "github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"rule-srv/src/util"
	"time"
)

type RedisCfg struct {
	Url string
}

type MongoCfg struct {
	Url string
}

type Config struct {
	Timeout int
	Redis RedisCfg
	Mongodb MongoCfg
}

var Cfg Config

func init () {
	if err := microConfig.Load(file.NewSource(file.WithPath("./config/default.yml"))); err != nil {
		panic(err)
	}
	if err := microConfig.Scan(&Cfg); err != nil {
		panic(err)
	}
	util.Sugar.Infow("config inited.", "time", time.Now().String())
}
