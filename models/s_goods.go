package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

//结构体名字、查询名、对应的表明必须一样(不含大小写)
//否则bego认为不存在相应的表名
type Goods struct {
	Id 						int
	Goods_name 				string
	Goods_classify_id 		int
	Goods_saled_nums 		int
	Goods_prices 			float64
	Goods_inventory			float64		//库存
	Goods_picture_address 	string
	//Cart					[]*Cart		`orm:"reverse(many)"`//一对多反向关系
}

func GetAllGoods() []*Goods {
	o := orm.NewOrm()
	o.Using("default")
	var goods []*Goods
	q:= o.QueryTable("goods")
	q.All(&goods)
	return goods
}

func GetGoodsById(id int) Goods{
	u:=Goods{Id:id}
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
func AddGoods(goods *Goods) int{
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(goods)
	return goods.Id
}
func UpdateGoods(goods *Goods) {
	o := orm.NewOrm()
	o.Using("default")
	o.Update(goods)
}

func DeleteGoods(id int){
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Goods{Id:id})
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Goods))
}