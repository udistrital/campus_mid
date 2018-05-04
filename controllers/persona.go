package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/campus_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// PersonaController ...
type PersonaController struct {
	beego.Controller
}

// URLMapping ...
func (c *PersonaController) URLMapping() {
	c.Mapping("GuardarPersona", c.GuardarPersona)
}

// GuardarPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaCompleta	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /GuardarPersona [post]
func (c *PersonaController) GuardarPersona() {
	// persona datos que entran a la funcion GuardarPersona
	var persona models.PersonaCompleta

	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert

	//acumulado de alertas
	var alertas string

	//almacena resultado de la peticion POST /persona
	var resultado models.Persona

	//almacena resultado de la peticion POST /persona_estado_civil
	var resultado2 models.PersonaEstadoCivil

	//almacena resultado de la peticion POST /persona_genero
	var resultado3 models.PersonaGenero

	//modelo PersonaGenero
	var genero models.PersonaGenero

	//modelo  PersonaEstadoCivil
	var estadoCivil models.PersonaEstadoCivil

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {

		//funcion que realiza  de la  peticion POST /persona
		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona", "POST", &resultado, persona)

		if errPersona == nil && resultado.Id != 0 {

			alertas = alertas + " OK persona "
			alerta.Type = "OK"
			alerta.Code = "201"

			estadoCivil.Persona = &resultado
			estadoCivil.EstadoCivil = persona.EstadoCivil

			//funcion que realiza  de la  peticion POST /persona_estado_civil
			errEstadoCivil := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil", "POST", &resultado2, estadoCivil)
			if errEstadoCivil != nil || resultado2.Id == 0 {
				alertas = alertas + " ERROR persona_estado_civil "
				alerta.Type = "error"
				alerta.Code = "400"
			} else {
				alertas = alertas + " OK persona_estado_civil "
			}

			genero.Persona = &resultado
			genero.Genero = persona.Genero

			//funcion que realiza  de la  peticion POST /persona_genero
			errGenero := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero", "POST", &resultado3, genero)
			if errGenero != nil || resultado3.Id == 0 {
				alertas = alertas + " ERROR persona_genero "
				alerta.Type = "error"
				alerta.Code = "400"
			} else {
				alertas = alertas + " OK persona_genero "
			}

			alerta.Body = alertas
			c.Data["json"] = alerta
			c.ServeJSON()
		} else {
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = " ERROR persona "
			c.Data["json"] = alerta
			c.ServeJSON()
		}

	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = "ERROR formato incorrecto"
		c.Data["json"] = alerta
		c.ServeJSON()
	}

	c.ServeJSON()

}
