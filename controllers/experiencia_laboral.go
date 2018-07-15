package controllers

import (
    "encoding/json"
    "fmt"
    "strconv"

    "github.com/astaxie/beego"
    "github.com/udistrital/campus_mid/models"
    "github.com/udistrital/utils_oas/request"
)

// FormacionController ...
type ExperienciaLaboralController struct {
    beego.Controller
}

// URLMapping ...
func (c *ExperienciaLaboralController) URLMapping() {
    c.Mapping("PostExperienciaLaboral", c.PostExperienciaLaboral)
    c.Mapping("PutExperienciaLaboral", c.PutExperienciaLaboral)
    c.Mapping("GetExperienciaLaboral", c.GetExperienciaLaboral)
    c.Mapping("DeleteExperienciaLaboral", c.DeleteExperienciaLaboral)
}

// PostExperienciaLaboral ...
// @Title PostExperienciaLaboral
// @Description Agregar Experiencia Laboral
// @Param   body        body    {}  true        "body Agregar EXperiencia Laboral content"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *ExperienciaLaboralController) PostExperienciaLaboral() {
    //experiencia laboral
    var experiencia map[string]interface{}
    //alerta que retorna la funcion PostExperienciaLaboral
    var alerta models.Alert
    //cadena de alertas
    alertas := append([]interface{}{"Cadena de respuestas: "})
    //resultado formacion academica
    var resultado map[string]interface{}
    //resultado soporte experiencia laboral
    var resultado2 map[string]interface{}

    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &experiencia); err == nil {
        experienciaLaboral := map[string]interface{}{
            "Persona":           experiencia["Ente"].(map[string]interface{})["Id"],
            "Actividades":       experiencia["Actividades"],
            "FechaInicio":       experiencia["FechaInicio"],
            "FechaFinalizacion": experiencia["FechaFinalizacion"],
            "Organizacion":      experiencia["Organizacion"].(map[string]interface{})["Id"],
            "TipoDedicacion":    experiencia["TipoDedicacion"],
            "Cargo":             experiencia["Cargo"],
            "TipoVinculacion":   experiencia["TipoVinculacion"],
        }
        errExperienciaLaboral := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/experiencia_laboral", "POST", &resultado, experienciaLaboral)

        if errExperienciaLaboral == nil && resultado["Type"] != "error" {
            alertas = append(alertas, "se agrego la experiencia laboral")

            //si se envía algún soporte en la experiencia laboral
            if experiencia["Soporte"] != nil {
                experienciaLaboralSoporte := map[string]interface{}{
                    "Documento":          experiencia["Soporte"].(map[string]interface{})["Documento"],
                    "Descripcion":        experiencia["Soporte"].(map[string]interface{})["Descripcion"],
                    "ExperienciaLaboral": resultado["Body"],
                }
                errExperienciaLaboralSoporte := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral", "POST", &resultado2, experienciaLaboralSoporte)
                if errExperienciaLaboralSoporte == nil && resultado2["Type"] != "error" {
                    alerta.Type = "success"
                    alerta.Code = "200"
                    alertas = append(alertas, "se agrego el soporte correctamente")
                } else {
                    alerta.Type = "error"
                    alerta.Code = "400"
                    alertas = append(alertas, errExperienciaLaboralSoporte.Error())
                }
            }
        } else {
            alerta.Type = "error"
            alerta.Code = "400"
            if errExperienciaLaboral != nil {
                alertas = append(alertas, errExperienciaLaboral.Error())
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

// PutExperienciaLaboral ...
// @Title PutExperienciaLaboral
// @Description Modificar Experiencia Laboral
// @Param   id      path    string  true        "el id de la experiencia laboral a modificar"
// @Param   body        body    {}  true        "body Modificar Experiencia Laboral content"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [put]
func (c *ExperienciaLaboralController) PutExperienciaLaboral() {
    idStr := c.Ctx.Input.Param(":id")
    //experiencia laboral
    var experiencia map[string]interface{}
    //alerta que retorna la funcion PutExperienciaLaboral
    var alerta models.Alert
    //cadena de alertas
    alertas := append([]interface{}{"Cadena de respuestas: "})
    //resultado experiencia laboral
    var resultado map[string]interface{}
    //resultado dato adicional experiencia laboral
    var resultado2 map[string]interface{}
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &experiencia); err == nil {

        experienciaLaboral := map[string]interface{}{
            "Actividades":       experiencia["Actividades"],
            "FechaInicio":       experiencia["FechaInicio"],
            "FechaFinalizacion": experiencia["FechaFinalizacion"],
            "Organizacion":      experiencia["Organizacion"].(map[string]interface{})["Id"],
            "TipoDedicacion":    experiencia["TipoDedicacion"],
            "Cargo":             experiencia["Cargo"],
            "TipoVinculacion":   experiencia["TipoVinculacion"],
        }

        errExperienciaLaboral := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/experiencia_laboral/"+idStr, "PUT", &resultado, experienciaLaboral)
        if errExperienciaLaboral == nil && resultado["Type"] == "success" {

            alertas = append(alertas, "OK UPDATE experiencia laboral")
            alerta.Code = "200"
            alerta.Type = "success"

            //si se envía algún soporte en la experiencia laboral a modificar
            if experiencia["Soporte"] != nil {

                //buscar el soporte de la experiencia laboral
                var soporte []map[string]interface{}
                errSoportes := request.GetJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/?query=ExperienciaLaboral:"+idStr, &soporte)

                //si la experiencia laboral tiene soporte: Actualizarlo
                if errSoportes == nil && soporte != nil {
                    soporteExperienciaLaboral := map[string]interface{}{
                        "Documento":   experiencia["Soporte"].(map[string]interface{})["Documento"],
                        "Descripcion": experiencia["Soporte"].(map[string]interface{})["Descripcion"],
                    }

                    errSoporteExperienciaLaboral := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/"+fmt.Sprintf("%.f", soporte[0]["Id"]), "PUT", &resultado2, soporteExperienciaLaboral)
                    if errSoporteExperienciaLaboral == nil {
                        if resultado2["Type"] == "success" {
                            alertas = append(alertas, "OK UPDATE Soporte")
                            alerta.Code = "200"
                            alerta.Type = "success"
                        } else {
                            alertas = append(alertas, resultado2["Body"])
                            alerta.Code = "400"
                            alerta.Type = "error"
                        }
                    } else {
                        alertas = append(alertas, errSoporteExperienciaLaboral.Error())
                        alerta.Code = "400"
                        alerta.Type = "error"
                    }
                } else {
                    //si no tiene soporte, se registra
                    id, _ := strconv.Atoi(idStr)
                    experienciaLaboralSoporte := map[string]interface{}{
                        "Documento":          experiencia["Soporte"].(map[string]interface{})["Documento"],
                        "Descripcion":        experiencia["Soporte"].(map[string]interface{})["Descripcion"],
                        "ExperienciaLaboral": map[string]interface{}{"Id": id},
                    }
                    errExperienciaLaboralSoporte := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral", "POST", &resultado2, experienciaLaboralSoporte)
                    if errExperienciaLaboralSoporte == nil {
                        if resultado2["Type"] != "error" {
                            alerta.Type = "success"
                            alerta.Code = "200"
                            alertas = append(alertas, "se agrego el soporte correctamente")
                        } else {
                            alertas = append(alertas, resultado2["Body"])
                            alerta.Code = "400"
                            alerta.Type = "error"
                        }
                    } else {
                        alerta.Type = "error"
                        alerta.Code = "400"
                        alertas = append(alertas, errExperienciaLaboralSoporte.Error())
                    }
                }
            }
        } else {
            alertas = append(alertas, errExperienciaLaboral.Error())
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

// GetExperienciaLaboral ...
// @Title GetExperienciaLaboral
// @Description consultar Experiencia Laboral por userid
// @Param   id      path    string  true        "The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ExperienciaLaboralController) GetExperienciaLaboral() {
    //Id de la experiencia a consultar
    idStr := c.Ctx.Input.Param(":id")
    //alerta que retorna la funcion GetExperienciaLaboral
    var alerta models.Alert
    //cadena de alertas
    alertas := append([]interface{}{"Cadena de respuestas: "})
    //resultado experiencia laboral
    var resultado map[string]interface{}

    errExperienciaLaboral := request.GetJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/experiencia_laboral/"+idStr, &resultado)

    if errExperienciaLaboral == nil && resultado != nil {
        if resultado["Type"] != "error" {
            //buscar soporte_experiencia_laboral
            var soporte []map[string]interface{}
            errSoportes := request.GetJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/?query=ExperienciaLaboral:"+idStr+"&fields=Id,Documento,Descripcion", &soporte)

            if errSoportes == nil {
                if soporte != nil {
                    resultado["Soporte"] = soporte[0]
                }
                c.Data["json"] = resultado
            } else {
                alertas = append(alertas, errSoportes.Error())
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
        alertas = append(alertas, errExperienciaLaboral.Error())
        alerta.Code = "400"
        alerta.Type = "error"
        alerta.Body = alertas
        c.Data["json"] = alerta
    }

    c.ServeJSON()
}

// DeleteExperienciaLaboral ...
// @Title DeleteExperienciaLaboral
// @Description eliminar Experiencia Laboral por id
// @Param   id      path    string  true        "Id de la Experiencia Laboral"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [delete]
func (c *ExperienciaLaboralController) DeleteExperienciaLaboral() {
    idStr := c.Ctx.Input.Param(":id")
    var alerta models.Alert
    //cadena de alertas
    alertas := append([]interface{}{"Cadena de respuestas: "})
    //resultado experiencia laboral
    var resultado []map[string]interface{}
    var resultado2 map[string]interface{}
    var resultado3 map[string]interface{}

    errSoporteExperienciaLaboral := request.GetJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/?query=ExperienciaLaboral:"+idStr, &resultado)
    if errSoporteExperienciaLaboral == nil {
        errDeleteSoporte := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/"+fmt.Sprintf("%.f", resultado[0]["Id"].(float64)), "DELETE", &resultado2, nil)
        if errDeleteSoporte == nil {
            alertas = append(alertas, "OK DELETE soporte_experiencia_laboral")
        }
        errDeleteExperiencia := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/experiencia_laboral/"+idStr, "DELETE", &resultado3, nil)
        if errDeleteExperiencia == nil {
            fmt.Println(resultado3)
            alertas = append(alertas, "OK DELETE experiencia_laboral")
        }
        alerta.Code = "200"
        alerta.Type = "success"
    } else {
        alertas = append(alertas, errSoporteExperienciaLaboral.Error())
        alerta.Code = "400"
        alerta.Type = "error"
    }
    alerta.Body = alertas
    c.Data["json"] = alerta
    c.ServeJSON()
}
