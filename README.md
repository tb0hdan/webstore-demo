# Intro

This demo application is a simple REST API server that allows to create, read and sale products. Written in Go,
it uses [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) to generate API bindings and documentation from OpenAPI 3.0 specification.
HTTP server is implemented using [labstack echo](https://github.com/labstack/echo)

## API documentation

This project uses single source of truth for API documentation. The API documentation is written in OpenAPI 3.0 format and is located in `./api/api.yaml`.
Generated documentation is available at [http://localhost:8080/docs](http://localhost:8080/docs).

## Running server

```shell
make run
```

## Making requests

See `./examples`

## Updating API bindings and documentation

```shell
make api
```
