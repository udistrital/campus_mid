


Feature: Validate API responses
  CAMPUS_MID
  probe JSON reponses


Scenario Outline: To probe route code response  /evaluacion_inscripcion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>" 

  Examples: 
  |method|route                     |bodyreq               |codres       |
  |GET   |/v1/evaluacion_inscripcion|./files/req/Vacio.json|404 Not Found|
  |GET   |/v1/evaluacion_inscripcio |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/evaluacion_inscripcio |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/evaluacion_inscripcio |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/evaluacion_inscripcio |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /evaluacion_inscripcion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"

  Examples:
  |method|route                      |bodyreq               |codres|bodyres               |
  |GETID |/v1/evaluacion_inscripcion/|./files/req/Vacio.json|200 OK|./files/res2/Vok1.json|


# /v1/evaluacion_inscripcion/:inscripcion_id [get]
