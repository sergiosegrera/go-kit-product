# Go-Kit Product
A little microservice I made to learn go-kit.

The product microservice has a single endpoint that lists products.  
The endpoint is accessible by doing a get request at `/products`.  

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
products/service/ -- Business logic
products/transport/ -- Protocols for interfacing with the endpoints (http, gRPC, ...)
products/transport/http/ -- HTTP server to interface with endpoints
```

## Procedures to add a new endpoint
* Create models if needed in `/product/models/`.
* Add business logic in `product/service/`.
* Add endpoints in `product/endpoints/`.
* Add correct handlers for transports in `product/transport/`.

