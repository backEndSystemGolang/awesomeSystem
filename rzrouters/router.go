package routers

import (
	"github.com/astaxie/beego"
	"github.com/backEndSystemGolang/awesomeSystem/rzcontrollers"
)

func init() {
	//haha233
	beego.Router("/", &controllers.MainController{})
}
