package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// PersonaController ...
type PersonaController struct {
	beego.Controller
}

// URLMapping ...
func (c *PersonaController) URLMapping() {
	c.Mapping("ActualizarDatosComplementarios", c.ActualizarDatosComplementarios)
	c.Mapping("ActualizarDatosContacto", c.ActualizarDatosContacto)
	c.Mapping("ActualizarPersona", c.ActualizarPersona)
	c.Mapping("ConsultarDatosComplementarios", c.ConsultarDatosComplementarios)
	c.Mapping("ConsultarDatosContacto", c.ConsultarDatosContacto)
	c.Mapping("ConsultarPersona", c.ConsultarPersona)
	c.Mapping("ConsultarPersonaByUser", c.ConsultarPersonaByUser)
	c.Mapping("GuardarPersona", c.GuardarPersona)
	c.Mapping("GuardarDatosContacto", c.GuardarDatosContacto)
	c.Mapping("GuardarDatosComplementarios", c.GuardarDatosComplementarios)
}

// ActualizarPersona ...
// @Title ActualizarPersona
// @Description Actualizar Informacion Basica Persona
// @Param	body		body 	{}	true		"body for Actualizar Persona content"
// @Success 200 {}
// @Failure 400 the request contains incorrect syntax
// @router /actualizar_persona [put]
func (c *PersonaController) ActualizarPersona() {
	//resultado informacion persona
	var resultado map[string]interface{}
	//persona
	var persona map[string]interface{}
	var personaPut map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		actualizarpersona := map[string]interface{}{
			"Id":              persona["Id"],
			"Ente":            persona["Ente"],
			"PrimerNombre":    persona["PrimerNombre"],
			"SegundoNombre":   persona["SegundoNombre"],
			"PrimerApellido":  persona["PrimerApellido"],
			"SegundoApellido": persona["SegundoApellido"],
			"FechaNacimiento": persona["FechaNacimiento"],
			"Foto":            persona["Foto"],
			"Usuario":         persona["Usuario"],
		}

		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+
			"/persona/"+fmt.Sprintf("%v", persona["Id"]), "PUT", &personaPut, actualizarpersona)
		if errPersona == nil && fmt.Sprintf("%v", personaPut["System"]) != "map[]" && personaPut["Id"] != nil {
			if personaPut["Status"] != 400 {
				//identificacion de la persona
				var identificacion []map[string]interface{}
				var identificacionPut map[string]interface{}

				errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+fmt.Sprintf("%v", persona["Ente"]), &identificacion)
				if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
					if identificacion[0]["Status"] != 404 {
						identificacion[0]["NumeroIdentificacion"] = persona["NumeroIdentificacion"]
						identificacion[0]["TipoIdentificacion"] = persona["TipoIdentificacion"]
						identificacion[0]["Soporte"] = persona["SoporteDocumento"]

						errIdentificacionPut := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/"+
							fmt.Sprintf("%v", identificacion[0]["Id"]), "PUT", &identificacionPut, identificacion[0])
						if errIdentificacionPut == nil && fmt.Sprintf("%v", identificacionPut["System"]) != "map[]" && identificacionPut["Id"] != nil {
							if identificacionPut["Status"] != 400 {
								fmt.Println("La nueva identificacion es:", identificacionPut)
								//estado de la persona
								var estado []map[string]interface{}
								var estadoPut map[string]interface{}
								//genero de la persona
								var genero []map[string]interface{}
								var generoPut map[string]interface{}

								errEstado := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+
									"/persona_estado_civil/?query=Persona:"+fmt.Sprintf("%v", persona["Id"]), &estado)
								if errEstado == nil && fmt.Sprintf("%v", estado[0]["System"]) != "map[]" {
									if estado[0]["Status"] != 404 {
										estado[0]["EstadoCivil"] = persona["EstadoCivil"]

										errEstadoPut := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/"+
											fmt.Sprintf("%v", estado[0]["Id"]), "PUT", &estadoPut, estado[0])
										if errEstadoPut == nil && fmt.Sprintf("%v", estadoPut["System"]) != "map[]" && estadoPut["Id"] != nil {
											if estadoPut["Status"] != 400 {
												fmt.Println("El nuevo estado es:", estadoPut)
											} else {
												logs.Error(errEstadoPut)
												//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
												c.Data["system"] = estadoPut
												c.Abort("400")
											}
										} else {
											logs.Error(errEstadoPut)
											//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
											c.Data["system"] = estadoPut
											c.Abort("400")
										}
									} else {
										if estado[0]["Message"] == "Not found resource" {
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

								errGenero := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+
									"/persona_genero/?query=Persona:"+fmt.Sprintf("%v", persona["Id"]), &genero)
								if errGenero == nil && fmt.Sprintf("%v", genero[0]["System"]) != "map[]" {
									if genero[0]["Status"] != 404 {
										genero[0]["Genero"] = persona["Genero"]

										errGeneroPut := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero/"+
											fmt.Sprintf("%v", genero[0]["Id"]), "PUT", &generoPut, genero[0])
										if errGeneroPut == nil && fmt.Sprintf("%v", generoPut["System"]) != "map[]" && generoPut["Id"] != nil {
											if generoPut["Status"] != 400 {
												fmt.Println("El nuevo estado es:", generoPut)
											} else {
												logs.Error(errGeneroPut)
												//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
												c.Data["system"] = generoPut
												c.Abort("400")
											}
										} else {
											logs.Error(errGeneroPut)
											//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
											c.Data["system"] = generoPut
											c.Abort("400")
										}
									} else {
										if genero[0]["Message"] == "Not found resource" {
											c.Data["json"] = nil
										} else {
											logs.Error(errGenero)
											//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
											c.Data["system"] = genero
											c.Abort("404")
										}
									}
								} else {
									logs.Error(errGenero)
									//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
									c.Data["system"] = genero
									c.Abort("404")
								}

								resultado = persona
								c.Data["json"] = resultado
							} else {
								logs.Error(errIdentificacionPut)
								//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = identificacionPut
								c.Abort("400")
							}
						} else {
							logs.Error(errIdentificacionPut)
							//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
							c.Data["system"] = identificacionPut
							c.Abort("400")
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
			} else {
				logs.Error(errPersona)
				//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = personaPut
				c.Abort("400")
			}
		} else {
			logs.Error(errPersona)
			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = personaPut
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

// ActualizarDatosComplementarios ...
// @Title ActualizarDatosComplementarios
// @Description Actualizar Informacion Datos Complementarios Persona
// @Param	body		body 	{}	true		"body for Actualizar Persona content"
// @Success 200 {}
// @Failure 400 the request contains incorrect syntax
// @router /actualizar_complementarios [put]
func (c *PersonaController) ActualizarDatosComplementarios() {
	//resultado informacion persona
	var resultado map[string]interface{}
	//persona
	var persona map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		var grupoEtnico []map[string]interface{}

		errGrupoEtnico := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/?query=Persona:"+
			fmt.Sprintf("%v", persona["Persona"]), &grupoEtnico)
		if errGrupoEtnico == nil && fmt.Sprintf("%v", grupoEtnico[0]["System"]) != "map[]" {
			if grupoEtnico[0]["Status"] != 404 {
				var grupoEtnicoPut map[string]interface{}
				grupoEtnico[0]["GrupoEtnico"] = persona["GrupoEtnico"]

				errGrupoEtnicoPut := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/"+
					fmt.Sprintf("%v", grupoEtnico[0]["Id"]), "PUT", &grupoEtnicoPut, grupoEtnico[0])
				if errGrupoEtnicoPut == nil && fmt.Sprintf("%v", grupoEtnicoPut["System"]) != "map[]" && grupoEtnicoPut["Id"] != nil {
					if grupoEtnicoPut["Status"] != 400 {
						var grupoSanguineo []map[string]interface{}
						fmt.Println("El grupoEtnico es: " + fmt.Sprintf("%v", grupoEtnicoPut))

						errGrupoSanguineo := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/?query=Persona:"+
							fmt.Sprintf("%v", persona["Persona"]), &grupoSanguineo)
						if errGrupoSanguineo == nil && fmt.Sprintf("%v", grupoSanguineo[0]["System"]) != "map[]" {
							if grupoSanguineo[0]["Status"] != 404 {
								var grupoSanguineoPut map[string]interface{}
								grupoSanguineo[0]["GrupoSanguineo"] = persona["GrupoSanguineo"]
								grupoSanguineo[0]["FactorRh"] = persona["Rh"]

								errGrupoSanguineoPut := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/"+
									fmt.Sprintf("%v", grupoSanguineo[0]["Id"]), "PUT", &grupoSanguineoPut, grupoSanguineo[0])
								if errGrupoSanguineoPut == nil && fmt.Sprintf("%v", grupoSanguineoPut["System"]) != "map[]" && grupoSanguineoPut["Id"] != nil {
									if grupoSanguineoPut["Status"] != 400 {
										var ubicacion []map[string]interface{}
										fmt.Println("El grupoSanguineo es: " + fmt.Sprintf("%v", grupoSanguineoPut))

										errUbicacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/?query=Ente:"+
											fmt.Sprintf("%v", persona["Ente"])+",TipoRelacionUbicacionEnte:1,Activo:true", &ubicacion)
										if errUbicacion == nil && fmt.Sprintf("%v", ubicacion[0]["System"]) != "map[]" {
											if ubicacion[0]["Status"] != 404 {
												var ubicacionPut map[string]interface{}
												ubicacion[0]["Lugar"] = persona["Lugar"].(map[string]interface{})["Id"]

												errUbicacionPut := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/"+
													fmt.Sprintf("%v", ubicacion[0]["Id"]), "PUT", &ubicacionPut, ubicacion[0])
												if errUbicacionPut == nil && fmt.Sprintf("%v", ubicacionPut["System"]) != "map[]" && ubicacionPut["Id"] != nil {
													if ubicacionPut["Status"] != 400 {
														fmt.Println("El ubicacion es: " + fmt.Sprintf("%v", ubicacionPut))
														var personaDiscapacidades []map[string]interface{}
														discapacidades := persona["TipoDiscapacidad"].([]interface{})

														for i := 0; i < len(discapacidades); i++ {
															var encuentraDiscapacidad []map[string]interface{}
															discapacidad := discapacidades[i].(map[string]interface{})

															errDiscapacidad := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/?query=Persona:"+
																fmt.Sprintf("%v", persona["Persona"])+",TipoDiscapacidad.Id:"+
																fmt.Sprintf("%v", discapacidad["Id"]), &encuentraDiscapacidad)
															if errDiscapacidad == nil && fmt.Sprintf("%v", encuentraDiscapacidad[0]["System"]) != "map[]" {
																if encuentraDiscapacidad[0]["Status"] != 404 {
																	if encuentraDiscapacidad[0]["Id"] != nil {
																		fmt.Println("Esta " + fmt.Sprintf("%v", discapacidad["Id"]))

																		if encuentraDiscapacidad[0]["Activo"] == false {
																			var discapacidadPut map[string]interface{}
																			encuentraDiscapacidad[0]["Activo"] = true

																			errDiscapacidadPut := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/"+
																				fmt.Sprintf("%v", encuentraDiscapacidad[0]["Id"]), "PUT", &discapacidadPut, encuentraDiscapacidad[0])
																			if errDiscapacidadPut == nil && fmt.Sprintf("%v", discapacidadPut["System"]) != "map[]" && discapacidadPut["Id"] != nil {
																				if discapacidadPut["Status"] != 400 {
																					fmt.Println("El ajuste es:", discapacidadPut)
																				} else {
																					logs.Error(errDiscapacidadPut)
																					//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
																					c.Data["system"] = discapacidadPut
																					c.Abort("400")
																				}
																			} else {
																				logs.Error(errDiscapacidadPut)
																				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
																				c.Data["system"] = discapacidadPut
																				c.Abort("400")
																			}
																		}
																	} else {
																		fmt.Println("No esta " + fmt.Sprintf("%v", discapacidad["Id"]))

																		var discapacidadPost map[string]interface{}
																		//si no tiene la discapacidad se agrega
																		nuevadiscapacidad := map[string]interface{}{
																			"Persona":          map[string]interface{}{"Id": persona["Persona"]},
																			"Activo":           true,
																			"TipoDiscapacidad": discapacidad,
																		}

																		errDiscapacidadPost := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad", "POST", &discapacidadPost, nuevadiscapacidad)
																		if errDiscapacidadPost == nil && fmt.Sprintf("%v", discapacidadPost["System"]) != "map[]" && discapacidadPost["Id"] != nil {
																			if discapacidadPost["Status"] != 400 {
																				fmt.Println("El nueva discapacidad es: " + fmt.Sprintf("%v", discapacidadPost))
																			} else {
																				logs.Error(errDiscapacidadPost)
																				//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
																				c.Data["system"] = discapacidadPost
																				c.Abort("400")
																			}
																		} else {
																			logs.Error(errDiscapacidadPost)
																			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
																			c.Data["system"] = discapacidadPost
																			c.Abort("400")
																		}
																	}
																} else {
																	if encuentraDiscapacidad[0]["Message"] == "Not found resource" {
																		c.Data["json"] = nil
																	} else {
																		logs.Error(errDiscapacidad)
																		//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
																		c.Data["system"] = encuentraDiscapacidad
																		c.Abort("404")
																	}
																}
															} else {
																logs.Error(errDiscapacidad)
																//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
																c.Data["system"] = encuentraDiscapacidad
																c.Abort("404")
															}
														}

														errDiscapacidadB := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/?query=Persona:"+
															fmt.Sprintf("%v", persona["Persona"]), &personaDiscapacidades)
														if errDiscapacidadB == nil && fmt.Sprintf("%v", personaDiscapacidades[0]["System"]) != "map[]" {
															if personaDiscapacidades[0]["Status"] != 404 {
																for j := 0; j < len(personaDiscapacidades); j++ {
																	activar := false
																	personaDiscapacidad := personaDiscapacidades[j]
																	for k := 0; k < len(discapacidades); k++ {
																		d := personaDiscapacidad["TipoDiscapacidad"].(map[string]interface{})
																		discapacidad := discapacidades[k].(map[string]interface{})
																		if d["Id"] == discapacidad["Id"] {
																			activar = true
																		}
																	}
																	if activar == false {
																		fmt.Println("No esta " + fmt.Sprintf("%v", personaDiscapacidades[j]))
																		var discapacidadPutB map[string]interface{}
																		personaDiscapacidad["Activo"] = false

																		errDiscapacidadPutB := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/"+
																			fmt.Sprintf("%v", personaDiscapacidad["Id"]), "PUT", &discapacidadPutB, personaDiscapacidad)
																		if errDiscapacidadPutB == nil && fmt.Sprintf("%v", discapacidadPutB["System"]) != "map[]" && discapacidadPutB["Id"] != nil {
																			if discapacidadPutB["Status"] != 400 {
																				fmt.Println("El ajuste es:", discapacidadPutB)
																			} else {
																				logs.Error(errDiscapacidadPutB)
																				//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
																				c.Data["system"] = discapacidadPutB
																				c.Abort("400")
																			}
																		} else {
																			logs.Error(errDiscapacidadPutB)
																			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
																			c.Data["system"] = discapacidadPutB
																			c.Abort("400")
																		}
																	}
																}
															} else {
																if personaDiscapacidades[0]["Message"] == "Not found resource" {
																	c.Data["json"] = nil
																} else {
																	logs.Error(errDiscapacidadB)
																	//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
																	c.Data["system"] = personaDiscapacidades
																	c.Abort("404")
																}
															}
														} else {
															logs.Error(errDiscapacidadB)
															//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
															c.Data["system"] = personaDiscapacidades
															c.Abort("404")
														}

														resultado = persona
														c.Data["json"] = resultado

													} else {
														logs.Error(errUbicacionPut)
														//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
														c.Data["system"] = ubicacionPut
														c.Abort("400")
													}
												} else {
													logs.Error(errUbicacionPut)
													//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
													c.Data["system"] = ubicacionPut
													c.Abort("400")
												}
											} else {
												if ubicacion[0]["Message"] == "Not found resource" {
													c.Data["json"] = nil
												} else {
													logs.Error(errUbicacion)
													//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
													c.Data["system"] = ubicacion
													c.Abort("404")
												}
											}
										} else {
											logs.Error(errUbicacion)
											//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
											c.Data["system"] = ubicacion
											c.Abort("404")
										}
									} else {
										logs.Error(errGrupoSanguineoPut)
										//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = grupoSanguineoPut
										c.Abort("400")
									}
								} else {
									logs.Error(errGrupoSanguineoPut)
									//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
									c.Data["system"] = grupoSanguineoPut
									c.Abort("400")
								}
							} else {
								if grupoSanguineo[0]["Message"] == "Not found resource" {
									c.Data["json"] = nil
								} else {
									logs.Error(errGrupoSanguineo)
									//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
									c.Data["system"] = grupoSanguineo
									c.Abort("404")
								}
							}
						} else {
							logs.Error(errGrupoSanguineo)
							//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
							c.Data["system"] = grupoSanguineo
							c.Abort("404")
						}
					} else {
						logs.Error(errGrupoEtnicoPut)
						//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = grupoEtnicoPut
						c.Abort("400")
					}
				} else {
					logs.Error(errGrupoEtnicoPut)
					//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = grupoEtnicoPut
					c.Abort("400")
				}
			} else {
				if grupoEtnico[0]["Message"] == "Not found resource" {
					c.Data["json"] = nil
				} else {
					logs.Error(errGrupoEtnico)
					//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = grupoEtnico
					c.Abort("404")
				}
			}
		} else {
			logs.Error(errGrupoEtnico)
			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = grupoEtnico
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

// ActualizarDatosContacto ...
// @Title ActualizarDatosContacto
// @Description Actualizar Informacion Datos Contacto Persona
// @Param	body		body 	{}	true		"body for Actualizar Persona content"
// @Success 200 {}
// @Failure 400 the request contains incorrect syntax
// @router /actualizar_contacto [put]
func (c *PersonaController) ActualizarDatosContacto() {
	//resultado informacion contacto
	var resultado map[string]interface{}
	//persona
	var persona map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		var ubicacionEnte []map[string]interface{}
		contactos := persona["ContactoEnte"].([]interface{})
		for i := 0; i < len(contactos); i++ {
			var contactoEnte []map[string]interface{}
			contacto := contactos[i].(map[string]interface{})

			errContacto := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente/?query=Ente.Id:"+
				fmt.Sprintf("%v", persona["Ente"])+",TipoContacto.Id:"+
				fmt.Sprintf("%v", contacto["TipoContacto"].(map[string]interface{})["Id"]), &contactoEnte)
			if errContacto == nil && fmt.Sprintf("%v", contactoEnte[0]["System"]) != "map[]" {
				if contactoEnte[0]["Status"] != 404 {
					var contactoEntePut map[string]interface{}
					contactoEnte[0]["Valor"] = contacto["Valor"]

					errContactoEntePut := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente/"+
						fmt.Sprintf("%v", contactoEnte[0]["Id"]), "PUT", &contactoEntePut, contactoEnte[0])
					if errContactoEntePut == nil && fmt.Sprintf("%v", contactoEntePut["System"]) != "map[]" && contactoEntePut["Id"] != nil {
						if contactoEntePut["Status"] != 400 {
							fmt.Println("El ajuste es:", contactoEntePut)
						} else {
							logs.Error(errContactoEntePut)
							//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
							c.Data["system"] = contactoEntePut
							c.Abort("400")
						}
					} else {
						logs.Error(errContactoEntePut)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = contactoEntePut
						c.Abort("400")
					}
				} else {
					if contactoEnte[0]["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(contactoEnte)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errContacto
						c.Abort("404")
					}
				}
			} else {
				logs.Error(contactoEnte)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errContacto
				c.Abort("404")
			}
		}

		errUbicacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/?query=Ente:"+
			fmt.Sprintf("%v", persona["Ente"])+",TipoRelacionUbicacionEnte:2,Activo:true", &ubicacionEnte)
		if errUbicacion == nil && fmt.Sprintf("%v", ubicacionEnte[0]["System"]) != "map[]" {
			if ubicacionEnte[0]["Status"] != 404 {
				//actualización ubicaciones
				var ubicacionPut map[string]interface{}
				ubicacion := persona["UbicacionEnte"].(map[string]interface{})
				lugar := ubicacion["Lugar"].(map[string]interface{})
				ubicacionEnte[0]["Lugar"] = lugar["Id"]

				errUbicacionPut := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/"+
					fmt.Sprintf("%v", ubicacionEnte[0]["Id"]), "PUT", &ubicacionPut, ubicacionEnte[0])
				if errUbicacionPut == nil && fmt.Sprintf("%v", ubicacionPut["System"]) != "map[]" && ubicacionPut["Id"] != nil {
					if ubicacionPut["Status"] != 400 {
						var atributos []interface{}
						ubicacionPersona := persona["UbicacionEnte"].(map[string]interface{})
						atributos = ubicacionPersona["Atributos"].([]interface{})

						for i := 0; i < len(atributos); i++ {
							var atributosEnte []map[string]interface{}
							atributo := atributos[i].(map[string]interface{})

							errAtributos := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion/?query=UbicacionEnte.Id:"+
								fmt.Sprintf("%v", ubicacionEnte[0]["Id"])+",AtributoUbicacion.Id:"+
								fmt.Sprintf("%v", atributo["AtributoUbicacion"].(map[string]interface{})["Id"]), &atributosEnte)
							if errAtributos == nil && fmt.Sprintf("%v", atributosEnte[0]["System"]) != "map[]" {
								if atributosEnte[0]["Status"] != 404 {
									var atributoPut map[string]interface{}
									atributosEnte[0]["Valor"] = atributo["Valor"]

									errAtributoPut := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion/"+
										fmt.Sprintf("%v", atributosEnte[0]["Id"]), "PUT", &atributoPut, atributosEnte[0])
									if errAtributoPut == nil && fmt.Sprintf("%v", atributoPut["System"]) != "map[]" && atributoPut["Id"] != nil {
										if atributoPut["Status"] != 400 {
											fmt.Println("Ajuste del atributo de ubicacion: " + fmt.Sprintf("%v", atributoPut))
										} else {
											logs.Error(errAtributoPut)
											//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
											c.Data["system"] = atributoPut
											c.Abort("400")
										}
									} else {
										logs.Error(errAtributoPut)
										//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = atributoPut
										c.Abort("400")
									}
								} else {
									if atributosEnte[0]["Message"] == "Not found resource" {
										c.Data["json"] = nil
									} else {
										logs.Error(atributosEnte)
										//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = errAtributos
										c.Abort("404")
									}
								}
							} else {
								logs.Error(atributosEnte)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errAtributos
								c.Abort("404")
							}
						}
						resultado = persona
						c.Data["json"] = resultado
					} else {
						logs.Error(errUbicacionPut)
						//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = ubicacionPut
						c.Abort("400")
					}
				} else {
					logs.Error(errUbicacionPut)
					//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = ubicacionPut
					c.Abort("400")
				}
			} else {
				if ubicacionEnte[0]["Message"] == "Not found resource" {
					c.Data["json"] = nil
				} else {
					logs.Error(ubicacionEnte)
					//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = errUbicacion
					c.Abort("404")
				}
			}

		} else {
			logs.Error(ubicacionEnte)
			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = errUbicacion
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

// ConsultarDatosContacto ...
// @Title ConsultarDatosContacto
// @Description get ConsultarDatosContacto by id
// @Param	ente_id	path	int	true	"Id del ente"
// @Success 200 {}
// @Failure 404 not found resource
// @router /consultar_contacto/:ente_id [get]
func (c *PersonaController) ConsultarDatosContacto() {
	//Id de la persona
	idStr := c.Ctx.Input.Param(":ente_id")
	fmt.Println("El id es: " + idStr)
	//resultado datos complementarios persona
	var resultado map[string]interface{}
	var persona []map[string]interface{}

	errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Ente:"+idStr, &persona)
	if errPersona == nil && fmt.Sprintf("%v", persona[0]["System"]) != "map[]" {
		if persona[0]["Status"] != 404 {
			var contactoEnte []map[string]interface{}
			resultado = map[string]interface{}{"Ente": persona[0]["Ente"], "Persona": persona[0]["Id"]}

			errContacto := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente/?query=Ente.Id:"+idStr+"&fields=Id,TipoContacto,Valor", &contactoEnte)
			if errContacto == nil && fmt.Sprintf("%v", contactoEnte[0]["System"]) != "map[]" {
				if contactoEnte[0]["Status"] != 404 {
					var ubicacionEnte []map[string]interface{}
					resultado["ContactoEnte"] = contactoEnte

					errUbicacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/?query=Ente:"+idStr+
						",TipoRelacionUbicacionEnte:2,Activo:true&fields=Id,TipoRelacionUbicacionEnte,Lugar", &ubicacionEnte)
					if errUbicacion == nil && fmt.Sprintf("%v", ubicacionEnte[0]["System"]) != "map[]" {
						if ubicacionEnte[0]["Status"] != 404 {
							var atributosEnte []map[string]interface{}

							errAtributos := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion/?query=UbicacionEnte.Id:"+
								fmt.Sprintf("%v", ubicacionEnte[0]["Id"])+"&fields=Id,AtributoUbicacion,Valor", &atributosEnte)

							if errAtributos == nil && fmt.Sprintf("%v", atributosEnte) != "[]" {
								if errAtributos == nil && fmt.Sprintf("%v", atributosEnte[0]["System"]) != "map[]" {
									if atributosEnte[0]["Status"] != 404 {
										var lugar map[string]interface{}
										ubicacionEnte[0]["Atributos"] = atributosEnte

										errLugar := request.GetJson("http://"+beego.AppConfig.String("UbicacionesService")+"/relacion_lugares/jerarquia_lugar/"+
											fmt.Sprintf("%v", ubicacionEnte[0]["Lugar"]), &lugar)
										if errLugar == nil && fmt.Sprintf("%v", lugar["System"]) != "map[]" {
											if lugar["Status"] != 404 {
												ubicacionEnte[0]["Lugar"] = lugar
												resultado["UbicacionEnte"] = ubicacionEnte[0]
												c.Data["json"] = resultado
											} else {
												if lugar["Message"] == "Not found resource" {
													c.Data["json"] = nil
												} else {
													logs.Error(lugar)
													//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
													c.Data["system"] = errLugar
													c.Abort("404")
												}
											}
										} else {
											logs.Error(lugar)
											//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
											c.Data["system"] = errLugar
											c.Abort("404")
										}
									} else {
										if atributosEnte[0]["Message"] == "Not found resource" {
											c.Data["json"] = nil
										} else {
											logs.Error(atributosEnte)
											//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
											c.Data["system"] = errAtributos
											c.Abort("404")
										}
									}
								} else {
									logs.Error(atributosEnte)
									//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
									c.Data["system"] = errAtributos
									c.Abort("404")
								}
							} else {
								if errAtributos == nil && fmt.Sprintf("%v", atributosEnte) == "[]" {
									fmt.Println("El error esta aqui")
									atributosEnte = append(atributosEnte, map[string]interface{}{})
									c.Data["json"] = atributosEnte
								} else {
									logs.Error(atributosEnte)
									//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
									c.Data["system"] = errAtributos
									c.Abort("404")
								}
							}
						} else {
							if ubicacionEnte[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(ubicacionEnte)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errUbicacion
								c.Abort("404")
							}
						}
					} else {
						logs.Error(ubicacionEnte)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errUbicacion
						c.Abort("404")
					}
				} else {
					logs.Error(contactoEnte)
					//c.Data["Development"] = map[string]interface{}{"Code": "404", "Body": "", "Type": "error"}
					c.Data["system"] = errContacto
					c.Abort("404")
				}
			} else {
        			logs.Error(contactoEnte)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = errContacto
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
	c.ServeJSON()
}

// ConsultarDatosComplementarios ...
// @Title ConsultarDatosComplementarios
// @Description get ConsultarDatosComplementarios by id
// @Param	ente_id	path	int	true	"Id del ente"
// @Success 200 {}
// @Failure 404 not found resource
// @router /consultar_complementarios/:ente_id [get]
func (c *PersonaController) ConsultarDatosComplementarios() {
	//Id de la persona
	idStr := c.Ctx.Input.Param(":ente_id")
	fmt.Println("El id es: " + idStr)
	//resultado datos complementarios persona
	var resultado map[string]interface{}
	var persona []map[string]interface{}

	errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Ente:"+idStr, &persona)
	if errPersona == nil && fmt.Sprintf("%v", persona[0]) != "map[]" {
		if persona[0]["Status"] != 404 {
			var grupoEtnico []map[string]interface{}
			resultado = map[string]interface{}{"Ente": persona[0]["Ente"], "Persona": persona[0]["Id"]}

			errGrupoEtnico := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/?query=Persona:"+
				fmt.Sprintf("%v", persona[0]["Id"]), &grupoEtnico)
			if errGrupoEtnico == nil && fmt.Sprintf("%v", grupoEtnico[0]) != "map[]" {
				if grupoEtnico[0]["Status"] != 404 {
					var grupoSanguineo []map[string]interface{}
					resultado["GrupoEtnico"] = grupoEtnico[0]["GrupoEtnico"]

					errGrupoSanguineo := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/?query=Persona:"+
						fmt.Sprintf("%v", persona[0]["Id"]), &grupoSanguineo)
					if errGrupoSanguineo == nil && fmt.Sprintf("%v", grupoSanguineo[0]) != "map[]" {
						if grupoSanguineo[0]["Status"] != 404 {
							var discapacidades []map[string]interface{}
							resultado["GrupoSanguineo"] = grupoSanguineo[0]["GrupoSanguineo"]
							resultado["Rh"] = grupoSanguineo[0]["FactorRh"]

							errDiscapacidad := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad/?query=Persona:"+
								fmt.Sprintf("%v", persona[0]["Id"])+",Activo:true", &discapacidades)
							if errDiscapacidad == nil && fmt.Sprintf("%v", discapacidades[0]) != "map[]" {
								if discapacidades[0]["Status"] != 404 {
									var tipoDiscapacidad []map[string]interface{}
									var ubicacionEnte []map[string]interface{}
									for i := 0; i < len(discapacidades); i++ {
										if len(discapacidades) > 0 {
											discapacidad := discapacidades[i]["TipoDiscapacidad"].(map[string]interface{})
											tipoDiscapacidad = append(tipoDiscapacidad, discapacidad)
										}
									}
									resultado["TipoDiscapacidad"] = tipoDiscapacidad

									errUbicacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente/?query=Ente:"+idStr+
										",TipoRelacionUbicacionEnte:1,Activo:true&fields=Id,TipoRelacionUbicacionEnte,Lugar", &ubicacionEnte)
									if errUbicacion == nil && fmt.Sprintf("%v", ubicacionEnte[0]) != "map[]" {
										if ubicacionEnte[0]["Status"] != 404 {
											var lugar map[string]interface{}

											errLugar := request.GetJson("http://"+beego.AppConfig.String("UbicacionesService")+"/relacion_lugares/jerarquia_lugar/"+
												fmt.Sprintf("%v", ubicacionEnte[0]["Lugar"]), &lugar)
											if errLugar == nil && fmt.Sprintf("%v", lugar["System"]) != "map[]" {
												if lugar["Status"] != 404 {
													ubicacionEnte[0]["Lugar"] = lugar
													resultado["Lugar"] = ubicacionEnte[0]
													c.Data["json"] = resultado
												} else {
													if lugar["Message"] == "Not found resource" {
														c.Data["json"] = nil
													} else {
														logs.Error(lugar)
														//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
														c.Data["system"] = errLugar
														c.Abort("404")
													}
												}
											} else {
												logs.Error(lugar)
												//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
												c.Data["system"] = errLugar
												c.Abort("404")
											}
										} else {
											if ubicacionEnte[0]["Message"] == "Not found resource" {
												c.Data["json"] = nil
											} else {
												logs.Error(ubicacionEnte)
												//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
												c.Data["system"] = errUbicacion
												c.Abort("404")
											}
										}
									} else {
										logs.Error(ubicacionEnte)
										//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = errUbicacion
										c.Abort("404")
									}
								} else {
									if discapacidades[0]["Message"] == "Not found resource" {
										c.Data["json"] = nil
									} else {
										logs.Error(discapacidades)
										//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = errDiscapacidad
										c.Abort("404")
									}
								}
							} else {
								logs.Error(discapacidades)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errDiscapacidad
								c.Abort("404")
							}
						} else {
							if grupoSanguineo[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(grupoSanguineo)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errGrupoSanguineo
								c.Abort("404")
							}
						}
					} else {
						logs.Error(grupoSanguineo)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errGrupoSanguineo
						c.Abort("404")
					}
				} else {
					if grupoEtnico[0]["Message"] == "Not found resource" {
						c.Data["json"] = nil
					} else {
						logs.Error(grupoEtnico)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errGrupoEtnico
						c.Abort("404")
					}
				}
			} else {
				if errGrupoEtnico == nil && fmt.Sprintf("%v", grupoEtnico[0]) == "map[]" {
					fmt.Println("El error esta aqui")
					c.Data["json"] = grupoEtnico
				} else {
					logs.Error(grupoEtnico)
					//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = errGrupoEtnico
					c.Abort("404")
				}
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
	c.ServeJSON()
}

// ConsultarPersona ...
// @Title ConsultarPersona
// @Description get ConsultaPersona by id
// @Param	ente_id	path	int	true	"Id del ente"
// @Success 200 {}
// @Failure 404 not found resource
// @router /consultar_persona/:ente_id [get]
func (c *PersonaController) ConsultarPersona() {
	//Id de la persona
	idStr := c.Ctx.Input.Param(":ente_id")
	fmt.Println("El id es: " + idStr)
	//resultado informacion basica persona
	var resultado map[string]interface{}
	var persona []map[string]interface{}

	errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Ente:"+idStr, &persona)
	if errPersona == nil && fmt.Sprintf("%v", persona[0]) != "map[]" {
		if persona[0]["Status"] != 404 {
			var identificacion []map[string]interface{}

			errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+idStr, &identificacion)
			if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
				if identificacion[0]["Status"] != 404 {
					var estado []map[string]interface{}
					var genero []map[string]interface{}

					resultado = persona[0]
					resultado["NumeroIdentificacion"] = identificacion[0]["NumeroIdentificacion"]
					resultado["TipoIdentificacion"] = identificacion[0]["TipoIdentificacion"]
					resultado["SoporteDocumento"] = identificacion[0]["Soporte"]

					errEstado := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/?query=Persona:"+
						fmt.Sprintf("%v", persona[0]["Id"]), &estado)
					if errEstado == nil && fmt.Sprintf("%v", estado[0]) != "map[]" {
						if estado[0]["Status"] != 404 {
							resultado["EstadoCivil"] = estado[0]["EstadoCivil"]
						} else {
							if estado[0]["Message"] == "Not found resource" {
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

					errGenero := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero/?query=Persona:"+
						fmt.Sprintf("%v", persona[0]["Id"]), &genero)
					if errGenero == nil && fmt.Sprintf("%v", genero[0]) != "map[]" {
						if genero[0]["Status"] != 404 {
							resultado["Genero"] = genero[0]["Genero"]
						} else {
							if genero[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(genero)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errGenero
								c.Abort("404")
							}
						}
					} else {
						logs.Error(genero)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errGenero
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
	c.ServeJSON()
}

// ConsultarPersonaByUser ...
// @Title ConsultarPersonaByUser
// @Description get ConsultarPersonaByUser by username
// @Param	User	query	string	true	"Usuario de la persona"
// @Success 200 {}
// @Failure 404 not found resource
// @router /consultar_persona/ [get]
func (c *PersonaController) ConsultarPersonaByUser() {
	//Usuario de la persona
	user := c.GetString("User")
	fmt.Println("El usuario es: " + user)
	//resultado informacion basica persona
	var resultado map[string]interface{}
	var persona []map[string]interface{}

	errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/?query=Usuario:"+user, &persona)
	fmt.Println(errPersona)
	fmt.Println(persona)
	if errPersona == nil && fmt.Sprintf("%v", persona[0]) != "map[]" {
		if persona[0]["Status"] != 404 {
			var identificacion []map[string]interface{}

			errIdentificacion := request.GetJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion/?query=Ente:"+
				fmt.Sprintf("%v", persona[0]["Ente"]), &identificacion)
			if errIdentificacion == nil && fmt.Sprintf("%v", identificacion[0]["System"]) != "map[]" {
				if identificacion[0]["Status"] != 404 {
					var estado []map[string]interface{}
					var genero []map[string]interface{}

					resultado = persona[0]
					resultado["NumeroIdentificacion"] = identificacion[0]["NumeroIdentificacion"]
					resultado["TipoIdentificacion"] = identificacion[0]["TipoIdentificacion"]
					resultado["SoporteDocumento"] = identificacion[0]["Soporte"]

					errEstado := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/?query=Persona:"+
						fmt.Sprintf("%v", persona[0]["Id"]), &estado)
					if errEstado == nil && fmt.Sprintf("%v", estado[0]) != "map[]" {
						if estado[0]["Status"] != 404 {
							resultado["EstadoCivil"] = estado[0]["EstadoCivil"]
						} else {
							if estado[0]["Message"] == "Not found resource" {
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

					errGenero := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero/?query=Persona:"+
						fmt.Sprintf("%v", persona[0]["Id"]), &genero)
					if errGenero == nil && fmt.Sprintf("%v", genero[0]) != "map[]" {
						if genero[0]["Status"] != 404 {
							resultado["Genero"] = genero[0]["Genero"]
						} else {
							if genero[0]["Message"] == "Not found resource" {
								c.Data["json"] = nil
							} else {
								logs.Error(genero)
								//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = errGenero
								c.Abort("404")
							}
						}
					} else {
						logs.Error(genero)
						//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = errGenero
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
	c.ServeJSON()
}

// GuardarPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	{}	true		"body for Guardar Persona content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router /guardar_persona [post]
func (c *PersonaController) GuardarPersona() {
	//resultado solicitud de descuento
	var resultado map[string]interface{}
	//solicitud de descuento
	var persona map[string]interface{}
	var personaPost map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		guardarpersona := map[string]interface{}{
			"PrimerNombre":    persona["PrimerNombre"],
			"SegundoNombre":   persona["SegundoNombre"],
			"PrimerApellido":  persona["PrimerApellido"],
			"SegundoApellido": persona["SegundoApellido"],
			"FechaNacimiento": persona["FechaNacimiento"],
			"Foto":            persona["Foto"],
			"Usuario":         persona["Usuario"],
		}

		errPersona := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona", "POST", &personaPost, guardarpersona)
		if errPersona == nil && fmt.Sprintf("%v", personaPost["System"]) != "map[]" && personaPost["Id"] != nil {
			if personaPost["Status"] != 400 {
				//identificacion
				var identificacion map[string]interface{}

				identificacionpersona := map[string]interface{}{
					"NumeroIdentificacion": persona["NumeroIdentificacion"],
					"TipoIdentificacion":   persona["TipoIdentificacion"],
					"Soporte":              persona["SoporteDocumento"],
					"Ente":                 map[string]interface{}{"Id": personaPost["Ente"]},
				}

				errIdentificacion := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion", "POST", &identificacion, identificacionpersona)
				if errIdentificacion == nil && fmt.Sprintf("%v", identificacion["System"]) != "map[]" && identificacion["Id"] != nil {
					if identificacion["Status"] != 400 {
						var estado map[string]interface{}

						estadopersona := map[string]interface{}{
							"EstadoCivil": persona["EstadoCivil"],
							"Persona":     personaPost,
						}

						errEstado := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil", "POST", &estado, estadopersona)
						if errEstado == nil && fmt.Sprintf("%v", estado["System"]) != "map[]" && estado["Id"] != nil {
							if estado["Status"] != 400 {
								var genero map[string]interface{}

								generopersona := map[string]interface{}{
									"Genero":  persona["Genero"],
									"Persona": personaPost,
								}

								errGenero := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero", "POST", &genero, generopersona)
								if errGenero == nil && fmt.Sprintf("%v", genero["System"]) != "map[]" && genero["Id"] != nil {
									if genero["Status"] != 400 {

										resultado = personaPost
										resultado["NumeroIdentificacion"] = identificacion["NumeroIdentificacion"]
										resultado["TipoIdentificacion"] = identificacion["TipoIdentificacion"]
										resultado["SoporteDocumento"] = identificacion["Soporte"]
										resultado["EstadoCivil"] = estado["EstadoCivil"]
										resultado["Genero"] = genero["Genero"]
										c.Data["json"] = resultado

									} else {
										//resultado solicitud de descuento
										var resultado2 map[string]interface{}
										request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil/%.f", estado["Id"]), "DELETE", &resultado2, nil)
										request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/identificacion/%.f", identificacion["Id"]), "DELETE", &resultado2, nil)
										request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("PersonaService")+"/persona/%.f", personaPost["Id"]), "DELETE", &resultado2, nil)
										request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/ente/%.f", personaPost["Ente"]), "DELETE", &resultado2, nil)
										logs.Error(errGenero)
										//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = genero
										c.Abort("400")
									}
								} else {
									logs.Error(errGenero)
									//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
									c.Data["system"] = genero
									c.Abort("400")
								}
							} else {
								//resultado solicitud de descuento
								var resultado2 map[string]interface{}
								request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/identificacion/%.f", identificacion["Id"]), "DELETE", &resultado2, nil)
								request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("PersonaService")+"/persona/%.f", personaPost["Id"]), "DELETE", &resultado2, nil)
								request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/ente/%.f", personaPost["Ente"]), "DELETE", &resultado2, nil)
								logs.Error(errEstado)
								//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
								c.Data["system"] = estado
								c.Abort("400")
							}
						} else {
							logs.Error(errEstado)
							//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
							c.Data["system"] = estado
							c.Abort("400")
						}
					} else {
						//resultado solicitud de descuento
						var resultado2 map[string]interface{}
						request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("PersonaService")+"/persona/%.f", personaPost["Id"]), "DELETE", &resultado2, nil)
						request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/ente/%.f", personaPost["Ente"]), "DELETE", &resultado2, nil)
						logs.Error(errIdentificacion)
						//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = identificacion
						c.Abort("400")
					}
				} else {
					logs.Error(errIdentificacion)
					//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = identificacion
					c.Abort("400")
				}
			} else {
				logs.Error(errPersona)
				//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = personaPost
				c.Abort("400")
			}
		} else {
			logs.Error(errPersona)
			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = personaPost
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

// GuardarDatosContacto ...
// @Title GuardarDatosContacto
// @Description Guardar Datos Contacto Persona
// @Param	body		body 	{}	true		"body for Guardar Datos Contacto Persona content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router /guardar_contacto [post]
func (c *PersonaController) GuardarDatosContacto() {
	//resultado solicitud de descuento
	var resultado map[string]interface{}
	//solicitud de descuento
	var persona map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		contactos := persona["ContactoEnte"].([]interface{})
		var ubicacionPost map[string]interface{}

		for i := 0; i < len(contactos); i++ {
			var contactoPost map[string]interface{}
			contacto := contactos[i].(map[string]interface{})
			contactoEnte := map[string]interface{}{
				"Ente":         map[string]interface{}{"Id": persona["Ente"]},
				"TipoContacto": contacto["TipoContacto"].(map[string]interface{}),
				"Valor":        contacto["Valor"],
			}

			errContactoPost := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente", "POST", &contactoPost, contactoEnte)
			if errContactoPost == nil && fmt.Sprintf("%v", contactoPost["System"]) != "map[]" && contactoPost["Id"] != nil {
				if contactoPost["Status"] != 400 {
					fmt.Println("Nuevo dato de contacto:", contactoPost)
				} else {
					logs.Error(errContactoPost)
					//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = contactoPost
					c.Abort("400")
				}
			} else {
				logs.Error(errContactoPost)
				//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = contactoPost
				c.Abort("400")
			}
		}

		ubicacionEnte := map[string]interface{}{
			"Activo":                    true,
			"Ente":                      map[string]interface{}{"Id": persona["Ente"]},
			"Lugar":                     persona["UbicacionEnte"].(map[string]interface{})["Lugar"].(map[string]interface{})["Id"].(float64),
			"TipoRelacionUbicacionEnte": map[string]interface{}{"Id": 2},
		}

		errUbicacionPost := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente", "POST", &ubicacionPost, ubicacionEnte)
		if errUbicacionPost == nil && fmt.Sprintf("%v", ubicacionPost["System"]) != "map[]" && ubicacionPost["Id"] != nil {
			if ubicacionPost["Status"] != 400 {
				fmt.Println("Nueva ubicacion:", ubicacionPost)
				ubicacionPersona := persona["UbicacionEnte"].(map[string]interface{})
				atributos := ubicacionPersona["Atributos"].([]interface{})

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

				resultado = persona

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

		c.Data["json"] = resultado
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GuardarDatosComplementarios ...
// @Title GuardarDatosComplementarios
// @Description Guardar Datos Complementarios Persona
// @Param	body		body 	{}	true		"body for Guardar Datos Complementarios Persona content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router /guardar_complementarios [post]
func (c *PersonaController) GuardarDatosComplementarios() {
	//resultado solicitud de descuento
	var resultado map[string]interface{}
	//solicitud de descuento
	var persona map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		var grupoEtnicoPost map[string]interface{}
		grupoEtnico := map[string]interface{}{
			"GrupoEtnico": persona["GrupoEtnico"],
			"Persona":     map[string]interface{}{"Id": persona["Persona"].(float64)},
		}

		errGrupoEtnicoPost := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico", "POST", &grupoEtnicoPost, grupoEtnico)
		if errGrupoEtnicoPost == nil && fmt.Sprintf("%v", grupoEtnicoPost["System"]) != "map[]" && grupoEtnicoPost["Id"] != nil {
			if grupoEtnicoPost["Status"] != 400 {
				var grupoSanguineoPost map[string]interface{}
				fmt.Println("Grupo etnico: " + fmt.Sprintf("%v", grupoEtnicoPost))
				grupoSanguineo := map[string]interface{}{
					"FactorRh":       persona["Rh"],
					"GrupoSanguineo": persona["GrupoSanguineo"],
					"Persona":        map[string]interface{}{"Id": persona["Persona"].(float64)},
				}

				errGrupoSanguineoPost := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona", "POST", &grupoSanguineoPost, grupoSanguineo)
				if errGrupoSanguineoPost == nil && fmt.Sprintf("%v", grupoSanguineoPost["System"]) != "map[]" && grupoSanguineoPost["Id"] != nil {
					if grupoSanguineoPost["Status"] != 400 {
						var ubicacionPost map[string]interface{}
						fmt.Println("Grupo sanguineo: " + fmt.Sprintf("%v", grupoSanguineoPost))
						ubicacionEnte := map[string]interface{}{
							"Activo":                    true,
							"Ente":                      map[string]interface{}{"Id": persona["Ente"]},
							"Lugar":                     persona["Lugar"].(map[string]interface{})["Lugar"].(map[string]interface{})["Id"].(float64),
							"TipoRelacionUbicacionEnte": map[string]interface{}{"Id": 1},
						}

						errUbicacionPost := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/ubicacion_ente", "POST", &ubicacionPost, ubicacionEnte)
						if errUbicacionPost == nil && fmt.Sprintf("%v", ubicacionPost["System"]) != "map[]" && ubicacionPost["Id"] != nil {
							if ubicacionPost["Status"] != 400 {
								discapacidades := persona["TipoDiscapacidad"].([]interface{})
								fmt.Println("Nueva ubicacion:" + fmt.Sprintf("%v", ubicacionPost))

								for i := 0; i < len(discapacidades); i++ {
									var discapacidadPost map[string]interface{}
									discapacidad := discapacidades[i].(map[string]interface{})
									nuevadiscapacidad := map[string]interface{}{
										"Persona":          map[string]interface{}{"Id": persona["Persona"]},
										"Activo":           true,
										"TipoDiscapacidad": discapacidad,
									}

									errDiscapacidadPost := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_tipo_discapacidad", "POST", &discapacidadPost, nuevadiscapacidad)
									if errDiscapacidadPost == nil && fmt.Sprintf("%v", discapacidadPost["System"]) != "map[]" && discapacidadPost["Id"] != nil {
										if discapacidadPost["Status"] != 400 {
											fmt.Println("El nueva discapacidad es: " + fmt.Sprintf("%v", discapacidadPost))
										} else {
											logs.Error(errDiscapacidadPost)
											//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
											c.Data["system"] = discapacidadPost
											c.Abort("400")
										}
									} else {
										logs.Error(errDiscapacidadPost)
										//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
										c.Data["system"] = discapacidadPost
										c.Abort("400")
									}
								}

								resultado = persona
								c.Data["json"] = resultado

							} else {
								var resultado2 map[string]interface{}
								request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/grupo_sanguineo_persona/"+fmt.Sprintf("%v", grupoSanguineoPost["Id"]), "DELETE", &resultado2, nil)
								request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/"+fmt.Sprintf("%v", grupoEtnicoPost["Id"]), "DELETE", &resultado2, nil)
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
						var resultado2 map[string]interface{}
						request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_grupo_etnico/"+fmt.Sprintf("%v", grupoEtnicoPost["Id"]), "DELETE", &resultado2, nil)
						logs.Error(errGrupoSanguineoPost)
						//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
						c.Data["system"] = grupoSanguineoPost
						c.Abort("400")
					}
				} else {
					logs.Error(errGrupoSanguineoPost)
					//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
					c.Data["system"] = grupoSanguineoPost
					c.Abort("400")
				}
			} else {
				logs.Error(errGrupoEtnicoPost)
				//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
				c.Data["system"] = grupoEtnicoPost
				c.Abort("400")
			}
		} else {
			logs.Error(errGrupoEtnicoPost)
			//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = grupoEtnicoPost
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
