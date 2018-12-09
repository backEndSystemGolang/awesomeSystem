package routers

import (
	"beego_workspace/awesomeSystem/rzcontrollers"
	"github.com/astaxie/beego"
)

func init() {
    //haha233
    beego.Router("/", &controllers.MainController{})
}
