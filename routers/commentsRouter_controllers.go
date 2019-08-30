package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PolicyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PublicController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:PublicController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:TestController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/iamMarkchu/goldfish/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
