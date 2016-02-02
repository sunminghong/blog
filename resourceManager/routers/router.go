package routers

import (
	"resourceManager/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/upload", &controllers.UploadController{})
}
