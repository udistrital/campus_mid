package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/planesticud/campus_mid/models"
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
	c.Mapping("ActualizarDatosComplementarios", c.ActualizarDatosComplementarios)
	c.Mapping("DatosContacto", c.DatosContacto)
	c.Mapping("consultadatoscomplementarios", c.ConsultaDatosComplementarios)
}

// GuardarPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	{}	true		"body for Guardar Persona content"
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
	var resultadoIdentificacion map[string]interface{}
	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert
	//acumulado de alertas
	//var alertas interface{}

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		alertas := append([]interface{}{"Response:"})

		//funcion que realiza  de la  peticion POST /persona
		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona", "POST", &resultado, persona)
		//fmt.Println("este es el resultado",resultado["Type"])
		if resultado["Type"] != "error" || errPersona != nil {

			alertas = append(alertas, "OK persona Id "+fmt.Sprintf("%.f", resultado["Body"].(map[string]interface{})["Id"]))
			alerta.Type = "success"
			alerta.Code = "201"

			var identificacion map[string]interface{}
			identificacion = make(map[string]interface{})
			identificacion["Ente"] = map[string]interface{}{"Id": resultado["Body"].(map[string]interface{})["Ente"]}
			identificacion["TipoIdentificacion"] = map[string]interface{}{"Id": persona["TipoIdentificacion"]}
			identificacion["NumeroIdentificacion"] = persona["NumeroDocumento"]
			identificacion["Soporte"] = persona["SoporteDocumento"]

			errIdentificacion := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion", "POST", &resultadoIdentificacion, identificacion)
			if resultadoIdentificacion["Type"] == "error" || errIdentificacion != nil {
				alertas = append(alertas, resultadoIdentificacion)
				alerta.Type = "error"
				alerta.Code = "400"
			} else {
				alertas = append(alertas, "OK identificacion")
			}

			var estadoCivil map[string]interface{}
			estadoCivil = make(map[string]interface{})
			estadoCivil["Persona"] = resultado["Body"]
			//fmt.Println("estado", estadoCivil)
			estadoCivil["EstadoCivil"] = persona["EstadoCivil"]

			//funcion que realiza  de la  peticion POST /persona_estado_civil
			errEstadoCivil := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil", "POST", &resultado2, estadoCivil)
			if resultado2["Type"] == "error" || errEstadoCivil != nil {
				alertas = append(alertas, resultado2)

				alerta.Type = "error"
				alerta.Code = "400"
			} else {
				alertas = append(alertas, "OK persona_estado_civil")
			}

			var genero map[string]interface{}
			genero = make(map[string]interface{})
			genero["Persona"] = resultado["Body"]
			genero["Genero"] = persona["Genero"]

			//funcion que realiza  de la  peticion POST /persona_genero
			errGenero := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero", "POST", &resultado3, genero)
			//fmt.Println("el genero", resultado3)
			if resultado3["Type"] == "error" || errGenero != nil {
				alertas = append(alertas, resultado3)
				alerta.Type = "error"
				alerta.Code = "400"
			} else {
				alertas = append(alertas, "OK persona_genero")
			}
			//var resultado5 map[string]interface{}
			//var admision map[string]interface{}
			//admision = make(map[string]interface{})
			//admision = map[string]interface{}{
			//	"Aspirante":          resultado["Body"].(map[string]interface{})["Ente"],
			//	"AceptaTerminos":     persona["Admision"].(map[string]interface{})["AceptaTerminos"],
			//	"EstadoAdmision":     map[string]interface{}{"Id": persona["Admision"].(map[string]interface{})["EstadoAdmision"]},
			//	"Enfasis":            map[string]interface{}{"Id": persona["Admision"].(map[string]interface{})["Enfasis"]},
			//	"LineaInvestigacion": map[string]interface{}{"Id": persona["Admision"].(map[string]interface{})["LineaInvestigacion"]},
			//	"Periodo":            persona["Admision"].(map[string]interface{})["Periodo"],
			//	"ProgramaAcademico":  persona["Admision"].(map[string]interface{})["ProgramaAcademico"],
			//}
			////fmt.Println("la admision es: ", admision)

			////funcion que realiza  de la  peticion POST /admisiones
			//errAdmision := request.SendJson("http://"+beego.AppConfig.String("AdmisionService")+"/admision", "POST", &resultado5, admision)
			//fmt.Println("la admision", resultado5)
			//if resultado5["Type"] == "error" || errAdmision != nil {
			//	alertas = append(alertas, resultado5)
			//	alerta.Type = "error"
			//	alerta.Code = "400"
			//} else {
			//	alertas = append(alertas, "OK admision")
			//}

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
	//resultado de la identifiacion
	var resultado4 []map[string]interface{}
	//nueva identifiacion
	var resultado5 map[string]interface{}
	var newidentificacion map[string]interface{}
	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert
	//acumulado de alertas
	alertas := append([]interface{}{"Response:"})

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		//funcion que realiza  de la  peticion PUT /persona
		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/"+fmt.Sprintf("%.f", persona["Id"].(float64)), "PUT", &resultado, persona)
		if errPersona == nil && resultado["Type"] == "success" {

			alertas = append(alertas, "UPDATE Persona")
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
					alertas = append(alertas, "UPDATE persona_genero")
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
					alertas = append(alertas, "UPDATE persona_estado_civil")

				}
			}

			errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+fmt.Sprintf("%.f", persona["Ente"].(float64)), &resultado4)
			if errIdentificacion == nil && resultado4 != nil {

				newidentificacion = map[string]interface{}{
					"TipoIdentificacion":   persona["TipoIdentificacion"],
					"NumeroIdentificacion": persona["NumeroDocumento"],
					"Soporte":              persona["SoporteDocumento"],
					"Ente":                 map[string]interface{}{"Id": persona["Ente"]},
					"Id":                   resultado4[0]["Id"],

					//
				}
				errPutIdentificacion := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/"+fmt.Sprintf("%.f", resultado4[0]["Id"].(float64)), "PUT", &resultado5, newidentificacion)
				if errPutIdentificacion != nil || resultado5["Type"] == "error" {
					alertas = append(alertas, resultado5)
					alerta.Type = "error"
					alerta.Code = "400"
				} else {
					alertas = append(alertas, "UPDATE ente_identifiacion")
				}
				fmt.Println("el la nueva identifiacion es:", newidentificacion)
			}
			alerta.Body = alertas
			c.Data["json"] = alerta
			c.ServeJSON()
		} else {
			if errPersona != nil {
				alerta.Body = append(alertas, resultado)
			} else {
				alerta.Body = append(alertas, resultado)
			}
			alerta.Type = "error"
			alerta.Code = "400"

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
// @Param	id	query	string	false	"busca por  id del ente"
// @Param	userid	query	string	false	"busca por nombre de usuario"
// @Success 200 {object} interface{}
// @Failure 403 :id is empty
// @router /ConsultaPersona/ [get]
func (c *PersonaController) ConsultaPersona() {
	// alerta que retorna la funcion ConsultaPersona

	//var alerta models.Alert
	//idStr := c.Ctx.Input.Param(":id")
	var resultado []map[string]interface{}
	var resultado2 []map[string]interface{}
	var EstadoCivil []map[string]interface{}
	var Genero []map[string]interface{}
	var GrupoEtnico []map[string]interface{}
	var Discapacidad []map[string]interface{}
	var GrupoSanguineo []map[string]interface{}
	var newpersona map[string]interface{}
	newpersona = make(map[string]interface{})
	//alertas := append([]interface{}{"acumulado de alertas"})

	var id = 0
	var uid = ""
	id, _ = c.GetInt("id")
	uid = c.GetString("userid")
	var errPersona error

	if id != 0 && uid == "" {
		id := c.GetString("id")
		errPersona = request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Ente:"+id, &resultado)
		if resultado != nil {
			newpersona = map[string]interface{}{
				"FechaNacimiento": resultado[0]["FechaNacimiento"],
				"Foto":            resultado[0]["Foto"],
				"PrimerApellido":  resultado[0]["PrimerApellido"],
				"PrimerNombre":    resultado[0]["PrimerNombre"],
				"SegundoApellido": resultado[0]["SegundoApellido"],
				"SegundoNombre":   resultado[0]["SegundoNombre"],
				"Usuario":         resultado[0]["Usuario"],
				"Id":              resultado[0]["Id"],
				"Ente":            resultado[0]["Ente"],
				//
			}
		}

	} else if id == 0 && uid != "" {
		errPersona = request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Usuario:"+uid, &resultado)

		if resultado != nil {
			newpersona = map[string]interface{}{
				"FechaNacimiento": resultado[0]["FechaNacimiento"],
				"Foto":            resultado[0]["Foto"],
				"PrimerApellido":  resultado[0]["PrimerApellido"],
				"PrimerNombre":    resultado[0]["PrimerNombre"],
				"SegundoApellido": resultado[0]["SegundoApellido"],
				"SegundoNombre":   resultado[0]["SegundoNombre"],
				"Usuario":         resultado[0]["Usuario"],
				"Id":              resultado[0]["Id"],
				"Ente":            resultado[0]["Ente"],
				//
			}
		}
	}
	if errPersona == nil && resultado != nil {

		errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+fmt.Sprintf("%.f", resultado[0]["Ente"].(float64))+"&fields=TipoIdentificacion,NumeroIdentificacion,FechaExpedicion,LugarExpedicion,Id,Soporte", &resultado2)
		if errIdentificacion == nil && resultado2 != nil {
			newpersona["TipoIdentificacion"] = resultado2[0]["TipoIdentificacion"]
			newpersona["NumeroDocumento"] = resultado2[0]["NumeroIdentificacion"]
			newpersona["SoporteDocumento"] = resultado2[0]["Soporte"]

		}
		errEstadoCivil := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/?query=Persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64))+"&fields=EstadoCivil,Id", &EstadoCivil)
		if errEstadoCivil == nil && EstadoCivil != nil {
			newpersona["EstadoCivil"] = EstadoCivil[0]["EstadoCivil"]
		}
		errGenero := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero/?query=Persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64))+"&fields=Genero,Id", &Genero)
		if errGenero == nil && Genero != nil {
			newpersona["Genero"] = Genero[0]["Genero"]
		}
		errGrupoEtnico := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/?query=Persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64))+"&fields=GrupoEtnico,Id", &GrupoEtnico)
		if errGrupoEtnico == nil && GrupoEtnico != nil {
			newpersona["GrupoEtnico"] = GrupoEtnico[0]["GrupoEtnico"]
		}
		errDiscapacidad := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/?query=Persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64))+"&fields=TipoDiscapacidad,Id", &Discapacidad)
		if errDiscapacidad == nil && Discapacidad != nil {
			newpersona["Discapacidad"] = Discapacidad
		}
		errGrupoSanguineo := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/?query=Persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64))+"&fields=GrupoSanguineo,Id,FactorRh", &GrupoSanguineo)
		if errGrupoSanguineo == nil && GrupoSanguineo != nil {
			newpersona["GrupoSanguineo"] = GrupoSanguineo[0]["GrupoSanguineo"]
			newpersona["Rh"] = GrupoSanguineo[0]["FactorRh"]
		}

	} else {
		newpersona = nil
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = newpersona

	c.ServeJSON()

}

// RegistroUbicaciones ...
func RegistroUbicaciones(ubicaciones map[string]interface{}) models.Alert {
	errores := append([]interface{}{"acumulado de alertas"})
	ubicacionesPersona := ubicaciones

	var ubicacionPersona map[string]interface{}
	var ubicacion map[string]interface{}
	var valorAtributoUbicacion map[string]interface{}

	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert

	ubicacionPersona = ubicacionesPersona

	lugar, err := ubicacionPersona["Lugar"].(map[string]interface{})
	if err == false {
		ubicacion = make(map[string]interface{})
		ubicacion["Ente"] = map[string]interface{}{"Id": ubicacionPersona["Ente"]}
		ubicacion["Lugar"] = lugar["Id"]
		ubicacion["TipoRelacionUbicacionEnte"] = map[string]interface{}{"Id": ubicacionPersona["TipoRelacionUbicacionEnte"]}
		ubicacion["Activo"] = true

		// resultado registro ubicacion_ente
		var resultado map[string]interface{}
		var resultado2 map[string]interface{}

		//funcion que realiza  de la  peticion POST /ubicacion_ente

		if errUbicacionEnte := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente", "POST", &resultado, ubicacion); errUbicacionEnte == nil {
			if resultado["Type"] == "success" {
				errores = append(errores, []interface{}{" OK ubicacion_ente "})
				//recorrer arreglo de atributos y registrarlos
				if ubicacionPersona["Atributos"] != nil {
					atributos := ubicacionPersona["Atributos"].([]interface{})
					if len(atributos) > 0 {
						for i := 0; i < len(atributos); i++ {
							atributo := atributos[i].(map[string]interface{})
							valorAtributoUbicacion = make(map[string]interface{})
							valorAtributoUbicacion["UbicacionEnte"] = resultado["Body"]
							valorAtributoUbicacion["AtributoUbicacion"] = map[string]interface{}{"Id": atributo["AtributoUbicacion"]}
							valorAtributoUbicacion["Valor"] = atributo["Valor"]

							errAtributoUbicacion := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion", "POST", &resultado2, valorAtributoUbicacion)
							if errAtributoUbicacion == nil && resultado2["Type"] == "success" {
								errores = append(errores, []interface{}{"OK atributo_ubicacion "})
								alerta.Type = resultado2["Type"].(string)
								alerta.Code = resultado2["Code"].(string)
							} else {
								errores = append(errores, " ERROR atributo_ubicacion: "+resultado2["Body"].(string))
								alerta.Type = resultado2["Type"].(string)
								alerta.Code = resultado2["Code"].(string)
							}

						}
					}
				}

			} else {
				errores = append(errores, " ERROR ubicacion_ente: "+resultado["Body"].(string))
				alerta.Type = "error"
				alerta.Code = "400"
			}
		} else {
			errores = append(errores, " ERROR ubicacion_ente: "+errUbicacionEnte.Error())
			alerta.Type = "error"
			alerta.Code = "400"
		}

	} else {
		errores = append(errores, " ERROR lugar")
		alerta.Type = "error"
		alerta.Code = "400"
	}

	alerta.Body = errores
	return alerta
}

// RegistrarUbicaciones ...
// @Title RegistrarUbicaciones
// @Description Registrar Ubicaciones
// @Param	body		body 	map[string]interface{}	true		"body for Registrar Ubicaciones content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /RegistrarUbicaciones [post]
func (c *PersonaController) RegistrarUbicaciones() {
	var datos map[string]interface{}
	var rta models.Alert
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err == nil {
		rta = RegistroUbicaciones(datos)
	}
	c.Data["json"] = rta
	c.ServeJSON()

}

// GuardarDatosContacto ...
// @Title PostDatosContacto
// @Description Guardar Datos de Contacto
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar datos contacto content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /DatosContacto [post]
func (c *PersonaController) GuardarDatosContacto() {
	alertas := append([]interface{}{"Response:"})
	// datos de contacto de la persona
	var datos map[string]interface{}
	var contactoEnte map[string]interface{}
	//reultado de la creacion de la persona
	var resultado map[string]interface{}
	//var resultado2 map[string]interface{}
	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err == nil {
		contactoEnte = make(map[string]interface{})
		contactos := datos["ContactoEnte"].([]interface{})

		for i := 0; i < len(contactos); i++ {
			contacto := contactos[i].(map[string]interface{})
			contactoEnte["Ente"] = map[string]interface{}{"Id": datos["Ente"]}
			contactoEnte["TipoContacto"] = map[string]interface{}{"Id": contacto["TipoContacto"]}
			contactoEnte["Valor"] = contacto["Valor"]

			errContacto := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente", "POST", &resultado, contactoEnte)

			if errContacto == nil && resultado["Type"] == "success" {
				alertas = append(alertas, "OK contacto_ente ")

			} else {
				alertas = append(alertas, resultado["Body"].(string))
			}
			alerta.Type = resultado["Type"].(string)
			alerta.Code = resultado["Code"].(string)
		}
		//guardar las ubicaciones
		var ubicacion map[string]interface{}
		ubicacion = make(map[string]interface{})

		UbicacionEnte := datos["UbicacionEnte"].(map[string]interface{})
		ubicacion["Ente"] = datos["Ente"]
		ubicacion["Lugar"] = UbicacionEnte["Lugar"]
		ubicacion["TipoRelacionUbicacionEnte"] = UbicacionEnte["TipoRelacionUbicacionEnte"]
		ubicacion["Atributos"] = UbicacionEnte["Atributos"]

		errUbicaciones := RegistroUbicaciones(ubicacion)
		if errUbicaciones.Type != "success" {
			alertas = append(alertas, errUbicaciones)
			alerta.Code = errUbicaciones.Code
		} else {
			alertas = append(alertas, "OK RegistroUbicaciones")
			alerta.Code = errUbicaciones.Code
		}
		alerta.Body = alertas
		c.Data["json"] = alerta
		c.ServeJSON()

	} else {
		alertas = append(alertas, []interface{}{err.Error()})
		c.Ctx.Output.SetStatus(400)
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = alertas
		c.Data["json"] = alerta
		c.ServeJSON()
	}

}

// ActualizarDatosContacto ...
// @Title ActualizarDatosContacto
// @Description Actualizar Datos de Contacto
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Actualizar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /DatosContacto [put]
func (c *PersonaController) ActualizarDatosContacto() {
	alertas := append([]interface{}{"Response: "})
	// datos de contacto de la persona
	var datos map[string]interface{}
	//var contactoEnte map[string]interface{}
	//reultado de la creacion de la persona
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}
	var resultado3 map[string]interface{}
	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err == nil {

		contactos := datos["ContactoEnte"].([]interface{})
		for i := 0; i < len(contactos); i++ {
			contacto := contactos[i].(map[string]interface{})
			if errContacto := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente/"+fmt.Sprintf("%.f", contacto["Id"].(float64)), "PUT", &resultado, contacto); errContacto == nil {
				if resultado["Type"].(string) == "error" {
					alertas = append(alertas, "Error UPDATE datos_contacto: "+resultado["Body"].(string))
				} else {
					alertas = append(alertas, "UPDATE datos_contacto")
				}
				alerta.Type = resultado["Type"].(string)
				alerta.Code = resultado["Code"].(string)
			} else {
				alertas = append(alertas, "ERROR contacto_ente: "+errContacto.Error())
				alerta.Type = "error"
				alerta.Code = "400"
			}
			alerta.Body = alertas
			c.Data["json"] = alerta
		}

		//actualización ubicaciones
		UbicacionEnte := datos["UbicacionEnte"].(map[string]interface{})
		lugar := UbicacionEnte["Lugar"].(map[string]interface{})
		UbicacionEnte["Lugar"] = lugar["Id"]
		if errUbicacionEnte := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/"+fmt.Sprintf("%.f", UbicacionEnte["Id"].(float64)), "PUT", &resultado2, UbicacionEnte); errUbicacionEnte == nil {

			if resultado2["Type"].(string) == "error" {
				alertas = append(alertas, "Error ubicacion_ente: "+resultado2["Body"].(string))
			} else {
				alertas = append(alertas, "UPDATE ubicacion_ente")

				//actualización atributos
				var ubicacion map[string]interface{}
				ubicacion = make(map[string]interface{})
				ubicacion["Atributos"] = UbicacionEnte["Atributos"]
				atributos := ubicacion["Atributos"].([]interface{})

				for i := 0; i < len(atributos); i++ {
					atributo := atributos[i].(map[string]interface{})
					if errAtributoUbicacion := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion/"+fmt.Sprintf("%.f", atributo["Id"].(float64)), "PUT", &resultado3, atributo); errAtributoUbicacion == nil {
						if resultado3["Type"].(string) == "error" {
							alertas = append(alertas, "Error UPDATE atributos: "+resultado3["Body"].(string))
						} else {
							alertas = append(alertas, "UPDATE  atributos")
						}
						alerta.Type = resultado3["Type"].(string)
						alerta.Code = resultado3["Code"].(string)
					} else {
						alertas = append(alertas, "Error UPDATE atributos: "+errAtributoUbicacion.Error())
						alerta.Type = "error"
						alerta.Code = "400"
					}
					alerta.Body = alertas
					c.Data["json"] = alerta
					c.ServeJSON()
				}
			}

		} else {
			alertas = append(alertas, "Error UPDATE ubicación: "+errUbicacionEnte.Error())
			alerta.Type = "error"
			alerta.Code = "400"
		}
		alerta.Body = alertas
		c.Data["json"] = alerta
		c.ServeJSON()

	} else {
		alertas = append(alertas, err.Error())
		c.Ctx.Output.SetStatus(400)
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = alertas
		c.Data["json"] = alerta
		c.ServeJSON()
	}
}

// ConsultaDatosComplementarios ...
// @Title Getdatoscomplementarios
// @Description conultar datos complementarios
// @Param	id		path 	string	true		"consulta de datos complementarios por el Id del ente"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /DatosComplementarios/:id [get]
func (c *PersonaController) ConsultaDatosComplementarios() {
	var query = make(map[string]string)
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = models.Alert{Type: "error", Code: "S_400", Body: "Error: invalid query key/value pair"}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v

		}
	}

	var alerta models.Alert
	idStr := c.Ctx.Input.Param(":id")
	var Persona []map[string]interface{}
	var GrupoEtnico []map[string]interface{}
	var TipoGrupoEtnico interface{}
	var TipoGrupoSanguineo interface{}
	var TipoRh interface{}

	var Discapacidades []map[string]interface{}
	var Lugar map[string]interface{}
	var GrupoSanguineo []map[string]interface{}
	var UbicacionEnte []map[string]interface{}
	var TipoDiscapacidad []map[string]interface{}
	var IdentificacionEnte []map[string]interface{}
	errores := append([]interface{}{"Response:"})

	//buscar persona con el ente
	errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Ente:"+idStr, &Persona)
	var s string
	if query != nil {
		for key, val := range query {
			s = fmt.Sprintf(",%s:%s", key, val)
		}
	}

	if errPersona == nil && Persona != nil {
		errGrupoEtnico := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/?query=Persona:"+fmt.Sprintf("%.f", Persona[0]["Id"].(float64)), &GrupoEtnico)
		errDiscapacidades := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/?query=Persona:"+fmt.Sprintf("%.f", Persona[0]["Id"].(float64))+",Activo:true", &Discapacidades)
		errUbicacionEnte := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/?query=Ente:"+idStr+s+"&fields=Id,TipoRelacionUbicacionEnte,Lugar", &UbicacionEnte)
		errGrupoSanguineo := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/?query=Persona:"+fmt.Sprintf("%.f", Persona[0]["Id"].(float64)), &GrupoSanguineo)
		/*errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+idStr+"&fields=Id,TipoIdentificacion,NumeroIdentificacion,FechaExpedicion,LugarExpedicion", &IdentificacionEnte)
		if errIdentificacion != nil {
			fmt.Println("error de identificacion", errIdentificacion.Error())
		} else {
			fmt.Println(IdentificacionEnte)
		}*/
		var nuevapersona map[string]interface{}
		if GrupoEtnico == nil && Discapacidades == nil && UbicacionEnte == nil && GrupoSanguineo == nil && IdentificacionEnte == nil {
			nuevapersona = nil
		} else {
			if UbicacionEnte != nil {
				for i := 0; i < len(UbicacionEnte); i++ {
					//buscar relaciones del lugar
					l := fmt.Sprintf("%.f", UbicacionEnte[i]["Lugar"].(float64))

					if errJerarquiaLugar := request.GetJson("http://"+beego.AppConfig.String("UbicacionesService")+"/relacion_lugares/jerarquia_lugar/"+l, &Lugar); errJerarquiaLugar == nil {
						UbicacionEnte[i]["Lugar"] = Lugar
					} else {
						errores = append(errores, errJerarquiaLugar.Error())
						alerta.Type = "error"
						alerta.Code = "400"
						alerta.Body = errores
						c.Data["json"] = alerta
						c.ServeJSON()
					}
				}
			}

			for i := 0; i < len(Discapacidades); i++ {
				d := Discapacidades[i]["TipoDiscapacidad"].(map[string]interface{})
				TipoDiscapacidad = append(TipoDiscapacidad, d)
			}

			if errDiscapacidades == nil && errGrupoEtnico == nil && errGrupoSanguineo == nil && errUbicacionEnte == nil {
				if GrupoEtnico == nil {
					TipoGrupoEtnico = nil
				} else {
					TipoGrupoEtnico = GrupoEtnico[0]["GrupoEtnico"]
				}
				if GrupoSanguineo == nil {
					TipoGrupoSanguineo = nil
				} else {
					TipoGrupoSanguineo = GrupoSanguineo[0]["GrupoSanguineo"]
					TipoRh = GrupoSanguineo[0]["FactorRh"]
				}

				nuevapersona = map[string]interface{}{
					"GrupoEtnico":      TipoGrupoEtnico,
					"TipoDiscapacidad": TipoDiscapacidad,
					"Lugar":            UbicacionEnte,
					"GrupoSanguineo":   TipoGrupoSanguineo,
					"Rh":               TipoRh,
					"Identificacion":   IdentificacionEnte,
					"Ente":             Persona[0]["Ente"],
					//"Foto":            resultado["Persona"].(map[string]interface{})["Foto"],
				}
			} else {
				errores = append(errores, "ERROR")
				alerta.Type = "error"
				alerta.Code = "500"
				alerta.Body = errores
				c.Data["json"] = alerta
				c.ServeJSON()
			}
		}
		c.Data["json"] = nuevapersona
		c.ServeJSON()

	} else {
		errores = append(errores, "No found Persona", Persona)
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = errores
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
// @router /DatosComplementarios [post]
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
					errores = append(errores, "error grupo etnico: "+errGrupoEtnico.Error())
				}
				if resultado2["Type"] == "error" {
					errores = append(errores, resultado2["Body"])
				}

			} else {
				errores = append(errores, "OK persona_grupo_etnico")
			}
			if (persona["GrupoSanguineo"] == "O" || persona["GrupoSanguineo"] == "A" || persona["GrupoSanguineo"] == "AB" || persona["GrupoSanguineo"] == "B") && (persona["Rh"] == "+" || persona["Rh"] == "-") {

				GrupoSanguineo["Persona"] = resultado[0]
				GrupoSanguineo["FactorRh"] = persona["Rh"]
				GrupoSanguineo["GrupoSanguineo"] = persona["GrupoSanguineo"]
				errGrupoSanguineo := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona", "POST", &resultado3, GrupoSanguineo)

				if errGrupoSanguineo == nil {
					errores = append(errores, "OK grupo_sanquineo_persona")
				} else {
					errores = append(errores, errGrupoSanguineo.Error())
				}
			} else {

				errores = append(errores, "el grupo sanguineo es incorrecto")
			}

			discapacidad := persona["TipoDiscapacidad"].([]interface{})

			for i := 0; i < len(discapacidad); i++ {

				Discapacidad["Persona"] = resultado[0]
				Discapacidad["TipoDiscapacidad"] = discapacidad[i]
				Discapacidad["Activo"] = true
				errDiscapacidad := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad", "POST", &resultado2, Discapacidad)

				if errDiscapacidad != nil || resultado2["Type"] == "error" {
					errores = append(errores, "La discapacidad con Id "+fmt.Sprintf("%.f", discapacidad[i].(map[string]interface{})["Id"])+" obtuvo "+fmt.Sprintf("%v", resultado2["Type"])+" "+fmt.Sprintf("%v", resultado2["Body"]))
					alerta.Type = "error"
					alerta.Code = "400"
				} else {
					errores = append(errores, "La discapacidad con Id "+fmt.Sprintf("%.f", discapacidad[i].(map[string]interface{})["Id"])+" obtuvo "+fmt.Sprintf("%v", resultado2["Type"])+" "+fmt.Sprintf("%v", resultado2["Body"]))
					alerta.Type = "sucess"
					alerta.Code = "200"
				}
			}

			//agregar identifiacion
			//fmt.Println("la identificacion es: ", persona["Identificacion"])
			var identificacion map[string]interface{}
			identificacion = make(map[string]interface{})
			var ente2 map[string]interface{}
			ente2 = make(map[string]interface{})
			ente2["Id"] = persona["Ente"]
			identificacion["Ente"] = ente2

			/*evaluar si guardar la libreta militar
			identificacion["FechaExpedicion"] = persona["Identificacion"].(map[string]interface{})["FechaExpedicion"]
			identificacion["LugarExpedicion"] = persona["Identificacion"].(map[string]interface{})["LugarExpedicion"]
			identificacion["NumeroIdentificacion"] = persona["Identificacion"].(map[string]interface{})["NumeroIdentificacion"]
			identificacion["TipoIdentificacion"] = persona["Identificacion"].(map[string]interface{})["TipoIdentificacion"]
			fmt.Println("la identifiacion nueva es: ", identificacion)
			var resultadoid map[string]interface{}
			if errIdentificacionEnte := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/", "POST", &resultadoid, identificacion); errIdentificacionEnte == nil {
				fmt.Println("el resultado de la identificacion es: ", resultadoid)
				if resultadoid["Type"] == "error" {
					errores = append(errores, []interface{}{"La Identificacion tuvo el error ", resultadoid["Type"]})

				} else {
					errores = append(errores, []interface{}{"OK identificacion"})
				}

			} else {
				errores = append(errores, []interface{}{"error identifiacion: ", errIdentificacionEnte.Error()})
			}
			*/

			//registro del lugar
			var ubicacion map[string]interface{}
			ubicacion = make(map[string]interface{})
			ubicacion["Ente"] = persona["Ente"]
			ubicacion["Lugar"] = persona["Lugar"]
			ubicacion["TipoRelacionUbicacionEnte"] = persona["TipoRelacionUbicacionEnte"]

			errUbicaciones := RegistroUbicaciones(ubicacion)
			if errUbicaciones.Type != "success" {
				errores = append(errores, errUbicaciones.Body)
			} else {
				errores = append(errores, "succes: ")
			}

			alerta.Body = errores
			c.Data["json"] = alerta
			c.ServeJSON()
		} else {

			if errPersona != nil {
				errores = append(errores, "error persona: ", errPersona)
			}
			if len(resultado) == 0 {
				errores = append(errores, "NO existe ninguna persona con este ente")
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

// ActualizarDatosComplementarios ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /DatosComplementarios [put]
func (c *PersonaController) ActualizarDatosComplementarios() {
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

	var idpersona_grupo_etnico []map[string]interface{}
	var idpersona_grupo_sanguineo []map[string]interface{}
	var persona_discapacidades []interface{}

	var id_ubicacion_ente []map[string]interface{}
	//var id_identificacion_ente []map[string]interface{}
	//Resultado de agregar grupo sanguineo y discapacidades
	var resultado2 map[string]interface{}
	//Resultado de agregar grupo sanguineo y discapacidades
	var resultado3 map[string]interface{}

	//var resultado4 map[string]interface{}
	//acumulado de errores
	errores := append([]interface{}{"Response:"})
	//comprobar que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Ente:"+fmt.Sprintf("%.f", persona["Ente"].(float64))+"&fields=Id", &resultado)

		if errPersona == nil && resultado != nil {

			GrupoEtnico["GrupoEtnico"] = persona["GrupoEtnico"]
			GrupoEtnico["Persona"] = resultado[0]

			request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/?query=persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64)), &idpersona_grupo_etnico)
			fmt.Println("persona grupo etnico ", idpersona_grupo_etnico)
			errGrupoEtnico := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/"+fmt.Sprintf("%.f", idpersona_grupo_etnico[0]["Id"].(float64)), "PUT", &resultado2, GrupoEtnico)

			if errGrupoEtnico != nil || resultado2["Id"] == 0 || resultado2["Type"] == "error" {

				if errGrupoEtnico != nil {
					errores = append(errores, "ERROR grupo etnico: "+errGrupoEtnico.Error())
				}
				if resultado2["Type"] == "error" {
					errores = append(errores, resultado2)
				}

			} else {
				errores = append(errores, []interface{}{"UPDATE persona_grupo_etnico"})
			}
			if (persona["GrupoSanguineo"] == "O" || persona["GrupoSanguineo"] == "A" || persona["GrupoSanguineo"] == "AB" || persona["GrupoSanguineo"] == "B") && (persona["Rh"] == "+" || persona["Rh"] == "-") {
				GrupoSanguineo["Persona"] = resultado[0]
				GrupoSanguineo["FactorRh"] = persona["Rh"]
				GrupoSanguineo["GrupoSanguineo"] = persona["GrupoSanguineo"]
				request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/?query=persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64)), &idpersona_grupo_sanguineo)
				errGrupoSanguineo := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/"+fmt.Sprintf("%.f", idpersona_grupo_sanguineo[0]["Id"].(float64)), "PUT", &resultado3, GrupoSanguineo)

				if errGrupoSanguineo == nil {
					errores = append(errores, "UPDATE grupo_sanquineo_persona")
				} else {
					errores = append(errores, "ERROR grupo_sanquineo_persona"+errGrupoSanguineo.Error())
				}
			} else {
				errores = append(errores, "el grupo sanguineo es incorrecto:"+fmt.Sprintf("%v", persona["GrupoSanguineo"])+fmt.Sprintf("%v", persona["Rh"]))
			}

			/*####################################################################################*/
			//discapacidades enviadas en el form
			discapacidades := persona["TipoDiscapacidad"].([]interface{})
			//get discapacidades persona
			request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/?query=persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64)), &persona_discapacidades)

			//se pueden actualizar(cambiar de estado) o agregar nuevas
			var resultado2 map[string]interface{}

			for i := 0; i < len(discapacidades); i++ {
				discapacidad := discapacidades[i].(map[string]interface{})

				for j := 0; j < len(persona_discapacidades); j++ {
					persona_discapacidad := persona_discapacidades[j].(map[string]interface{})
					d := persona_discapacidad["TipoDiscapacidad"].(map[string]interface{})

					//está!!
					if d["Id"] == discapacidad["Id"] {
						//verificar que esté activa, si no lo está activarla
						if persona_discapacidad["Activo"] == false {
							Discapacidad["Persona"] = resultado[0]
							Discapacidad["TipoDiscapacidad"] = persona_discapacidad["TipoDiscapacidad"]
							Discapacidad["Activo"] = true
							errDiscapacidad := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/"+fmt.Sprintf("%.f", persona_discapacidad["Id"].(float64)), "PUT", &resultado2, Discapacidad)

							if errDiscapacidad != nil || resultado2["Type"] == "error" {
								errores = append(errores, "La discapacidad con Id "+fmt.Sprintf("%.f", persona_discapacidad["Id"])+" obtuvo "+fmt.Sprintf("%v", resultado2["Type"])+" "+fmt.Sprintf("%v", resultado2["Body"]))
								alerta.Type = "error"
								alerta.Code = "400"
							} else {
								errores = append(errores, "La discapacidad con Id "+fmt.Sprintf("%.f", persona_discapacidad["Id"])+" obtuvo "+fmt.Sprintf("%v", resultado2["Type"]))
								alerta.Type = "sucess"
								alerta.Code = "200"
							}

						}
					} else {
						//si no tiene la discapacidad se agrega
						Discapacidad["Persona"] = resultado[0]
						Discapacidad["TipoDiscapacidad"] = discapacidad
						Discapacidad["Activo"] = true
						errDiscapacidad := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad", "POST", &resultado2, Discapacidad)

						if errDiscapacidad != nil || resultado2["Type"] == "error" {
							errores = append(errores, "La discapacidad con Id "+fmt.Sprintf("%.f", persona_discapacidad["Id"])+" obtuvo "+fmt.Sprintf("%v", resultado2["Type"])+" "+fmt.Sprintf("%v", resultado2["Body"]))
							alerta.Type = "error"
							alerta.Code = "400"
						} else {
							errores = append(errores, "La discapacidad con Id "+fmt.Sprintf("%.f", persona_discapacidad["Id"])+" obtuvo "+fmt.Sprintf("%v", resultado2["Type"]))
							alerta.Type = "sucess"
							alerta.Code = "200"
						}
					}

				}
			}

			//verificar si se han quitado discapacidades
			for j := 0; j < len(persona_discapacidades); j++ {
				persona_discapacidad := persona_discapacidades[j].(map[string]interface{})
				d := persona_discapacidad["TipoDiscapacidad"].(map[string]interface{})
				iguales := false

				for i := 0; i < len(discapacidades); i++ {
					discapacidad := discapacidades[i].(map[string]interface{})
					if d["Id"] == discapacidad["Id"] {
						iguales = true
					}
				}
				if iguales == false {
					//inactivar discapacidad
					Discapacidad["Persona"] = resultado[0]
					Discapacidad["TipoDiscapacidad"] = d
					Discapacidad["Activo"] = false
					errDiscapacidad := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/"+fmt.Sprintf("%.f", persona_discapacidad["Id"].(float64)), "PUT", &resultado2, Discapacidad)

					if errDiscapacidad != nil || resultado2["Type"] == "error" {
						errores = append(errores, []interface{}{"La discapacidad con Id ", persona_discapacidad["Id"], " obtuvo ", resultado2["Type"], " ", resultado2["Body"]})
						alerta.Type = "error"
						alerta.Code = "400"
					} else {
						errores = append(errores, []interface{}{"La discapacidad con Id ", persona_discapacidad["Id"], " obtuvo el resultado ", resultado2["Type"]})
						alerta.Type = "sucess"
						alerta.Code = "200"
					}
				}
			}
			/*####################################################################################*/

			request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/?query=Ente:"+fmt.Sprintf("%.f", persona["Ente"].(float64))+"&fields=Id", &id_ubicacion_ente)
			fmt.Println("el id de la ubicacion ente: ", id_ubicacion_ente)
			var ubicacion map[string]interface{}
			ubicacion = make(map[string]interface{})
			ubicacion["Ente"] = map[string]interface{}{"Id": persona["Ente"]}
			lugar := persona["Lugar"].(map[string]interface{})
			ubicacion["Lugar"] = lugar["Id"]
			ubicacion["TipoRelacionUbicacionEnte"] = map[string]interface{}{"Id": persona["TipoRelacionUbicacionEnte"]}
			ubicacion["Activo"] = true
			if errUbicacionEnte := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/"+fmt.Sprintf("%.f", id_ubicacion_ente[0]["Id"].(float64)), "PUT", &resultado2, ubicacion); errUbicacionEnte == nil {
				if resultado2["Type"].(string) == "error" {
					errores = append(errores, resultado2["Body"])
				} else {
					errores = append(errores, "UPDATE ubicacion_ente")
				}
			}
			/*errIdentifiacionEnte := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+fmt.Sprintf("%.f", persona["Ente"].(float64))+"&fields=Id", &id_identificacion_ente)
			if errIdentifiacionEnte == nil && id_identificacion_ente[0]["Id"] != nil {

				var Identificacion map[string]interface{}
				Identificacion = make(map[string]interface{})
				Identificacion["Ente"] = map[string]interface{}{"Id": persona["Ente"]}
				Identificacion["TipoIdentificacion"] = persona["Identificacion"].(map[string]interface{})["TipoIdentificacion"]
				Identificacion["NumeroIdentificacion"] = persona["Identificacion"].(map[string]interface{})["NumeroIdentificacion"]
				Identificacion["FechaExpedicion"] = persona["Identificacion"].(map[string]interface{})["FechaExpedicion"]
				Identificacion["LugarExpedicion"] = persona["Identificacion"].(map[string]interface{})["LugarExpedicion"]
				//fmt.Println("la identificacion que se va en el PUT es: ", Identificacion)
				if errIdentificacion := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/"+fmt.Sprintf("%.f", id_identificacion_ente[0]["Id"].(float64)), "PUT", &resultado4, Identificacion); errIdentificacion == nil {
					if resultado4["Type"].(string) == "error" {
						errores = append(errores, resultado4["Body"])
					} else {
						errores = append(errores, []interface{}{"OK update identificacion"})
					}
				}
			} else {
				errores = append(errores, []interface{}{"error identificacion: ", errIdentifiacionEnte.Error()})
			}*/

			alerta.Body = errores
			c.Data["json"] = alerta
			c.ServeJSON()
		} else {

			if errPersona != nil {
				errores = append(errores, "Error persona: "+errPersona.Error())
			}
			if len(resultado) == 0 {
				errores = append(errores, "NO existe ninguna persona con este ente")

			}
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = errores
			c.Data["json"] = alerta
			c.ServeJSON()
		}
	} else {

		errores = append(errores, err.Error())
		c.Ctx.Output.SetStatus(200)
		alerta.Type = "error"
		alerta.Code = "401"
		alerta.Body = errores
		c.Data["json"] = alerta
		c.ServeJSON()
	}
}

// DatosContacto ...
// @Title DatosContacto
// @Description Datos de contacto
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {}
// @Failure 403 body is empty
// @router /DatosContacto/:id [get]
func (c *PersonaController) DatosContacto() {
	var query = make(map[string]string)
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = models.Alert{Type: "error", Code: "S_400", Body: "Error: invalid query key/value pair"}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v

		}
	}

	var alerta models.Alert
	idStr := c.Ctx.Input.Param(":id")
	alertas := append([]interface{}{"acumulado de alertas"})

	var ContactoEnte []map[string]interface{}
	var UbicacionEnte []map[string]interface{}

	var s string
	if query != nil {
		for key, val := range query {
			s = fmt.Sprintf(",%s:%s", key, val)
		}
	}

	if errContactoEnte := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente/?query=Ente.Id:"+idStr+"&fields=Id,TipoContacto,Valor", &ContactoEnte); errContactoEnte == nil {

		if errUbicacionEnte := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/?query=Ente.Id:"+idStr+s+"&fields=Id,Lugar,TipoRelacionUbicacionEnte", &UbicacionEnte); errUbicacionEnte == nil {
			var persona map[string]interface{}
			if ContactoEnte == nil && UbicacionEnte == nil {
				persona = nil
			} else {
				//buscar atributos de la ubicacion
				var AtributosEnte []map[string]interface{}
				var Lugar map[string]interface{}
				for i := 0; i < len(UbicacionEnte); i++ {
					s := fmt.Sprintf("%.f", UbicacionEnte[i]["Id"].(float64))
					if errAtributosUbicacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion?query=UbicacionEnte.Id:"+s+"&fields=Id,AtributoUbicacion,Valor", &AtributosEnte); errAtributosUbicacion == nil {
						UbicacionEnte[i]["Atributos"] = AtributosEnte
					} else {
						alertas = append(alertas, errAtributosUbicacion.Error())
						alerta.Type = "error"
						alerta.Code = "400"
						alerta.Body = alertas
						c.Data["json"] = alerta
						c.ServeJSON()
					}

					//buscar relaciones del lugar
					l := fmt.Sprintf("%.f", UbicacionEnte[i]["Lugar"].(float64))
					if errJerarquiaLugar := request.GetJson("http://"+beego.AppConfig.String("UbicacionesService")+"/relacion_lugares/jerarquia_lugar/"+l, &Lugar); errJerarquiaLugar == nil {
						UbicacionEnte[i]["Lugar"] = Lugar
					} else {
						alertas = append(alertas, errJerarquiaLugar.Error())
						alerta.Type = "error"
						alerta.Code = "400"
						alerta.Body = alertas
						c.Data["json"] = alerta
						c.ServeJSON()
					}
				}

				persona = map[string]interface{}{
					"ContactoEnte":  ContactoEnte,
					"UbicacionEnte": UbicacionEnte,
				}
			}

			c.Data["json"] = persona
			c.ServeJSON()
		} else {
			alertas = append(alertas, errUbicacionEnte.Error())
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = alertas
			c.Data["json"] = alerta
			c.ServeJSON()
		}
	} else {
		alertas = append(alertas, errContactoEnte.Error())
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = alertas
		c.Data["json"] = alerta
		c.ServeJSON()
	}

}
