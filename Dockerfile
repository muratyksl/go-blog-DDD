FROM golang:latest

WORKDIR /app

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

EXPOSE 8080

# Create a script to run migrations and start the app
COPY start.sh /start.sh
RUN chmod +x /start.sh

CMD ["/start.sh"]