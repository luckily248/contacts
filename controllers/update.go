package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	."contacts/models"
)

type UpdateController struct {
	beego.Controller
}
func (c *UpdateController) Get() {
	contact:=new(Contact)
	contacts,err:=GetAllContact(*contact)
	if err!=nil{
		c.Ctx.WriteString(err.Error())
	}
	fmt.Printf("%v\n",contacts)
	c.Data["contacts"]=contacts
	c.TplNames = "manager.tpl"
}


func (c *UpdateController) Add(){
	id:=c.Ctx.Input.Param(":id")
	fmt.Println("%s\n",id)
	c.Ctx.WriteString(id)
}
func (c *UpdateController) Del(){
	id:=c.Ctx.Input.Param(":id")
	fmt.Println("%s\n",id)
}
func (c *UpdateController) Update(){
	id:=c.Ctx.Input.Param(":id")
	fmt.Println("%s\n",id)
}
func (c *UpdateController) GetAll(){
	
}