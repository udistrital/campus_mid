package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type PersonaController struct {
	beego.Controller
}

func (c *PersonaController) URLMapping() {
	c.Mapping("GuardarPersona", c.GuardarPersona)
}

// Get ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaCompleta	true		"body for Guardar Persona content"
// @Success 201 {int} models.PersonaCompleta
// @Failure 403 body is empty
// @router /GuardarPersona [post]
func (c *PersonaController) GuardarPersona() {
	fmt.Println("guardando persona")

	//c.Data["json"] = m
	c.ServeJSON()

}
