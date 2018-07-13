package controllers

import (
    "encoding/json"
    "fmt"

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
// @router /ExperienciaLaboral [post]
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
            "Organizacion":      experiencia["Organizacion"],
            "TipoDedicacion":    experiencia["TipoDedicacion"].(map[string]interface{})["Id"],
            "Cargo":             experiencia["Cargo"].(map[string]interface{})["Id"],
            "TipoVinculacion":   experiencia["TipoVinculacion"].(map[string]interface{})["Id"],
        }
        errExperienciaLaboral := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/experiencia_laboral", "POST", &resultado, experienciaLaboral)
        //fmt.Println("el resultado es: ", resultado)
        if errExperienciaLaboral == nil && resultado["Type"] != "error" {
            alertas = append(alertas, "se agrego la experiencia laboral")

            experienciaLaboralSoporte := map[string]interface{}{
                "Soporte":            experiencia["Soporte"].(map[string]interface{})["Id"],
                "FormacionAcademica": resultado["Body"],
            }
            fmt.Println("el soporte es:", experienciaLaboralSoporte)
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
// @router /ExperienciaLaboral/:id [put]
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
            "Persona":           experiencia["Ente"].(map[string]interface{})["Id"],
            "Actividades":       experiencia["Actividades"],
            "FechaInicio":       experiencia["FechaInicio"],
            "FechaFinalizacion": experiencia["FechaFinalizacion"],
            "Organizacion":      experiencia["Organizacion"],
            "TipoDedicacion":    experiencia["TipoDedicacion"].(map[string]interface{})["Id"],
            "Cargo":             experiencia["Cargo"].(map[string]interface{})["Id"],
            "TipoVinculacion":   experiencia["TipoVinculacion"].(map[string]interface{})["Id"],
        }

        errExperienciaLaboral := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/experiencia_laboral/"+idStr, "PUT", &resultado, experienciaLaboral)
        if errExperienciaLaboral == nil {
            if resultado["Type"] == "success" {
                alertas = append(alertas, "OK UPDATE formacion_academica")
                alerta.Code = "200"
                alerta.Type = "success"
                alerta.Body = alertas
                c.Data["json"] = alerta
            }

            soporteExperienciaLaboral := map[string]interface{}{
                "Documento": experiencia["Soporte"].(map[string]interface{})["Documento"],
            }

            errSoporteExperienciaLaboral := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/"+fmt.Sprintf("%.f", experiencia["Soporte"].(map[string]interface{})["Documento"]), "PUT", &resultado2, soporteExperienciaLaboral)
            if errSoporteExperienciaLaboral == nil && resultado2["Type"] == "success" {
                alertas = append(alertas, "OK UPDATE Documento ")
            }

        } else {
            //fmt.Println("error de formacion", errFormacion2.Error())
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
// @router /ExperienciaLaboral/:id [get]
func (c *ExperienciaLaboralController) GetExperienciaLaboral() {
    //Id de la persona
    idStr := c.Ctx.Input.Param(":id")
    //formacion academica
    //var formacion map[string]interface{}
    //alerta que retorna la funcion GetExperienciaLaboral
    //var alerta models.Alert
    //cadena de alertas
    //alertas := append([]interface{}{"Cadena de respuestas: "})
    //resultado experiencia laboral
    var resultado map[string]interface{}
    //var resultado2 map[string]interface{}

    errExperienciaLaboral := request.GetJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/experiencia_laboral/"+idStr, &resultado)

    if errExperienciaLaboral == nil && resultado != nil {
        if resultado["Type"] != "error" {
            //soporte_experiencia_laboral
        }
    } else {

    }

    c.ServeJSON()
}

// DeleteExperienciaLaboral ...
// @Title DeleteExperienciaLaboral
// @Description eliminar Experiencia Laboral por id
// @Param   id      path    string  true        "Id de la Experiencia Laboral"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /ExperienciaLaboral/:id [delete]
func (c *ExperienciaLaboralController) DeleteExperienciaLaboral() {
    idStr := c.Ctx.Input.Param(":id")
    var alerta models.Alert
    //cadena de alertas
    alertas := append([]interface{}{"Cadena de respuestas: "})
    //resultado experiencia laboral
    var resultado map[string]interface{}
    var resultado4 []map[string]interface{}
    var resultado5 map[string]interface{}
    var resultado6 map[string]interface{}

    errExperienciaLaboral := request.GetJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/experiencia_laboral/"+idStr, &resultado)
    if errExperienciaLaboral == nil {

        if errExperienciaLaboral == nil {
            errExperienciaLaboral := request.GetJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/?query=ExperienciaLaboral:"+idStr, &resultado4)
            if errExperienciaLaboral == nil {
                //fmt.Println("el documento a borrar es ", resultado4[0]["Id"])
                err := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/"+fmt.Sprintf("%.f", resultado4[0]["Id"].(float64)), "DELETE", &resultado5, nil)
                if err == nil {
                    //fmt.Println("el resultado  del delete es:", resultado5)
                    alertas = append(alertas, "OK DELETE soporte_experiencia_laboral")
                }
            } else {
                alertas = append(alertas, errExperienciaLaboral.Error())
            }
            err := request.SendJson("http://"+beego.AppConfig.String("ExperienciaLaboralService")+"/soporte_experiencia_laboral/"+idStr, "DELETE", &resultado6, nil)
            if err == nil {
                fmt.Println("el resultado del DELETE es: ", resultado6)
                alertas = append(alertas, "OK DELETE experiencia_laboral")
            }
            alerta.Body = alertas
            alerta.Code = "200"
            alerta.Type = "success"
            c.Data["json"] = alerta
        }
    } else {
        alertas = append(alertas, errExperienciaLaboral.Error())
        alerta.Body = alertas
        alerta.Code = "400"
        alerta.Type = "error"
        c.Data["json"] = alerta

    }
    c.ServeJSON()
}
