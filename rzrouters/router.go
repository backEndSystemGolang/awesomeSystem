package routers

import (
	"beego_workspace/awesomeSystem/rzcontrollers"
	"github.com/astaxie/beego"
)

func init() {
    //haha	
    beego.Router("/", &controllers.MainController{})
}
