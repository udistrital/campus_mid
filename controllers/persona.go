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
}

// GuardarPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	models.PersonaCompleta	true		"body for Guardar Persona content"
// @Success 200 {string} models.Persona.Id
// @Failure 403 body is empty
// @router /GuardarPersona [post]
func (c *PersonaController) GuardarPersona() {
	var persona models.PersonaCompleta
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &persona); err == nil {
		fmt.Println(persona)
		var resultado models.Persona
		err := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona", "POST", &resultado, persona)
		fmt.Println("este es el resultado", resultado.Id)
		fmt.Println("este es el error", err)

		var estadoCivil models.PersonaEstadoCivil
		estadoCivil.Persona = &resultado
		estadoCivil.EstadoCivil = persona.EstadoCivil

		var resultado2 interface{}
		errEstadoCivil := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_estado_civil", "POST", &resultado2, estadoCivil)
		fmt.Println("este es el resultado", resultado2)
		fmt.Println("este es el error", errEstadoCivil)

		var genero models.PersonaGenero
		genero.Persona = &resultado
		genero.Genero = persona.Genero

		var resultado3 interface{}
		errGenero := request.SendJson("http://"+beego.AppConfig.String("PersonaService")+"/persona_genero", "POST", &resultado3, genero)
		fmt.Println("este es el resultado", resultado3)
		fmt.Println("este es el error", errGenero)

	} else {
		fmt.Println(err)
	}

	fmt.Println("guardando persona")

	//c.Data["json"] = m
	c.ServeJSON()

}
