package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// OrganizacionController operations for Organizacion
type OrganizacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrganizacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	// c.Mapping("GetOne", c.GetOne)
	// c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Organizacion
// @Param	body		body 	interface	true		"body for Organizaciona scontent"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *OrganizacionController) Post() {
	var organizacion map[string]interface{}
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &organizacion); err == nil {
		o := map[string]interface{}{
			"TipoOrganizacion": organizacion["TipoOrganizacion"],
			"Nombre":           organizacion["Nombre"],
		}
		if err := request.SendJson("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion", "POST", &resultado, o); err == nil {

			p := map[string]interface{}{
				"LugarExpedicion":      organizacion["LugarExpedicion"],
				"FechaExpedicion":      organizacion["FechaExpedicion"],
				"TipoIdentificacion":   organizacion["TipoIdentificacion"], // asegurando que 5 es el ID para NIT
				"NumeroIdentificacion": organizacion["NumeroIdentificacion"],
				"Ente":                 map[string]interface{}{"Id": resultado["Ente"]},
			}
			if err := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion", "POST", &resultado2, p); err == nil && resultado2["Type"] != "error" {
				c.Data["json"] = resultado2
			} else {
				request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion/%.f", resultado["Id"]), "DELETE", &resultado2, nil)
				c.Data["json"] = []interface{}{err, resultado2}
			}
		} else {
			c.Data["json"] = err
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

/*
// GetOne ...
// @Title GetOne
// @Description get Organizacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Organizacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrganizacionController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Organizacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Organizacion
// @Failure 403
// @router / [get]
func (c *OrganizacionController) GetAll() {
	fmt.Print("aasdsad")
}

// Put ...
// @Title Put
// @Description update the Organizacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Organizacion	true		"body for Organizacion content"
// @Success 200 {object} models.Organizacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrganizacionController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Organizacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrganizacionController) Delete() {

}
*/
