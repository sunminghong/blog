package main

import (
	_ "resourceManager/routers"
	"github.com/astaxie/beego"
)

func main() {
    beego.SetStaticPath("/sounds","resource/sounds")
    //beego.SetStaticPath("/sounds","resource/sounds")
	beego.Run()
}

