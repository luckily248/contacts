package test

import (
	//"net/http"
	//"net/http/httptest"
	. "../models"
	_ "../routers"
	"github.com/astaxie/beego"
	"path/filepath"
	"runtime"
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestMain is a sample to run an endpoint test
//func TestMain(t *testing.T) {
//	r, _ := http.NewRequest("GET", "/", nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)

//	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

//	Convey("Subject: Test Station Endpoint\n", t, func() {
//	        Convey("Status Code Should Be 200", func() {
//	                So(w.Code, ShouldEqual, 200)
//	        })
//	        Convey("The Result Should Not Be Empty", func() {
//	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
//	        })
//	})
//}
func TestDB(t *testing.T) {
	contact1 := &Contact{*new(BaseDBmodel), "", "me", "student", "123444 234234234 2342342525 23429342"}
	t.Log("start add!")
	err := AddContact(*contact1)
	if err != nil {
		t.Error(err)
	}

	contacts, err := GetAllContact(*new(Contact))
	if len(contacts) != 0 {
		t.Logf("%v\n", contacts)
		t.Logf("%s\n", contacts[0].Id)
		contact1ex := new(Contact)
		contact1ex.SetId(contacts[0].GetId())
		err = GetContactById(*contact1ex)
		if err != nil {
			t.Error(err)
		}
		t.Logf("findbyid :%s\n", contact1ex.Name)
		contact2 := &Contact{*new(BaseDBmodel), contacts[0].Id, "new", "pig", "12312312 1231231231 123123"}
		t.Log("start updata!")
		err2 := UpdataContact(*contact2)
		if err2 != nil {
			t.Error(err2)
		}
	}
	contacts, err = GetAllContact(*new(Contact))
	if err != nil {
		t.Error(err)
	}
	if len(contacts) != 0 {
		t.Logf("%v\n", contacts)
		contactdel := &Contact{*new(BaseDBmodel), contacts[0].Id, "", "", ""}
		t.Log("start del!")
		err3 := DelContactById(*contactdel)
		if err3 != nil {
			t.Error(err3)
		}
	}
	contacts, err = GetAllContact(*new(Contact))
	if len(contacts) == 0 {
		t.Logf("%s\n", "len=0")
	} else {
		t.Logf("%v\n", contacts)
	}

}
