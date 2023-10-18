# GO-MICRO

This application was created for learning about microservices architecture.

## Features

- ✅ Using `Microservices` and `Vertical Slice Architecture` as a high level architecture
- ✅ Using `Event Driven Architecture` on top of RabbitMQ Message Broker and MassTransit library
- ✅ Using `gRPC` create gRPC api tranposrt
- ✅ Using `REST API` create REST api tranposrt

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

  ```bash
  brew install golang-migrate
  ```

### Setup infrastructure

- Run db migration create :

  ```bash
  make migrateup
  ```

- Run db migration up :

  ```bash
  make migrate.up
  ```

- Run db migration down :

  ```bash
  make migrate.down
  ```

### Compiler protobuf

- Protobuf compiler :

  ```bash
  make proto
  ```

### How to run

- Run Server :

  ```bash
  make migrateup
  ```
