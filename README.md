# GO User Management RESTful API

### Developer Note
A simple CRUD API built with Go's standard [`net/http`](https://pkg.go.dev/net/http) library. This is my first Go API project. I didn't look at any Go best practices or examples beforehand and just implement this based on my experience building APIs in other languages. 

I built this project to get used to Go's syntax and to reflect on what I have learned throughout years of coding and try to set my own standard of doing things such as coding structure, response format, error handling, etc.

I wanted to play around with [Goroutines](https://go.dev/tour/concurrency/1), [Channels](https://go-tour-th.appspot.com/tour/concurrency/2), [WaitGroup](https://pkg.go.dev/sync#WaitGroup) , and [Mutex](https://go.dev/tour/concurrency/9) in this project, but couldn't think of any features that would actually benefit from concurrency. Maybe next time.

### Structure
- `constants`: app configs and constant values such as messages
- `db`: database scripts
- `dto`: data transfer objects
- `handlers`: business logic
- `models`: data structures to map database relations
- `repository`: methods to interact with database
- `routes`: API endpoints
- `utils`: helper functions and libraries (I was going to name it `libs` but too lazy to change it)
- `main.go`: startup server

### ENV template
```
APP_ENV=dev
APP_PORT=8080
APP_TOKEN_KEY=thisistokenkey
APP_HOST=localhost
APP_URL=http://localhost:8080

DB_HOST=db
DB_PORT=5432
DB_NAME=user_management
DB_USER=username
DB_PASS=password

# google gamil server
MAIL_USER=example@gmail.com
MAIL_PASS=this is the example
MAIL_HOST=smtp.gmail.com
MAIL_PORT=587
```

### Start Project

#### Development (Docker Compose)
1. APP_ENV=dev
2. Run with `docker compose`
```sh
docker compose --profile dev up --build
```

#### Production (Docker Compose)
1. APP_ENV=prod
2. Run with `docker compose`
```sh
docker compose up --build
```
