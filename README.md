# Management System Workspaces Service

This is the workspaces service for my management system.

# Used Technologies

- Golang
- GoFiber
- Postgresql

# Quick Start

Started Manually

```console
$ export DB_USER="postgres"
$ export DB_PASSWORD="password"
$ export DB_HOST="172.17.0.1"
$ export DB_NAME="workspaces"
$ export JWT_SECRET="JWT_SECRET" # Must be the same as the accounts service

$ go run ./src/main.go
```

Using Docker-compose

```
$ docker-compose up --build
```

# Documentation

You can find the documentation for each route in the `docs` directory.
