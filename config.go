package redigo_pack

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config struct {
	Host      string `ini:"host"`
	Port      int64  `ini:"port"`
	Db        int64  `ini:"db"`
	Password  string `ini:"password"`
	MaxIdle   int    `ini:"max_idle"`   //default 10  最大连接数
	MaxActive int    `ini:"max_active"` //default 10000 连接池最大数
	Wait      bool   `ini:"wait"`
}

func getConfigWithFile(tagName, path string) (*Config, error) {
	f, err := ini.Load(path)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	f.BlockMode = false
	config := new(Config)
	if err = f.Section(tagName).MapTo(&config); err != nil {
		return nil, err
	}
	if config.MaxActive == 0 {
		config.MaxActive = 10000
	}
	if config.MaxIdle == 0 {
		config.MaxIdle = 10
	}
	return config, nil
}

func getConfig(host string, port int64, password string, db int64, args ...interface{}) (*Config, error) {
	config := &Config{
		Host:     host,
		Port:     port,
		Db:       db,
		Password: password,
	}
	if len(args) > 0 {
		var ok bool
		config.MaxIdle, ok = args[0].(int)
		if !ok {
			return config, fmt.Errorf("参数错误")
		}
	}
	if len(args) > 1 {
		var ok bool
		config.MaxActive, ok = args[1].(int)
		if !ok {
			return config, fmt.Errorf("参数错误")
		}
	}
	if len(args) > 2 {
		var ok bool
		config.Wait, ok = args[2].(bool)
		if !ok {
			return config, fmt.Errorf("参数错误")
		}
	}
	return config, nil
}
