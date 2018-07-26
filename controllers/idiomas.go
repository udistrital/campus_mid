package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/campus_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// IdiomasController ...
type IdiomasController struct {
	beego.Controller
}

// URLMapping ...
func (c *IdiomasController) URLMapping() {
	c.Mapping("PostIdiomas", c.PostIdiomas)
	c.Mapping("PutIdiomas", c.PutIdiomas)
	c.Mapping("GetIdiomas", c.GetIdiomas)
	c.Mapping("DeleteIdiomas", c.DeleteIdiomas)
}

// PostIdiomas ...
// @Title PostIdiomas
// @Description Agregar Idioma
// @Param	body		body 	{}	true		"body Agregar Idioma content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /idioma [post]
func (c *IdiomasController) PostIdiomas() {
	//idioma a agregar
	var idioma map[string]interface{}
	//persona a la que corresponde ese idioma
	var persona map[string]interface{}
	//alerta que retorna la funcion PostIdiomas
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado idiomas
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}
	var soporte map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &idioma); err == nil {

		errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/"+fmt.Sprintf("%.f", idioma["Persona"]), &persona)
		if errPersona == nil && persona["Type"] != "error" {

			errIdiomas := request.SendJson("http://"+beego.AppConfig.String("IdiomaService")+"/conocimiento_idioma", "POST", &resultado, idioma)

			if errIdiomas == nil && resultado != nil {
				if resultado["Id"] != 0 {
					soporte = map[string]interface{}{
						"Descripcion":        idioma["Soporte"].(map[string]interface{})["Descripcion"],
						"Documento":          idioma["Soporte"].(map[string]interface{})["Documento"],
						"Institucion":        idioma["Soporte"].(map[string]interface{})["Institucion"],
						"ConocimientoIdioma": map[string]interface{}{"Id": resultado["Id"]},

						//
					}

					errSoporteIdiomas := request.SendJson("http://"+beego.AppConfig.String("IdiomaService")+"/soporte_conocimiento_idioma", "POST", &resultado2, soporte)
					if errSoporteIdiomas == nil {
						alertas = append(alertas, "OK soporte_idioma")
					}
				}

				alerta.Type = "success"
				alerta.Code = "200"
				alertas = append(alertas, "OK idioma")
			} else {
				alerta.Type = "error"
				alerta.Code = "400"
				if errIdiomas != nil {
					alertas = append(alertas, errIdiomas.Error())
				} else {
					alertas = append(alertas, "Error Idiomas")
				}
			}

		} else {
			alerta.Type = "error"
			alerta.Code = "400"
			alertas = append(alertas, "La persona no existe")
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

// PutIdiomas ...
// @Title PutIdiomas
// @Description Modificar Idioma
// @Param	id		path 	string	true		"el id de la Idioma a modificar"
// @Param	body		body 	{}	true		"body Modificar Idioma content"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /idioma/:id [put]
func (c *IdiomasController) PutIdiomas() {
	//Id del idioma a modificar
	idStr := c.Ctx.Input.Param(":id")
	//idioma a agregar
	var idioma map[string]interface{}
	//alerta que retorna la funcion PostIdiomas
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado modiciacion idioma a agregar
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}
	var soporte map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &idioma); err == nil {
		i64, _ := strconv.ParseInt(idStr, 10, 32)
		i := int32(i64)
		idioma["Id"] = i

		errIdioma := request.SendJson("http://"+beego.AppConfig.String("IdiomaService")+"/conocimiento_idioma/"+idStr, "PUT", &resultado, idioma)
		if idioma["Soporte"] != nil {
			fmt.Println("el soporte es: ")
			soporte = map[string]interface{}{
				"Descripcion":        idioma["Soporte"].(map[string]interface{})["Descripcion"],
				"Documento":          idioma["Soporte"].(map[string]interface{})["Documento"],
				"Institucion":        idioma["Soporte"].(map[string]interface{})["Institucion"],
				"ConocimientoIdioma": map[string]interface{}{"Id": idioma["Id"]},

				//
			}
			errSoporteIdioma := request.SendJson("http://"+beego.AppConfig.String("IdiomaService")+"/soporte_conocimiento_idioma/"+fmt.Sprintf("%.f", idioma["Soporte"].(map[string]interface{})["Id"].(float64)), "PUT", &resultado2, soporte)
			if errSoporteIdioma == nil {
				alertas = append(alertas, "OK UPDATE soporte_idioma")
			} else {
				alertas = append(alertas, "ERROR UPDATE soporte_idioma")
			}
		}
		if errIdioma == nil && resultado["Type"] == "success" {
			alertas = append(alertas, "OK UPDATE idioma")
			alerta.Code = "200"
			alerta.Type = "success"
			alerta.Body = alertas
			c.Data["json"] = alerta
		} else {
			if errIdioma == nil {
				alertas = append(alertas, resultado["Body"])
			} else {
				alertas = append(alertas, errIdioma.Error())
			}
			alerta.Code = "400"
			alerta.Type = "error"

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

// GetIdiomas ...
// @Title GetIdiomas
// @Description consultar Idiomas por userid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /idioma/:id [get]
func (c *IdiomasController) GetIdiomas() {
	//Id la persona a la que se consultan los idiomas
	idStr := c.Ctx.Input.Param(":id")
	//alerta que retorna la funcion GetIdiomas
	//var alerta models.Alert
	//acumulado de alertas
	//alertas := append([]interface{}{"Cadena de respuestas: "})
	//persona a la que corresponde ese idioma
	var persona map[string]interface{}
	var soporte []map[string]interface{}
	var idiomas []map[string]interface{}
	errPersona := request.GetJson("http://"+beego.AppConfig.String("PersonaService")+"/persona/"+idStr, &persona)
	if errPersona == nil && persona["Type"] != "error" {
		errIdioma := request.GetJson("http://"+beego.AppConfig.String("IdiomaService")+"/conocimiento_idioma/?query=Persona:"+idStr+"&fields=Id,ClasificacionNivelIdioma,Idioma,Nativo,NivelEscribe,NivelEscucha,NivelHabla,NivelLee", &idiomas)
		if errIdioma == nil && idiomas != nil {
			for i := 0; i < len(idiomas); i++ {
				errSoporteIdioma := request.GetJson("http://"+beego.AppConfig.String("IdiomaService")+"/soporte_conocimiento_idioma/?query=ConocimientoIdioma:"+fmt.Sprintf("%.f", idiomas[i]["Id"].(float64))+"&fields=Id,Descripcion,Documento,Institucion", &soporte)
				if errSoporteIdioma == nil && soporte != nil {
					idiomas[i]["Soporte"] = soporte[0]
				}
			}

			residioma := map[string]interface{}{
				"Persona": persona,
				"Idiomas": idiomas,
			}

			c.Data["json"] = residioma
		} else {
			c.Data["json"] = nil
		}

	} else {

		c.Data["json"] = nil
	}

	c.ServeJSON()
}

// DeleteIdiomas ...
// @Title DeleteIdiomas
// @Description eliminar Idioma por id de la formacion
// @Param	id		path 	string	true		"Id de la Idioma"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /idioma/:id [delete]
func (c *IdiomasController) DeleteIdiomas() {
	idStr := c.Ctx.Input.Param(":id")
	var alerta models.Alert
	var resultado map[string]interface{}
	//var soporte map[string]interface{}
	var idioma []map[string]interface{}
	//var soportes map[string]interface{}
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	errSoporteIdioma := request.GetJson("http://"+beego.AppConfig.String("IdiomaService")+"/soporte_conocimiento_idioma/?query=ConocimientoIdioma:"+idStr, &idioma)
	if errSoporteIdioma == nil {
		fmt.Println("el idioma es: ", idioma)
		for i := 0; i < len(idioma); i++ {
			err := request.SendJson("http://"+beego.AppConfig.String("IdiomaService")+"/soporte_conocimiento_idioma/"+fmt.Sprintf("%.f", idioma[i]["Id"].(float64)), "DELETE", &resultado, nil)
			if err == nil {
				alertas = append(alertas, "OK DELETE soporte_idioma")
			}
		}
	} else {
		alertas = append(alertas, errSoporteIdioma.Error())
	}

	err := request.SendJson("http://"+beego.AppConfig.String("IdiomaService")+"/conocimiento_idioma/"+idStr, "DELETE", &resultado, nil)
	if err == nil && resultado["Type"] == "success" {
		alertas = append(alertas, "OK DELETE idioma")
		alerta.Body = alertas
		alerta.Code = "200"
		alerta.Type = "success"
		c.Data["json"] = alerta
	} else {
		if err == nil {
			alertas = append(alertas, resultado["Body"])
		} else {
			alertas = append(alertas, err.Error())
		}
		alerta.Body = alertas
		alerta.Code = "400"
		alerta.Type = "error"
		c.Data["json"] = alerta
	}
	c.ServeJSON()
}
