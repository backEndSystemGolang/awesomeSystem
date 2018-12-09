package routers

import (
	"beego_workspace/awesomeSystem/rzcontrollers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
