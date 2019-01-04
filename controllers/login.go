package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

// 小程序相关信息
type User struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

var (
	appid = beego.AppConfig.String("appid")
	secret = beego.AppConfig.String("secret")
)

func (this *LoginController) Get() {

	code := this.GetString("code")
	fmt.Println("临时登录凭证：" + code)

	//获取openid接口
	te_uri := "https://api.weixin.qq.com/sns/jscode2session?appid="+appid+"&secret="+secret+"&js_code=" + code + "&grant_type=authorization_code"
	res, _ := http.Get(te_uri)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("用户openid等相关数据:" + string(body))
	defer res.Body.Close()

	//json数据的获取与返回
	var userInfo User
	json.Unmarshal(body, &userInfo) //反序列化

	//新用户注册(无论新老用户尝试一次openid注册)
	o := orm.NewOrm()
	o.Raw("insert `user`(openid) value(?)", userInfo.Openid).Exec()

	////反射遍历结构体，赋值给json
	////要求是可导出的字段大写
	//k := reflect.ValueOf(userInfo)
	//v := reflect.TypeOf(userInfo)
	////返回的json集合
	//jsonMap := make(map[string]string)
	//for i := 0; i < v.NumField(); i++{
	//	typeX := v.Field(i).Type;
	//	switch typeX.String() { //多选语句switch
	//		case "string":
	//			jsonMap[v.Field(i).Name] = k.Field(i).Interface().(string)
	//		//case "int":
	//		//	this.Data["json"] = map[string]int{v.Field(i).Name: k.Field(i).Interface().(int)}
	//	}
	//}
	this.Data["json"] = userInfo
	//只返回了最后一个json
	fmt.Println(this.Data["json"])
	this.ServeJSON()
}
