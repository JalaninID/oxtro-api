# Dealls Job Interview

Backend system for a simple dating app. The system is built using Golang and PostgreSQL. in this project I use multiple protocols, GRPC and RestFull API. This project represents speed in feature development

## Prerequisites

- Golang version 1.21 or higher
- PostgreSQL
- Docker
- Docker Compose
- Makefile
- Golang Migrate
- Golang Mockery
- Buf Protoc

## Tech Stack

- Connect RPC

  Connect RPC is a library for building RPC servers and clients in Go. It's built on the gRPC framework, and it's designed to be easy to use and highly performant.

  Connect RPC support multiple protocols, such as HTTP/1.1, HTTP/2, and gRPC.

- Gin Gonic

  Gin is a web framework written in Go, Gin support RESTful API and http2 it's easy to use.

- Gorm

  Gorm is a library for building ORM in Go, it's easy to use and support multiple databases. Using gorm we can build a database model easily and fast, then focus for the business logic.

- Go Cron

  Go Cron is a library for building cron jobs in Go, it's easy to use and support multiple time format.
  I use this library to build a cron job for deleting the table interest that has been show for yesterday.

## How to Run

- Clone this repository
- Create a new file `.env` in the root directory
- Copy the content of `.env.example` to `.env`
- Fill the `.env` file with your configuration
- Run the mode development following command
  ```bash
  make gow
  ```
- Run in docker compose
  ```bash
  docker-compose up
  ```
- Connect your database and Run the migration
  ```bash
  make migration-up version=1
  ```

## How to Run Test

Run the test following command

```bash
make test
```

## ERD Diagram

![alt](./assets/ERD.png)

## Structure Project

- `assets` contains the assets for the project
- `config` contains the configuration for the project
- `constant` contains the constant for the project
- `crons` contains the cron job for the project
- `domain` contains the contract interface for the project
- `dto` contains the data transfer object for the project
- `gen` contains the generated file from the protobuf
- `injector` contains the dependency injection for the project
- `middleware` contains the middleware for the project
- `migration` contains the database migration for the project
- `model` contains the database model for the project
- `pkg` contains the package or utils for the project
- `proto` contains the protobuf file for the project
- `repository` contains the repository to access the database
- `routers` contains the router handler for the project
- `service` contains the service to handle business logic for the project
- `.env` contains the environment configuration for the project
- `.env.example` contains the example of the environment configuration for the project
- `.gitignore` contains the git ignore for the project
- `go.mod` contains the go module for the project
- `go.sum` contains the go sum for the project
- `main.go` contains the main file for the project
- `Makefile` contains the makefile for the project
- `buf.gen.yaml` contains the buf configuration for generate proto the project
- `buf.yaml` contains the buf configuration for the project
- `Dockerfile` contains the docker file for the project
- `docker-compose.yaml` contains the docker compose file for the project
- `README.md` contains the readme and documentation for the project