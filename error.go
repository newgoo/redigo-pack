package redigo_pack

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	ErrorConfig error = fmt.Errorf("请先初始化redis配置")
)

var NilError = redis.ErrNil
