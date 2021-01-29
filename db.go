package redigo_pack

type dbRds struct {
}

func (d *dbRds) SelectDb(db int64) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("select", db))
}

func (d *dbRds) Do(commend string, args ...interface{}) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do(commend, args...))
}