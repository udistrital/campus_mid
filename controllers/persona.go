package controllers

import (
	"encoding/json"
	"fmt"

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
	c.Mapping("GuardarPersonai", c.GuardarPersonai)
}

// GuardarPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /GuardarPersona [post]
/*func (c *PersonaController) GuardarPersona() {
	// persona datos que entran a la funcion GuardarPersona
	var persona models.PersonaDatosBasicos

	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert

	//acumulado de alertas
	var alertas string

	//almacena resultado de la peticion POST /persona
	var resultado map[string]interface{}

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
		myId, _ := (resultado["Id"]).(float64)
		fmt.Println("acaaaaaaaaaaa", myId, resultado["Id"], resultado)
		if errPersona == nil && resultado["Id"] != 0 {

			alertas = alertas + " OK persona "
			alerta.Type = "OK"
			alerta.Code = "201"

			estadoCivil.Persona = &models.Persona{Id: myId}
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

			genero.Persona = &models.Persona{Id: myId}
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

}*/

// GuardarPersonai ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /GuardarPersona [post]
func (c *PersonaController) GuardarPersonai() {
	// persona datos que entran a la funcion GuardarPersona
	var persona map[string]interface{}
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}
	var resultado3 map[string]interface{}
	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert

	//acumulado de alertas
	var alertas string

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {

		//funcion que realiza  de la  peticion POST /persona
		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona", "POST", &resultado, persona)
		fmt.Println("acaaaaaaaaaaa", resultado["Id"], resultado)
		if errPersona == nil && resultado["Id"] != 0 {

			alertas = alertas + " OK persona "
			alerta.Type = "OK"
			alerta.Code = "201"

			var estadoCivil map[string]interface{}
			estadoCivil = make(map[string]interface{})
			estadoCivil["Persona"] = resultado
			fmt.Println("estado", estadoCivil)
			estadoCivil["EstadoCivil"] = persona["EstadoCivil"]

			//funcion que realiza  de la  peticion POST /persona_estado_civil
			errEstadoCivil := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil", "POST", &resultado2, estadoCivil)
			if errEstadoCivil != nil || resultado2["Id"] == 0 {
				alertas = alertas + " ERROR persona_estado_civil "
				alerta.Type = "error"
				alerta.Code = "400"
			} else {
				alertas = alertas + " OK persona_estado_civil "
			}

			var genero map[string]interface{}
			genero = make(map[string]interface{})
			genero["Persona"] = resultado
			genero["Genero"] = persona["Genero"]

			//funcion que realiza  de la  peticion POST /persona_genero
			errGenero := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero", "POST", &resultado3, genero)
			if errGenero != nil || resultado3["Id"] == 0 {
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
