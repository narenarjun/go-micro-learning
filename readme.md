# Product API

This is a go rest api app, which uses gorilla-mux to sever a simple product api. The app is named as `go-micro-learning`. 

## Generating Swagger Documentation

Swagger documentation is generated from the code annotations inside the source using go-swagger.

Go swagger can be installed with the following command:

```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

You can generate the documentation using the command:

```
swagger generate spec -o ./swagger.yaml --scan-models
```

After running the application:

```
go run main.go
```

Docs are servered up with [Redoc](https://github.com/Redocly/redoc)

Swagger documentation can be viewed using the ReDoc UI in your browser at [http://localhost:9090/docs](http://localhost:9090/docs).


## Containerized 

The app is containerized and hosted in the [Quay Container registry](quay.io). 

To get the image:

```

docker pull quay.io/narendev/go-micro-learning

```

