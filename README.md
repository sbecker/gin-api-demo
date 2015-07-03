= Gin API Demo

An example of using the Gin framework in Go to create a JSON api with most of the core functionality a JSON api needs.


== How to run

In one terminal, run the server:

```
$ go run main.go
```

In another terminal, try making requests:

```
# Get all users as a non-admin - see only one user, with a subset of json attributes
$ curl -H 'Authorization: Bearer userA' -i http://localhost:8080/users
```

```
# Get all users as an admin - see all users, with all json attributes
$ curl -H 'Authorization: Bearer userB' -i http://localhost:8080/users
```

```
# Try to get all users as a non-user - get 401 Unauthorized
$ curl -i http://localhost:8080/users
```

```
# Get a single user as a non-admin - see a subset of json attributes
$ curl -H 'Authorization: Bearer userA' -i http://localhost:8080/users/1
```

```
# Get a single user as an admin - see a subset of json attributes
$ curl -H 'Authorization: Bearer userB' -i http://localhost:8080/users/1
```

```
# Try to get a different user than self - get 404 Not Found
$ curl -H 'Authorization: Bearer userA' -i http://localhost:8080/users/2
```

```
# Try to get a single user as a non-user, get 401 Unauthorized
$ curl -i http://localhost:8080/users/2
```


== Done:
- [x] basic restful routing - GET /users, GET /users/:id
- [x] token authorization middleware
- [x] multi-level json serialization - display subset of json objects for non-admin users
- [x] separation of concerns - models, dao, resources, serializers, authorization

== Todo:
- [ ] UUIDs instead of integers for primary keys
- [ ] JWT token authorization
- [ ] Pagination
- [ ] Integration with a real database
- [ ] Swagger live documentation
- [ ] Godoc
- [ ] A test suite
- [ ] Database migrations
- [ ] Continuous integration
- [ ] Continuous deployment
- [ ] Command line interface - run server, run migrations, start db console, etc
- [ ] Godep for dependency management
- [ ] Env var configuration
