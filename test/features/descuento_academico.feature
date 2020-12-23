Feature: Validate API responses
  CAMPUS_MID
  probe JSON reponses


Scenario Outline: To probe route code response  /descuento_academico
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>" 

  Examples: 
  |method|route                  |bodyreq               |codres       |
  |GET   |/v1/descuento_academico|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/descuento_academic |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/descuento_academic |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/descuento_academic |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/descuento_academic |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /descuento_academico
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"

  Examples:
  |method        |route                                                      |bodyreq               |codres         |bodyres                |
  |GETID         |/v1/descuento_academico/                                   |./files/req/Vacio.json|200 OK         |./files/res1/Vok1.json |
  |GETDEPENDENCIA|/v1/descuento_academico/descuentodependenciaperiodo/       |./files/req/Vacio.json|200 OK         |./files/res1/Vok1.json |
  |GETPERSONA    |/v1/descuento_academico/descuentopersonaperiododependencia/|./files/req/Vacio.json|200 OK         |./files/res1/Vok1.json |
  |GETSOLICITUD  |/v1/descuento_academico/                                   |./files/req/Vacio.json|200 OK         |./files/res1/Vok2.json |
  |POST          |/v1/descuento_academico/                                   |./files/req/Vacio.json|400 Bad Request|./files/res1/Ierr1.json|
  |PUT           |/v1/descuento_academico/                                   |./files/req/Vacio.json|400 Bad Request|./files/res1/Ierr1.json|
  |DELETE        |/v1/descuento_academico/                                   |./files/req/Vacio.json|200 OK         |./files/res1/Ierr2.json|


# /v1/descuento_academico/:persona_id [get]
# /v1/descuento_academico/?PersonaId=int&SolicitudId=int [get]
# /v1/descuento_academico/descuentodependenciaperiodo/?DependenciaId=int&PeriodoId=int [get]
# /v1/descuento_academico/descuentopersonaperiododependencia/?PersonaId=int&DependenciaId=int&PeriodoId=int [get]
# /v1/descuento_academico/ [post]
# /v1/descuento_academico/:id [put]
# /v1/descuento_academico/:id [delete]
