# WebCrawler

A high-performance web crawling service built with Go that provides REST API endpoints for web scraping and crawling operations.

## Features

- REST API for web crawling operations
- MongoDB integration for data persistence
- Docker containerization support
- Configuration management with Viper
- Dependency injection with Google Wire
- HTTP request middleware for request tracking

## Prerequisites

- Go 1.25.5 or later
- MongoDB
- Docker (optional, for containerized deployment)

## Project Structure

```
WebCrawler/
├── appinit/           # Application initialization
├── common/            # Shared utilities and common code
├── configs/           # Configuration management
├── internal/
│   ├── crawler/       # Crawler module (controller, service, repository)
│   ├── interfaces/    # Service interfaces
│   └── models/        # Data models
├── Dockerfile         # Docker container configuration
├── compose.yaml       # Docker Compose configuration
├── go.mod            # Go module dependencies
└── main.go           # Application entry point

crawler-core/         # Core crawling library (dependency)
```

## Quick Start

### 1. Clone the Repository

```bash
git clone <repository-url>
cd WebCrawler
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Configuration

Copy the example configuration file:

```bash
cp configs/config.example.yaml configs/config.yaml
```

Update the configuration values in `configs/config.yaml`:

```yaml
env: "local"
environment: "WebCrawler"

server:
  httpServerAddress: ":8086"

db:
  loglevel: "info"
  mongoDB:
    mongoURI: ${MONGODB_URI}
```

### 4. Set Environment Variables

Set the MongoDB connection string:

```bash
export MONGODB_URI="mongodb://username:password@localhost:27017"
```

### 5. Start MongoDB

#### Using Docker Compose (Recommended)

```bash
docker-compose up -d
```

This will start MongoDB with the default credentials configured in `compose.yaml`.

#### Or Using Local MongoDB

Ensure MongoDB is running locally on port 27017 with authentication configured.

### 6. Run the Application

#### Local Development

```bash
go run main.go
```

#### Using Docker

Build and run the application:

```bash
docker build -t webcrawler .
docker run -p 8086:8080 --env-file .env webcrawler
```

## API Usage

### Crawl Endpoint

**POST** `/crawl`

Starts a web crawling operation from a seed URL.

**Request Body:**
```json
{
  "seedurl": "https://example.com"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8086/crawl \
  -H "Content-Type: application/json" \
  -d '{"seedurl": "https://example.com"}'
```

**Response:**
```json
{
  "links": ["https://example.com/page1", "https://example.com/page2", ...]
}
```

## Configuration

The application uses a YAML configuration file located at `configs/config.yaml`. Key configuration options:

- `server.httpServerAddress`: Server listening address (default: ":8086")
- `db.mongoDB.mongoURI`: MongoDB connection string (set via environment variable)
- `env`: Environment name ("local", "development", "production")
- `db.loglevel`: Database logging level

## Docker Deployment

### Build the Image

```bash
docker build -t webcrawler .
```

### Run with Docker Compose

```bash
docker-compose up --build
```

This will start both the MongoDB database and the WebCrawler application.

### Environment Variables for Docker

Create a `.env` file in the project root:

```
MONGODB_URI=mongodb://daksh:password1234@mongodb:27017
```

## Development

### Code Generation

The project uses Google Wire for dependency injection. To regenerate the wire_gen.go file:

```bash
cd internal/crawler
wire
```

### Adding New Features

1. Define interfaces in `internal/interfaces/`
2. Implement services in `internal/crawler/service.go`
3. Add repository methods in `internal/crawler/repository.go`
4. Create controller methods in `internal/crawler/controller.go`
5. Wire dependencies in `internal/crawler/wire.go`

## Dependencies

- **Gin**: HTTP web framework
- **MongoDB Go Driver**: Database connectivity
- **Viper**: Configuration management
- **Google Wire**: Dependency injection
- **crawler-core**: Core crawling functionality (local module)

