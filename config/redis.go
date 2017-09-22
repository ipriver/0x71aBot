package config

import (
	"github.com/garyburd/redigo/redis"
)

var Rc redis.Conn

func init() {
	Rc, err = redis.DialURL("redis://user:@localhost:6379/0")
	if err != nil {
		panic(err)
	}
}
