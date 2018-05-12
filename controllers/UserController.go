package controllers

import (
	"encoding/json"
	"yq_shop/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title 获得所有用户
// @Description 返回所有的用户数据
// @Success 200 {object} models.User
// @router /user [get,post]
func (u *UserController) GetAll() {
	ss := models.GetAllUser()
	u.Data["json"] = ss
	u.ServeJSON()
}
// @Title 获得一个用户
// @Description 返回某用户数据
// @Success 200 {object} models.User
// @router /user/:id [get]
func (u *UserController) GetById() {
	id ,_:= u.GetInt(":id")
	s := models.GetUserById(id)
	u.Data["json"] = s
	u.ServeJSON()
}
// @Title 创建用户
// @Description 创建用户的描述
// @Param      body          body   models.User true          "body for user content"
// @Success 200 {int} models.Student.Id
// @Failure 403 body is empty
// @router /user [post]
func (u *UserController) Post() {
	var s models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	uid := models.AddUser(&s)
	u.Data["json"] = uid
	u.ServeJSON()
}
// @Title 修改用户
// @Description 修改用户的内容
// @Param      body          body   models.User true          "body for user content"
// @Success 200 {int} models.Student
// @Failure 403 body is empty
// @router /user [put]
func (u *UserController) Update() {
	var s models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateUser(&s)
	u.Data["json"] = s
	u.ServeJSON()
}
// @Title 删除一个用户
// @Description 删除某用户数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.User
// @router /user/:id [delete]
func (u *UserController) Delete() {
	id ,_:= u.GetInt(":id")
	models.DeleteUser(id)
	u.Data["json"] = true
	u.ServeJSON()
}