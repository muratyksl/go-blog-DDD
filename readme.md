# Post Management API

This is a Go-based RESTful API for managing posts. It provides endpoints for creating, retrieving, and deleting posts, with a PostgreSQL database backend.

## Features

- CRUD operations for posts
- PostgreSQL database integration
- Docker and Docker Compose support
- Database migrations
- Environment variable configuration
- Retry mechanism for database connection

## Project Structure

```
.
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── post
│   │   ├── handler
│   │   ├── repository
│   │   └── service
│   └── server
├── pkg
│   └── database
│       └── postgres.go
├── migrations
├── Dockerfile
├── docker-compose.yml
├── start.sh
├── go.mod
├── go.sum
└── README.md
```

## Prerequisites

- Go 1.16 or later
- Docker and Docker Compose
- PostgreSQL (if running without Docker)

## Environment Variables

The application uses the following environment variables:

- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port
- `DB_USER`: PostgreSQL user
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name

These are set in the `docker-compose.yml` file for containerized deployment.

## Getting Started

### Running with Docker Compose

1. Clone the repository:

   ```
   git clone https://github.com//muratyksl/go-blog-DDD.git
   cd go-blog-DDD
   ```

2. Start the application and database:
   ```
   docker-compose up --build
   ```

The API will be available at `http://localhost:8080`.

### Running Locally

1. Set up a PostgreSQL database and note the connection details.

2. Set the environment variables (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME).

3. Run the migrations:

   ```
   migrate -path ./migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up
   ```

4. Build and run the application:
   ```
   go build -o main cmd/api/main.go
   ./main
   ```

## Database Migrations

Migrations are automatically run when the application starts in Docker. The `start.sh` script handles running migrations before starting the main application.

To run migrations manually:

```
migrate -path ./migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up
```

## API Endpoints

- `GET /posts`: Retrieve all posts
- `GET /posts/{id}`: Retrieve a specific post
- `POST /posts`: Create a new post
- `DELETE /posts/delete?ids=1,2,3`: Delete multiple posts

## Development

### Adding New Migrations

To add a new migration:

```
migrate create -ext sql -dir migrations -seq your_migration_name
```

This will create up and down migration files in the `migrations` directory.

### Modifying the Database Connection

The database connection is managed in `pkg/database/postgres.go`. It includes a retry mechanism to handle initial connection failures.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
