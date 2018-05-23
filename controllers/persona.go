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
	c.Mapping("GuardarPersona", c.GuardarPersona)
	c.Mapping("ActualizarPersona", c.ActualizarPersona)
	c.Mapping("ConsultaPersona", c.ConsultaPersona)
	c.Mapping("RegistrarUbicaciones", c.RegistrarUbicaciones)
}

// GuardarPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar Persona content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /GuardarPersona [post]
func (c *PersonaController) GuardarPersona() {
	// persona datos que entran a la funcion GuardarPersona
	var persona map[string]interface{}
	//reultado de la creacion de la persona
	var resultado map[string]interface{}
	// resultado de la adicion del estado civil
	var resultado2 map[string]interface{}
	// reultado de la adicion del genero
	var resultado3 map[string]interface{}
	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert
	//acumulado de alertas
	var alertas string

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {

		//funcion que realiza  de la  peticion POST /persona
		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona", "POST", &resultado, persona)

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
		alerta.Body = "ERROR formato incorrecto" + err.Error()
		c.Data["json"] = alerta
		c.ServeJSON()
	}

	c.ServeJSON()

}

// ActualizarPersona ...
// @Title ActualizarPersona
// @Description Actualizar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Actualizar Persona content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /ActualizarPersona [put]
func (c *PersonaController) ActualizarPersona() {
	// persona datos que entran a la funcion ActualizarPersona
	var persona map[string]interface{}
	var personaGenero []map[string]interface{}
	var personaEstadoCivil []map[string]interface{}
	//reultado de la creacion de la persona
	var resultado map[string]interface{}
	// resultado de la adicion del estado civil
	var resultado2 map[string]interface{}
	// reultado de la adicion del genero
	var resultado3 map[string]interface{}
	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert
	//acumulado de alertas
	var alertas string

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		//funcion que realiza  de la  peticion PUT /persona
		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/"+fmt.Sprintf("%.f", persona["Id"].(float64)), "PUT", &resultado, persona)
		if errPersona == nil && resultado["Type"] == "success" {

			alertas = alertas + " OK persona "
			alerta.Type = "OK"
			alerta.Code = "200"

			//obtener id persona_genero
			if err := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero/?query=persona:"+fmt.Sprintf("%.f", persona["Id"].(float64)), &personaGenero); err == nil {
				//actualizar genero
				personaGenero[0]["Genero"] = persona["Genero"]
				errGenero := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero/"+fmt.Sprintf("%.f", personaGenero[0]["Id"].(float64)), "PUT", &resultado2, personaGenero[0])
				if errGenero != nil {
					alertas = alertas + " ERROR persona_genero "
					alerta.Type = "error"
					alerta.Code = "400"
				} else {
					alertas = alertas + " OK persona_genero "
				}
			}

			//obtener id persona_estado_civil
			if err := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/?query=persona:"+fmt.Sprintf("%.f", persona["Id"].(float64)), &personaEstadoCivil); err == nil {
				//actualizar estado_civil
				personaEstadoCivil[0]["EstadoCivil"] = persona["EstadoCivil"]
				errEstadoCivil := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/"+fmt.Sprintf("%.f", personaEstadoCivil[0]["Id"].(float64)), "PUT", &resultado3, personaEstadoCivil[0])
				if errEstadoCivil != nil {
					alertas = alertas + " ERROR persona_estado_civil "
					alerta.Type = "error"
					alerta.Code = "400"
				} else {
					alertas = alertas + " OK persona_estado_civil "
				}
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
		alerta.Body = "ERROR formato incorrecto" + err.Error()
		c.Data["json"] = alerta
		c.ServeJSON()
	}

	c.ServeJSON()

}

// ConsultaPersona ...
// @Title Get One
// @Description get ConsultaPersona by userid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /ConsultaPersona/:id [get]
func (c *PersonaController) ConsultaPersona() {
	// alerta que retorna la funcion ConsultaPersona

	var alerta models.Alert
	idStr := c.Ctx.Input.Param(":id")
	var resultado map[string]interface{}

	errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/full/?userid="+idStr, &resultado)
	if errPersona == nil {
		nuevapersona := map[string]interface{}{
			"FechaNacimiento": resultado["Persona"].(map[string]interface{})["FechaNacimiento"],
			"Foto":            resultado["Persona"].(map[string]interface{})["Foto"],
			"PrimerApellido":  resultado["Persona"].(map[string]interface{})["PrimerApellido"],
			"PrimerNombre":    resultado["Persona"].(map[string]interface{})["PrimerNombre"],
			"SegundoApellido": resultado["Persona"].(map[string]interface{})["SegundoApellido"],
			"SegundoNombre":   resultado["Persona"].(map[string]interface{})["SegundoNombre"],
			"Usuario":         resultado["Persona"].(map[string]interface{})["Usuario"],
			"Id":              resultado["Persona"].(map[string]interface{})["Id"],
			"EstadoCivil":     resultado["EstadoCivil"],
			"Genero":          resultado["Genero"],
		}

		c.Data["json"] = nuevapersona
		c.ServeJSON()
	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = errPersona
		c.Data["json"] = alerta
		c.ServeJSON()
	}
}

// RegistrarUbicaciones ...
// @Title RegistrarUbicaciones
// @Description Registrar Ubicaciones
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Registrar Ubicaciones content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /RegistrarUbicaciones [post]
func (c *PersonaController) RegistrarUbicaciones() {
	// persona datos que entran a la funcion ActualizarPersona
	var ubicacionPersona map[string]interface{}
	var ubicacion map[string]interface{}
	var valorAtributoUbicacion map[string]interface{}

	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert
	//acumulado de alertas
	var alertas string

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ubicacionPersona); err == nil {

		ubicacion = make(map[string]interface{})
		ubicacion["Ente"] = map[string]interface{}{"Id": ubicacionPersona["Ente"]}
		ubicacion["Lugar"] = ubicacionPersona["Lugar"]
		ubicacion["TipoRelacionUbicacionEnte"] = map[string]interface{}{"Id": ubicacionPersona["TipoRelacionUbicacionEnte"]}
		ubicacion["Activo"] = true

		// resultado registro ubicacion_ente
		var resultado map[string]interface{}
		var resultado2 map[string]interface{}

		//funcion que realiza  de la  peticion POST /ubicacion_ente
		errUbicacionEnte := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/ubicacion_ente", "POST", &resultado, ubicacion)
		if errUbicacionEnte == nil && resultado["Type"] == "success" {
			alertas = alertas + " OK ubicacion_ente "
			//recorrer arreglo de atributos y registrarlos
			atributos := ubicacionPersona["Atributos"].([]interface{})
			if len(atributos) > 0 {
				for i := 0; i < len(atributos); i++ {

					atributo := atributos[i].(map[string]interface{})
					valorAtributoUbicacion = make(map[string]interface{})
					valorAtributoUbicacion["UbicacionEnte"] = resultado["Body"]
					valorAtributoUbicacion["AtributoUbicacion"] = map[string]interface{}{"Id": atributo["AtributoUbicacion"]}
					valorAtributoUbicacion["Valor"] = atributo["Valor"]

					//funcion que realiza  de la  peticion POST /ubicacion_ente
					errAtributoUbicacion := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/valor_atributo_ubicacion", "POST", &resultado2, valorAtributoUbicacion)
					fmt.Println(errAtributoUbicacion)
					fmt.Println(resultado2)
					if errAtributoUbicacion == nil && resultado2["Type"] == "success" {

						alertas = alertas + " OK atributo_ubicacion "
					} else {
						alertas = alertas + " ERROR atributo_ubicacion: " + resultado2["Body"].(string)
						alerta.Type = "error"
						alerta.Code = "400"
					}

				}
			}

			c.Data["json"] = alertas
			c.ServeJSON()

		} else {
			alertas = alertas + " ERROR ubicacion_ente: " + resultado["Body"].(string)
			alerta.Type = "error"
			alerta.Code = "400"
		}
		c.Data["json"] = alertas
		c.ServeJSON()

	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = "ERROR formato incorrecto " + err.Error()
		c.Data["json"] = alerta
		c.ServeJSON()
	}

	c.ServeJSON()

}
