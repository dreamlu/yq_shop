package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type User struct {
	Id    		int
	Openid  	string//OpenId不行,只能首字母大写,beego...
	Contact 	string
}

func GetAllUser() []*User {
	o := orm.NewOrm()
	o.Using("default")
	var user []*User
	q:= o.QueryTable("user")
	q.All(&user)
	return user
}

func GetUserById(id int) User{
	u:=User{Id:id}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	return u
}
func AddUser(user *User) int{
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(user)
	return user.Id
}
func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Using("default")
	o.Update(user)
}

func DeleteUser(id int){
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&User{Id:id})
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}
