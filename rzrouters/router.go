package rzrouters

import (
	"awesomeSystem/rzcontrollers"
	"github.com/astaxie/beego"
)

func init() {
	//haha233
	beego.Router("/", &rzcontrollers.MainController{})
}
