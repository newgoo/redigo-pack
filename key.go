package redigo_pack

import "github.com/garyburd/redigo/redis"

type keyRds struct {
}

// 	查找键 [*模糊查找]
func (k *keyRds) Keys(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("keys", key))
}

// 	判断key是否存在
func (k *keyRds) Exists(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("exists", key))
}

// 随机返回一个key
func (k *keyRds) RandomKey() *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("randomkey"))
}

// 返回值类型
func (k *keyRds) Type(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("type", key))
}

// 删除key
func (k *keyRds) Del(keys ...string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("del", redis.Args{}.AddFlat(keys)...))
}

//重命名
func (k *keyRds) Rename(key, newKey string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("rename", key, newKey))
}

// 仅当newkey不存在时重命名
func (k *keyRds) RenameNX(key, newKey string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("renamenx", key, newKey))
}

//	序列化key
func (k *keyRds) Dump(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("dump", key))
}

// 反序列化
func (k *keyRds) Restore(key string, ttl, serializedValue interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("restore", key, ttl, serializedValue))
}

// 秒
func (k *keyRds) Expire(key string, seconds int64) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("expire", key, seconds))
}

// 秒
func (k *keyRds) ExpireAt(key string, timestamp int64) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("expireat", key, timestamp))
}

// 毫秒
func (k *keyRds) Persist(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("persist", key))
}

// 毫秒
func (k *keyRds) PersistAt(key string, milliSeconds int64) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("persistat", key, milliSeconds))
}

// 秒
func (k *keyRds) TTL(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("ttl", key))
}

// 毫秒
func (k *keyRds) PTTL(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("pttl", key))
}

//	同实例不同库间的键移动
func (k *keyRds) Move(key string, db int64) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("move", key, db))
}
