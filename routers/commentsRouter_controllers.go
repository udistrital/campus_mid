package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "PostExperienciaLaboral",
			Router: `/ExperienciaLaboral`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "PutExperienciaLaboral",
			Router: `/ExperienciaLaboral/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "GetExperienciaLaboral",
			Router: `/ExperienciaLaboral/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method: "DeleteExperienciaLaboral",
			Router: `/ExperienciaLaboral/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method: "PostFormacionAcademica",
			Router: `/formacionacademica`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method: "PutFormacionAcademica",
			Router: `/formacionacademica/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method: "GetFormacionAcademica",
			Router: `/formacionacademica/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method: "DeleteFormacionAcademica",
			Router: `/formacionacademica/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

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
			Router: `/ConsultaPersona/`,
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
