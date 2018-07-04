package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/campus_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// FormacionController ...
type FormacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *FormacionController) URLMapping() {
	c.Mapping("PostFormacionAcademica", c.PostFormacionAcademica)
	c.Mapping("PutFormacionAcademica", c.PutFormacionAcademica)
	c.Mapping("GetFormacionAcademica", c.GetFormacionAcademica)
	c.Mapping("DeleteFormacionAcademica", c.DeleteFormacionAcademica)
}

// PostFormacionAcademica ...
// @Title PostFormacionAcademica
// @Description Agregar Formacion Academica
// @Param	body		body 	{}	true		"body Agregar Formacion Academica content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /formacionacademica [post]
func (c *FormacionController) PostFormacionAcademica() {
	//formacion academica
	var formacion map[string]interface{}
	//alerta que retorna la funcion PostFormacionAcademica
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado formacion academica
	var resultado map[string]interface{}
	//resultado dato adicional formacion academica
	var resultado2 map[string]interface{}
	//resultado dato adicional formacion academica
	var resultado3 map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &formacion); err == nil {
		formacionacademica := map[string]interface{}{
			"Persona":           formacion["Persona"].(map[string]interface{})["Id"],
			"Titulacion":        formacion["ProgramaAcademico"].(map[string]interface{})["Id"],
			"FechaInicio":       formacion["FechaInicio"],
			"FechaFinalizacion": formacion["FechaFinalizacion"],
		}
		errFormacion := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica", "POST", &resultado, formacionacademica)
		//fmt.Println("el resultado es: ", resultado)
		if errFormacion == nil && resultado["Type"] != "error" {
			alertas = append(alertas, "se agrego la formacion correctamente")
			formaciondatoadicional := map[string]interface{}{
				"Activo":             true,
				"FormacionAcademica": resultado["Body"],
				"TipoDatoAdicional":  1,
				"Valor":              formacion["TituloTrabajoGrado"],
			}

			errFormacionAdicional := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica", "POST", &resultado2, formaciondatoadicional)

			if errFormacionAdicional == nil && resultado2["Type"] != "error" {
				alerta.Type = "success"
				alerta.Code = "200"
				alertas = append(alertas, "se agrego el titulo del trabajo de grado correctamente")
			} else {
				alerta.Type = "error"
				alerta.Code = "400"
				alertas = append(alertas, errFormacionAdicional.Error())
			}
			formaciondatoadicional2 := map[string]interface{}{
				"Activo":             true,
				"FormacionAcademica": resultado["Body"],
				"TipoDatoAdicional":  2,
				"Valor":              formacion["DescripcionTrabajoGrado"],
			}

			errFormacionAdicional2 := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica", "POST", &resultado3, formaciondatoadicional2)

			if errFormacionAdicional2 == nil && resultado2["Type"] != "error" {
				alerta.Type = "success"
				alerta.Code = "200"
				alertas = append(alertas, "se agrego la descripcion del trabajo de grado correctamente")
			} else {
				alerta.Type = "error"
				alerta.Code = "400"
				alertas = append(alertas, errFormacionAdicional2.Error())
			}
		} else {
			alerta.Type = "error"
			alerta.Code = "400"
			if errFormacion != nil {
				alertas = append(alertas, errFormacion.Error())
			}
			if resultado["Type"] == "error" {
				alertas = append(alertas, resultado["Body"])
			}
		}

	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alertas = append(alertas, err.Error())
	}
	alerta.Body = alertas
	c.Data["json"] = alerta
	c.ServeJSON()
}

// PutFormacionAcademica ...
// @Title PutFormacionAcademica
// @Description Modificar Formacion Academica
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body Modificar Formacion Academica content"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /formacionacademica/:id [put]
func (c *FormacionController) PutFormacionAcademica() {
	idStr := c.Ctx.Input.Param(":id")
	//formacion academica
	var formacion map[string]interface{}
	//alerta que retorna la funcion PostFormacionAcademica
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado formacion academica
	var resultado []map[string]interface{}
	//resultado dato adicional formacion academica
	var resultado2 map[string]interface{}
	//resultado dato adicional formacion academica
	//var resultado3 map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &formacion); err == nil {
		errFormacion := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica/?query=Persona:"+idStr, &resultado)
		if errFormacion == nil {
			for i := 0; i < len(resultado); i++ {

				if resultado[i]["Id"] == formacion["Id"] {
					formacionacademica := map[string]interface{}{
						"Persona":           formacion["Persona"].(map[string]interface{})["Id"],
						"Titulacion":        formacion["ProgramaAcademico"].(map[string]interface{})["Id"],
						"FechaInicio":       formacion["FechaInicio"],
						"FechaFinalizacion": formacion["FechaFinalizacion"],
					}
					errFormacion2 := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica/"+fmt.Sprintf("%.f", resultado[i]["Id"].(float64)), "PUT", &resultado2, formacionacademica)
					if errFormacion2 == nil {
						if resultado2["Type"] == "success" {
							alertas = append(alertas, "actualizada la formacion academica")
							alerta.Code = "200"
							alerta.Type = "success"
							alerta.Body = alertas
							c.Data["json"] = alerta
						}
					} else {
						fmt.Println("error de formacion", errFormacion2.Error())
						alertas = append(alertas, errFormacion2.Error())
						alerta.Code = "400"
						alerta.Type = "error"
						alerta.Body = alertas
						c.Data["json"] = alerta
					}

				}
			}

		} else {
			alertas = append(alertas, errFormacion.Error())
			alerta.Code = "400"
			alerta.Type = "error"
			alerta.Body = alertas
			c.Data["json"] = alerta
		}

	} else {
		alertas = append(alertas, err.Error())
		alerta.Code = "400"
		alerta.Type = "error"
		alerta.Body = alertas
		c.Data["json"] = alerta
	}
	c.ServeJSON()
}

// GetFormacionAcademica ...
// @Title GetFormacionAcademica
// @Description consultar Fromacion Academica por userid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /formacionacademica/:id [get]
func (c *FormacionController) GetFormacionAcademica() {
	//Id de la persona
	idStr := c.Ctx.Input.Param(":id")
	//formacion academica
	//var formacion map[string]interface{}
	//alerta que retorna la funcion PostFormacionAcademica
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado formacion academica
	var resultado []map[string]interface{}
	//resultado dato adicional formacion academica
	var resultado2 []map[string]interface{}

	var resultado3 []map[string]interface{}
	//resultado dato adicional formacion academica
	//var resultado3 map[string]interface{}
	errFormacion := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica/?query=Persona:"+idStr, &resultado)

	//fmt.Println("el resultado de la consulta es: ", resultado)
	if errFormacion == nil && resultado != nil {
		for u := 0; u < len(resultado); u++ {

			errTitulacion := request.GetJson("http://"+beego.AppConfig.String("ProgramaAcademicoService")+"/programa_academico/?query=Id:"+fmt.Sprintf("%.f", resultado[u]["Titulacion"].(float64)), &resultado2)
			if errTitulacion == nil {

				resultado[u]["Titulacion"] = resultado2[0]
				//fmt.Println("la titulacion de esa persona es: ", resultado)

			} else {
				fmt.Println("el error de la titulacion es: ", errTitulacion.Error())
			}

			errFormacionAdicional := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica/?query=FormacionAcademica:"+fmt.Sprintf("%.f", resultado[u]["Id"].(float64))+"&fields=TipoDatoAdicional,Valor,Id", &resultado3)

			if errFormacionAdicional == nil {
				fmt.Println("los datos adicionales de la formacion son: ", resultado3)
				for i := 0; i < len(resultado3); i++ {

					if resultado3[i]["TipoDatoAdicional"].(float64) == 1 {

						resultado[u]["TituloTrabajoGrado"] = resultado3[i]["Valor"]
					}
					if resultado3[i]["TipoDatoAdicional"].(float64) == 2 {
						resultado[u]["DescripcionTrabajoGrado"] = resultado3[i]["Valor"]
					}
				}

				c.Data["json"] = resultado
			} else {
				//fmt.Println("el error de adicional formacion academica es: ", errFormacionAdicional.Error())
			}
		}
	} else {
		alertas = append(alertas, errFormacion.Error())
		alerta.Code = "400"
		alerta.Type = "error"
		alerta.Body = alertas
		c.Data["json"] = alerta
	}

	c.ServeJSON()
}

// DeleteFormacionAcademica ...
// @Title DeleteFormacionAcademica
// @Description elimonar Fromacion Academica por userid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /formacionacademica/:id [delete]
func (c *FormacionController) DeleteFormacionAcademica() {
}
