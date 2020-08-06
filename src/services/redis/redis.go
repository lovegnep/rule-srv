package redis

import(
	redigo "github.com/gomodule/redigo/redis"
	"rule-srv/src/config"
	"rule-srv/src/util"
	"time"
)

var Redis *redigo.Pool

func init() {

	Redis = &redigo.Pool{
		MaxIdle:     config.Cfg.Redis.MaxIdle,
		MaxActive:   config.Cfg.Redis.MaxActive,
		IdleTimeout: time.Duration(config.Cfg.Redis.MaxIdleTimeout) * time.Second,
		Wait:        config.Cfg.Redis.Wait,
		Dial: func() (redigo.Conn, error) {
			con, err := redigo.DialURL(config.Cfg.Redis.Url)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	util.Sugar.Infow("redis inited.")
}

func GetClient () redigo.Conn {
	return Redis.Get()
}
