package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "ActualizarPersona",
			Router: `/ActualizarPersona`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GuardarPersona",
			Router: `/GuardarPersona`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
