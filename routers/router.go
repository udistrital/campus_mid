// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/campus_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),
		beego.NSNamespace("/formacion",
			beego.NSInclude(
				&controllers.FormacionController{},
			),
		),
		beego.NSNamespace("/experiencia_laboral",
			beego.NSInclude(
				&controllers.ExperienciaLaboralController{},
			),
		),
		beego.NSNamespace("/organizacion",
			beego.NSInclude(
				&controllers.OrganizacionController{},
			),
		),
		beego.NSNamespace("/idiomas",
			beego.NSInclude(
				&controllers.IdiomasController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
