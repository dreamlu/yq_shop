package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"yq_shop/models"
)

type GoodsController struct {
	beego.Controller
}
// @Title 获得所有商品
// @Description 返回所有的商品数据
// @Success 200 {object} models.Goods
// @router /goods [get]
func (u *GoodsController) GetAll() {
	ss := models.GetAllGoods()
	u.Data["json"] = ss
	u.ServeJSON()
}
// @Title 获得一个商品
// @Description 返回某商品数据
// @Success 200 {object} models.Goods
// @router /goods/:id [get]
func (u *GoodsController) GetById() {
	id ,_:= u.GetInt(":id")
	s := models.GetGoodsById(id)
	u.Data["json"] = s
	u.ServeJSON()
}
// @Title 创建商品
// @Description 创建商品的描述
// @Param      body          body   models.Goods true          "body for user content"
// @Success 200 {int} models.Student.Id
// @Failure 403 body is empty
// @router /goods [post]
func (u *GoodsController) Post() {
	var s models.Goods
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	uid := models.AddGoods(&s)
	u.Data["json"] = uid
	u.ServeJSON()
}
// @Title 修改用户
// @Description 修改用户的内容
// @Param      body          body   models.Goods true          "body for user content"
// @Success 200 {int} models.Student
// @Failure 403 body is empty
// @router /goods [put]
func (u *GoodsController) Update() {
	var s models.Goods
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateGoods(&s)
	u.Data["json"] = s
	u.ServeJSON()
}
// @Title 删除一个商品
// @Description 删除某商品数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Goods
// @router /goods/:id [delete]
func (u *GoodsController) Delete() {
	id ,_:= u.GetInt(":id")
	models.DeleteGoods(id)
	u.Data["json"] = true
	u.ServeJSON()
}
