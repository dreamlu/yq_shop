package controllers

import (
	"github.com/astaxie/beego"
	"yq_shop/models"
	"strconv"
	"yq_shop/util"
	"fmt"
)

type CartController struct {
	beego.Controller
}

// @Title 获得一个购物车
// @Description 返回某购物车数据
// @Success 200 {object} models.Cart
// @router /cart/:openid [get]
func (u *CartController) GetById() {
	openid := u.GetString(":openid")
	s := models.GetCartById(openid)
	u.Data["json"] = s
	u.ServeJSON()
}
//添加商品到购物车
// @router /cart [post]
func (u *CartController) Post() {
	openid := u.GetString("openid")
	goods_id,_ := u.GetInt("goods_id")
	goods_nums,_:= u.GetInt("goods_nums")
	userid := util.GetIdByOpenid(openid)
	fmt.Println("cart数据:openid:"+openid+"数量:"+strconv.Itoa(goods_nums))

	models.AddCart(userid,goods_id,goods_nums)
}
// @Title 修改购物车
// @Description 修改购物车的内容
// @router /cartUpdate [post]
func (u *CartController) Update() {
	openid := u.GetString("openid")
	goods_id,_ := u.GetInt("goods_id")
	goods_nums,_:= u.GetInt("goods_nums")
	userid := util.GetIdByOpenid(openid)
	fmt.Println("cart数据:openid:"+openid+"数量:"+strconv.Itoa(goods_nums))

	models.UpdateCart(userid,goods_id,goods_nums)
}
// @Title 删除一个购物车商品
// @Description 删除某购物车数据
// @router /cartDelete [post]
func (u *CartController) Delete() {
	openid := u.GetString("openid")
	goods_id,_ := u.GetInt("goods_id")
	userid := util.GetIdByOpenid(openid)
	fmt.Println("删除cart数据:openid:"+openid+"goods_id:"+strconv.Itoa(goods_id))

	models.DeleteCart(userid,goods_id)

}
