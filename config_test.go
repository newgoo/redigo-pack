package redigo_pack

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetConfigWithFilr(t *testing.T) {
	logrus.Info(getConfigWithFile("redis", "./config/config.ini"))
}

func TestNewConnection(t *testing.T) {
	//logrus.Info(NewConnectionWithFile("redis", "./config/config.ini"))
	RedigoConn.String.Set("1", 2, 10)
	logrus.Info(RedigoConn.String.Get("1").Int64())
}
