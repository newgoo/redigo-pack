package redigo_pack

import "github.com/sirupsen/logrus"

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
		//return err
		logrus.Error("没有找到配置文件,默认文件./conf/app.conf")
		config, err = getConfigWithFile(tagName, "./conf/app.conf")
		if err != nil || config == nil {
			return err
		}
	}

	initPool(config)
	logrus.Info("redis 配置为: ", config)
	return nil
}

func NewConnection() {

}
