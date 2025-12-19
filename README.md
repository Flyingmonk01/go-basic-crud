# Go Basic CRUD API

A simple RESTful API built with Go that provides CRUD (Create, Read, Update, Delete) operations for managing a movie collection.

## Features

- In-memory data storage for movies
- RESTful API endpoints
- JSON request/response format
- Health check endpoint

## Tech Stack

- **Language**: Go 1.25.1
- **Router**: [Gorilla Mux](https://github.com/gorilla/mux) v1.8.1

## Data Models

### Movie
```json
{
  "id": "string",
  "idbn": "string",
  "title": "string",
  "director": {
    "firstname": "string",
    "lastname": "string"
  }
}
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Health check |
| GET | `/movies` | Get all movies |
| GET | `/movie/{id}` | Get a single movie by ID |
| POST | `/movies` | Create a new movie |
| PUT | `/movie/{id}` | Update a movie by ID |
| DELETE | `/movie/{id}` | Delete a movie by ID |

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Flyingmonk01/go-basic-crud.git
cd go-basic-crud
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8000`

## Usage Examples

### Health Check
```bash
curl http://localhost:8000/
```

### Get All Movies
```bash
curl http://localhost:8000/movies
```

### Get Movie by ID
```bash
curl http://localhost:8000/movie/{id}
```

### Create a New Movie
```bash
curl -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{
    "idbn": "123456",
    "title": "Inception",
    "director": {
      "firstname": "Christopher",
      "lastname": "Nolan"
    }
  }'
```

### Update a Movie
```bash
curl -X PUT http://localhost:8000/movie/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "idbn": "123456",
    "title": "Inception Updated",
    "director": {
      "firstname": "Christopher",
      "lastname": "Nolan"
    }
  }'
```

### Delete a Movie
```bash
curl -X DELETE http://localhost:8000/movie/{id}
```

## Response Status Codes

- `200 OK` - Successful GET/PUT request
- `201 Created` - Successful POST request
- `204 No Content` - Successful DELETE request
- `400 Bad Request` - Invalid request body
- `404 Not Found` - Movie not found

## Notes

- Movie IDs are auto-generated using Unix nanosecond timestamps
- Data is stored in-memory and will be lost when the server restarts
- No database persistence is implemented

## License

This project is open source and available for educational purposes.
