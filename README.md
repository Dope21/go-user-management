# GO User Management RESTful API

### Structure
- `constants`: app configs and constant values such as messages
- `db`: database scripts
- `dto`: data transfer objects
- `handlers`: business logic
- `models`: data structures to map database relations
- `repository`: methods to interact with database
- `routes`: API endpoints
- `libs`: helper functions and libraries
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
docker compose up dev db --build
```

#### Production (Docker Compose)
1. APP_ENV=prod
2. Run with `docker compose`
```sh
docker compose up prod db --build
```
