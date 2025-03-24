# QR Generator.

Este es un ejemplo de como usar Swagger y Echo en Golang.

## Instalación para poder correr Swagger:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Y luego, por cada cambio que se haga en el código, se debe correr el siguiente comando para que se actualice la documentación:

```bash
swag init
```

Y se puede formatear la documentación con el siguiente comando:

```bash
swag fmt
```

## Instalación de dependencias:

```bash
go mod tidy
```

## Correr el proyecto:

```bash
go run main.go
```