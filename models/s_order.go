package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"yq_shop/util"
)

/*订单以及订单详情,返回时合并重复键值对*/
type Order struct {
	Id				int			//订单id
	User_id			int
	Saler_id		int
	Payment			float64		//付款价格
	Status			string		//交易状态,待付款/已付款/已发货/未发货
	Buyer_message	string		//买家留言
	//一个订单对应多个订单商品详情
	Orderitem		[]*Orderitem
}

type Orderitem struct {		//与商品一一对应
	//商品购买数量
	Goods_id		int
	Goods_nums		int
	//订单中商品详情
	Goods
}

//根据openid返回order
func GetOrderById(openid string) []*Order{
	o := orm.NewOrm()
	o.Using("default")
	userid := util.GetIdByOpenid(openid)

	//获得用户所有未付款订单
	var u []*Order
	_,err := o.Raw("select * from `order` where status=0 and user_id=?",userid).QueryRows(&u)
	if err == nil {
		fmt.Println("订单查询成功")
	}
	//查询用户的订单中的每件商品详情
	for i := 0; i < len(u);  i++{
		o.Raw("select * from `orderitem` a inner join `goods` b on a.goods_id=b.id where order_id=?",u[i].Id).QueryRows(&(u[i].Orderitem))
	}
	return u
}