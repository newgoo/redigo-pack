package redigo_pack

import (
	"github.com/garyburd/redigo/redis"
)

type StringRds struct {
}

func (s *StringRds) Set(key string, value interface{}, expire ...int64) (err error) {
	c := pool.Get()
	defer c.Close()
	if len(expire) == 0 {
		_, err = redis.String(c.Do("set", key, value))
	} else {
		_, err = redis.String(c.Do("set", key, value, "ex", expire[0]))
	}
	return err
}

func (s *StringRds) Get(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("get", key))
}
