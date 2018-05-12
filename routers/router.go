package routers

import (
	"yq_shop/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//登录验证
	//beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)

    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{})
    //注解路由
    beego.Include(&controllers.GoodsController{})
    //beego.NSInclude()??区别呢
    beego.Include(&controllers.UserController{})
    beego.Include(&controllers.CartController{})
    beego.Include(&controllers.OrderController{})
}
//登录验证
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}