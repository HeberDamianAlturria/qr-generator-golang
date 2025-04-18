# QR Generator.

Este es un ejemplo de como usar Swagger y Echo en Golang. La idea del proyecto es poder generar un código QR a partir de una URL y poder también obtener la información a partir de un código QR.

## Cómo correr este proyecto:

Para poder correr este proyecto, debemos ejecutar los siguientes pasos:

1. Instalación para poder correr Swagger:

    El proyecto utiliza Swagger para documentar la API. Para poder correr Swagger, se debe instalar el siguiente paquete:

    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

    Y luego, se deben obtener los docs pertienentes de la siguiente forma:

    ```bash
    swag init
    ```

    por cada cambio que se haga en el código, se debe ejecutar el comando anterior para poder actualizar los docs.


2. Instalación de dependencias:

    Para poder instalar las dependencias del proyecto, se debe ejecutar el siguiente comando:

    ```bash
    go mod tidy
    ```

3. Correr el proyecto:

    Para poder correr el proyecto, se debe ejecutar el siguiente comando:

    ```bash
    go run main.go
    ```

    El proyecto se ejecutará en el puerto `8080` por defecto, por lo que podemos acceder a la API desde `http://localhost:8080`.

    También podemos acceder a la documentación de Swagger desde `http://localhost:8080/swagger/index.html`.