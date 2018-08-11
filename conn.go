package redigo_pack

type redigoPack struct {
	String stringRds
	List   listRds
	Hash   hashRds
	Key    keyRds
	Set    setRds
	ZSet   zSetRds
	Bit    bitRds
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
