package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Rc redis.Conn

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("No connection to redis DB, app can work with errors")
		}
	}()
	Rc, err = redis.DialURL("redis://user:@localhost:6379/0")
	if err != nil {
		panic(err)
	}
}
