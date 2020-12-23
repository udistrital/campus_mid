Feature: Validate API responses
  CAMPUS_MID
  probe JSON reponses


Scenario Outline: To probe route code response  /inscripcion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"

  Examples:
  |method|route          |bodyreq               |codres       |
  |GET   |/v1/inscripcion|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/inscripcio |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/inscripcio |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/inscripcio |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/inscripcio |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /inscripcion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"

  Examples:
  |method           |route                          |bodyreq               |codres         |bodyres                |
  |GETID            |/v1/inscripcion/               |./files/req/Vacio.json|200 OK         |./files/res5/Vok1.json |
  |GETIDENTIFICACION|/v1/inscripcion/identificacion/|./files/req/Vacio.json|200 OK         |./files/res5/Vok1.json |
  |GETPROGRAMA      |/v1/inscripcion/               |./files/req/Vacio.json|200 OK         |./files/res5/Vok1.json |
  |PUT              |/v1/inscripcion/               |./files/req/Vacio.json|400 Bad Request|./files/res5/Ierr1.json|


# /v1/inscripcion/:id [get]
# /v1/inscripcion/?Id=int&ProgramaId=int&PeriodoId=int  [get]
# /v1/inscripcion/identificacion/?Identificacion=string&ProgramaId=int&PeriodoId=int [get]
# /v1/inscripcion/:id [put]
