package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/astaxie/beego"
	"github.com/xeipuuv/gojsonschema"
)

// @resStatus codigo de respuesta a las solicitudes a la api
var resStatus string

//@resBody JSON de respuesta a las solicitudesde la api
var resBody []byte

//@opt opciones de godog
var opt = godog.Options{Output: colors.Colored(os.Stdout)}

var savepostres map[string]interface{}

//@TestMain para realizar la ejecucion con el comando go test ./test
func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

//@init inicia la aplicacion para realizar los test
func init() {
	run_bee()
	//pasa las banderas al comando godog
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

//@run_bee activa el servicio de la api para realizar los test
func run_bee() {
	parametros :=
		"CAMPUS_MID_HTTP_PORT=" + beego.AppConfig.String("httpport") +
			" PERSONAS_SERVICE=" + beego.AppConfig.String("PersonaService") +
			" UBICACIONES_SERVICE=" + beego.AppConfig.String("UbicacionesService") +
			" ENTE_SERVICE=" + beego.AppConfig.String("EnteService") +
			" ORGANIZACION_SERVICE=" + beego.AppConfig.String("OrganizacionService") +
			" FORMACION_ACADEMICA_SERVICE=" + beego.AppConfig.String("FormacionAcademicaService") +
			" EXPERIENCIA_LABORAL_SERVICE=" + beego.AppConfig.String("ExperienciaLaboralService") +
			" PROGRAMA_ACADEMICO_SERVICE=" + beego.AppConfig.String("ProgramaAcademicoService") +
			" INSCRIPCION_SERVICE=" + beego.AppConfig.String("InscripcionService") +
			" PRODUCCION_ACADEMICA_SERVICE=" + beego.AppConfig.String("ProduccionAcademicaService") +
			" DESCUENTO_ACADEMICO_SERVICE=" + beego.AppConfig.String("DescuentoAcademicoService") +
			" EVALUACION_INSCRIPCION_SERVICE=" + beego.AppConfig.String("EvaluacionInscripcionService") +
			" bee run"

	file, err := os.Create("script.sh")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "cd ..")
	fmt.Fprintln(file, parametros)

	wg := new(sync.WaitGroup)
	commands := []string{"sh script.sh &"}
	for _, str := range commands {
		wg.Add(1)
		go exe_cmd(str, wg)
	}
	time.Sleep(5 * time.Second)
	deleteFile("script.sh")
	wg.Done()
}

func deleteFile(path string) {
	// delete file
	err := os.Remove(path)
	if err != nil {
		fmt.Errorf("no se pudo eliminar el archivo")
	}
}

//@exe_cmd ejecuta comandos en la terminal
func exe_cmd(cmd string, wg *sync.WaitGroup) {
	//fmt.Println(cmd)
	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1]).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done()
}

//@AreEqualJSON comparar dos JSON si son iguales retorna true de lo contrario false
func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

//@toJson convierte string en JSON
func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

//@getPages convierte en un tipo el json
func getPages(ruta string) []byte {
	raw, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []byte
	c = raw
	return c
}

//@theResponseCodeShouldBe valida el codigo de respuesta
func theResponseCodeShouldBe(arg1 string) error {
	if resStatus != arg1 {
		return fmt.Errorf("se esperaba el codigo de respuesta .. %s .. y se obtuvo el codigo de respuesta .. %s .. ", arg1, resStatus)
	}
	return nil
}

//@theResponseShouldMatchJson valida el JSON de respuesta
func theResponseShouldMatchJson(arg1 string) error {
	div := strings.Split(arg1, "")

	pages := getPages(arg1)
	//areEqual, _ := AreEqualJSON(string(pages), string(resBody))
	if div[13] == "V" {
		schemaLoader := gojsonschema.NewStringLoader(string(pages))
		documentLoader := gojsonschema.NewStringLoader(string(resBody))
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if result.Valid() {
			return nil
		} else {
			return fmt.Errorf("Errores : %s", result.Errors())

			return nil
		}
	}
	if div[13] == "I" {
		areEqual, _ := AreEqualJSON(string(pages), string(resBody))
		if areEqual {
			return nil
		} else {
			return fmt.Errorf(" se esperaba el body de respuesta %s y se obtuvo %s", string(pages), resBody)
		}

	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is json "([^"]*)"$`, iSendRequestToWhereBodyIsJson)
	s.Step(`^the response code should be "([^"]*)"$`, theResponseCodeShouldBe)
	s.Step(`^the response should match json "([^"]*)"$`, theResponseShouldMatchJson)
}

//@iSendRequestToWhereBodyIsJson realiza la solicitud a la API
func iSendRequestToWhereBodyIsJson(method, endpoint, bodyreq string) error {
	var url string

	if method == "GET" || method == "POST" {
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint
	} else if method == "PUT" || method == "DELETE" {
		var Id float64

		str := strconv.FormatFloat(Id, 'f', 0, 64)
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + str
	} else if method == "GETID" {
		if endpoint == "/v1/descuento_academico/" {
			var PersonaId float64

			method = "GET"
			PersonaId = 27

			str := strconv.FormatFloat(PersonaId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + str
		} else if endpoint == "/v1/evaluacion_inscripcion/" {
			var InscripcionId float64

			method = "GET"
			InscripcionId = 3

			str := strconv.FormatFloat(InscripcionId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + str
		} else if endpoint == "/v1/inscripcion/" {
			var InscripcionId float64

			method = "GET"
			InscripcionId = 3

			str := strconv.FormatFloat(InscripcionId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + str
		} else if endpoint == "/v1/experiencia_laboral/" || endpoint == "/v1/formacion_academica/" {
			var Id float64

			method = "GET"
			Id = 1

			str := strconv.FormatFloat(Id, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + str
		}
	} else if method == "GETDEPENDENCIA" {
		var DependenciaId float64
		var PeriodoId float64

		method = "GET"
		DependenciaId = 21
		PeriodoId = 1

		dependencia := strconv.FormatFloat(DependenciaId, 'f', 0, 64)
		periodo := strconv.FormatFloat(PeriodoId, 'f', 0, 64)
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "?DependenciaId=" + dependencia + "&PeriodoId=" + periodo
	} else if method == "GETPERSONA" {
		if endpoint == "/v1/produccion_academica/" {
			var PersonaId float64

			method = "GET"
			PersonaId = 3

			persona := strconv.FormatFloat(PersonaId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + persona
		} else {
			var PersonaId float64
			var DependenciaId float64
			var PeriodoId float64

			method = "GET"
			PersonaId = 27
			DependenciaId = 21
			PeriodoId = 1

			persona := strconv.FormatFloat(PersonaId, 'f', 0, 64)
			dependencia := strconv.FormatFloat(DependenciaId, 'f', 0, 64)
			periodo := strconv.FormatFloat(PeriodoId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "?PersonaId=" + persona + "&DependenciaId=" + dependencia + "&PeriodoId=" + periodo
		}
	} else if method == "GETSOLICITUD" {
		var PersonaId float64
		var SolicitudId float64

		method = "GET"
		PersonaId = 27
		SolicitudId = 1

		persona := strconv.FormatFloat(PersonaId, 'f', 0, 64)
		solicitud := strconv.FormatFloat(SolicitudId, 'f', 0, 64)
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "?PersonaId=" + persona + "&SolicitudId=" + solicitud
	} else if method == "GETENTE" {
		if endpoint == "/v1/organizacion/" {
			var EnteId float64

			method = "GET"
			EnteId = 19

			ente := strconv.FormatFloat(EnteId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + ente
		} else if endpoint == "/v1/persona/consultar_persona/" || endpoint == "/v1/persona/consultar_contacto/" || endpoint == "/v1/persona/consultar_complementarios/" {
			var EnteId float64

			method = "GET"
			EnteId = 6

			ente := strconv.FormatFloat(EnteId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + ente
		} else {
			var EnteId float64

			method = "GET"
			EnteId = 13

			ente := strconv.FormatFloat(EnteId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "?Ente=" + ente
		}
	} else if method == "GETIDENTIFICACION" {
		if endpoint == "/v1/organizacion/identificacion/" {
			var TipoId float64

			method = "GET"
			TipoId = 5

			identificacion := "899999230"
			tipo := strconv.FormatFloat(TipoId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "?Id=" + identificacion + "&TipoId=" + tipo
		} else {
			var DependenciaId float64
			var PeriodoId float64

			method = "GET"
			DependenciaId = 21
			PeriodoId = 1

			identificacion := "7848932098"
			programa := strconv.FormatFloat(DependenciaId, 'f', 0, 64)
			periodo := strconv.FormatFloat(PeriodoId, 'f', 0, 64)
			url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "?Identificacion=" + identificacion + "&ProgramaId=" + programa + "&PeriodoId=" + periodo
		}
	} else if method == "GETPROGRAMA" {
		var InscripcionId float64
		var DependenciaId float64
		var PeriodoId float64

		method = "GET"
		InscripcionId = 3
		DependenciaId = 21
		PeriodoId = 1

		inscripcion := strconv.FormatFloat(InscripcionId, 'f', 0, 64)
		programa := strconv.FormatFloat(DependenciaId, 'f', 0, 64)
		periodo := strconv.FormatFloat(PeriodoId, 'f', 0, 64)
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "?Id=" + inscripcion + "&ProgramaId=" + programa + "&PeriodoId=" + periodo
	} else if method == "GETUSER" {
		method = "GET"

		user := "utest07"
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + "?User=" + user
	} else if method == "POSTUBICACION" {
		method = "POST"
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint
	} else if method == "PUTAUTOR" {
		var Id float64

		method = "PUT"

		str := strconv.FormatFloat(Id, 'f', 0, 64)
		url = "http://localhost:" + beego.AppConfig.String("httpport") + endpoint + str
	}

	pages := getPages(bodyreq)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(pages))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyr, _ := ioutil.ReadAll(resp.Body)
	resStatus = resp.Status
	resBody = bodyr

	return nil
}
