redis的使用更简单
---

### 说明
基于redigo的封装，写入好配置文件后，在任何地方随时调用，使用起来更简单，代码更优雅
支持和redigo同时使用

### 安装
`go get -u -v github.com/newgoo/redigo-pack`

### 使用
```go
package main

import (
	"fmt"

	"github.com/newgoo/redigo-pack"
)

func init() {
	redigo_pack.NewConnectionWithFile("redis", "./config/config.ini")
}

func main() {
	err := redigo_pack.RedigoConn.String.Set("1", 2, 10).Error()
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := redigo_pack.RedigoConn.String.Get("1").Int64()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
```


### License
The MIT License.

### 联系我
如果有不足的地方，需要添加的请联系我
newgoo: happs.lives@gmail.com

### 感谢
[vichuyongqiang](https://github.com/vichuyongqiang)提供的list的封装
