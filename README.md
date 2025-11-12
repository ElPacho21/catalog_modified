<!-- cSpell:language es -->

# Catalog Service en GO

Este es el microservicio de Catalogo del proyecto

[Microservicios](https://github.com/nmarsollier/ecommerce)

[Documentación de API](./README-API.md)

La documentación de las api también se pueden consultar desde el home del microservicio
que una vez levantado el servidor se puede navegar en [localhost:3002](http://localhost:3002/docs/index.html)

El servidor GraphQL puede navegar en [localhost:4002](http://localhost:4002/)

## Directorios

- **article:** Logica de negocio del agregado article
- **security:** Validaciones de usuario contra el MS de Auth
- **services:** Servicios de dominio para el negocio.
- **graph:** Servidor y Controllers GraphQL federation server
- **rabbit:** Servidor y Controllers RabbitMQ
- **rest:** Servidor y Controllers Rest
- **tools:** Herramientas varias

## Requisitos

Go 1.21 [golang.org](https://golang.org/doc/install)

## Configuración inicial

establecer variables de entorno (consultar documentación de la version instalada)

Para descargar el proyecto correctamente hay que ejecutar :

```bash
git clone https://github.com/nmarsollier/cataloggo $GOPATH/src/github.com/nmarsollier/cataloggo
```

Una vez descargado, tendremos el codigo fuente del proyecto en la carpeta

```bash
cd $GOPATH/src/github.com/nmarsollier/cataloggo
```

## Dependencias

### Auth

Las ordenes sólo puede usarse por usuario autenticados, ver la arquitectura de microservicios de [ecommerce](https://github.com/nmarsollier/ecommerce).

### Catálogo

Las ordenes tienen una fuerte dependencia del catalogo:

- Se validan los artículos contra el catalogo.
- Se descuentan los artículos necesarios.
- Se puede devolver articulos si la orden se cancela.

Ver la arquitectura de microservicios de [ecommerce](https://github.com/nmarsollier/ecommerce).

## Instalar Librerías requeridas

Configurar hooks de git

```bash
git config core.hooksPath .githooks
go install github.com/swaggo/gin-swagger/swaggerFiles
go install github.com/swaggo/gin-swagger
go install github.com/swaggo/swag/cmd/swag
go install github.com/golang/mock/mockgen@v1.6.0
go install github.com/99designs/gqlgen@v0.17.56
```

Build y ejecución

```bash
go install
cataloggo
```

## MongoDB

La base de datos se almacena en MongoDb.

Seguir las guías de instalación de mongo desde el sitio oficial [mongodb.com](https://www.mongodb.com/download-center#community)

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## RabbitMQ

Este microservicio notifica los logouts de usuarios con Rabbit.

Seguir los pasos de instalación en la pagina oficial [rabbitmq.com](https://www.rabbitmq.com/)

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## Swagger

Usamos [swaggo](https://github.com/swaggo/swag)

Requisitos

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

La documentacion la generamos con el comando

```bash
swag init
```

Para generar el archivo README-API.md

Requisito

```bash
sudo npm install -g swagger-markdown
```

y ejecutamos

```bash
npx swagger-markdown -i ./docs/swagger.yaml -o README-API.md
```

## Configuración del servidor

Este servidor usa las siguientes variables de entorno para configuración :

RABBIT_URL : Url de rabbit (default amqp://localhost)
MONGO_URL : Url de mongo (default mongodb://localhost:27017)
PORT : Puerto (default 3002)
GQL_PORT : Puerto GraphQL (default 4002)

## Docker

Estos comandos son para dockerizar el microservicio desde el codigo descargado localmente.

### Build

```bash
docker build -t dev-catalog-go .
```

### El contenedor

Mac | Windows

```bash
docker run -it --name dev-catalog-go -p 3002:3002 -p 4002:4002 -v $PWD:/go/src/github.com/nmarsollier/cataloggo dev-catalog-go
```

Linux

```bash
docker run -it --add-host host.docker.internal:172.17.0.1 --name dev-catalog-go -p 3002:3002 -p 4002:4002 -v $PWD:/go/src/github.com/nmarsollier/cataloggo dev-catalog-go
```
