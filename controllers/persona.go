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
	c.Mapping("ActualizarDatosComplementarios", c.ActualizarDatosComplementarios)
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
	alertas := append([]interface{}{"acumulado de alertas"})
	errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/full/?userid="+idStr, &resultado)
	if errPersona == nil && resultado["Type"] != "error" && resultado != nil {
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
		if errPersona != nil {
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = append(alertas, []interface{}{"entro al error de persona"}, []interface{}{errPersona.Error()})
			c.Data["json"] = alerta
		} else {
			alerta.Type = "error"
			alerta.Code = "401"
			alerta.Body = append(alertas, []interface{}{"entro al error de respuesta"}, []interface{}{resultado["Body"]})
			c.Data["json"] = alerta
		}
		c.ServeJSON()
	}
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
	ubicacion = make(map[string]interface{})
	ubicacion["Ente"] = map[string]interface{}{"Id": ubicacionPersona["Ente"]}
	ubicacion["Lugar"] = ubicacionPersona["Lugar"]
	ubicacion["TipoRelacionUbicacionEnte"] = map[string]interface{}{"Id": ubicacionPersona["TipoRelacionUbicacionEnte"]}
	ubicacion["Activo"] = true

	// resultado registro ubicacion_ente
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}

	//funcion que realiza  de la  peticion POST /ubicacion_ente

	if errUbicacionEnte := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/ubicacion_ente", "POST", &resultado, ubicacion); errUbicacionEnte == nil {
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

						//funcion que realiza  de la  peticion POST /ubicacion_ente
						errAtributoUbicacion := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/valor_atributo_ubicacion", "POST", &resultado2, valorAtributoUbicacion)
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
	alerta.Body = errores
	return alerta
}

// RegistrarUbicaciones ...
// @Title RegistrarUbicacionalertases
// @Description Registrar Ubalertasicaciones
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
// @router /GuardarDatosContacto [post]
func (c *PersonaController) GuardarDatosContacto() {
	errores := append([]interface{}{"acumulado de alertas"})
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

			errContacto := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/contacto_ente", "POST", &resultado, contactoEnte)

			if errContacto == nil && resultado["Type"] == "success" {
				errores = append(errores, []interface{}{"succes: "})

			} else {
				errores = append(errores, resultado["Body"].(string))
			}

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
			errores = append(errores, errUbicaciones)
		} else {
			errores = append(errores, []interface{}{"succes: "})
		}

	} else {
		errores = append(errores, err.Error())
	}
	alerta.Type = "error"
	alerta.Body = errores
	c.Data["json"] = alerta
	c.ServeJSON()

}

// ActualizarDatosContacto ...
// @Title ActualizarDatosContacto
// @Description Actualizar Datos de Contacto
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Actualizar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /ActualizarDatosContacto [put]
func (c *PersonaController) ActualizarDatosContacto() {
	// datos de contacto de la persona
	var datos map[string]interface{}
	//var contactoEnte map[string]interface{}
	//reultado de la creacion de la persona
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}
	var resultado3 map[string]interface{}
	// alerta que retorna la funcion Guardar persona
	var alerta models.Alert
	//acumulado de alertas
	var alertas string

	//valida que el JSON de entrada sea correcto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err == nil {

		contactos := datos["ContactoEnte"].([]interface{})
		for i := 0; i < len(contactos); i++ {
			contacto := contactos[i].(map[string]interface{})
			if errContacto := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/contacto_ente/"+fmt.Sprintf("%.f", contacto["Id"].(float64)), "PUT", &resultado, contacto); errContacto == nil {
				if resultado["Type"].(string) == "error" {
					alertas = alertas + "Error en la actualización del contacto: " + resultado["Body"].(string)
				} else {
					alertas = alertas + "OK actualización de contacto"
				}
				alerta.Type = resultado["Type"].(string)
				alerta.Code = resultado["Code"].(string)
				alerta.Body = alertas
			} else {
				alertas = alertas + " ERROR contacto_ente: " + errContacto.Error()
				alerta.Type = "error"
				alerta.Code = "400"
				alerta.Body = alertas
			}
			c.Data["json"] = alerta
		}

		//actualización ubicaciones
		UbicacionEnte := datos["UbicacionEnte"].(map[string]interface{})
		if errUbicacionEnte := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/ubicacion_ente/"+fmt.Sprintf("%.f", UbicacionEnte["Id"].(float64)), "PUT", &resultado2, UbicacionEnte); errUbicacionEnte == nil {
			fmt.Println(resultado2)
			if resultado2["Type"].(string) == "error" {
				alertas = alertas + "Error ubicacion_ente: " + resultado2["Body"].(string)
			} else {
				alertas = alertas + "OK ubicacion_ente"

				//actualización atributos
				var ubicacion map[string]interface{}
				ubicacion = make(map[string]interface{})
				ubicacion["Atributos"] = UbicacionEnte["Atributos"]
				atributos := ubicacion["Atributos"].([]interface{})

				for i := 0; i < len(atributos); i++ {
					atributo := atributos[i].(map[string]interface{})
					if errAtributoUbicacion := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/valor_atributo_ubicacion/"+fmt.Sprintf("%.f", atributo["Id"].(float64)), "PUT", &resultado3, atributo); errAtributoUbicacion == nil {
						if resultado3["Type"].(string) == "error" {
							alertas = alertas + "Error en la actualización de de los atributos: " + resultado3["Body"].(string)
							fmt.Println(resultado3["Code"].(string))
						} else {
							alertas = alertas + "OK actualización de atributos"
						}
						alerta.Type = resultado3["Type"].(string)
						alerta.Code = resultado3["Code"].(string)
						alerta.Body = alertas
					} else {
						alertas = alertas + "Error en la actualización de los atributos: " + errAtributoUbicacion.Error()
						alerta.Type = "error"
						alerta.Code = "400"
						alerta.Body = alertas
					}
					c.Data["json"] = alerta
					c.ServeJSON()
				}
			}

		} else {
			alertas = alertas + "Error en la actualización de la ubicación: " + errUbicacionEnte.Error()
			alerta.Type = "error"
			alerta.Code = "400"
			alerta.Body = alertas
		}
		c.Data["json"] = alerta
		c.ServeJSON()

	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = "ERROR formato incorrecto" + err.Error()
		c.Data["json"] = alerta

	}

	c.ServeJSON()
}

// ConsultaDatosComplementarios ...
// @Title Getdatoscomplementarios
// @Description conultar datos complementarios
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /ConsultaDatosComplementarios/:id [get]
func (c *PersonaController) ConsultaDatosComplementarios() {
	var alerta models.Alert
	idStr := c.Ctx.Input.Param(":id")
	var GrupoEtnico []map[string]interface{}
	var TipoGrupoEtnico interface{}
	var TipoGrupoSanguineo interface{}
	var TipoRh interface{}

	var Discapacidades []map[string]interface{}
	var Lugar map[string]interface{}
	var GrupoSanguineo []map[string]interface{}
	var UbicacionEnte []map[string]interface{}
	var TipoDiscapacidad [3]interface{}
	errores := append([]interface{}{"acumulado de alertas"})
	//var persona map[string]interface{}

	errGrupoEtnico := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/?query=Persona:"+idStr, &GrupoEtnico)
	errDiscapacidades := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/?query=Persona:"+idStr, &Discapacidades)
	errUbicacionEnte := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/ubicacion_ente/?query=Ente:"+fmt.Sprintf("%.f", GrupoEtnico[0]["Persona"].(map[string]interface{})["Ente"].(float64)), &UbicacionEnte)
	errGrupoSanguineo := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/?query=Persona:"+idStr, &GrupoSanguineo)

	if UbicacionEnte == nil {

	} else {

		errLugar := request.GetJson("http://"+beego.AppConfig.String("UbicacionesService")+"/lugar/"+fmt.Sprintf("%.f", UbicacionEnte[0]["Lugar"].(float64)), &Lugar)
		fmt.Println("la url: ", "http://"+beego.AppConfig.String("UbicacionesService")+"/lugar/"+fmt.Sprintf("%.f", UbicacionEnte[0]["Lugar"].(float64)))
		fmt.Println("el lugar: ", Lugar)
		if errLugar != nil {
			fmt.Println("el error de lugar", errLugar)
		}
	}

	for i := 0; i < len(Discapacidades); i++ {

		TipoDiscapacidad[i] = Discapacidades[i]["TipoDiscapacidad"]
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

		nuevapersona := map[string]interface{}{
			"GrupoEtnico":      TipoGrupoEtnico,
			"TipoDiscapacidad": TipoDiscapacidad,
			"PaisNacimiento":   Lugar,
			"GrupoSanguineo":   TipoGrupoSanguineo,
			"Rh":               TipoRh,

			//"Foto":            resultado["Persona"].(map[string]interface{})["Foto"],

		}

		c.Data["json"] = nuevapersona
		c.ServeJSON()
	} else {
		errores = append(errores, []interface{}{"otro error "})
		alerta.Type = "sucess"
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

			//registro del lugar
			var ubicacion map[string]interface{}
			ubicacion = make(map[string]interface{})
			ubicacion["Ente"] = persona["Ente"]
			ubicacion["Lugar"] = persona["Lugar"]
			ubicacion["TipoRelacionUbicacionEnte"] = persona["TipoRelacionUbicacionEnte"]

			errUbicaciones := RegistroUbicaciones(ubicacion)
			if errUbicaciones.Type != "success" {
				errores = append(errores, errUbicaciones)
			} else {
				errores = append(errores, []interface{}{"succes: "})
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

// ActualizarDatosComplementarios ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaDatosBasicos	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /ActualizarDatosComplementarios [put]
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

	var id_ubicacion_ente []map[string]interface{}
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

			request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/?query=persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64)), &idpersona_grupo_etnico)

			errGrupoEtnico := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/"+fmt.Sprintf("%.f", idpersona_grupo_etnico[0]["Id"].(float64)), "PUT", &resultado2, GrupoEtnico)

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
				request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/?query=persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64)), &idpersona_grupo_sanguineo)

				errGrupoSanguineo := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/"+fmt.Sprintf("%.f", idpersona_grupo_sanguineo[0]["Id"].(float64)), "PUT", &resultado3, GrupoSanguineo)
				//fmt.Println("el resultado del grupo sanguineo: ", resultado3)
				if errGrupoSanguineo == nil {
					errores = append(errores, []interface{}{"OK grupo_sanquineo_persona"})
				} else {
					errores = append(errores, []interface{}{"err grupo_sanquineo_persona", errGrupoSanguineo.Error()})
				}
			} else {

				errores = append(errores, []interface{}{"el grupo sanguineo es incorrecto:", persona["GrupoSanguineo"], persona["Rh"]})
			}

			discapacidad := persona["TipoDiscapacidad"].([]interface{})
			request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/?query=persona:"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64)), &idpersona_grupo_sanguineo)
			//fmt.Println("las discapacidades actuales son", idpersona_grupo_sanguineo)
			for i := 0; i < len(discapacidad); i++ {
				//fmt.Println("el id de la discapacidad ", i, " es ", idpersona_grupo_sanguineo[i]["Id"])
				Discapacidad["Persona"] = resultado[0]
				Discapacidad["TipoDiscapacidad"] = discapacidad[i]

				errDiscapacidad := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/"+fmt.Sprintf("%.f", idpersona_grupo_sanguineo[i]["Id"].(float64)), "PUT", &resultado2, Discapacidad)
				//fmt.Println("el error de la discapacidad ", i, " es ", errDiscapacidad)
				//fmt.Println("el resultado de la dicapacidad es", i, " es ", resultado2)
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

			request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/ubicacion_ente/?query=Ente:"+fmt.Sprintf("%.f", persona["Ente"].(float64))+"&fields=Id", &id_ubicacion_ente)
			//fmt.Println("el id de la ubicacion ente: ", id_ubicacion_ente)
			var ubicacion map[string]interface{}
			ubicacion = make(map[string]interface{})
			ubicacion["Ente"] = map[string]interface{}{"Id": persona["Ente"]}
			ubicacion["Lugar"] = persona["Lugar"]
			ubicacion["TipoRelacionUbicacionEnte"] = map[string]interface{}{"Id": persona["TipoRelacionUbicacionEnte"]}
			ubicacion["Activo"] = true
			fmt.Println("la ubicacion ente es:", ubicacion)
			if errUbicacionEnte := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/ubicacion_ente/"+fmt.Sprintf("%.f", id_ubicacion_ente[0]["Id"].(float64)), "PUT", &resultado2, ubicacion); errUbicacionEnte == nil {
				fmt.Println(resultado2)
				if resultado2["Type"].(string) == "error" {
					errores = append(errores, resultado2["Body"])
				} else {
					errores = append(errores, []interface{}{"OK update ubicacion_ente"})
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
