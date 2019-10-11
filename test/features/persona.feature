Feature: Validate API responses
  CAMPUS_MID
  probe JSON reponses


Scenario Outline: To probe route code response  /persona
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"

  Examples:
  |method|route      |bodyreq               |codres       |
  |GET   |/v1/persona|./files/req/Vacio.json|404 Not Found|
  |GET   |/v1/person |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/person |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/person |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/person |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /persona
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"

  Examples:
  |method|route                                  |bodyreq               |codres         |bodyres                |
  |GETENTE|/v1/persona/consultar_persona/        |./files/req/Vacio.json|200 OK         |./files/res7/Vok1.json |
  |GETENTE|/v1/persona/consultar_contacto/       |./files/req/Vacio.json|200 OK         |./files/res7/Vok2.json |
  |GETENTE|/v1/persona/consultar_complementarios/|./files/req/Vacio.json|200 OK         |./files/res7/Vok3.json |
  |GETUSER|/v1/persona/consultar_persona/        |./files/req/Vacio.json|200 OK         |./files/res7/Vok1.json |
  |POST   |/v1/persona/guardar_persona           |./files/req/Vacio.json|400 Bad Request|./files/res7/Ierr1.json|
  |POST   |/v1/persona/guardar_contacto          |./files/req/Vacio.json|400 Bad Request|./files/res7/Ierr1.json|
  |POST   |/v1/persona/guardar_complementarios   |./files/req/Vacio.json|400 Bad Request|./files/res7/Ierr1.json|
  |PUT    |/v1/persona/actualizar_persona        |./files/req/Vacio.json|404 Not Found  |./files/res7/Ierr2.json|
  |PUT    |/v1/persona/actualizar_contacto       |./files/req/Vacio.json|404 Not Found  |./files/res7/Ierr2.json|
  |PUT    |/v1/persona/actualizar_complementarios|./files/req/Vacio.json|404 Not Found  |./files/res7/Ierr2.json|


# /v1/persona/consultar_persona/:ente_id [get]
# /v1/persona/consultar_persona/?User=string [get]
# /v1/persona/consultar_contacto/:ente_id [get]
# /v1/persona/consultar_complementarios/:ente_id [get]
# /v1/persona/guardar_persona [post]
# /v1/persona/guardar_contacto [post]
# /v1/persona/guardar_complementarios [post]
# /v1/persona/actualizar_persona [put]
# /v1/persona/actualizar_contacto [put]
# /v1/persona/actualizar_complementarios [put]
