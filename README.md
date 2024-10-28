# Go Bookstore API

A simple RESTful API for managing a bookstore, built with Go and MySQL. Supports CRUD operations for book records.

## Features

- **Create**, **Retrieve**, **Update**, and **Delete** books.

## Technologies

- **Golang** - Backend
- **MySQL** - Database
- **Gorilla Mux** - Routing
- **GORM** - ORM

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-username/go-bookstore.git
   cd go-bookstore
   ```

2. **Set Up MySQL Database**

   - Create a MySQL database:
     ```sql
     CREATE DATABASE simplerest;
     ```
   - Update MySQL credentials in `pkg/config/app.go`:
     ```go
     user := "your_mysql_username"
     pass := "your_mysql_password"
     dbname := "your_mysql_dbname"
     ```

3. **Run the Application**

   ```bash
   go run cmd/main.go
   ```

## API Endpoints

| Method | Endpoint         | Description         |
|--------|-------------------|---------------------|
| GET    | `/book/`         | Get all books       |
| GET    | `/book/{bookId}` | Get a book by ID    |
| POST   | `/book/`         | Create a new book   |
| PUT    | `/book/{bookId}` | Update a book       |
| DELETE | `/book/{bookId}` | Delete a book       |

## Example Request

To create a book:

```http
POST http://localhost:9010/book/
Content-Type: application/json

{
    "name": "The Go Programming Language",
    "author": "Alan Donovan",
    "publication": "Addison-Wesley"
}
```

