# SPYcats

# Spy Cats Project

This project implements a microservice architecture in Go, which interacts with PostgreSQL for data storage. Docker is used for containerization to isolate the environment, automate service startups, and manage database migrations.

## Description

The project features a Go API for handling data, utilizing PostgreSQL as the database. Docker is used to ensure the consistency of the environment and automate service management (e.g., migrations).



## Steps to Run

### 1. Clone the Repository

Clone the repository to your local machine:
```bash
git clone https://github.com/FREEGREAT/SPYcats.git
cd SPYcats
```

### 2. Verify Docker and Docker Compose Installation

Ensure Docker and Docker Compose are installed. Run the following commands to check:
```bash
docker --version
docker-compose --version
```

If Docker or Docker Compose is not installed, follow the official installation guide at docker.com.

### 3. Build Docker Images and Start Containers

Navigate to the root directory of the project where the docker-compose.yml file is located.
Run the following command to build images and start the containers:
```bash
docker-compose up --build
```

This will build the images for services and start containers for:
- PostgreSQL
- Migrations (for creating and updating database tables)
- Your Go API service

The migration container will run only once to set up the database.

### 4. Run Migrations (if needed)

Migrations are applied automatically. However, you can also run them manually if needed:
```bash
docker-compose run --rm migrate
```

This will apply all migrations to the database.

### 5. API Documentation

Detailed API documentation is available in Postman:
[Spy Cats API Documentation](https://documenter.getpostman.com/view/31399546/2sAYdcqXRT)

### Available Endpoints
- ROOT `localhost:8080` 
#### Cats
- GET `/api/cats` - Get all cats
- GET `/api/cats/{id}` - Get cat by ID
- POST `/api/cats` - Create new cat
- PATCH `/api/cats/{id}` - Update cat
- DELETE `/api/cats/{id}` - Delete cat

#### Missions
- GET `/api/missions` - Get all missions
- GET `/api/missions/{id}` - Get mission by ID
- POST `/api/missions` - Create new mission
- PATCH `/api/missions/{id}` - Update mission
- DELETE `/api/missions/{id}` - Delete mission

#### Notes
- GET `/api/notes` - Get all notes
- GET `/api/notes/{id}` - Get note by ID
- POST `/api/notes` - Create new note
- PATCH `/api/notes/{id}` - Update note
- DELETE `/api/notes/{id}` - Delete note

#### Targets
- GET `/api/targets` - Get all targets
- GET `/api/targets/{id}` - Get target by ID
- POST `/api/targets` - Create new target
- PATCH `/api/targets/{id}` - Update target
- DELETE `/api/targets/{id}` - Delete target

For detailed request/response examples, query parameters, and testing instructions, please refer to the [Postman Documentation](https://documenter.getpostman.com/view/31399546/2sAYdcqXRT).

[Previous sections (Project Structure, etc.) remain the same...]


### 6. Clean Up

After finishing with the project, stop the containers and clean up resources:
- Stop all containers:
```bash
docker-compose down
```
- Clean up unused images, containers, and networks:
```bash
docker system prune -f
```

## File Structure
```
├── cmd
│   └── main.go
├── config
│   └── config.yml
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── handler
│   │   ├── cat.go
│   │   ├── handler.go
│   │   ├── middleware
│   │   │   └── logger.go
│   │   ├── mission.go
│   │   ├── note.go
│   │   ├── response.go
│   │   └── target.go
│   ├── models
│   │   ├── cats.go
│   │   ├── mission.go
│   │   ├── note.go
│   │   ├── request.go
│   │   └── target.go
│   ├── services
│   │   ├── api
│   │   │   └── cat_api.go
│   │   ├── cat.go
│   │   ├── mission.go
│   │   ├── note.go
│   │   ├── service.go
│   │   └── target.go
│   └── storage
│       ├── repo
│       │   ├── cat_table.go
│       │   ├── mission_table.go
│       │   ├── note_table.go
│       │   └── target_table.go
│       └── storage.go
├── local.Dockerfile
├── pkg
│   ├── logger
│   │   └── zap.go
│   ├── pg_connection
│   │   └── postgre.go
│   ├── server
│   │   └── http.go
│   └── utils
│       ├── connection_tries.go
│       └── init_cfg.go
├── README.md
└── schema
    ├── 000001_init.down.sql
    └── 000001_init.up.sqls
```
## Minimal Requirements
- Docker
