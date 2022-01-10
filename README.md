
# Bookings and Reservations
This is the repository for my bookings and reservations project.

- Built in Go version 1.17
- Uses the [chi](https://github.com/go-chi/chi) router
- Uses alex edwards [SCS](https://github.com/alexedwards/scs) session management
- Uses [no surf](https://github.com/justinas/nosurf) for CSRF prevention

To run the application and compile everything:

```bash
go run *.go
```

OR

```bash
go run cmd/web/*.go
```

To remove unused packages run command
```bash
go mod tidy
```

To run tests
```bash
go test
```


Command for checking for test coverage
```bash
go test -cover
```

```bash
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

To run all tests make sure you are on your root directory and run the following command:
```shell
go test -v ./...
```

### Steps to create an easy executable file to run our app
1. Create a `run.sh` file in the root directory
```
#!/bin/bash

go build -o bookings cmd/web/*.go && ./bookings
```
2. Run command on terminal, this should create a script called bookings.
```shell
chmod +x run.sh
```

3. To run our script/build and run our application:
```shell
./run.sh
```

### Database
To run migration:
```shell
soda migrate
```

To reverse migration
```shell
soda migrate down
```