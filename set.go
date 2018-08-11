package redigo_pack

type SetRds struct {
}

func (h *SetRds) SAdd(key string, fileds []interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("sadd", key, fileds))
}

func (h *SetRds) SCard(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("scard", key))
}

func (h *SetRds) SMembers(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("smembers", key))
}

func (h *SetRds) SisMember(key string, member interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("sismember", key, member))
}

func (h *SetRds) SPop(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("spop", key))
}

func (h *SetRds) SRandMember(key string, count ...int64) *Reply {
	c := pool.Get()
	defer c.Close()
	if len(count) > 0 {
		return getReply(c.Do("srandmember", key, count[0]))
	}
	return getReply(c.Do("srandmember", key))
}

func (h *SetRds) SRem(key string, members []interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("srem", key, members))
}

func (h *SetRds) SMove(sourceKey, destinationKey string, member interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("smove", sourceKey, destinationKey, member))
}

func (h *SetRds) SDiff(keys []string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("sdiff", keys))
}

func (h *SetRds) SDiffStore(destinationKey string, keys []string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("sdiffstore", destinationKey, keys))
}
