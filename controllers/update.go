package controllers

import (
	"bytes"
	//"bufio"
	//"strings"
	"html/template"
	"fmt"
	"github.com/astaxie/beego"
	."contacts/models"
)

type UpdateController struct {
	beego.Controller
}
// "/update" get默认页面
func (c *UpdateController) Index() {
	contact:=new(Contact)
	contacts,err:=GetAllContact(*contact)
	if err!=nil{
		c.Ctx.WriteString(err.Error())
	}
	fmt.Printf("%v\n",contacts)
	c.Data["contacts"]=contacts
	c.TplNames = "manager.tpl"
}

// "/contacts/add" 添加空行
func (c *UpdateController) Add(){
	rawt,err:=template.ParseFiles("/Developer/gopath/src/contacts/views/contactraw.tpl")
	if err!=nil{
		fmt.Printf("%v\n",err)
	}
	var buf bytes.Buffer
	rawt.Execute(&buf,nil)
	c.Data["contact"]=&Contact{}
	c.Data["json"]=map[string]interface{}{
		"html":buf.String(),
	}
	//fmt.Printf("%v\n",buf.String())
	c.Ctx.Output.Status=200
	c.ServeJson()
}

// "contacts/:id/del" 删除id对应对象
func (c *UpdateController) Del(){
	id:=c.Ctx.Input.Param(":id")
	fmt.Printf("%s\n",id)
	c.Ctx.Output.Status=200
	return
}

// "contacts/:id/update" 更新id对应对象
func (c *UpdateController) Update(){
	id:=c.Ctx.Input.Param(":id")
	fmt.Printf("%s\n",id)
	c.Ctx.Output.Status=200
	return
}

//"contacts/:id/preupdate" 准备更新目标
func (c *UpdateController)PreUpdate(){
	id:=c.Ctx.Input.Param(":id")
	rawt,err:=template.ParseFiles("/Developer/gopath/src/contacts/views/contactraw.tpl")
	if err!=nil{
		fmt.Printf("%v\n",err)
	}
	contact:=new(Contact)
	contact.SetId(id)
	fmt.Printf("%s\n",id)
	UpdataContact(*contact)
	var buf bytes.Buffer
	rawt.Execute(&buf,contact)
	c.Data["contact"]=&Contact{}
	c.Data["json"]=map[string]interface{}{
		"html":buf.String(),
	}
	fmt.Printf("%v\n",buf.String())
	c.Ctx.Output.Status=200
	c.ServeJson()
}

//"contacts/:id/cancel"取消修改
func (c *UpdateController)Cancel(){
	
}

// "contacts/getall" 获取所有
func (c *UpdateController) GetAll(){
	
}
//contacts/getalltable 获取表格
func (c *UpdateController)GetAllTable(){
	rawt,err:=template.ParseFiles("/Developer/gopath/src/contacts/views/contactstable.tpl")
	if err!=nil{
		fmt.Printf("%v\n",err)
	}
	contact:=new(Contact)
	contacts,err:=GetAllContact(*contact)
	if err!=nil{
		c.Ctx.WriteString(err.Error())
	}
	var buf bytes.Buffer
	rawt.Execute(&buf,contacts)
	c.Data["json"]=map[string]interface{}{
		"html":buf.String(),
	}
	fmt.Printf("%v\n",buf.String())
	c.Ctx.Output.Status=200
	c.ServeJson()
}