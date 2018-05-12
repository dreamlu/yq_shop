package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["json"] = map[string]string{"Author": "lucheng","wechat": "a_862362681"}
	c.ServeJSON()//返回json

	//整个管理界面是固定的，只会变化中间的部分,layout
	/*c.Layout="index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["基本信息"] = "unit/基本信息.html"
	c.LayoutSections["home1"] = "unit/警务信息.html"
	c.TplName = "index.html"*/
}

func (c *MainController) Post() {

}