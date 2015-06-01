package routers

import (
	"contacts/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/update",&controllers.UpdateController{})
	beego.Router("/contacts/:id/add",&controllers.UpdateController{},"post:Add")
	beego.Router("/contacts/:id/del",&controllers.UpdateController{},"post:Del")
	beego.Router("/contacts/:id/update",&controllers.UpdateController{},"post:Update")
	beego.Router("/contacts/:id/get",&controllers.UpdateController{},"post:Get")
	beego.Router("/contacts/getAll",&controllers.UpdateController{},"post:GetAll")
}
