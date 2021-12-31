
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

