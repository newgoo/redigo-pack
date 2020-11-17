package redigo_pack

import (
	"github.com/garyburd/redigo/redis"
)

type RedigoPack struct {
	String stringRds
	List   listRds
	Hash   hashRds
	Key    keyRds
	Set    setRds
	ZSet   zSetRds
	Bit    bitRds
	Db     dbRds
}

var RedigoConn = new(RedigoPack)

func NewConnectionWithFile(addr, password string) error {

	initPool(addr, password)
	return nil
}

func NewConnectionByPool(pool2 *redis.Pool) error {
	initPoolByOld(pool2)
	return nil
}
