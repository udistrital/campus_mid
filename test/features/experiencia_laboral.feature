Feature: Validate API responses
  CAMPUS_MID
  probe JSON reponses


Scenario Outline: To probe route code response  /experiencia_laboral
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>" 

  Examples: 
  |method|route                  |bodyreq               |codres       |
  |GET   |/v1/experiencia_laboral|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/experiencia_labora |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/experiencia_labora |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/experiencia_labora |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/experiencia_labora |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /experiencia_laboral
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"

  Examples:
  |method |route                   |bodyreq               |codres         |bodyres                |
  |GETENTE|/v1/experiencia_laboral/|./files/req/Vacio.json|200 OK         |./files/res3/Vok1.json |
  |GETID  |/v1/experiencia_laboral/|./files/req/Vacio.json|200 OK         |./files/res3/Vok2.json |
  |POST   |/v1/experiencia_laboral/|./files/req/Vacio.json|400 Bad Request|./files/res3/Ierr1.json|
  |PUT    |/v1/experiencia_laboral/|./files/req/Vacio.json|400 Bad Request|./files/res3/Ierr1.json|
  |DELETE |/v1/experiencia_laboral/|./files/req/Vacio.json|200 OK         |./files/res3/Ierr2.json|


# /v1/experiencia_laboral/:id [get]
# /v1/experiencia_laboral/?Ente=int [get]
# /v1/experiencia_laboral/ [post]
# /v1/experiencia_laboral/:id [put]
# /v1/experiencia_laboral/:id [delete]
