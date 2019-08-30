// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/iamMarkchu/goldfish/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/public",
		beego.NSInclude(
			&controllers.PublicController{},
		),
	)
	ns2 := beego.NewNamespace("/api",
		beego.NSBefore(
			auth,
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/policy",
			beego.NSInclude(
				&controllers.PolicyController{},
			),
		),
	)
	beego.AddNamespace(ns, ns2)
}
