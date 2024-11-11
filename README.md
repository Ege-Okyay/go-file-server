# Go File Server

Built with Fiber in Go, this file server is an easy-to-use and safe file hosting solution. It is intended for both personal and business projects where cloud-based hosting is either too expensive or not preferred.

## Features
- **UUID Naming**: Each uploaded file is renamed with a unique UUID to avoid filename conflicts.
- **CORS Protection**: Only allows uploads from specified domains.
- **Self Hosted**: Easily deploy on your own VPS for full control without relying on third-party cloud providers.
- **REST API**: Simple endpoint to upload files.

## Requirements
- **Go** 1.18 or higher
- **Docker** (optional) / **VPS Server** to host the application

## Installation
1. **Clone the Repository**
```bash
git clone https://github.com/Ege-Okyay/go-file-server.git
cd go-file-server
```

2. **Install Dependencies**
```bash
go mod tidy
```

3. **Create and `.env` File**
```env
ALLOWED_ORIGIN=<your_url>
UPLOAD_PATH=uploads/
UPLOAD_URL=<your_url>/uploads/
PORT=8080
```

4. **Create Upload Directory**
```bash
mkdir uploads
```

## Usage
1. **Run the Server**
```bash
go run main.go
```

2. **Upload Files**
```bash
curl -X POST -F "file=@path/to/yourfile.jpg" https://yourserver.com/upload
```

3. **Access Files**
```
https://yourserver.com/uploads/{uuid.extension}
```

## Alternative: Run with Docker

In the root directory, create a `DockerFile`:
```dockerfile
# Use the official Golang image as the base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o file-server

# Expose the port the server will run on
EXPOSE 8080

# Command to run the executable
CMD ["./file-server"]
```

Create a Docker Compose File `docker-compose.yml` (Optional)

```yaml
version: '3'
services:
  file-server:
    build: .
    ports:
      - "8080:8080"
    environment:
      - ALLOWED_ORIGIN=https://yourwebsite.com
      - UPLOAD_PATH=/app/uploads/
    volumes:
      - ./uploads:/app/uploads
```

Build and Run the Docker Container
```bash
# Build the Docker image
docker build -t file-server .

# Run the Docker container
docker run -d -p 8080:8080 --name file-server --env-file .env -v $(pwd)/uploads:/app/uploads file-server

# If you're using Docker Compose, simply run
docker-compose up -d
```

## Folder Structure
```
file-server/
├── config/         # Configuration setup and environment variables
├── handlers/       # Route handlers
├── middleware/     # CORS and other middleware configurations
├── services/       # File saving logic and UUID generation
├── uploads/        # Directory where uploaded files are stored
└── main.go         # Application entry point
```