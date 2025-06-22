# Go REST API Events Booking

This project is a RESTful API built in Go using the [Gin](https://github.com/gin-gonic/gin) framework and a SQLite database. It allows you to manage events and users.

## Project Structure

```
rest_api_project/
├── main.go           # Application entry point
├── db/               # Database connection and migration logic
├── models/           # Data models and access logic
├── api.db            # SQLite database (git-ignored)
├── go.mod, go.sum    # Project dependencies
└── .gitignore        # Ignored files and folders
```

## Requirements
- Go 1.21+

## Installation & Usage

1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd rest_api_project
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run main.go
   ```
4. The API will be available at: `http://localhost:8080`

## Main Endpoints
- `GET    /events`         - List all events
- `GET    /events/:id`     - Get an event by ID
- `POST   /events`         - Create a new event

## Notes
- The `api.db` database is created automatically and is ignored by git.
- You can modify models and migrations in the `models/` and `db/` folders respectively.

---

Contributions and suggestions are welcome! 