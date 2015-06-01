package controllers

import (
	"github.com/astaxie/beego"
	."contacts/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	contact1:=&Contact{*new(BaseDBmodel),"","me","student","123444 234234234 2342342525 23429342"}
	contacts:=map[string]*Contact{"1":contact1}
	c.Data["contacts"]=contacts
	c.TplNames = "index.tpl"
}
