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
	c.Mapping("DatosComplementariosPersona", c.DatosComplementariosPersona)
}

// GuardarPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
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
	//var alertas interface{}

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		alertas := append([]interface{}{"Cadena de respuestas: "})
		//funcion que realiza  de la  peticion POST /persona
		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona", "POST", &resultado, persona)
		//fmt.Println("este es el resultado",resultado["Type"])
		if resultado["Type"] != "error" || errPersona != nil {

			alertas = append(alertas, []interface{}{"Persona creada con Id "}, []interface{}{resultado["Id"]})
			alerta.Type = "OK"
			alerta.Code = "201"

			var estadoCivil map[string]interface{}
			estadoCivil = make(map[string]interface{})
			estadoCivil["Persona"] = resultado
			//fmt.Println("estado", estadoCivil)
			estadoCivil["EstadoCivil"] = persona["EstadoCivil"]

			//funcion que realiza  de la  peticion POST /persona_estado_civil
			errEstadoCivil := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil", "POST", &resultado2, estadoCivil)
			if resultado2["Type"] == "error" || errEstadoCivil != nil {
				alertas = append(alertas, resultado2)

				alerta.Type = "error"
				alerta.Code = "400"
			} else {
				alertas = append(alertas, []interface{}{" OK persona_estado_civil "})
			}

			var genero map[string]interface{}
			genero = make(map[string]interface{})
			genero["Persona"] = resultado
			genero["Genero"] = persona["Genero"]

			//funcion que realiza  de la  peticion POST /persona_genero
			errGenero := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero", "POST", &resultado3, genero)
			//fmt.Println("el genero", resultado3)
			if resultado3["Type"] == "error" || errGenero != nil {
				alertas = append(alertas, resultado3)
				alerta.Type = "error"
				alerta.Code = "400"
			} else {
				alertas = append(alertas, []interface{}{"OK persona_genero"})
			}

			alerta.Body = alertas
			c.Data["json"] = alerta
			c.ServeJSON()
		} else {
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = append(alertas, resultado)
			c.Data["json"] = alerta
			c.ServeJSON()
		}

	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = append([]interface{}{err.Error()})
		c.Data["json"] = alerta
		c.ServeJSON()
	}

	c.ServeJSON()

}

// ActualizarPersona ...
// @Title ActualizarPersona
// @Description Actualizar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Actualizar Persona content"
// @Success 200 {string} models.Persona.Id
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
	alertas := append([]interface{}{"Acumulado de respuestas"})

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		//funcion que realiza  de la  peticion PUT /persona
		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/"+fmt.Sprintf("%.f", persona["Id"].(float64)), "PUT", &resultado, persona)
		if errPersona == nil && resultado["Type"] == "success" {

			alertas = append(alertas, []interface{}{"Persona Actualizada"})
			alerta.Type = "OK"
			alerta.Code = "200"

			//obtener id persona_genero
			if err := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero/?query=persona:"+fmt.Sprintf("%.f", persona["Id"].(float64)), &personaGenero); err == nil {
				//actualizar genero
				personaGenero[0]["Genero"] = persona["Genero"]
				errGenero := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero/"+fmt.Sprintf("%.f", personaGenero[0]["Id"].(float64)), "PUT", &resultado2, personaGenero[0])
				//fmt.Println("el genero", resultado2)
				if errGenero != nil || resultado2["Type"] == "error" {
					alertas = append(alertas, resultado2)
					alerta.Type = "error"
					alerta.Code = "400"
				} else {
					alertas = append(alertas, []interface{}{"OK persona_genero"})
				}
			}

			//obtener id persona_estado_civil
			if err := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/?query=persona:"+fmt.Sprintf("%.f", persona["Id"].(float64)), &personaEstadoCivil); err == nil {
				//actualizar estado_civil
				personaEstadoCivil[0]["EstadoCivil"] = persona["EstadoCivil"]
				errEstadoCivil := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/"+fmt.Sprintf("%.f", personaEstadoCivil[0]["Id"].(float64)), "PUT", &resultado3, personaEstadoCivil[0])
				if errEstadoCivil != nil || resultado3["Type"] == "error" {
					alertas = append(alertas, resultado3)
					alerta.Type = "error"
					alerta.Code = "400"
				} else {
					alertas = append(alertas, []interface{}{"OK persona estado_civil"})

				}
			}
			alerta.Body = alertas
			c.Data["json"] = alerta
			c.ServeJSON()
		} else {
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = append(alertas, resultado)
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
// @Description get consultapersona by userid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {
// @Failure 403 :id is empty
// @router /consultapersona/:id [get]
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

// DatosComplementariosPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /DatosComplementariosPersona [post]
func (c *PersonaController) DatosComplementariosPersona() {
	// alerta que retorna la funcion ConsultaPersona
	var alerta models.Alert
	//Persona a la cual se van a agregar los datos complementarios
	var persona map[string]interface{}
	//Grupo etnico al que pertenece la persona
	var GrupoEtnico map[string]interface{}
	GrupoEtnico = make(map[string]interface{})
	//Discapacidades que tiene la persona
	var Discapacidad map[string]interface{}
	Discapacidad = make(map[string]interface{})
	//Grupo sanguineo de la persona
	var GrupoSanguineo map[string]interface{}
	GrupoSanguineo = make(map[string]interface{})
	//resultado de la consulta por ente
	var resultado []map[string]interface{}
	//Resultado de agregar grupo sanguineo y discapacidades
	var resultado2 map[string]interface{}
	//Resultado de agregar grupo sanguineo y discapacidades
	var resultado3 map[string]interface{}
	//acumulado de errores
	errores := append([]interface{}{"acumulado de alertas"})
	//comprobar que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {

		errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Ente:"+fmt.Sprintf("%.f", persona["Ente"].(float64))+"&fields=Id", &resultado)

		if errPersona == nil && resultado != nil {

			GrupoEtnico["GrupoEtnico"] = persona["GrupoEtnico"]
			GrupoEtnico["Persona"] = resultado[0]

			errGrupoEtnico := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico", "POST", &resultado2, GrupoEtnico)

			if errGrupoEtnico != nil || resultado2["Id"] == 0 || resultado2["Type"] == "error" {

				if errGrupoEtnico != nil {
					errores = append(errores, []interface{}{"error grupo etnico: ", errGrupoEtnico.Error()})
				}
				if resultado2["Type"] == "error" {

					errores = append(errores, resultado2)
				}

			} else {
				errores = append(errores, []interface{}{"OK persona_grupo_etnico"})
			}
			if (persona["GrupoSanguineo"] == "O" || persona["GrupoSanguineo"] == "A" || persona["GrupoSanguineo"] == "AB" || persona["GrupoSanguineo"] == "B") && (persona["Rh"] == "+" || persona["Rh"] == "-") {

				GrupoSanguineo["Persona"] = resultado[0]
				GrupoSanguineo["FactorRh"] = persona["Rh"]
				GrupoSanguineo["GrupoSanguineo"] = persona["GrupoSanguineo"]

				errGrupoSanguineo := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona", "POST", &resultado3, GrupoSanguineo)

				if errGrupoSanguineo == nil {
					errores = append(errores, []interface{}{"OK grupo_sanquineo_persona"})
				} else {
					errores = append(errores, []interface{}{"err grupo_sanquineo_persona", errGrupoSanguineo.Error()})
				}
			} else {

				errores = append(errores, []interface{}{"el grupo sanguineo es incorrecto:", persona["GrupoSanguineo"], persona["Rh"]})
			}

			discapacidad := persona["TipoDiscapacidad"].([]interface{})

			for i := 0; i < len(discapacidad); i++ {

				Discapacidad["Persona"] = resultado[0]
				Discapacidad["TipoDiscapacidad"] = discapacidad[i]

				errDiscapacidad := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad", "POST", &resultado2, Discapacidad)

				if errDiscapacidad != nil || resultado2["Type"] == "error" {

					errores = append(errores, []interface{}{"La discapacidad con Id ", discapacidad[i].(map[string]interface{})["Id"], " obtuvo ", resultado2["Type"], " ", resultado2["Body"]})
					alerta.Type = "error"
					alerta.Code = "400"
				} else {
					errores = append(errores, []interface{}{"La discapacidad con Id ", discapacidad[i].(map[string]interface{})["Id"], " obtuvo el resultado ", resultado2["Type"]})
					alerta.Type = "sucess"
					alerta.Code = "200"
				}
			}
			alerta.Body = errores
			c.Data["json"] = alerta
			c.ServeJSON()
		} else {

			if errPersona != nil {
				errores = append(errores, []interface{}{"error persona: ", errPersona})
			}
			if len(resultado) == 0 {
				errores = append(errores, []interface{}{"NO existe ninguna persona con este ente"})

			}
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = errores
			c.Data["json"] = alerta
			c.ServeJSON()
		}
	} else {

		errores = append(errores, []interface{}{err.Error()})
		c.Ctx.Output.SetStatus(200)
		alerta.Type = "error"
		alerta.Code = "401"
		alerta.Body = errores
		c.Data["json"] = alerta
		c.ServeJSON()
	}
}
