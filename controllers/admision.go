package controllers

import (
	//"encoding/json"
	//"fmt"

	"github.com/astaxie/beego"
	//"github.com/udistrital/campus_mid/models"
	//"github.com/udistrital/utils_oas/request"
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
	
}
