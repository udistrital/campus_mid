package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// EvaluacionInscripcionController ...
type EvaluacionInscripcionController struct {
	beego.Controller
}

// URLMapping ...
func (c *EvaluacionInscripcionController) URLMapping() {
	c.Mapping("GetEvaluacionInscripcionByIdInscripcion", c.GetEvaluacionInscripcionByIdInscripcion)
}

// GetEvaluacionInscripcionByIdInscripcion ...
// @Title GetEvaluacionInscripcion
// @Description consultar evaluacion por inscripcion_id
// @Param	inscripcion_id		path 	int	true		"The key for staticblock"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:inscripcion_id [get]
func (c *EvaluacionInscripcionController) GetEvaluacionInscripcionByIdInscripcion() {
	//Id de la inscripcion a consultar
	idStr := c.Ctx.Input.Param(":inscripcion_id")
	fmt.Println("Id de inscripcion: ", idStr)
	//resultado evaluacion
	var resultado []map[string]interface{}
	var resultadoEntrevista []map[string]interface{}
	var entrevista map[string]interface{}
	var resultado2 map[string]interface{}
	var resultado3 map[string]interface{}

	// Obtener notas por InscripcionId
	errEvaluacionInscripcion := request.GetJson("http://"+beego.AppConfig.String("EvaluacionInscripcionService")+"/evaluacion_inscripcion/?query=InscripcionId:"+idStr, &resultado)
	fmt.Println(resultado)
	fmt.Println(errEvaluacionInscripcion)
	if errEvaluacionInscripcion == nil && fmt.Sprintf("%v", resultado[0]) != "map[]" {
		fmt.Println("Entró")
		if resultado[0]["Status"] != 404 {
			var notaAcumulada = 0.00
			var notaEntrevista = 0.00
			var idEntrevista = ""
			var idEvaluacion = ""
			var entrevistadores = 0
			for u := 0; u < len(resultado); u++ {
				// fmt.Println("Nota: ", resultado[u]["NotaFinal"])
				var requisitoProgramaAcademico map[string]interface{}
				requisitoProgramaAcademico = resultado[u]["RequisitoProgramaAcademicoId"].(map[string]interface{})
				// Calcular nota si el criterio tiene entrevista
				if resultado[u]["EntrevistaId"] != nil {
					entrevista = resultado[u]["EntrevistaId"].(map[string]interface{})
					idEntrevista = fmt.Sprintf("%.f", entrevista["Id"].(float64))
					idEvaluacion = fmt.Sprintf("%.f", resultado[u]["Id"]) //new
					// fmt.Println("Id de la entrevista", entrevista["Id"])
					errEntrevistadorEntrevista := request.GetJson("http://"+beego.AppConfig.String("EvaluacionInscripcionService")+"/entrevistador_entrevista/?query=EntrevistaId:"+idEntrevista+"&EstadoEntrevistaId:3", &resultadoEntrevista)
					if errEntrevistadorEntrevista == nil && fmt.Sprintf("%v", resultadoEntrevista[0]["System"]) != "map[]" {
						for u := 0; u < len(resultadoEntrevista); u++ {
							notaEntrevista += resultadoEntrevista[u]["NotaParcial"].(float64)
						}
						entrevistadores = len(resultadoEntrevista)
						notaEntrevista = notaEntrevista / float64(entrevistadores)
						resultado[u]["NotaFinal"] = notaEntrevista //New
						var evaluacionUpdate map[string]interface{}
						// New Actualiza el valor de la nota de la entrevista en el registro
						evaluacionUpdate = resultado[u]
						errSolicitud := request.SendJson("http://"+beego.AppConfig.String("EvaluacionInscripcionService")+"/evaluacion_inscripcion/"+idEvaluacion, "PUT", &resultado2, evaluacionUpdate)
						if errSolicitud == nil && fmt.Sprintf("%v", resultado2["System"]) != "map[]" {
							if resultado2["Status"] != 400 {
								c.Ctx.Output.SetStatus(200)
								c.Data["json"] = notaAcumulada
								// alerta.Body = alertas
							} else {
								logs.Error(resultado2)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errSolicitud
								c.Abort("400")
							}
						} else {
							logs.Error(resultado2)
							//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
							c.Data["System"] = errSolicitud
							c.Abort("400")
						}
						notaAcumulada += notaEntrevista * requisitoProgramaAcademico["Porcentaje"].(float64) / 100.00
					}
				} else {
					notaAcumulada += resultado[u]["NotaFinal"].(float64) * requisitoProgramaAcademico["Porcentaje"].(float64) / 100.00
				}
				resultado[u]["NotaFinal"] = notaEntrevista
				var resultadoInscripcion map[string]interface{}
				errInscripcion := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/inscripcion/"+idStr, &resultadoInscripcion)
				if errInscripcion == nil && fmt.Sprintf("%v", resultadoInscripcion["System"]) != "map[]" {
					if resultadoInscripcion["Status"] != 404 {
						resultadoInscripcion["PuntajeTotal"] = notaAcumulada
						errSolicitud := request.SendJson("http://"+beego.AppConfig.String("InscripcionService")+"/inscripcion/"+fmt.Sprintf("%.f", resultadoInscripcion["Id"].(float64)), "PUT", &resultado3, resultadoInscripcion)
						if errSolicitud == nil && fmt.Sprintf("%v", resultado3["System"]) != "map[]" {
							if resultado3["Status"] != 400 {
								c.Ctx.Output.SetStatus(200)
								c.Data["json"] = notaAcumulada
								// alerta.Body = alertas
								// fmt.Println("Actualizado CORRECTAMENTE", resultadoInscripcion)
							} else {
								logs.Error(resultado3)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errSolicitud
								c.Abort("400")
							}
						} else {
							logs.Error(resultado3)
							//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
							c.Data["System"] = errSolicitud
							c.Abort("400")
						}

					} else {
						logs.Error(errInscripcion)
						//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
						c.Data["json"] = nil
						c.Data["System"] = errInscripcion
						c.Abort("404")
					}
				} else {
					logs.Error(errInscripcion)
					//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
					c.Data["json"] = nil
					c.Data["System"] = errInscripcion
					c.Abort("404")
				}
				c.Data["json"] = notaAcumulada
			}
		} else {
			logs.Error(errEvaluacionInscripcion)
			//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
			c.Data["json"] = nil
			c.Data["System"] = errEvaluacionInscripcion
			c.Abort("404")
		}
	} else {
		logs.Error(errEvaluacionInscripcion)
		//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Data["json"] = nil
		c.Data["System"] = errEvaluacionInscripcion
		c.Abort("404")
	}
	//c.Data["json"] = alerta
	c.ServeJSON()
}
