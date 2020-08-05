package redis

import(
	redigo "github.com/gomodule/redigo/redis"
	"rule-srv/src/config"
	"rule-srv/src/util"
)

var Redis redigo.Conn

func init() {
	var err error
	Redis, err = redigo.DialURL(config.Cfg.Redis.Url)
	if err != nil {
		panic(err)
	}
	util.Sugar.Infow("redis inited.")
}
