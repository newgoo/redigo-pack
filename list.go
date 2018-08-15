package redigo_pack

type listRds struct {
}

//向列表头插入元素
func (l *listRds) LPush(key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lpush", key, value))
}

//当列表存在则将元素插入表头
func (l *listRds) LPushx(key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lpushx", key, value))
}

//将指定元素插入列表末尾
func (l *listRds) RPush(key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("rpush", key, value))
}

//当列表存在则将元素插入表尾
func (l *listRds) RPushx(key string, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("rpushx", key, value))
}

//将元素插入指定位置position:BEFORE|AFTER,当 pivot 不存在于列表 key 时，不执行任何操作。当 key 不存在时， key 被视为空列表，不执行任何操作。
func (l *listRds) LInsert(key, position, pivot, value string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("linsert", key, position, pivot, value))
}

//返回列表头元素
func (l *listRds) LPop(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lpop", key))
}

//阻塞并弹出头元素
func (l *listRds) BLpop(key, timeout interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("blpop", key, timeout))
}

//返回列表尾元素
func (l *listRds) RPop(key string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("cpop", key))
}

//阻塞并弹出末尾元素
func (l *listRds) BRpop(key, timeout interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("brpop", key, timeout))
}

//返回指定位置的元素
func (l *listRds) LIndex(key string, index interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lindex", key, index))
}

//获取指定区间的元素
func (l *listRds) LRange(key string, start, stop interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lrange", key, start, stop))
}

//设置指定位元素
func (l *listRds) LSet(key string, index, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lset", key, index, value))
}

//弹出source尾元素并返回，将弹出元素插入destination列表的开头
func (l *listRds) RPoplpush(key, source, destination string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("rpoplpush ", key, source, destination))
}

//阻塞并弹出尾元素，将弹出元素插入另一列表的开头
func (l *listRds) BRpoplpush(key, source, destination string, timeout interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("brpoplpush ", key, source, destination, timeout))
}

//移除元素,count = 0 : 移除表中所有与 value 相等的值,count!=0,移除与 value 相等的元素，数量为 count的绝对值
func (l *listRds) LRem(key string, count, value interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("lrem", key, count, value))
}

//列表裁剪，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。-1 表示尾部
func (l *listRds) LTrim(key string, start, stop interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("ltrim", key, start, stop))
}
