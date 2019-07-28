package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/planesticud/campus_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// DescuentoController ...
type DescuentoController struct {
  beego.Controller
}

// URLMapping ...
func (c *DescuentoController) URLMapping() {
	c.Mapping("PostDescuentoAcademico", c.PostDescuentoAcademico)
	c.Mapping("PutDescuentoAcademico", c.PutDescuentoAcademico)
	c.Mapping("GetDescuentoAcademico", c.GetDescuentoAcademico)
	c.Mapping("GetDescuentoAcademicoByPersona", c.GetDescuentoAcademicoByPersona)
	c.Mapping("DeleteDescuentoAcademico", c.DeleteDescuentoAcademico)
}

// PostDescuentoAcademico ...
// @Title PostDescuentoAcademico
// @Description Agregar Descuento Academico ud
// @Param	body		body 	'hola'	true		"body Agregar Descuento Academico content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /descuentoacademico [post]
func (c *DescuentoController) PostDescuentoAcademico() {
  //solicitud de descuento
  var solicitud map[string]interface{}
  //alerta que retorna la funcion PostDescuentoAcademico
  var alerta models.Alert
  //cadena de alertas
  alertas := append([]interface{}{"Cadena de respuestas: "})
  //resultado solicitud de descuento
  var resultado map[string]interface{}
  //resultado de soporte descuento
  var resultado2 map[string]interface{}

  if err := json.Unmarshal(c.Ctx.Input.RequestBody, &solicitud); err == nil {
    solicituddescuento := map[string]interface{}{
      "PersonaId":              solicitud["Persona"],
      "Estado":                 "Por aprobar",
      "PeriodoId":              solicitud["Periodo"],
      "Activo":                 true,
      "DescuentosDependenciaId": solicitud["DescuentoDependencia"],//.(map[string]interface{})["Id"],
    }
    errSolicitud := request.SendJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/solicitud_descuento", "POST", &resultado, solicituddescuento)
    if errSolicitud == nil && resultado["Type"] != "error" {
      alertas = append(alertas, "se agrego la solicitud correctamente")
      soportedescuento := map[string]interface{}{
        "SolicitudDescuentoId": resultado["Body"],
        "Activo":               true,
        "DocumentoId":          solicitud["Documento"],
      }
      fmt.Println("el soporte es:", soportedescuento)
      errSoporteDescuento := request.SendJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/soporte_descuento", "POST", &resultado2, soportedescuento)
      if errSoporteDescuento == nil && resultado2["Type"] != "error" {
        alerta.Type = "success"
				alerta.Code = "200"
				alertas = append(alertas, "se agrego el soporte correctamente")
      } else {
				alerta.Type = "error"
				alerta.Code = "400"
				alertas = append(alertas, errSoporteDescuento.Error())
			}
    } else {
			fmt.Println(resultado)
			alerta.Type = "error"
			alerta.Code = "400"
			if errSolicitud != nil {
				alertas = append(alertas, errSolicitud.Error())
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

// PutDescuentoAcademico ...
// @Title PutDescuentoAcademico
// @Description Modificar Descuento Academico
// @Param	Id	path 	string	true		"el id de la solicitud de descuento a modificar"
// @Param	body		body 	{}	true		"body Modificar Descuento Academico content"
// @Success 200 {}
// @Failure 403 :id s empty
// @router /descuentoacademico/:Id [put]
func (c *DescuentoController) PutDescuentoAcademico() {
	idStr := c.Ctx.Input.Param(":Id")
	//solicitud descuento
	var solicitud map[string]interface{}
	//alerta que retorna la funcion PutDescuentoAcademico
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado solicitud descuento
	var resultado []map[string]interface{}
  var resultado2 map[string]interface{}
	//resultado soporte descuento
  var resultado3 []map[string]interface{}
  var resultado4 map[string]interface{}

  if err := json.Unmarshal(c.Ctx.Input.RequestBody, &solicitud); err == nil {
    errSolicitud := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/solicitud_descuento/?query=Id:"+idStr, &resultado)
	  if errSolicitud == nil {
      for i := 0; i < len(resultado); i++ {
        if resultado[i]["Id"] == solicitud["Id"] {
          solicituddescuento := map[string]interface{}{
						"Id":											 solicitud["Id"],
            "PersonaId":               solicitud["Persona"],
            "Activo":                  true,
            "Estado":                  solicitud["Estado"],
            "PeriodoId":               solicitud["Periodo"],
            "DescuentosDependenciaId": solicitud["DescuentoDependencia"],
          }
					errSolicitud2 := request.SendJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/solicitud_descuento/"+fmt.Sprintf("%.f", resultado[i]["Id"].(float64)), "PUT", &resultado2, solicituddescuento)
					if errSolicitud2 == nil {
            if resultado2["Type"] == "success" {
							alertas = append(alertas, "OK UPDATE solicitud_descuento")
							alerta.Code = "200"
							alerta.Type = "success"
							alerta.Body = alertas
						}
          } else {
  					alertas = append(alertas, errSolicitud2.Error())
  					alerta.Code = "400"
  					alerta.Type = "error"
					}
				  errSoporteDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/soporte_descuento/?query=SolicitudDescuentoId:"+fmt.Sprintf("%.f", resultado[i]["Id"].(float64)), &resultado3)
          if errSoporteDescuento == nil && resultado3 != nil {
            fmt.Println("el soporte actual es:", resultado3[0])
            resultado3[0]["DocumentoId"] = solicitud["Documento"]
            fmt.Println("el nuevo soporte es:", resultado3[0])
            errSoporteDescuento2 := request.SendJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/soporte_descuento/"+fmt.Sprintf("%.f", resultado3[0]["Id"].(float64)), "PUT", &resultado4, resultado3[0])
            if errSoporteDescuento2 == nil {
              alertas = append(alertas, "OK UPDATE Documento ")
            }
          } else {
						if errSoporteDescuento != nil {
							fmt.Println("error soporte: ", errSoporteDescuento.Error())
							alertas = append(alertas, errSoporteDescuento.Error())
						} else {
							alertas = append(alertas, "ERROR soporte no se encontro registro")
						}
						alerta.Code = "400"
						alerta.Type = "error"
					}
        }
      }
    } else {
			alertas = append(alertas, errSolicitud.Error())
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

// GetDescuentoAcademico ...
// @Title GetDescuentoAcademico
// @Description consultar Descuento Academico por userid
// @Param	Id		query 	string	true		"The key for staticblock"
// @Param	idsolicitud		query 	string	false		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /descuentoacademico/ [get]
func (c *DescuentoController) GetDescuentoAcademico() {
  //Id de la persona
	var idStr = ""
	var idSolitudDes float64 = 0
	idStr = c.GetString("Id")
	idSolitudDes, _ = c.GetFloat("idsolicitud")
	fmt.Println("el idSolitudDes es: ", idSolitudDes)
  //alerta que retorna la funcion GetDescuentoAcademico
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
  //resultado solicitud descuento
	var resultado []map[string]interface{}
	//resultado descuento dependencia
	var resultado2 []map[string]interface{}
  //resultado tipo descuento
	var resultado3 []map[string]interface{}
	//resultado soporte descuento
	var resultado4 []map[string]interface{}

  errSolicitud := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/solicitud_descuento/?query=PersonaId:"+idStr+",Id:"+fmt.Sprintf("%.f", idSolitudDes)+"&fields=Id,PersonaId,Estado,PeriodoId,DescuentosDependenciaId", &resultado)
  if errSolicitud == nil && resultado != nil {
    if resultado[0]["Type"] != "error" {
      resultado[0]["Persona"] = idStr
      for u := 0; u < len(resultado); u++ {
				var descuentoDependencia map[string]interface{}
				descuentoDependencia = resultado[u]["DescuentosDependenciaId"].(map[string]interface{})
        errDescuentoDependencia := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/descuentos_dependencia/?query=Id:"+fmt.Sprintf("%.f", descuentoDependencia["Id"].(float64)), &resultado2)
        if errDescuentoDependencia == nil {
					var tipoDescuento map[string]interface{}
					tipoDescuento = resultado2[0]["TipoDescuentoId"].(map[string]interface{})
          errTipoDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/tipo_descuento/?query=Id:"+fmt.Sprintf("%.f", tipoDescuento["Id"].(float64)), &resultado3)
          if errTipoDescuento == nil && resultado3[0] != nil {
            resultado2[0]["TipoDescuento"] = resultado3[0]
            resultado[u]["DescuentoDependencia"] = resultado2[0]
          } else {
            fmt.Println(errTipoDescuento)
          }
        } else {
  				alertas = append(alertas, errDescuentoDependencia.Error())
  				alerta.Code = "400"
  				alerta.Type = "error"
  				alerta.Body = alertas
  				c.Data["json"] = alerta
  			}
        errSoporteDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/soporte_descuento/?query=SolicitudDescuentoId:"+fmt.Sprintf("%.f", resultado[u]["Id"].(float64))+"&fields=DocumentoId", &resultado4)
      	if errSoporteDescuento == nil && resultado4 != nil {
					//fmt.Println("el resultado de los documentos es: ", resultado4)
					resultado[u]["Documento"] = resultado4[0]["DocumentoId"]
				} else {
					if errSoporteDescuento != nil {
						fmt.Println("el error es: ", errSoporteDescuento.Error())
					}
				}
      }
      if idSolitudDes != 0 {
				for u := 0; u < len(resultado); u++ {
					fmt.Println("El id de la solicitud es ", resultado[u]["Id"], "y el de la consulta: ", idSolitudDes)
					if resultado[u]["Id"].(float64) == idSolitudDes {
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
		if errSolicitud != nil {
			c.Data["json"] = nil
		} else {
			c.Data["json"] = nil
		}
	}
	c.ServeJSON()
}

// GetDescuentoAcademicoByPersona ...
// @Title GetDescuentoAcademicoByPersona
// @Description consultar Descuento Academico por userid
// @Param	Persona		query 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /descuentoacademicopersona/ [get]
func (c *DescuentoController) GetDescuentoAcademicoByPersona() {
	//Id de la persona
	var idStr = ""
	idStr = c.GetString("Persona")
	fmt.Println("el id es: ", idStr)
  //alerta que retorna la funcion GetDescuentoAcademico
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
  //resultado solicitud descuento
	var resultado []map[string]interface{}
	//resultado descuento dependencia
	var resultado2 []map[string]interface{}
  //resultado tipo descuento
	var resultado3 []map[string]interface{}
	//resultado soporte descuento
	var resultado4 []map[string]interface{}

  errSolicitud := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/solicitud_descuento/?query=PersonaId:"+idStr+"&fields=Id,PersonaId,Estado,PeriodoId,DescuentosDependenciaId", &resultado)
  if errSolicitud == nil && resultado != nil {
    if resultado[0]["Type"] != "error" {
      resultado[0]["Persona"] = idStr
      for u := 0; u < len(resultado); u++ {
				var descuentoDependencia map[string]interface{}
				descuentoDependencia = resultado[u]["DescuentosDependenciaId"].(map[string]interface{})
        errDescuentoDependencia := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/descuentos_dependencia/?query=Id:"+fmt.Sprintf("%.f", descuentoDependencia["Id"].(float64)), &resultado2)
        if errDescuentoDependencia == nil {
					var tipoDescuento map[string]interface{}
					tipoDescuento = resultado2[0]["TipoDescuentoId"].(map[string]interface{})
          errTipoDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/tipo_descuento/?query=Id:"+fmt.Sprintf("%.f", tipoDescuento["Id"].(float64)), &resultado3)
          if errTipoDescuento == nil && resultado3[0] != nil {
            resultado2[0]["TipoDescuento"] = resultado3[0]
            resultado[u]["DescuentoDependencia"] = resultado2[0]
          } else {
            fmt.Println(errTipoDescuento)
          }
        } else {
  				alertas = append(alertas, errDescuentoDependencia.Error())
  				alerta.Code = "400"
  				alerta.Type = "error"
  				alerta.Body = alertas
  				c.Data["json"] = alerta
  			}
        errSoporteDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/soporte_descuento/?query=SolicitudDescuentoId:"+fmt.Sprintf("%.f", resultado[u]["Id"].(float64))+"&fields=DocumentoId", &resultado4)
				if errSoporteDescuento == nil && resultado4 != nil {
					//fmt.Println("el resultado de los documentos es: ", resultado4)
					resultado[u]["Documento"] = resultado4[0]["DocumentoId"]
				} else {
					if errSoporteDescuento != nil {
						fmt.Println("el error es: ", errSoporteDescuento.Error())
					}
				}
				c.Data["json"] = resultado[u]
      }
    }
  } else {
		fmt.Println("entro al error")
		if errSolicitud != nil {
			c.Data["json"] = nil
		} else {
			c.Data["json"] = nil
		}
	}
	c.ServeJSON()
}

// GetDescuentoByDependenciaPeriodo ...
// @Title GetDescuentoByDependenciaPeriodo
// @Description consultar Descuento Academico por userid
// @Param	Dependencia		query 	string	true		"The key for staticblock"
// @Param	Periodo		query 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /descuentodependenciaperiodo/ [get]
func (c *DescuentoController) GetDescuentoByDependenciaPeriodo() {
	var dependencia = ""
	var periodo = ""
	dependencia = c.GetString("Dependencia")
	periodo = c.GetString("Periodo")
  //alerta que retorna la funcion GetDescuentoAcademico
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado descuento dependencia
	var resultado []map[string]interface{}
	//resultado tipo descuento
	var resultado2 []map[string]interface{}

	errDescuentoDependencia := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/descuentos_dependencia/?query=DependenciaId:"+dependencia+",PeriodoId:"+periodo, &resultado)
	if errDescuentoDependencia == nil && resultado != nil {
		if resultado[0]["Type"] != "error" {
			resultado[0]["Dependencia"] = dependencia
			resultado[0]["Periodo"]= periodo
      for u := 0; u < len(resultado); u++ {
			  errTipoDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/tipo_descuento/?query=Id:"+fmt.Sprintf("%.f", resultado[u]["TipoDescuentoId"].(map[string]interface{})["Id"].(float64)), &resultado2)
				if errTipoDescuento == nil && resultado2[0] != nil {
          resultado[u]["TipoDescuento"] = resultado2[0]
        } else {
  				alertas = append(alertas, errTipoDescuento.Error())
  				alerta.Code = "400"
  				alerta.Type = "error"
  				alerta.Body = alertas
  				c.Data["json"] = alerta
  			}
      }
			c.Data["json"] = resultado
    }
  } else {
		fmt.Println("entro al error")
		if errDescuentoDependencia != nil {
			c.Data["json"] = nil
		} else {
			c.Data["json"] = nil
		}
	}
	c.ServeJSON()
}

// GetDescuentoByPersonaPeriodoDependencia ...
// @Title GetDescuentoByPersonaPeriodoDependencia
// @Description consultar Descuento Academico por userid
// @Param	Persona		query 	string	true		"The key for staticblock"
// @Param	Dependencia		query 	string	true		"The key for staticblock"
// @Param	Periodo		query 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /descuentopersonaperiododependencia/ [get]
func (c *DescuentoController) GetDescuentoByPersonaPeriodoDependencia() {
	var dependencia = ""
	var periodo = ""
	var persona = ""
	dependencia = c.GetString("Dependencia")
	periodo = c.GetString("Periodo")
	persona = c.GetString("Persona")
  //alerta que retorna la funcion GetDescuentoAcademico
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado descuento dependencia
	var resultado []map[string]interface{}
	//resultado solicitud descuento
	var resultado2 []map[string]interface{}
  //resultado tipo descuento
	var resultado3 []map[string]interface{}
	//resultado soporte descuento
	var resultado4 []map[string]interface{}

	errDescuentoDependencia := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/descuentos_dependencia/?query=DependenciaId:"+dependencia+",PeriodoId:"+periodo, &resultado)
	if errDescuentoDependencia == nil && resultado != nil {
    if resultado[0]["Type"] != "error" {
      for u := 0; u < len(resultado); u++ {
				errSolicitud := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/solicitud_descuento/?query=DescuentosDependenciaId:"+fmt.Sprintf("%.f",resultado[u]["Id"].(float64))+",PersonaId:"+persona, &resultado2)
        if errSolicitud == nil {
					var tipoDescuento map[string]interface{}
					tipoDescuento = resultado[u]["TipoDescuentoId"].(map[string]interface{})
          errTipoDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/tipo_descuento/?query=Id:"+fmt.Sprintf("%.f", tipoDescuento["Id"].(float64)), &resultado3)
	        if errTipoDescuento == nil && resultado3[0] != nil {
						resultado[u]["TipoDescuento"] = resultado3[0]
						resultado2[0]["DescuentoDependencia"] = resultado[u]
						resultado[u] = resultado2[0]
          } else {
            fmt.Println(errTipoDescuento)
          }
        } else {
  				alertas = append(alertas, errSolicitud.Error())
  				alerta.Code = "400"
  				alerta.Type = "error"
  				alerta.Body = alertas
  				c.Data["json"] = alerta
  			}
        errSoporteDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/soporte_descuento/?query=SolicitudDescuentoId:"+fmt.Sprintf("%.f", resultado2[0]["Id"].(float64))+"&fields=DocumentoId", &resultado4)
				if errSoporteDescuento == nil && resultado4 != nil {
					//fmt.Println("el resultado de los documentos es: ", resultado4)
					resultado[u]["Documento"] = resultado4[0]["DocumentoId"]
				} else {
					if errSoporteDescuento != nil {
						fmt.Println("el error es: ", errSoporteDescuento.Error())
					}
				}
				c.Data["json"] = resultado
      }
    }
  } else {
		fmt.Println("entro al error")
		if errDescuentoDependencia != nil {
			c.Data["json"] = nil
		} else {
			c.Data["json"] = nil
		}
	}
	c.ServeJSON()
}

// DeleteDescuentoAcademico ...
// @Title DeleteDescuentoAcademico
// @Description eliminar Formacion Academica por id de la formacion
// @Param	Id		path 	string	true		"Id de la formacion academica"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /descuentoacademico/:Id [delete]
func (c *DescuentoController) DeleteDescuentoAcademico() {
  idStr := c.Ctx.Input.Param(":Id")
  var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
  //resultado solicitud descuento
  var resultado map[string]interface{}
  //resultado soporte descuento
  var resultado2 []map[string]interface{}
  //resultados eliminacion
  var resultado3 map[string]interface{}
  var resultado4 map[string]interface{}

  errSolicitud := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/solicitud_descuento/"+idStr, &resultado)
  if errSolicitud == nil {
    errSoporteDescuento := request.GetJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/soporte_descuento/?query=SolicitudDescuentoId:"+idStr, &resultado2)
    if errSoporteDescuento == nil {
      err := request.SendJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/soporte_descuento/"+fmt.Sprintf("%.f", resultado2[0]["Id"].(float64)), "DELETE", &resultado3, nil)
      if err == nil {
        alertas = append(alertas, "OK DELETE soporte_descueto")
      }
    } else {
      alertas = append(alertas, errSoporteDescuento.Error())
    }
    err2 := request.SendJson("http://"+beego.AppConfig.String("DescuentoAcademicoService")+"/solicitud_descuento/"+idStr, "DELETE", &resultado4, nil)
    if err2 == nil {
      alertas = append(alertas, "OK DELETE solicitud_descuento")
    }
    alerta.Body = alertas
    alerta.Code = "200"
    alerta.Type = "success"
    c.Data["json"] = alerta
  } else {
  	alertas = append(alertas, errSolicitud.Error())
  	alerta.Body = alertas
  	alerta.Code = "400"
  	alerta.Type = "error"
  	c.Data["json"] = alerta
	}
	c.ServeJSON()
}
