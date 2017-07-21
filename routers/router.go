// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/softfn/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/sys_grant",
			beego.NSInclude(
				&controllers.SysGrantController{},
			),
		),

		beego.NSNamespace("/sys_menu",
			beego.NSInclude(
				&controllers.SysMenuController{},
			),
		),

		beego.NSNamespace("/sys_org",
			beego.NSInclude(
				&controllers.SysOrgController{},
			),
		),

		beego.NSNamespace("/sys_org_staff",
			beego.NSInclude(
				&controllers.SysOrgStaffController{},
			),
		),

		beego.NSNamespace("/sys_role",
			beego.NSInclude(
				&controllers.SysRoleController{},
			),
		),

		beego.NSNamespace("/sys_role_user",
			beego.NSInclude(
				&controllers.SysRoleUserController{},
			),
		),

		beego.NSNamespace("/sys_staff",
			beego.NSInclude(
				&controllers.SysStaffController{},
			),
		),

		beego.NSNamespace("/sys_user",
			beego.NSInclude(
				&controllers.SysUserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
