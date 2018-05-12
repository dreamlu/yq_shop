package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/goods`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/goods`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/goods`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "GetById",
			Router: `/goods:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/goods:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
