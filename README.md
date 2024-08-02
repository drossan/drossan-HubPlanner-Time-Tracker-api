# HubPlanner Proxy API

## Descripción

Esta es una API para manejar las solicitudes a la API de HubPlanner mediante un proxy. La API está diseñada para ser
utilizada con una arquitectura hexagonal y puede ejecutarse en un entorno local o en AWS Lambda.

La motivación de este proyecto, es el poder consumir la API que no proporciona Hub Planner para insertar el tiempo de
trabajo.

## TODO

- [ ] Mejorar el manejo de errores.
- [ ] Añadir Swagger.
- [ ] Añadir soporte para autenticación avanzada.
- [ ] Optimizar el rendimiento de la API.
- [ ] Agregar pruebas unitarias.
- [ ] Documentar mejor el código.

## Características

- Proxy para solicitudes a la API de HubPlanner.
- Implementación con arquitectura hexagonal.
- Soporte para despliegue en AWS Lambda y ejecución local.

## Instalación

### Prerrequisitos

- Go instalado.

### Clonar el repositorio

```sh
git clone git@github.com:Secuoyas-Experience/HubPlanner-Time-Tracker-api-go.git
cd hubplanner-proxy-api
```

### Configurar el entorno

Crea un archivo `.env` en el directorio raíz con el siguiente contenido:

```env
ENV=
JWT_SECRET=
SERVER_ADDRESS=:1331

API_URI_COMPANY=
API_URL=
API_TOKEN=
```

> La variable API_URI_COMPANY tiene que hacer referencia a la url que se usa para el login en la propia página de
> HubPlanner
> La variable API_URL, será la que se nos proporciona en la documentación oficial:
> - https://hubplanner.com/hub-planner-api/
> - https://api-docs.hubplanner.com/
>

### Ejecutar la API localmente

```sh
go run cmd/server/main.go
```

### Desplegar en AWS Lambda

1. Compila el binario para Linux:

```sh
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap cmd/lambda/main.go
```

2. Crea un archivo ZIP:

```sh
zip function.zip bootstrap
```

3. Sube el archivo ZIP a AWS Lambda.

4. Configura la función Lambda para usar el manejador Go.

5. Configura las variables de entorno necesarias (por ejemplo JWT_SECRET).

#### Configura API Gateway

##### Crear una API

1. Ve a la consola de API Gateway.
2. Crea una nueva API REST.
3. Configura un recurso y método para tu API (por ejemplo, /api/v1/health).
    1. Hay que configura un recurso por cada unas de las tutas que tenga la API
   
##### Configurar el Método

1. Selecciona el recurso y método que has creado (por ejemplo, GET para /api/v1/health).

2. Haz clic en "Integration Request".
3. Activa el check de "Use Lambda Proxy integration".

##### Desplegar la API

1. En la consola de API Gateway, selecciona "Stages".
2. Despliega tu API en el stage correspondiente (por ejemplo, dev).

## Estructura del Proyecto

```
/hubplanner-proxy-api/
|-- cmd/
|   |-- lambda/
|   |   |-- main.go
|   |-- server/
|       |-- main.go
|-- go.mod
|-- go.sum
|-- .env
|-- adapters/
|   |-- controllers/
|   |   |-- hubplanner_controller.go
|   |   |-- health_controller.go
|   |-- routes/
|       |-- routes.go
|-- application/
|   |-- services/
|       |-- hubplanner_service.go
|-- domain/
|   |-- models/
|       |-- login_request.go
|       |-- user_response.go
|       |-- project.go
|       |-- category.go
|   |-- repositories/
|       |-- hubplanner_repository.go
|-- infrastructure/
|   |-- repositories/
|       |-- hubplanner_api_repository.go
|-- config/
|   |-- config.go
```

## Archivos Clave

- **main.go**: Punto de entrada para la API, tanto para AWS Lambda como para ejecución local.
- **hubplanner_controller.go**: Controlador para manejar las solicitudes a HubPlanner.
- **health_controller.go**: Controlador para verificar la salud de la API.
- **hubplanner_service.go**: Servicio que contiene la lógica de negocio.
- **hubplanner_repository.go**: Interfaz del repositorio de HubPlanner.
- **hubplanner_api_repository.go**: Implementación del repositorio para manejar las solicitudes a HubPlanner.

## Contribuciones

Las contribuciones son bienvenidas. Siéntete libre de abrir un issue o enviar un pull request.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.
