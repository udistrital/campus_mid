Feature: Validate API responses
  CAMPUS_MID
  probe JSON reponses


Scenario Outline: To probe route code response  /organizacion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"

  Examples:
  |method|route           |bodyreq               |codres       |
  |GET   |/v1/organizacion|./files/req/Vacio.json|404 Not Found|
  |GET   |/v1/organizacio |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/organizacio |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/organizacio |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/organizacio |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /organizacion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"

  Examples:
  |method           |route                              |bodyreq               |codres         |bodyres                |
  |GETENTE          |/v1/organizacion/                  |./files/req/Vacio.json|200 OK         |./files/res6/Vok1.json |
  |GETIDENTIFICACION|/v1/organizacion/identificacion/   |./files/req/Vacio.json|200 OK         |./files/res6/Vok1.json |
  |POST             |/v1/organizacion/                  |./files/req/Vacio.json|400 Bad Request|./files/res6/Ierr1.json|
  |POSTUBICACION    |/v1/organizacion/registar_ubicacion|./files/req/Vacio.json|400 Bad Request|./files/res6/Ierr1.json|


# /v1/organizacion/:ente [get]
# /v1/organizacion/identificacion/?Id=int&TipoId=int [get]
# /v1/organizacion/ [post]
# /v1/organizacion/registar_ubicacion [post]
