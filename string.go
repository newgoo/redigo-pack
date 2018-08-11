package redigo_pack

type StringRds struct {
}

func (s *StringRds) Set(key string, value interface{}, expire ...int64) *Reply {
	c := pool.Get()
	defer c.Close()
	if len(expire) == 0 {
		return getReply(c.Do("set", key, value))
	}
	return getReply(c.Do("set", key, value, "ex", expire[0]))
}

func (s *StringRds) Get(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("get", key))
}
