package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// InscripcionController ...
type InscripcionController struct {
	beego.Controller
}

// URLMapping ...
func (c *InscripcionController) URLMapping() {
	c.Mapping("PutInscripcion", c.PutInscripcion)
	c.Mapping("GetInscripcion", c.GetInscripcion)
	c.Mapping("GetInscripcionByPeriodoPrograma", c.GetInscripcionByPeriodoPrograma)
	c.Mapping("GetByIdentificacion", c.GetByIdentificacion)
}

// PutInscripcion ...
// @Title PutInscripcion
// @Description Modificar datos de la admisión
// @Param	id		path 	int	true		"Id de la inscripción a modificar"
// @Param	body		body 	{}	true		"body Modificar inscripción content"
// @Success 200 {}
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *InscripcionController) PutInscripcion() {
	idStr := c.Ctx.Input.Param(":id")
	// datos de inscripción de la persona
	var datos map[string]interface{}
	//reultado de la actualizacion de la inscripcion
	var resultado map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err == nil {
		errInscripcionPut := request.SendJson("http://"+beego.AppConfig.String("InscripcionService")+"/inscripcion/"+idStr, "PUT", &resultado, datos)
		if errInscripcionPut == nil && fmt.Sprintf("%v", resultado["System"]) != "map[]" {
			if resultado["Status"] != 400 {
				c.Ctx.Output.SetStatus(200)
				c.Data["json"] = resultado
			} else {
				logs.Error(resultado)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errInscripcionPut
				c.Abort("400")
			}
		} else {
			logs.Error(resultado)
			//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
			c.Data["System"] = errInscripcionPut
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Data["System"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetInscripcion ...
// @Title GetInscripcion
// @Description consultar Inscripcion por id
// @Param	id		path 	int	true		"Id de la inscripción"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:id [get]
func (c *InscripcionController) GetInscripcion() {
	//Id de la inscripción a consultar
	idStr := c.Ctx.Input.Param(":id")
	//resultado Inscripcion
	var resultado map[string]interface{}

	errInscripcion := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/inscripcion/"+idStr, &resultado)
	if errInscripcion == nil && fmt.Sprintf("%v", resultado["System"]) != "map[]" {
		if resultado["Status"] != 404 {
			//resultado del Estado de inscripcion
			var estado map[string]interface{}
			idEstado := resultado["EstadoInscripcionId"].(map[string]interface{})
			//resultado del Tipo de Inscripcion
			var tipo map[string]interface{}
			idTipo := resultado["TipoInscripcionId"].(map[string]interface{})
			//resultado del Aspirante
			var aspirante map[string]interface{}

			errEstado := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/estado_inscripcion/"+fmt.Sprintf("%v", idEstado["Id"]), &estado)
			if errEstado == nil && fmt.Sprintf("%v", estado["System"]) != "map[]" {
				if estado["Status"] != 404 {
					resultado["EstadoInscripcionId"] = estado
				} else {
					if estado["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(estado)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errEstado
						c.Abort("404")
					}
				}
			} else {
				logs.Error(estado)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errEstado
				c.Abort("404")
			}

			errTipo := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/tipo_inscripcion/"+fmt.Sprintf("%v", idTipo["Id"]), &tipo)
			if errTipo == nil && fmt.Sprintf("%v", tipo["System"]) != "map[]" {
				if tipo["Status"] != 404 {
					resultado["TipoInscripcionId"] = tipo
				} else {
					if tipo["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(tipo)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errTipo
						c.Abort("404")
					}
				}
			} else {
				logs.Error(tipo)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errTipo
				c.Abort("404")
			}

			errAspirante := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/"+fmt.Sprintf("%v", resultado["PersonaId"]), &aspirante)
			if errAspirante == nil && fmt.Sprintf("%v", aspirante["System"]) != "map[]" {
				if aspirante["Status"] != 404 {
					//resultado de la Identificacion del Aspirante
					var identificacion []map[string]interface{}
					errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+fmt.Sprintf("%v", aspirante["Ente"]), &identificacion)
					if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
						if identificacion[0]["Status"] != 404 {
							aspirante["TipoIdentificacion"] = identificacion[0]["TipoIdentificacion"]
							aspirante["NumeroDocumento"] = identificacion[0]["NumeroIdentificacion"]
							resultado["PersonaId"] = aspirante
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
					if aspirante["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(aspirante)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errAspirante
						c.Abort("404")
					}
				}
			} else {
				logs.Error(aspirante)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errAspirante
				c.Abort("404")
			}
		} else {
			if resultado["Message"] == "Not found resource" {
				c.Data["json"] = nil
			} else {
				logs.Error(resultado)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errInscripcion
				c.Abort("404")
			}
		}
	} else {
		logs.Error(resultado)
		//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = errInscripcion
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetInscripcionByPeriodoPrograma ...
// @Title GetInscripcionByPeriodoPrograma
// @Description consultar Inscripcion por id
// @Param	Id		query 	int	true		"Id de la inscripción"
// @Param	ProgramaId		query 	int	true		"Id del programa académico"
// @Param	PeriodoId		query 	int	true		"Id del periodo académico"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *InscripcionController) GetInscripcionByPeriodoPrograma() {
	//Ids de la inscripción a consultar
	idStr := c.GetString("Id")
	idPrograma := c.GetString("ProgramaId")
	idPeriodo := c.GetString("PeriodoId")

	//resultado Inscripcion
	var resultado map[string]interface{}
	var inscripcion []map[string]interface{}

	errInscripcion := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/inscripcion/?query=Id:"+idStr+",PeriodoId:"+idPeriodo+",ProgramaAcademicoId:"+idPrograma, &inscripcion)
	if errInscripcion == nil && fmt.Sprintf("%v", inscripcion[0]["System"]) != "map[]" {
		if inscripcion[0]["Status"] != 404 {
			//resultado del Estado de inscripcion
			var estado map[string]interface{}
			idEstado := inscripcion[0]["EstadoInscripcionId"].(map[string]interface{})
			//resultado del Tipo de Inscripcion
			var tipo map[string]interface{}
			idTipo := inscripcion[0]["TipoInscripcionId"].(map[string]interface{})
			//resultado del Aspirante
			var aspirante map[string]interface{}

			errEstado := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/estado_inscripcion/"+fmt.Sprintf("%v", idEstado["Id"]), &estado)
			if errEstado == nil && fmt.Sprintf("%v", estado["System"]) != "map[]" {
				if estado["Status"] != 404 {
					inscripcion[0]["EstadoInscripcionId"] = estado
				} else {
					if estado["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(estado)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errEstado
						c.Abort("404")
					}
				}
			} else {
				logs.Error(estado)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errEstado
				c.Abort("404")
			}

			errTipo := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/tipo_inscripcion/"+fmt.Sprintf("%v", idTipo["Id"]), &tipo)
			if errTipo == nil && fmt.Sprintf("%v", tipo["System"]) != "map[]" {
				if tipo["Status"] != 404 {
					inscripcion[0]["TipoInscripcionId"] = tipo
				} else {
					if tipo["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(tipo)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errTipo
						c.Abort("404")
					}
				}
			} else {
				logs.Error(tipo)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errTipo
				c.Abort("404")
			}

			errAspirante := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/"+fmt.Sprintf("%v", inscripcion[0]["PersonaId"]), &aspirante)
			if errAspirante == nil && fmt.Sprintf("%v", aspirante["System"]) != "map[]" {
				if aspirante["Status"] != 404 {
					//resultado de la Identificacion del Aspirante
					var identificacion []map[string]interface{}
					errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+fmt.Sprintf("%v", aspirante["Ente"]), &identificacion)
					if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
						if identificacion[0]["Status"] != 404 {
							aspirante["TipoIdentificacion"] = identificacion[0]["TipoIdentificacion"]
							aspirante["NumeroDocumento"] = identificacion[0]["NumeroIdentificacion"]
							inscripcion[0]["PersonaId"] = aspirante
							resultado = inscripcion[0]
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
					if aspirante["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(aspirante)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errAspirante
						c.Abort("404")
					}
				}
			} else {
				logs.Error(aspirante)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errAspirante
				c.Abort("404")
			}
		} else {
			if inscripcion[0]["Message"] == "Not found resource" {
				c.Data["json"] = nil
			} else {
				logs.Error(inscripcion)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errInscripcion
				c.Abort("404")
			}
		}
	} else {
		logs.Error(inscripcion)
		//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = errInscripcion
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetByIdentificacion ...
// @Title GetByIdentificacion
// @Description get Inscripcion by Identificacion
// @Param	Identificacion		query 	string	true		"Identificación de la persona"
// @Param	ProgramaId		query 	int	true		"Id del programa académico"
// @Param	PeriodoId		query 	int	true		"Id del periodo académico"
// @Success 200 {}
// @Failure 404 not found resource
// @router /identificacion/ [get]
func (c *InscripcionController) GetByIdentificacion() {
	//Ids de la inscripción a consultar
	idStr := c.GetString("Identificacion")
	idPrograma := c.GetString("ProgramaId")
	idPeriodo := c.GetString("PeriodoId")
	var resultado map[string]interface{}
	var identificacion []map[string]interface{}
	errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=NumeroIdentificacion:"+idStr, &identificacion)

	if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
		if identificacion[0]["Status"] != 404 {
			var persona []map[string]interface{}

			errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Ente:"+fmt.Sprintf("%v", identificacion[0]["Ente"].(map[string]interface{})["Id"]), &persona)
			if errPersona == nil && fmt.Sprintf("%v", persona[0]["System"]) != "map[]" {
				if persona[0]["Status"] != 404 {
					var inscripcion []map[string]interface{}

					errInscripcion := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/inscripcion/?query=PersonaId:"+fmt.Sprintf("%v", persona[0]["Id"])+
						",PeriodoId:"+idPeriodo+",ProgramaAcademicoId:"+idPrograma, &inscripcion)
					if errInscripcion == nil && fmt.Sprintf("%v", inscripcion[0]["System"]) != "map[]" {
						if inscripcion[0]["Status"] != 404 {
							//resultado del Estado de inscripcion
							var estado map[string]interface{}
							idEstado := inscripcion[0]["EstadoInscripcionId"].(map[string]interface{})
							//resultado del Tipo de Inscripcion
							var tipo map[string]interface{}
							idTipo := inscripcion[0]["TipoInscripcionId"].(map[string]interface{})
							//resultado del Aspirante
							var aspirante map[string]interface{}

							errEstado := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/estado_inscripcion/"+fmt.Sprintf("%v", idEstado["Id"]), &estado)
							if errEstado == nil && fmt.Sprintf("%v", estado["System"]) != "map[]" {
								if estado["Status"] != 404 {
									inscripcion[0]["EstadoInscripcionId"] = estado
								} else {
									if estado["Message"] == "Not found resource" {
										c.Data["json"] = nil
									} else {
										logs.Error(estado)
										//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = errEstado
										c.Abort("404")
									}
								}
							} else {
								logs.Error(estado)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errEstado
								c.Abort("404")
							}

							errTipo := request.GetJson("http://"+beego.AppConfig.String("InscripcionService")+"/tipo_inscripcion/"+fmt.Sprintf("%v", idTipo["Id"]), &tipo)
							if errTipo == nil && fmt.Sprintf("%v", tipo["System"]) != "map[]" {
								if tipo["Status"] != 404 {
									inscripcion[0]["TipoInscripcionId"] = tipo
								} else {
									if tipo["Message"] == "Not found resource" {
										c.Data["json"] = nil
									} else {
										logs.Error(tipo)
										//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = errTipo
										c.Abort("404")
									}
								}
							} else {
								logs.Error(tipo)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errTipo
								c.Abort("404")
							}

							errAspirante := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/"+fmt.Sprintf("%v", inscripcion[0]["PersonaId"]), &aspirante)
							if errAspirante == nil && fmt.Sprintf("%v", aspirante["System"]) != "map[]" {
								if aspirante["Status"] != 404 {
									//resultado de la Identificacion del Aspirante
									var identificacion []map[string]interface{}
									errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+fmt.Sprintf("%v", aspirante["Ente"]), &identificacion)
									if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
										if identificacion[0]["Status"] != 404 {
											aspirante["TipoIdentificacion"] = identificacion[0]["TipoIdentificacion"]
											aspirante["NumeroDocumento"] = identificacion[0]["NumeroIdentificacion"]
											inscripcion[0]["PersonaId"] = aspirante
											resultado = inscripcion[0]
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
									if aspirante["Message"] == "Not found resource" {
										c.Data["json"] = nil
									} else {
										logs.Error(aspirante)
										//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = errAspirante
										c.Abort("404")
									}
								}
							} else {
								logs.Error(aspirante)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errAspirante
								c.Abort("404")
							}
						} else {
							if inscripcion[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(inscripcion)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errInscripcion
								c.Abort("404")
							}
						}
					} else {
						logs.Error(inscripcion)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errInscripcion
						c.Abort("404")
					}
				} else {
					if persona[0]["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(persona)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errPersona
						c.Abort("404")
					}
				}
			} else {
				logs.Error(persona)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errPersona
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
