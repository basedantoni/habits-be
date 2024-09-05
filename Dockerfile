FROM golang:1.23

WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Install the required packages
RUN apt-get update && apt-get install -y sqlite3
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the source code into the container
COPY . .

# Compile the application and name the binary as 'app'
RUN go build -v -o app .

# Set the environment variable for the database URL
ENV DATABASE_URL="sqlite:/etc/habits/habits.db"

# Expose the port the app listens on
EXPOSE 8080

# Grant execute permission to entrypoint.sh
RUN chmod +x ./entrypoint.sh

# Set entrypoint
ENTRYPOINT ["./entrypoint.sh"]