package controllers

import (
	"bytes"
	//"bufio"
	//"strings"
	"html/template"
	"fmt"
	"github.com/astaxie/beego"
	."contacts/models"
	"beego/validation"
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
		fmt.Errorf("%s\n",err.Error())
	}
	var buf bytes.Buffer
	contact:=NewContact()
	rawt.Execute(&buf,contact)
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
	contact:=Contact{}
	contact.SetId(id)
	err:=DelContactById(contact)
	if err!=nil{
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.Output.Status=200
	return
}

// "contacts/:id/update" 更新id对应对象
func (c *UpdateController) Update(){
	id:=c.Ctx.Input.Param(":id")
	contact:=Contact{}
	err:=c.ParseForm(&contact)
	valid:=validation.Validation{}
	b,err := valid.Valid(&contact)
	if err!=nil{
		fmt.Errorf("valid:%s\n",err.Error())
	}
	if !b{
		for _, err:= range valid.Errors{
			fmt.Printf("errkey:%s,errmessge:%s\n",err.Key,err.Message)
		}
		c.Ctx.Output.Status=400
		return
		
	}
	//fmt.Printf("id:%s\n",contact.GetId())
	//fmt.Printf("name:%s\n",contact.Name)
	if err!=nil{
			fmt.Errorf("parseerr:%s\n",err.Error())
			c.Ctx.Output.Status=400
			return
	}
	contact.SetId(id)
	err=UpsertContact(contact)
	if err!=nil{
		fmt.Errorf("update%s\n",err.Error())
		c.Ctx.Output.Status=400
		return
	}
	c.Ctx.Output.Status=200
	return
}

//"contacts/:id/preupdate" 准备更新目标
func (c *UpdateController)PreUpdate(){
	id:=c.Ctx.Input.Param(":id")
	rawt,err:=template.ParseFiles("/Developer/gopath/src/contacts/views/contactraw.tpl")
	if err!=nil{
		c.Ctx.WriteString(err.Error())
	}
	contact:=new(Contact)
	contact.SetId(id)
	contact,err =GetContactById(*contact)
	if err!=nil{
		c.Ctx.WriteString(err.Error())
		//fmt.Errorf("%v\n",err)
	}
	fmt.Printf("id:%s\n",contact.GetId())
	fmt.Printf("name:%s\n",contact.Name)
	var buf bytes.Buffer
	rawt.Execute(&buf,contact)
	c.Data["json"]=map[string]interface{}{
		"html":buf.String(),
	}
	//fmt.Printf("%v\n",buf.String())
	c.Ctx.Output.Status=200
	c.ServeJson()
}


// "contacts/getall" 获取所有
func (c *UpdateController) GetAll(){
	
}
//contacts/getalltable 获取表格
func (c *UpdateController)GetAllTable(){
	rawt,err:=template.ParseFiles("/Developer/gopath/src/contacts/views/contactstable.tpl")
	if err!=nil{
		c.Ctx.WriteString(err.Error())
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
	//fmt.Printf("%v\n",buf.String())
	c.Ctx.Output.Status=200
	c.ServeJson()
}