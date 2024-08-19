# Todo App Service

## Overview

This repository contains a Go-based backend service for a todo app. It is using a MySQL database. The service provides RESTful APIs to handle user auth, and CRUD todo.

## Project Structure
```
.
├── api
│   ├── handler
│   │   ├── todo_handler.go
│   │   └── user_handler.go
│   ├── presenter
│   │   ├── common.go
│   │   ├── todo.go
│   │   └── user.go
│   └── router
│       ├── todo_router.go
│       └── user_router.go
├── config
│   └── config.go
├── doc
│   └── 
├── middleware
│   └── auth.go
├── pkg
│   ├── todo
│   │   ├── entity.go
│   │   ├── repository.go
│   │   └── service.go
│   └── user
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Getting Started

1. Install Go: Download and install the latest version of Go from https://golang.org/dl/.

2. Install MySQL: Follow the instructions for your operating system to install MySQL. Create a new database and tables, the query migration is available on `database/migrations.sql`
3. Clone the repository:

```bash
git clone https://github.com/your-username/todo-app-service.git
cd todo-app-service
go run main.go
```
