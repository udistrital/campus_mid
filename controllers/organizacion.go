package controllers

import (
	"encoding/json"
	"fmt"
	"sync"

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
	c.Mapping("GetByIdentificacion", c.GetByIdentificacion)
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
	var resultado []map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &organizacion); err == nil {

		if err := request.GetJson(
			fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/identificacion?query=NumeroIdentificacion:%s,TipoIdentificacion.Id:%.f",
				organizacion["NumeroIdentificacion"], organizacion["TipoIdentificacion"].(map[string]interface{})["Id"]),
			&resultado); resultado != nil {
			c.Data["json"] = resultado[0]
		} else if err == nil {
			if res, errores := CrearOrganizacion(organizacion); errores == nil {
				c.Data["json"] = res
			} else {
				c.Data["json"] = errores
			}
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// CrearOrganizacion Funcion que valida si existe la organizacion si no la crea
func CrearOrganizacion(organizacion map[string]interface{}) (res map[string]interface{}, errores []interface{}) {
	var resultado map[string]interface{}
	var resultado2 map[string]interface{}
	o := map[string]interface{}{
		"TipoOrganizacion": organizacion["TipoOrganizacion"],
		"Nombre":           organizacion["Nombre"],
	}
	if err := request.SendJson("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion", "POST", &resultado, o); err == nil && resultado["Type"] != "error" {

		p := map[string]interface{}{
			"LugarExpedicion":      organizacion["LugarExpedicion"],
			"FechaExpedicion":      organizacion["FechaExpedicion"],
			"TipoIdentificacion":   organizacion["TipoIdentificacion"], // asegurando que 5 es el ID para NIT
			"NumeroIdentificacion": organizacion["NumeroIdentificacion"],
			"Ente":                 map[string]interface{}{"Id": resultado["Ente"]},
		}
		if err := request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/identificacion", "POST", &resultado2, p); err == nil && resultado2["Type"] != "error" {
			res = resultado2
			res["Nombre"] = organizacion["Nombre"]
			if organizacion["Contacto"] != nil {
				array := organizacion["Contacto"].([]interface{})
				for _, c := range array {
					contacto := c.(map[string]interface{})
					contacto["Ente"] = map[string]interface{}{"Id": resultado["Ente"]}
					var r interface{}
					request.SendJson("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente", "POST", &r, c)
				}
			}

		} else {
			request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion/%.f", resultado["Id"]), "DELETE", &resultado2, nil)
			errores = []interface{}{err, resultado2}
		}
	} else {
		errores = []interface{}{err, resultado}
	}
	return
}

// GetByIdentificacion ...
// @Title GetByIdentificacion
// @Description get Organizacion by id
// @Param	id		query 	string	true		"Identification number as id"
// @Param	tipoid		query 	string	true		"TipoIdentificacion number as nit"
// @Success 200 {object} models.Organizacion
// @Failure 403 :id is empty
// @router /identificacion/ [get]
func (c *OrganizacionController) GetByIdentificacion() {
	uid := c.GetString("id")
	tid := c.GetString("tipoid")
	var resultado map[string]interface{}

	var resId []map[string]interface{}
	// var res_ontacto []map[string]interface{}
	// var res_ubicacion []map[string]interface{}
	if uid != "" && tid != "" {
		if err := request.GetJson(
			fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/identificacion?query=NumeroIdentificacion:%s,TipoIdentificacion.Id:%s",
				uid, tid),
			&resId); err == nil {
			if resId != nil {
				resultado = resId[0]
				delete(resultado, "Id")
				var wg sync.WaitGroup
				wg.Add(3)

				go func() {
					var resOrg []map[string]interface{}
					if err := request.GetJson(
						fmt.Sprintf("http://"+beego.AppConfig.String("OrganizacionService")+"/organizacion/?query=Ente:%.f", resultado["Ente"].(map[string]interface{})["Id"]),
						&resOrg); err == nil {
						if resOrg != nil {
							resultado["Nombre"] = resOrg[0]["Nombre"]
							resultado["TipoOrganizacion"] = resOrg[0]["TipoOrganizacion"]
						}
					}
					// Do work
					wg.Done()
				}()

				go func() {
					var resContacto []map[string]interface{}
					if err := request.GetJson(
						fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/contacto_ente/?query=Ente:%.f&fields=TipoContacto,Valor", resultado["Ente"].(map[string]interface{})["Id"]),
						&resContacto); err == nil {
						if resContacto != nil {
							resultado["Contacto"] = resContacto
						}
					}
					// Do work
					wg.Done()
				}()

				go func() {
					var resUbicacion []map[string]interface{}
					if err := request.GetJson(
						fmt.Sprintf("http://"+beego.AppConfig.String("EnteService")+"/valor_atributo_ubicacion/?query=UbicacionEnte.Ente.Id:%.f", resultado["Ente"].(map[string]interface{})["Id"]),
						&resUbicacion); err == nil {
						if resUbicacion != nil {
							resultado["Ubicacion"] = resUbicacion
						}
					}
					// Do work
					wg.Done()
				}()

				wg.Wait()

			}
			c.Data["json"] = resultado

		} else {
			c.Data["json"] = err
		}
	}
	c.ServeJSON()
}

/*
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
