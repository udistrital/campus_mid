Feature: Validate API responses
  CAMPUS_MID
  probe JSON reponses


Scenario Outline: To probe route code response  /produccion_academica
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"

  Examples:
  |method|route                   |bodyreq               |codres       |
  |GET   |/v1/produccion_academica|./files/req/Vacio.json|404 Not Found|
  |GET   |/v1/produccion_academic |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/produccion_academic |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/produccion_academic |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/produccion_academic |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /produccion_academica
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"


  Examples:
  |method    |route                                            |bodyreq               |codres         |bodyres                |
  |GETPERSONA|/v1/produccion_academica/                        |./files/req/Vacio.json|200 OK         |./files/res8/Vok1.json |
  |POST      |/v1/produccion_academica/                        |./files/req/Vacio.json|400 Bad Request|./files/res8/Ierr1.json|
  |PUT       |/v1/produccion_academica/                        |./files/req/Vacio.json|400 Bad Request|./files/res8/Ierr1.json|
  |PUTAUTOR  |/v1/produccion_academica/estado_autor_produccion/|./files/req/Vacio.json|400 Bad Request|./files/res8/Ierr1.json|
  |DELETE    |/v1/produccion_academica/                        |./files/req/Vacio.json|404 Not Found  |./files/res8/Ierr2.json|


# /v1/produccion_academica/:persona [get]
# /v1/produccion_academica/ [post]
# /v1/produccion_academica/:id [put]
# /v1/produccion_academica/estado_autor_produccion/:id [put]
# /v1/produccion_academica/:id [delete]
