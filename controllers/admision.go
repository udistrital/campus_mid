package controllers

import (
	//"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/planesticud/campus_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// AdmisionController ...
type AdmisionController struct {
	beego.Controller
}

// URLMapping ...
func (c *AdmisionController) URLMapping() {
	c.Mapping("PutAdmision", c.PutAdmision)
	c.Mapping("GetAdmision", c.GetAdmision)
}


// PutAdmision ...
// @Title PutAdmision
// @Description Modificar datos de la admisión
// @Param	id		path 	string	true		"el id de la admisión a modificar"
// @Param	body		body 	{}	true		"body Modificar Admisión content"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [put]
func (c *AdmisionController) PutAdmision() {
	
}

// GetAdmision ...
// @Title GetAdmision
// @Description consultar admision por id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AdmisionController) GetAdmision() {
	//Id de la experiencia a consultar
	idStr := c.Ctx.Input.Param(":id")
	//alerta que retorna la funcion GetAdmision
	var alerta models.Alert
	//cadena de alertas
	alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado admision
	var resultado map[string]interface{}

	errAdmision := request.GetJson("http://"+beego.AppConfig.String("AdmisionService")+"/admision/"+idStr, &resultado)
	
	if errAdmision == nil && resultado != nil {
		if resultado["Type"] != "error" {
			//buscar programa académico y nombre del aspirante
			var programa_academico []map[string]interface{}
			errProgramaAcademico := request.GetJson("http://"+beego.AppConfig.String("ProgramaAcademicaService")+"/programa_academico/?query=Id:"+idStr+"&fields=Id,Codigo,Nombre", &programa_academico)
			if errProgramaAcademico == nil {
				resultado["ProgramaAcademico"] = programa_academico
				c.Data["json"] = resultado
				fmt.Println(resultado["Soporte"])
			} else {
				alertas = append(alertas, errProgramaAcademico.Error())
				alerta.Code = "400"
				alerta.Type = "error"
				alerta.Body = alertas
				c.Data["json"] = alerta
			}
		} else {
			if resultado["Body"] == "<QuerySeter> no row found" {
				c.Data["json"] = nil
			} else {
				alertas = append(alertas, resultado["Body"])
				alerta.Code = "400"
				alerta.Type = "error"
				alerta.Body = alertas
				c.Data["json"] = alerta
			}
		}
	} else {
		alertas = append(alertas, errAdmision.Error())
		alerta.Code = "400"
		alerta.Type = "error"
		alerta.Body = alertas
		c.Data["json"] = alerta
	}
	c.ServeJSON()
}
