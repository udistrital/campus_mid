package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"],
		beego.ControllerComments{
			Method:           "PostDescuentoAcademico",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"],
		beego.ControllerComments{
			Method:           "GetDescuentoAcademico",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"],
		beego.ControllerComments{
			Method:           "PutDescuentoAcademico",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"],
		beego.ControllerComments{
			Method:           "DeleteDescuentoAcademico",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"],
		beego.ControllerComments{
			Method:           "GetDescuentoAcademicoByPersona",
			Router:           `/:persona_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"],
		beego.ControllerComments{
			Method:           "GetDescuentoByDependenciaPeriodo",
			Router:           `/descuentodependenciaperiodo/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:DescuentoController"],
		beego.ControllerComments{
			Method:           "GetDescuentoByPersonaPeriodoDependencia",
			Router:           `/descuentopersonaperiododependencia/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:EvaluacionInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:EvaluacionInscripcionController"],
		beego.ControllerComments{
			Method:           "GetEvaluacionInscripcionByIdInscripcion",
			Router:           `/:inscripcion_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method:           "PostExperienciaLaboral",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method:           "GetExperienciaLaboralByEnte",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method:           "PutExperienciaLaboral",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method:           "GetExperienciaLaboral",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ExperienciaLaboralController"],
		beego.ControllerComments{
			Method:           "DeleteExperienciaLaboral",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method:           "PostFormacionAcademica",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method:           "GetFormacionAcademicaByEnte",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method:           "PutFormacionAcademica",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method:           "GetFormacionAcademica",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:FormacionController"],
		beego.ControllerComments{
			Method:           "DeleteFormacionAcademica",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:InscripcionController"],
		beego.ControllerComments{
			Method:           "GetInscripcionByPeriodoPrograma",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:InscripcionController"],
		beego.ControllerComments{
			Method:           "PutInscripcion",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:InscripcionController"],
		beego.ControllerComments{
			Method:           "GetInscripcion",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:InscripcionController"],
		beego.ControllerComments{
			Method:           "GetByIdentificacion",
			Router:           `/identificacion/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method:           "GetByEnte",
			Router:           `/:ente`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method:           "GetByIdentificacion",
			Router:           `/identificacion/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method:           "RegistrarUbicacion",
			Router:           `/registar_ubicacion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "ActualizarDatosComplementarios",
			Router:           `/actualizar_complementarios`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "ActualizarDatosContacto",
			Router:           `/actualizar_contacto`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "ActualizarPersona",
			Router:           `/actualizar_persona`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "ConsultarDatosComplementarios",
			Router:           `/consultar_complementarios/:ente_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "ConsultarDatosContacto",
			Router:           `/consultar_contacto/:ente_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "ConsultarPersonaByUser",
			Router:           `/consultar_persona/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "ConsultarPersona",
			Router:           `/consultar_persona/:ente_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "GuardarDatosComplementarios",
			Router:           `/guardar_complementarios`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "GuardarDatosContacto",
			Router:           `/guardar_contacto`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:PersonaController"],
		beego.ControllerComments{
			Method:           "GuardarPersona",
			Router:           `/guardar_persona`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "PostProduccionAcademica",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "PutProduccionAcademica",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "DeleteProduccionAcademica",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetProduccionAcademica",
			Router:           `/:persona`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/campus_mid/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "PutEstadoAutorProduccionAcademica",
			Router:           `/estado_autor_produccion/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
