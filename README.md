# campus_mid
API de documentos, Integración con CI
campus_mid master/develop
 ## Requirements
Go version >= 1.8.
 ## Preparation:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/campus_mid
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `CAMPUS_MID_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `CAMPUS_MID__PGUSER`: Usuario de la base de datos
 - `CAMPUS_MID__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `CAMPUS_MID__PGURLS`: Host de conexión
 - `CAMPUS_MID__PGDB`: Nombre de la base de datos
 - `CAMPUS_MID__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
CAMPUS_MID_HTTP_PORT=8088 PERSONAS_SERVICE=localhost:8080/v1 UBICACIONES_SERVICE=localhost:8085/v1 ENTE_SERVICE=localhost:8089/v1 FORMACION_ACADEMICA_SERVICE=localhost:8095/v1 bee run
