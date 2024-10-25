# Go CRUD API

A RESTful API built in Go using Gin and GORM, following clean architecture principles. This API supports basic CRUD operations for managing a collection of people, with validation, error handling, and CORS enabled for cross-origin requests.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Folder Structure](#folder-structure)
- [Setup Instructions](#setup-instructions)

## Overview
This project is a simple CRUD API for managing a collection of `Person` records, including attributes like name, age, and hobbies. It adheres to clean architecture principles, providing a scalable and maintainable codebase structure. The API also supports pagination and sorting and includes Swagger documentation for ease of use.

## Features
- **Create, Read, Update, Delete (CRUD)**: Full support for CRUD operations on `Person` records.
- **Validation**: Input validation for fields such as name and age.
- **Pagination & Sorting**: Paginated and sortable results for fetching multiple records.
- **Error Handling**: Advanced error handling with structured responses.
- **CORS Support**: Cross-Origin Resource Sharing enabled for broader accessibility.
- **Swagger Documentation**: Detailed API documentation generated using Swagger.

## Technologies Used
- **Go**: Core language
- **Gin**: HTTP web framework
- **GORM**: ORM for database interactions
- **SQLite**: In-memory database
- **Swagger**: API documentation
- **UUID**: For unique identifier generation

## Folder Structure
```
go_crud/
├── cmd/                   # Application entry point
├── internal/
│   ├── domain/            # Domain models
│   ├── repository/        # Database interaction (repository layer)
│   ├── service/           # Request validation structs
│   ├── usecase/           # Business logic (use case layer)
│   ├── handler/           # HTTP handlers
│   └── routes/            # Routes setup
└── docs/                  # Swagger documentation files
```

## Setup Instructions
### Prerequisites
- Go 1.16+
- SQLite or compatible in-memory database

### Installation
1. **Clone the Repository**
   ```bash
   git clone https://github.com/your-username/go_crud.git
   cd go_crud
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Run Migrations**
   - Use GORM’s `AutoMigrate` in `main.go` to set up the `Person` schema in the in-memory database.

4. **Run the Application**
   ```bash
   go run cmd/main.go
   ```

5. **View Swagger Documentation**
   - Once the server is running, navigate to `http://localhost:8080/swagger/index.html` to explore the API documentation.

### Endpoints
- **GET /person**: Retrieve paginated and sorted list of persons.(e.g: localhost:8080/person?page=1&limit=10&sortedBy=age&sortedOrder=asc)
- **GET /person/{id}**: Retrieve a single person by ID.
- **POST /person**: Create a new person.
- **PUT /person/{id}**: Update an existing person.
- **DELETE /person/{id}**: Delete a person by ID.
