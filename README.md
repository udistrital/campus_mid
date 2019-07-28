# campus_mid
API mid para integración de servicios para proyecto de Campus virtual

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `campus_mid master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/campus_mid`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `CAMPUS_MID_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `[SERVICIO]_SERVICE`: Host de conexión con el (los) API a integrar

## Ejemplo
CAMPUS_MID_HTTP_PORT=8095 FORMACION_ACADEMICA_SERVICE=localhost:8098/v1 PERSONAS_SERVICE=localhost:8083/v1 ADMISION_SERVICE=localhost:8887/v1 EXPERIENCIA_LABORAL_SERVICE=localhost:8099/v1 ORGANIZACION_SERVICE=localhost:8097/v1 IDIOMAS_SERVICE=localhost:8103/v1 DOCUMENTOS_SERVICE=localhost:8094/v1 ENTE_SERVICE=localhost:8096/v1 UBICACIONES_SERVICE=localhost:8085/v1 CORE_SERVICE=localhost:8102/v1 SESIONES_SERVICE=localhost:8081/v1 PROGRAMA_ACADEMICO_SERVICE=localhost:8101/v1 FORMS_MANAGEMENT_SERVICE=localhost:9011/v1 NOTIFICACION_SERVICE=localhost:8081/v1 DESCUENTO_ACADEMICO_SERVICE=localhost:9013/v1 CONFIGURACION_SERVICE=localhost:8080/v1 bee run
