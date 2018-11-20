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
// @Description Agregar Formacion Academica ud
// @Param	body		body 	'hola'	true		"body Agregar Formacion Academica content"
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
	// resultado soporte fromacion
	var resultado4 map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &formacion); err == nil {
		formacionacademica := map[string]interface{}{
			"Persona":           formacion["Ente"].(map[string]interface{})["Id"],
			"Titulacion":        formacion["ProgramaAcademico"].(map[string]interface{})["Id"],
			"FechaInicio":       formacion["FechaInicio"],
			"FechaFinalizacion": formacion["FechaFinalizacion"],
		}
		errFormacion := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica", "POST", &resultado, formacionacademica)
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
			formacionsoporte := map[string]interface{}{
				"Documento":          formacion["Documento"],
				"FormacionAcademica": resultado["Body"],
			}
			fmt.Println("el soporte es:", formacionsoporte)
			errFormacionSoporte := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/soporte_formacion_academica", "POST", &resultado4, formacionsoporte)
			if errFormacionSoporte == nil && resultado4["Type"] != "error" {
				alerta.Type = "success"
				alerta.Code = "200"
				alertas = append(alertas, "se agrego el soporte correctamente")
			} else {
				alerta.Type = "error"
				alerta.Code = "400"
				alertas = append(alertas, errFormacionSoporte.Error())
			}
		} else {
			fmt.Println(resultado)
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
// @Param	id		path 	string	true		"el id de la formacion academica a modificar"
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
	var resultado3 []map[string]interface{}
	var resultado4 map[string]interface{}
	var resultado5 []map[string]interface{}
	var resultado6 map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &formacion); err == nil {
		errFormacion := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica/?query=Id:"+idStr, &resultado)
		if errFormacion == nil {
			for i := 0; i < len(resultado); i++ {

				if resultado[i]["Id"] == formacion["Id"] {
					formacionacademica := map[string]interface{}{
						"Persona":           formacion["Ente"].(map[string]interface{})["Id"],
						"Titulacion":        formacion["ProgramaAcademico"].(map[string]interface{})["Id"],
						"FechaInicio":       formacion["FechaInicio"],
						"FechaFinalizacion": formacion["FechaFinalizacion"],
					}
					errFormacion2 := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica/"+fmt.Sprintf("%.f", resultado[i]["Id"].(float64)), "PUT", &resultado2, formacionacademica)
					if errFormacion2 == nil {
						if resultado2["Type"] == "success" {
							alertas = append(alertas, "OK UPDATE formacion_academica")
							alerta.Code = "200"
							alerta.Type = "success"
							alerta.Body = alertas

						}
					} else {
						//fmt.Println("error de formacion", errFormacion2.Error())
						alertas = append(alertas, errFormacion2.Error())
						alerta.Code = "400"
						alerta.Type = "error"

					}

					errFormacionAdd := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica/?query=FormacionAcademica:"+fmt.Sprintf("%.f", resultado[i]["Id"].(float64)), &resultado3)
					fmt.Println("lo adicional es: ", resultado3)
					if errFormacionAdd == nil && resultado3 != nil {
						fmt.Println("lo adicional es: ", resultado3)
						for u := 0; u < len(resultado3); u++ {
							if resultado3[u]["TipoDatoAdicional"].(float64) == 1 {
								resultado3[u]["Valor"] = formacion["TituloTrabajoGrado"]
								//fmt.Println("el nuevo titulo es: ", resultado3[u])
								errFormacionadd2 := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica/"+fmt.Sprintf("%.f", resultado3[u]["Id"].(float64)), "PUT", &resultado4, resultado3[u])
								if errFormacionadd2 == nil {
									alertas = append(alertas, "OK UPDATE TituloTrabajoGrado ")
									//fmt.Println("el resultado de actualizar lo adicional es: ", resultado4)
								} else {
									//fmt.Println("el error de actualizar lo adicional es:", errFormacionadd2.Error())
									alertas = append(alertas, errFormacionadd2.Error())
									alerta.Code = "400"
									alerta.Type = "error"
								}
							}
							if resultado3[u]["TipoDatoAdicional"].(float64) == 2 {
								resultado3[u]["Valor"] = formacion["DescripcionTrabajoGrado"]
								errFormacionadd2 := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica/"+fmt.Sprintf("%.f", resultado3[u]["Id"].(float64)), "PUT", &resultado4, resultado3[u])
								if errFormacionadd2 == nil {
									alertas = append(alertas, "OK UPDATE DescripcionTrabajoGrado ")
									//fmt.Println("el resultado de actualizar lo adicional es: ", resultado4)
								} else {
									//fmt.Println("el error de actualizar lo adicional es:", errFormacionadd2.Error())
									alertas = append(alertas, errFormacionadd2.Error())
									alerta.Code = "400"
									alerta.Type = "error"
								}
							}
						}
						errSoporteFormacion := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/soporte_formacion_academica/?query=FormacionAcademica:"+fmt.Sprintf("%.f", resultado[i]["Id"].(float64)), &resultado5)
						if errSoporteFormacion == nil && resultado5 != nil {
							fmt.Println("el soporte actual es:", resultado5[0])
							resultado5[0]["Documento"] = formacion["Documento"]
							fmt.Println("el nuevo soporte es:", resultado5[0])
							errFormacionsoporte2 := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/soporte_formacion_academica/"+fmt.Sprintf("%.f", resultado5[0]["Id"].(float64)), "PUT", &resultado6, resultado5[0])
							if errFormacionsoporte2 == nil {
								alertas = append(alertas, "OK UPDATE Documento ")
							}
						}
					} else {
						if errFormacionAdd != nil {
							fmt.Println("error adicional: ", errFormacionAdd.Error())
							alertas = append(alertas, errFormacion.Error())

						} else {
							alertas = append(alertas, "ERROR formacion adicional no se encontro registro")

						}
						alerta.Code = "400"
						alerta.Type = "error"

					}

				}

			}

		} else {
			alertas = append(alertas, errFormacion.Error())
			alerta.Code = "400"
			alerta.Type = "error"

		}

	} else {
		alertas = append(alertas, err.Error())
		alerta.Code = "400"
		alerta.Type = "error"

	}
	alerta.Body = alertas
	c.Data["json"] = alerta
	c.ServeJSON()
}

// GetFormacionAcademica ...
// @Title GetFormacionAcademica
// @Description consultar Fromacion Academica por userid
// @Param	id		query 	string	true		"The key for staticblock"
// @Param	idformacion		query 	string	false		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /formacionacademica/ [get]
func (c *FormacionController) GetFormacionAcademica() {
	//Id de la persona
	var idStr = ""
	var idFor float64 = 0
	idStr = c.GetString("id")
	idFor, _ = c.GetFloat("idformacion")
	fmt.Println("el id for es: ", idFor)
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

	var resultado4 []map[string]interface{}
	var resultadof []map[string]interface{}
	//resultado dato adicional formacion academica
	//var resultado3 map[string]interface{}

	errFormacion := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica/?query=Persona:"+idStr+"&fields=Id,FechaInicio,FechaFinalizacion,Titulacion", &resultado)

	//fmt.Println("el resultado de la consulta es: ", resultado)
	if errFormacion == nil && resultado != nil {
		if resultado[0]["Type"] != "error" {

			resultado[0]["Ente"] = idStr
			for u := 0; u < len(resultado); u++ {

				errTitulacion := request.GetJson("http://"+beego.AppConfig.String("ProgramaAcademicoService")+"/programa_academico/?query=Id:"+fmt.Sprintf("%.f", resultado[u]["Titulacion"].(float64)), &resultado2)
				if errTitulacion == nil {

					//fmt.Println("la titulacion de esa persona es: ", resultado2[0]["Institucion"])

					errOrganizacion := request.GetJson("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion/?query=id:"+fmt.Sprintf("%.f", resultado2[0]["Institucion"].(float64)), &resultadof)
					if errOrganizacion == nil && resultadof != nil {
						//fmt.Println("la universidad es: ", resultadof[0])
						resultado[u]["Institucion"] = resultadof[0]
					} else {
						//fmt.Println(errOrganizacion)
					}
					resultado[u]["Titulacion"] = resultado2
				} else {
					alertas = append(alertas, errTitulacion.Error())
					alerta.Code = "400"
					alerta.Type = "error"
					alerta.Body = alertas
					c.Data["json"] = alerta
				}

				errFormacionAdicional := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica/?query=FormacionAcademica:"+fmt.Sprintf("%.f", resultado[u]["Id"].(float64))+"&fields=TipoDatoAdicional,Valor,Id", &resultado3)

				if errFormacionAdicional == nil {
					//fmt.Println("los datos adicionales de la formacion son: ", resultado3)
					for i := 0; i < len(resultado3); i++ {

						if resultado3[i]["TipoDatoAdicional"].(float64) == 1 {

							resultado[u]["TituloTrabajoGrado"] = resultado3[i]["Valor"]
						}
						if resultado3[i]["TipoDatoAdicional"].(float64) == 2 {
							resultado[u]["DescripcionTrabajoGrado"] = resultado3[i]["Valor"]
						}
					}

				} else {
					//fmt.Println("el error de adicional formacion academica es: ", errFormacionAdicional.Error())
				}

				errDatoAdicional := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/soporte_formacion_academica/?query=FormacionAcademica:"+fmt.Sprintf("%.f", resultado[u]["Id"].(float64))+"&fields=Documento", &resultado4)
				if errDatoAdicional == nil && resultado4 != nil {
					//fmt.Println("el resultado de los documentos es: ", resultado4)
					resultado[u]["Documento"] = resultado4[0]["Documento"]
				} else {
					if errDatoAdicional != nil {
						//fmt.Println("el error es: ", errDatoAdicional.Error())
					}

				}
			}
			if idFor != 0 {
				for u := 0; u < len(resultado); u++ {

					fmt.Println("el id de la formacion es ", resultado[u]["Id"], "y el de la consulta: ", idFor)
					if resultado[u]["Id"].(float64) == idFor {
						c.Data["json"] = resultado[u]
					} else {

					}

				}

			} else {
				c.Data["json"] = resultado
			}

		}
	} else {
		fmt.Println("entro al error")
		if errFormacion != nil {
			c.Data["json"] = nil
		} else {

			c.Data["json"] = nil
		}

	}

	c.ServeJSON()
}

// DeleteFormacionAcademica ...
// @Title DeleteFormacionAcademica
// @Description eliminar Formacion Academica por id de la formacion
// @Param	id		path 	string	true		"Id de la formacion academica"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /formacionacademica/:id [delete]
func (c *FormacionController) DeleteFormacionAcademica() {
	idStr := c.Ctx.Input.Param(":id")
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado formacion academica
	var resultado map[string]interface{}
	//resultado dato adiconal formacion academica
	var resultado2 []map[string]interface{}
	var resultado3 map[string]interface{}
	var resultado4 []map[string]interface{}
	var resultado5 map[string]interface{}
	var resultado6 map[string]interface{}
	//fmt.Println("el id de la formacion a borrar es: ", idStr)
	errFormacion := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica/"+idStr, &resultado)
	if errFormacion == nil {
		//fmt.Println("la formacion es: ", resultado)
		errFormacion := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica/?query=FormacionAcademica:"+idStr, &resultado2)
		if errFormacion == nil {
			//fmt.Println("el dato adicional es: ", resultado2)
			for i := 0; i < len(resultado2); i++ {
				//fmt.Println("el id del dato adicional a borrar es: ", resultado2[i]["Id"])
				err := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/dato_adicional_formacion_academica/"+fmt.Sprintf("%.f", resultado2[i]["Id"].(float64)), "DELETE", &resultado3, nil)

				if err == nil {
					//fmt.Println("el resultado  del delete es:", resultado3)
					alertas = append(alertas, "OK DELETE dato_adicional_formacion_academica")
				}

			}

			errFormacion := request.GetJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/soporte_formacion_academica/?query=FormacionAcademica:"+idStr, &resultado4)
			if errFormacion == nil {
				//fmt.Println("el documento a borrar es ", resultado4[0]["Id"])
				err := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/soporte_formacion_academica/"+fmt.Sprintf("%.f", resultado4[0]["Id"].(float64)), "DELETE", &resultado5, nil)
				if err == nil {
					//fmt.Println("el resultado  del delete es:", resultado5)
					alertas = append(alertas, "OK DELETE soporte_formacion_academica")
				}
			} else {
				alertas = append(alertas, errFormacion.Error())
			}
			err := request.SendJson("http://"+beego.AppConfig.String("FormacionAcademicaService")+"/formacion_academica/"+idStr, "DELETE", &resultado6, nil)
			if err == nil {
				//fmt.Println("el resultado del DELETE es: ", resultado6)
				alertas = append(alertas, "OK DELETE formacion_academica")
			}
			alerta.Body = alertas
			alerta.Code = "200"
			alerta.Type = "success"
			c.Data["json"] = alerta
		}
	} else {
		alertas = append(alertas, errFormacion.Error())
		alerta.Body = alertas
		alerta.Code = "400"
		alerta.Type = "error"
		c.Data["json"] = alerta

	}
	c.ServeJSON()
}
