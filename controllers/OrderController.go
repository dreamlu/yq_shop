package controllers

import (
	"github.com/astaxie/beego"
	"yq_shop/models"
)

type OrderController struct {
	beego.Controller
}

// @router /orders/:openid [get]
func (u *GoodsController) GetOrderById() {
	openid := u.GetString(":openid")
	ss := models.GetOrderById(openid)
	u.Data["json"] = ss
	u.ServeJSON()
}


