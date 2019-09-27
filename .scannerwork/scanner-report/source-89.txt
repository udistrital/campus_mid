package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// OrganizacionController operations for Organizacion
type OrganizacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrganizacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("RegistrarUbicacion", c.RegistrarUbicacion)
	c.Mapping("GetByIdentificacion", c.GetByIdentificacion)
	c.Mapping("GetByEnte", c.GetByEnte)
}

// RegistrarUbicacion ...
// @Title RegistrarUbicacion
// @Description Registrar Ubicacion Organizacion
// @Param	body		body 	interface	true		"body for Ubicacion Organizacion content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router /registar_ubicacion [post]
func (c *OrganizacionController) RegistrarUbicacion() {
	var ubicacion map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ubicacion); err == nil {
		var ubicacionPost map[string]interface{}
		ubicacionEnte := map[string]interface{}{
			"Activo":                    true,
			"Ente":                      map[string]interface{}{"Id": ubicacion["Ente"]},
			"Lugar":                     ubicacion["Lugar"].(map[string]interface{})["Id"].(float64),
			"TipoRelacionUbicacionEnte": map[string]interface{}{"Id": 3},
		}

		errUbicacionPost := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente", "POST", &ubicacionPost, ubicacionEnte)
		if errUbicacionPost == nil && fmt.Sprintf("%v", ubicacionPost["System"]) != "map[]" && ubicacionPost["Id"] != nil {
			if ubicacionPost["Status"] != 400 {
				fmt.Println("Nueva ubicacion:", ubicacionPost)
				atributos := ubicacion["Atributos"].([]interface{})
				ubicacionPost["Atributos"] = atributos

				for i := 0; i < len(atributos); i++ {
					var atributoPost map[string]interface{}
					atributo := atributos[i].(map[string]interface{})

					nuevoAtributo := map[string]interface{}{
						"UbicacionEnte":     ubicacionPost,
						"Valor":             atributo["Valor"],
						"AtributoUbicacion": atributo["AtributoUbicacion"],
					}

					errAtributoPost := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion", "POST", &atributoPost, nuevoAtributo)
					if errAtributoPost == nil && fmt.Sprintf("%v", atributoPost["System"]) != "map[]" && atributoPost["Id"] != nil {
						if atributoPost["Status"] != 400 {
							fmt.Println("El atributo es:", atributoPost)
							ubicacionPost["Atributos"].([]interface{})[i] = atributoPost
						} else {
							logs.Error(errAtributoPost)
							//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
							c.Data["system"] = atributoPost
							c.Abort("400")
						}
					} else {
						logs.Error(errAtributoPost)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = atributoPost
						c.Abort("400")
					}

				}

				c.Data["json"] = ubicacionPost

			} else {
				logs.Error(errUbicacionPost)
				//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = ubicacionPost
				c.Abort("400")
			}
		} else {
			logs.Error(errUbicacionPost)
			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = ubicacionPost
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Post ...
// @Title Create
// @Description create Organizacion
// @Param	body		body 	interface	true		"body for Organizacion content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *OrganizacionController) Post() {
	var organizacion map[string]interface{}
	//var resultado []map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &organizacion); err == nil {
		var identificacion []map[string]interface{}
		errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=NumeroIdentificacion:"+
			fmt.Sprintf("%v", organizacion["NumeroIdentificacion"])+",TipoIdentificacion.Id:"+
			fmt.Sprintf("%v", organizacion["TipoIdentificacion"].(map[string]interface{})["Id"]), &identificacion)
		if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" && identificacion[0]["NumeroIdentificacion"] != nil {
			if identificacion[0]["Status"] != 404 {
				c.Ctx.Output.SetStatus(200)
				c.Data["json"] = identificacion[0]["Ente"].(map[string]interface{})["Id"]
			} else {
				if identificacion[0]["Message"] == "Not found resource" {
					c.Data["json"] = nil
				} else {
					logs.Error(identificacion)
					//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = errIdentificacion
					c.Abort("404")
				}
			}
		} else if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
			if res, errores := CrearOrganizacion(organizacion); errores == nil {
				c.Data["json"] = res
			} else {
				logs.Error(errores)
				//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
				c.Data["json"] = errores
				c.Data["system"] = errores
				c.Abort("400")
			}
		} else {
			logs.Error(identificacion)
			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = errIdentificacion
			c.Abort("404")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// CrearOrganizacion Funcion que valida si existe la organizacion si no la crea
func CrearOrganizacion(organizacion map[string]interface{}) (res map[string]interface{}, errores []interface{}) {
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}
	org := map[string]interface{}{
		"TipoOrganizacion": organizacion["TipoOrganizacion"],
		"Nombre":           organizacion["Nombre"],
	}
	errOrganizacionPost := request.SendJson("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion", "POST", &resultado, org)
	if errOrganizacionPost == nil && fmt.Sprintf("%v", resultado["System"]) != "map[]" && resultado["Id"] != nil {
		if resultado["Status"] != 400 {
			iden := map[string]interface{}{
				"TipoIdentificacion":   organizacion["TipoIdentificacion"], // asegurando que 5 es el ID para NIT
				"NumeroIdentificacion": organizacion["NumeroIdentificacion"],
				"Ente":                 map[string]interface{}{"Id": resultado["Ente"]},
			}

			errIdentificacionPost := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion", "POST", &resultado2, iden)
			if errIdentificacionPost == nil && fmt.Sprintf("%v", resultado2["System"]) != "map[]" && resultado2["Id"] != nil {
				if resultado2["Status"] != 400 {
					res = resultado
					res["TipoIdentificacion"] = resultado2["TipoIdentificacion"]
					res["NumeroIdentificacion"] = resultado2["NumeroIdentificacion"]

					if organizacion["Contacto"] != nil {
						array := organizacion["Contacto"].([]interface{})
						for _, cont := range array {
							contacto := cont.(map[string]interface{})
							contacto["Ente"] = map[string]interface{}{"Id": resultado["Ente"]}
							var r interface{}
							request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente", "POST", &r, cont)
						}
					}
				} else {
					request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion/%.f", resultado["Id"]), "DELETE", &resultado2, nil)
					logs.Error(resultado2)
					errores = []interface{}{errIdentificacionPost, resultado2}
				}
			} else {
				logs.Error(resultado2)
				errores = []interface{}{errIdentificacionPost, resultado2}
			}
		} else {
			logs.Error(resultado)
			errores = []interface{}{errOrganizacionPost, resultado}
		}
	} else {
		logs.Error(resultado)
		errores = []interface{}{errOrganizacionPost, resultado}
	}
	return
}

// GetByIdentificacion ...
// @Title GetByIdentificacion
// @Description get Organizacion by id
// @Param	Id		query 	int	true		"Identification number as id"
// @Param	TipoId		query 	int	true		"TipoIdentificacion number as nit"
// @Success 200 {}
// @Failure 404 not found resource
// @router /identificacion/ [get]
func (c *OrganizacionController) GetByIdentificacion() {
	uid := c.GetString("Id")
	tid := c.GetString("TipoId")
	var resultado map[string]interface{}
	var identificacion []map[string]interface{}
	errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=NumeroIdentificacion:"+uid+",TipoIdentificacion.Id:"+tid, &identificacion)
	if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
		if identificacion[0]["Status"] != 404 {
			var organizacion []map[string]interface{}
			errOrganizacion := request.GetJson("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion/?query=Ente:"+fmt.Sprintf("%v", identificacion[0]["Ente"].(map[string]interface{})["Id"]), &organizacion)
			if errOrganizacion == nil && fmt.Sprintf("%v", organizacion[0]["System"]) != "map[]" && fmt.Sprintf("%v", organizacion[0]) != "map[]" {
				if organizacion[0]["Status"] != 404 {
					resultado = organizacion[0]
					resultado["TipoIdentificacion"] = identificacion[0]["TipoIdentificacion"]
					resultado["NumeroIdentificacion"] = identificacion[0]["NumeroIdentificacion"]

					var contactos []map[string]interface{}
					errContacto := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente/?query=Ente:"+fmt.Sprintf("%v", identificacion[0]["Ente"].(map[string]interface{})["Id"]), &contactos)
					fmt.Println("la respuesta contactos es:", contactos)

					if errContacto == nil && fmt.Sprintf("%v", contactos[0]["System"]) != "map[]" {
						if contactos[0]["Status"] != 404 {
							resultado["Contacto"] = contactos
						} else {
							if contactos[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(contactos)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errContacto
								c.Abort("404")
							}
						}
					} else {
						logs.Error(contactos)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errContacto
						c.Abort("404")
					}

					var ubicacion []map[string]interface{}
					errUbicacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion/?query=UbicacionEnte.Ente.Id:"+fmt.Sprintf("%v", identificacion[0]["Ente"].(map[string]interface{})["Id"]), &ubicacion)
					fmt.Println("la respuesta ubicacion es:", ubicacion)

					if errUbicacion == nil && fmt.Sprintf("%v", ubicacion[0]["System"]) != "map[]" {
						if ubicacion[0]["Status"] != 404 {
							resultado["Ubicacion"] = ubicacion[0]
						} else {
							if ubicacion[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(ubicacion)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errUbicacion
								c.Abort("404")
							}
						}
					} else {
						logs.Error(ubicacion)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errUbicacion
						c.Abort("404")
					}

					c.Data["json"] = resultado

				} else {
					if organizacion[0]["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(organizacion)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errOrganizacion
						c.Abort("404")
					}
				}
			} else {
				logs.Error(organizacion)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errOrganizacion
				c.Abort("404")
			}
		} else {
			if identificacion[0]["Message"] == "Not found resource" {
				c.Data["json"] = nil
			} else {
				logs.Error(identificacion)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errIdentificacion
				c.Abort("404")
			}
		}
	} else {
		logs.Error(identificacion)
		//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = errIdentificacion
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetByEnte ...
// @Title GetByEnte
// @Description get Organizacion by id
// @Param	ente		path 	int	true		"The key for staticblock"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:ente [get]
func (c *OrganizacionController) GetByEnte() {
	uid := c.Ctx.Input.Param(":ente")
	var resultado map[string]interface{}
	var organizacion []map[string]interface{}
	errOrganizacion := request.GetJson("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion/?query=Ente:"+uid, &organizacion)
	if errOrganizacion == nil && fmt.Sprintf("%v", organizacion[0]["System"]) != "map[]" && fmt.Sprintf("%v", organizacion[0]) != "map[]" {
		if organizacion[0]["Status"] != 404 {
			resultado = organizacion[0]
			var identificacion []map[string]interface{}
			errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+uid, &identificacion)
			if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
				if identificacion[0]["Status"] != 404 {
					resultado["TipoIdentificacion"] = identificacion[0]["TipoIdentificacion"]
					resultado["NumeroIdentificacion"] = identificacion[0]["NumeroIdentificacion"]

					var contactos []map[string]interface{}
					errContacto := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente/?query=Ente:"+uid, &contactos)
					fmt.Println("la respuesta contactos es:", contactos)

					if errContacto == nil && fmt.Sprintf("%v", contactos[0]["System"]) != "map[]" {
						if contactos[0]["Status"] != 404 {
							resultado["Contacto"] = contactos
						} else {
							if contactos[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(contactos)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errContacto
								c.Abort("404")
							}
						}
					} else {
						logs.Error(contactos)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errContacto
						c.Abort("404")
					}

					var ubicacion []map[string]interface{}
					errUbicacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion/?query=UbicacionEnte.Ente.Id:"+uid, &ubicacion)
					fmt.Println("la respuesta ubicacion es:", ubicacion)

					if errUbicacion == nil && fmt.Sprintf("%v", ubicacion[0]["System"]) != "map[]" {
						if ubicacion[0]["Status"] != 404 {
							resultado["Ubicacion"] = ubicacion[0]
						} else {
							if ubicacion[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(ubicacion)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errUbicacion
								c.Abort("404")
							}
						}
					} else {
						logs.Error(ubicacion)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errUbicacion
						c.Abort("404")
					}

					c.Data["json"] = resultado
				} else {
					if identificacion[0]["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(identificacion)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errIdentificacion
						c.Abort("404")
					}
				}
			} else {
				logs.Error(identificacion)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errIdentificacion
				c.Abort("404")
			}
		} else {
			if organizacion[0]["Message"] == "Not found resource" {
				c.Data["json"] = nil
			} else {
				logs.Error(organizacion)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errOrganizacion
				c.Abort("404")
			}
		}
	} else {
		logs.Error(organizacion)
		//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = errOrganizacion
		c.Abort("404")
	}
	c.ServeJSON()
}
