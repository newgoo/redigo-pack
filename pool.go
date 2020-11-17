package redigo_pack

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initPool(addr, password string) {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   200,
		IdleTimeout: 180 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return setDialog(addr, password)
		},
	}
}

func initPoolByOld(old *redis.Pool) {
	pool = old
}

func setDialog(addr, password string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	if len(password) != 0 {
		if _, err := conn.Do("AUTH", password); err != nil {
			_ = conn.Close()
			return nil, err
		}
	}

	r, err := redis.String(conn.Do("PING"))
	if err != nil || r != "PONG" {
		log.Fatalf("failed to connect redis: %v, err: %v", addr, RedigoConn)
	}

	return conn, nil
}
