// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/planesticud/campus_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),
		beego.NSNamespace("/formacion_academica",
			beego.NSInclude(
				&controllers.FormacionController{},
			),
		),
		beego.NSNamespace("/descuento_academico",
			beego.NSInclude(
				&controllers.DescuentoController{},
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
		beego.NSNamespace("/inscripcion",
			beego.NSInclude(
				&controllers.InscripcionController{},
			),
		),
		beego.NSNamespace("/produccion_academica",
			beego.NSInclude(
				&controllers.ProduccionAcademicaController{},
			),
		),
		beego.NSNamespace("/evaluacion_inscripcion",
			beego.NSInclude(
				&controllers.EvaluacionInscripcionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
