package controllers

import (
	. "../models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	contact1 := &Contact{*new(BaseDBmodel), "", "me", "student", "123444", "1231231231"}
	contacts := map[string]*Contact{"1": contact1}
	c.Data["contacts"] = contacts
	c.TplNames = "index.tpl"
}
