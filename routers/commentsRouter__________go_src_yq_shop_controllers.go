package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["yq_shop/controllers:CartController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:CartController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/cart`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:CartController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:CartController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/cart/:openid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:CartController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:CartController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/cartDelete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:CartController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:CartController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/cartUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/goods`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/goods`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/goods`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/goods/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/goods/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetOrderById",
            Router: `/orders/:openid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:UserController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/user`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:UserController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:UserController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/user`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:UserController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["yq_shop/controllers:UserController"] = append(beego.GlobalControllerRouter["yq_shop/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
