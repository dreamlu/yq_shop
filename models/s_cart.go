package models

import (
	"github.com/astaxie/beego/orm"
	"yq_shop/util"
	"fmt"
	"strconv"
)

type Cart struct {
	Id    					int
	//Openid  	string
	Goods_id				int
	Goods_name 				string
	Goods_classify_id 		int
	Goods_prices 			float64
	Goods_nums				int
	Goods_picture_address	string
	//Goods 		*Goods	`orm:"rel(fk)"`//一对多,不可少
}
//根据openid返回cart
func GetCartById(openid string) []Cart{
	o := orm.NewOrm()
	o.Using("default")
	userid := util.GetIdByOpenid(openid)

	u:=[]Cart{}
	_,err := o.Raw("select a.id as id,b.id as goods_id,goods_name,goods_classify_id,goods_prices,goods_nums,goods_picture_address from `cart` a inner join `goods` b" +
		" on a.goods_id=b.id where user_id=?",userid).QueryRows(&u)
	if err == nil {
		fmt.Println("购物车查询成功")
	}

	return u
}
//cart加入数据
func AddCart(userid,goods_id,goods_nums int) {
	var maps []orm.Params
	o := orm.NewOrm()
	//判断购物车中是否已经拥有该商品
	num,err := o.Raw("select goods_nums from `cart` where user_id=? and goods_id=?",userid,goods_id).Values(&maps)
	if err == nil && num > 0 {
		//fmt.Println(maps[0]["user_name"]) // slene
		strGoodsNum := maps[0]["goods_nums"].(string)
		n,_ := strconv.Atoi(strGoodsNum)
		goods_nums += n

		o.Raw("update `cart` set goods_nums=? where user_id=? and goods_id=?",goods_nums,userid,goods_id).Exec()
	} else {
		o.Raw("insert `cart`(user_id,goods_id,goods_nums) value(?,?,?)",userid,goods_id,goods_nums).Exec()
	}
}
func UpdateCart(userid,goods_id,goods_nums int) {
	o := orm.NewOrm()
	_,err := o.Raw("update `cart` set goods_nums=? where user_id=? and goods_id=?",goods_nums,userid,goods_id).Exec()
	if(err != nil) {
		fmt.Println("购物车商品数量更新失败")
	}
}

func DeleteCart(userid,goods_id int){
	o := orm.NewOrm()
	_,err := o.Raw("delete from `cart` where user_id=? and goods_id=?",userid,goods_id).Exec()
	if(err != nil) {
		fmt.Println("购物车商品删除失败")
	}
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Cart))
}

