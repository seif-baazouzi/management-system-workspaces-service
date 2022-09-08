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
$ export DB_HOST="127.0.0.1"
$ export DB_NAME="workspaces"
$ export PORT="3001"
$ export ACCOUNTS_SERVICE="http://127.0.0.1:3000"

$ go run ./src/main.go
```

Using Docker-compose

```
$ docker-compose up --build
```

# Documentation

You can find the documentation for each route in the `docs` directory.
