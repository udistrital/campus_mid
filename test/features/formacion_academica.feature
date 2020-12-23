Feature: Validate API responses
  CAMPUS_MID
  probe JSON reponses


Scenario Outline: To probe route code response  /formacion_academica
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"

  Examples:
  |method|route                  |bodyreq               |codres       |
  |GET   |/v1/formacion_academica|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/formacion_academic |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/formacion_academic |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/formacion_academic |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/formacion_academic |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /formacion_academica
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"

  Examples:
  |method |route                   |bodyreq               |codres         |bodyres                |
  |GETENTE|/v1/formacion_academica/|./files/req/Vacio.json|200 OK         |./files/res4/Vok1.json |
  |GETID  |/v1/formacion_academica/|./files/req/Vacio.json|200 OK         |./files/res4/Vok2.json |
  |POST   |/v1/formacion_academica/|./files/req/Vacio.json|400 Bad Request|./files/res4/Ierr1.json|
  |PUT    |/v1/formacion_academica/|./files/req/Vacio.json|400 Bad Request|./files/res4/Ierr1.json|
  |DELETE |/v1/formacion_academica/|./files/req/Vacio.json|404 Not Found  |./files/res4/Ierr2.json|


# /v1/formacion_academica/ [post]
# /v1/formacion_academica/:id [put]
# /v1/formacion_academica/:id [get]
# /v1/formacion_academica/?Ente=int [get]
# /v1/formacion_academica/:id [delete]
