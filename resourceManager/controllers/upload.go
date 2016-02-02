package upload

import (
    "fmt"
    "github.com/astaxie/beego"
    "io"
    "os"
)

type UploadController struct {
    beego.Controller
}


func (this *UploadController) Get() {
    this.TplNames = "upload.tpl"
}

func (this *UploadController) Post() {

    this.Ctx.Request.ParseMultipartForm(32 << 20)
    file, handler, err := this.GetFile("uploadfile")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    f, err := os.OpenFile("./static/files/"+handler.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    io.Copy(f, file)

    this.Ctx.Redirect(302, "/upload")

}
