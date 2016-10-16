# Gin API Demo

An example of using the Gin framework in Go to create a JSON api with most of the core functionality a JSON api needs.

## Installation

```bash
$ go get github.com/sbecker/gin-api-demo
$ go get github.com/tools/godep
```

## How to run

In one terminal, run the server:

```bash
$ cd $GOPATH/src/github.com/sbecker/gin-api-demo
$ godep go run main.go
```

In another terminal, try making requests:

```bash
# Get all users as a non-admin - see only one user, with a subset of json attributes
$ curl -H 'Authorization: Bearer userA' -i http://localhost:8080/users
```

```bash
# Get all users as an admin - see all users, with all json attributes
$ curl -H 'Authorization: Bearer userB' -i http://localhost:8080/users
```

```bash
# Try to get all users as a non-user - get 401 Unauthorized
$ curl -i http://localhost:8080/users
```

```bash
# Get a single user as a non-admin - see a subset of json attributes
$ curl -H 'Authorization: Bearer userA' -i http://localhost:8080/users/1
```

```bash
# Get a single user as an admin - see a subset of json attributes
$ curl -H 'Authorization: Bearer userB' -i http://localhost:8080/users/1
```

```bash
# Try to get a different user than self - get 404 Not Found
$ curl -H 'Authorization: Bearer userA' -i http://localhost:8080/users/2
```

```bash
# Try to get a single user as a non-user, get 401 Unauthorized
$ curl -i http://localhost:8080/users/2
```


## Done:
- [x] basic restful routing - GET /users, GET /users/:id
- [x] token authorization middleware
- [x] multi-level json serialization - display subset of json objects for non-admin users
- [x] separation of concerns - models, dao, resources, serializers, authorization
- [x] Glide for dependency management
- [x] UUIDs instead of integers for primary keys
- [x] `Request-Id` headers: Each API request has an associated request identifier
- [x] Cross-origin-request headers
- [x] Continuous integration
- [x] JSON logging
- [x] Env var configuration

## Todo:
- [ ] JWT token authorization
- [ ] Pagination
- [ ] Versioning
- [ ] Standardize error responses
- [ ] Integration with a real database
- [ ] Swagger live documentation
- [ ] Godoc
- [ ] A test suite
- [ ] Database migrations
- [ ] Continuous deployment
- [ ] Command line interface - run server, run migrations, start db console, etc
- [ ] CLI with commands for migration, db drop/create/reset

## Research:
- [ ] Client side query language - a la GraphQL/Falkor/Datomic
