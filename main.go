package main

import (
	_ "testbeego/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	//beego.Run("127.0.0.1:8089")带参启动方式
}

