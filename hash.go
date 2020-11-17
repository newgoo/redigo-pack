package redigo_pack

import "github.com/gomodule/redigo/redis"

type hashRds struct {
}

//exist 为true 表示字段不存则设置其值
func (h *hashRds) HSet(key string, filed, value interface{}, exist ...bool) *Reply {
	c := pool.Get()
	defer c.Close()
	if len(exist) > 0 && exist[0] {
		return getReply(c.Do("hsetex", key, filed, value))
	}
	return getReply(c.Do("hset", key, filed, value))
}

//获取指定字段值
func (h *hashRds) HGet(key string, filed interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hget", key, filed))
}

//获取所有字段及值
func (h *hashRds) HGetAll(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hgetall", key))
}

//设置多个字段及值 [map]
func (h *hashRds) HMSetFromMap(key string, mp map[interface{}]interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hmset", redis.Args{}.Add(key).AddFlat(mp)...))
}

//设置多个字段及值 [struct]
func (h *hashRds) HMSetFromStruct(key string, obj interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hmset", redis.Args{}.Add(key).AddFlat(obj)...))
}

//返回多个字段值
func (h *hashRds) HMGet(key string, fileds interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hmget", redis.Args{}.Add(key).AddFlat(fileds)...))
}

//字段删除
func (h *hashRds) HDel(key string, fileds interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hdel", redis.Args{}.Add(key).AddFlat(fileds)...))
}

//判断字段是否存在
func (h *hashRds) HExists(key string, filed interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hexists", key, filed))
}

//返回所有字段
func (h *hashRds) HKeys(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hkeys", key))
}

//返回字段数量
func (h *hashRds) HLen(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hlen", key))
}

//返回所有字段值
func (h *hashRds) HVals(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hvals", key))
}

//为指定字段值增加
func (h *hashRds) HIncrBy(key string, filed interface{}, increment interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hincrby", key, filed, increment))
}

//为指定字段值增加浮点数
func (h *hashRds) HIncrByFloat(key string, filed interface{}, increment float64) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("hincrbyfloat", key, filed, increment))
}
