package redigo_pack

import "testing"

func TestStringRds_Get(t *testing.T) {
	v, err := RedigoConn.String.Get("namssss").Int64()
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}
