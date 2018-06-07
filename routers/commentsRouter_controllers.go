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
			Method: "ConsultaPersona",
			Router: `/ConsultaPersona/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "DatosComplementariosPersona",
			Router: `/DatosComplementarios`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "ActualizarDatosComplementarios",
			Router: `/DatosComplementarios`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "ConsultaDatosComplementarios",
			Router: `/DatosComplementarios/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GuardarDatosContacto",
			Router: `/DatosContacto`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "ActualizarDatosContacto",
			Router: `/DatosContacto`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "DatosContacto",
			Router: `/DatosContacto/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GuardarPersona",
			Router: `/GuardarPersona`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "RegistrarUbicaciones",
			Router: `/RegistrarUbicaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
