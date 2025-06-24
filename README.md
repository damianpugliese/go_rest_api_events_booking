# Go REST API Events Booking

This project is a RESTful API built in Go using the [Gin](https://github.com/gin-gonic/gin) framework and a SQLite database. It provides a complete event booking system with user authentication and event management capabilities.

## Project Structure

```
go_rest_api_events_booking/
├── main.go           # Application entry point
├── db/               # Database connection and migration logic
├── models/           # Data models and database operations
├── handlers/         # HTTP request handlers
├── routes/           # Route definitions
├── utils/            # Utility functions
├── api.db            # SQLite database (git-ignored)
├── go.mod, go.sum    # Project dependencies
└── .gitignore        # Ignored files and folders
```

## Requirements
- Go 1.23+

## Installation & Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/damianpugliese/go_rest_api_events_booking.git
   cd go_rest_api_events_booking
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run .
   ```

4. The API will be available at: `http://localhost:8080`

## API Endpoints

### User Management
- `POST   /signup`         - Register a new user
- `POST   /login`          - User authentication

### Event Management
- `GET    /events`         - Get all events
- `GET    /events/:id`     - Get an event by ID
- `POST   /events`         - Create a new event
- `PUT    /events/:id`     - Update an existing event
- `DELETE /events/:id`     - Delete an event

## Features

- **User Authentication**: Secure user registration and login with password hashing
- **Event Management**: Full CRUD operations for events
- **SQLite Database**: Lightweight, file-based database
- **RESTful API**: Clean, REST-compliant endpoints
- **Gin Framework**: High-performance HTTP web framework
- **Input Validation**: Request validation and error handling

## Database Schema

The application automatically creates the following tables:
- **users**: User accounts with email and hashed passwords
- **events**: Event information including title, description, date, and location

## Notes
- The `api.db` database is created automatically and is ignored by git
- Passwords are securely hashed using bcrypt
- All endpoints return JSON responses with appropriate HTTP status codes
- The API includes comprehensive error handling and validation

---

Contributions and suggestions are welcome! 