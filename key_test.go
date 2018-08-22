package redigo_pack

import "testing"

func TestKeyRds_Del(t *testing.T) {
	NewConnectionWithFile("redis", "./config/config.ini")
	key := "1"
	err := RedigoConn.String.Set(key, 2, 10).error
	if err != nil {
		t.Error(err)
	}
	exist, err := RedigoConn.Key.Exists(key).Bool()
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Error(err)
	}

	err = RedigoConn.Key.Del([]string{key}).error
	if err != nil {
		t.Error(err)
	}
	exist, err = RedigoConn.Key.Exists(key).Bool()
	if err != nil {
		t.Error(err)
	}
	if exist {
		t.Error(err)
	}

}
