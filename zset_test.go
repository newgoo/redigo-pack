package redigo_pack

import (
	"testing"
)

func Test_zSetRds_ZAdd(t *testing.T) {
	NewConnectionWithFile("redis", "./config/config.ini")

	err := RedigoConn.ZSet.ZAdd("key", map[interface{}]interface{}{3: 3, 1: 4, 5: 5}).error
	if err != nil {
		t.Error(err)
	}
}
