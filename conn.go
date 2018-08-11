package redigo_pack

type redigoPack struct {
	String StringRds
	List   ListRds
	Hash   hashRds
	Key    KeyRds
	Set    SetRds
	ZSet   ZSetRds
	Bit    BitRds
}

var RedigoConn = new(redigoPack)

func NewConnectionWithFile(tagName, path string) error {
	config, err := getConfigWithFile(tagName, path)
	if err != nil {
		return err
	}

	initPool(config)
	return nil
}

func NewConnection() {

}
