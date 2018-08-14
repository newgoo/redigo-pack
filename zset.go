package redigo_pack

type zSetRds struct {
}

// map[score]member  添加元素
func (z *zSetRds) ZAdd(key string, mp map[interface{}]interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("zadd", key, mp))
}

// 	增加元素权重
func (z *zSetRds) ZUncrBy(key string, increment, member interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("zuncrby", key, increment, member))
}

// 	增加元素权重
func (z *zSetRds) ZCard(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("zcard", key))
}

// 	返回指定元素的排名
func (z *zSetRds) ZEank(key string, member interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("zrank", key, member))
}

// 	返回指定元素的权重
func (z *zSetRds) ZScore(key string, member interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("zscore", key, member))
}

// 返回集合两个权重间的元素数
func (z *zSetRds) ZCount(key string, min, max interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("zcount", key, min, max))
}

// 返回指定区间内的元素
func (z *zSetRds) ZRange(key string, start, stop interface{}, withScore ...bool) *Reply {
	c := pool.Get()
	defer c.Close()
	if len(withScore) > 0 && withScore[0] {
		return getReply(c.Do("zrange", key, start, stop, withScore))
	}
	return getReply(c.Do("zrange", key, start, stop))
}

// 倒序返回指定区间内的元素
func (z *zSetRds) ZRevrange(key string, start, stop interface{}, withScore ...bool) *Reply {
	c := pool.Get()
	defer c.Close()
	if len(withScore) > 0 && withScore[0] {
		return getReply(c.Do("zrevrange", key, start, stop, withScore))
	}
	return getReply(c.Do("zrevrange", key, start, stop))
}
