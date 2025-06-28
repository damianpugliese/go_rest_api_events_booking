# Go REST API Events Booking

This project is a RESTful API built in Go using the [Gin](https://github.com/gin-gonic/gin) framework and a SQLite database. It provides a complete event booking system with user authentication and event management capabilities.

## Project Structure

```
go_rest_api_events_booking/
├── main.go           # Application entry point
├── db/               # Database connection and migration logic
│   └── db.go
├── models/           # Data models and database operations
│   ├── event.go
│   └── user.go
├── handlers/         # HTTP request handlers
│   ├── events.go
│   └── users.go
├── routes/           # Route definitions
│   ├── events.go
│   └── users.go
├── middlewares/      # HTTP middleware functions
│   └── auth.go
├── utils/            # Utility functions
│   ├── hash.go
│   └── jwt.go
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

3. Set up environment variables:
   
   Create a `.env` file in the project root with the following content:
   ```env
   JWT_SECRET_KEY=your-super-secret-jwt-key-here
   ```
   
   **Important**: Replace `your-super-secret-jwt-key-here` with a strong, unique secret key. This key is used to sign and verify JWT tokens for user authentication.
   
   Alternatively, you can set the environment variable directly:
   ```bash
   export JWT_SECRET_KEY="your-super-secret-jwt-key-here"
   ```

4. Run the application:
   ```bash
   go run .
   ```

5. The API will be available at: `http://localhost:8080`

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
- **JWT Authentication**: JSON Web Token-based authentication for secure API access
- **Route Protection**: Protected routes using JWT middleware for event management operations
- **Event Management**: Full CRUD operations for events
- **SQLite Database**: Lightweight, file-based database using [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
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