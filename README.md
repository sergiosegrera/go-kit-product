# Go-Kit Product
A little microservice I made to learn go-kit.


## HTTP Endpoints
The product microservice has a 2 endpoints that list products.

The `GET /products` endpoint lists all the products.  
The `GET /product/1` endpoint gives more information about a specific product.

## Installation
Install Docker compose first then:
```
$ git clone https://github.com/sergiosegrera/go-kit-product
$ cd go-kit-product
$ docker-compose up
```
Docker compose will start the database at port `5432` and the product microservice at port `8080`.

## Structure

```
product/ -- Microservice root
product/main.go
product/db/ -- Database initialization, Schema creation
product/endpoints/ -- Go-kit endpoints including request and response models
product/models/ -- Business logic data models
product/service/ -- Business logic
product/transport/ -- Protocols for interfacing with the endpoints (http, gRPC, ...)
product/transport/http/ -- HTTP server to interface with endpoints
product/transport/http/handlers/ -- Endpoint wrappers
```

## Procedures to add a new endpoint
* Create models if needed in `product/models/`.
* Add business logic in `product/service/`.
* Add endpoints in `product/endpoints/`.
* Add correct handlers for transports in `product/transport/{type}/handlers/`.

## TODO
* Finish implementing endpoints for product-manager
* Request input verification
* Resolve database models
* Auth server (gRPC, redis, JWT?)
* Easier way to convert models to proto and back?
