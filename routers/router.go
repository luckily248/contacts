package routers

import (
	"contacts/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/contacts/add",&controllers.UpdateController{},"post:Add")
	beego.Router("/contacts/:id/del",&controllers.UpdateController{},"post:Del")
	beego.Router("/contacts/:id/preupdate",&controllers.UpdateController{},"post:PreUpdate")
	beego.Router("/contacts/:id/update",&controllers.UpdateController{},"post:Update")
	beego.Router("/contacts/:id/cancel",&controllers.UpdateController{},"post:Cancel")
	beego.Router("/contacts/:id/get",&controllers.UpdateController{},"post:Get")
	beego.Router("/contacts/getall",&controllers.UpdateController{},"post:GetAll")
	beego.Router("/contacts/getalltable",&controllers.UpdateController{},"get:GetAllTable")
	beego.Router("/update",&controllers.UpdateController{},"get:Index")
	
}
