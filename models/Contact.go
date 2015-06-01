package models

import(
	"gopkg.in/mgo.v2/bson"
	)
	
type Contact struct{
	BaseDBmodel
	Id  bson.ObjectId `bson:"_id"`
	Name string 
	Remark string 
	Phone string 
}
func (this *Contact) Tablename() string {
	return "contacts"
}
func (this *Contact) init() {
	this.BaseDBmodel.init()
	this.c = this.db.C(this.Tablename())
}

func AddContact(contact Contact)error{
	contact.init()
	defer contact.session.Close()
	contact.Id=bson.NewObjectId()
	return contact.c.Insert(contact)
}

func DelContactById(contact Contact)error{
	contact.init()
	defer contact.session.Close()
	return contact.c.Remove(bson.M{"_id":contact.Id})
}

func UpdataContact(contact Contact)error{
	contact.init()
	defer contact.session.Close()
	return contact.c.Update(bson.M{"_id":contact.Id},bson.M{"$set":bson.M{
		"name":contact.Name,
		"remark":contact.Remark,
		"phone":contact.Phone,
	}})
}

func GetAllContact(contact Contact)(contacts []Contact,err error){
	contact.init()
	defer contact.session.Close()
	contact.c.Find(nil).All(&contacts)
	return
}