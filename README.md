# Go-Joke Worker Service

## go-joke-w-service

This project is a **Golang-based worker service** that fetches programming and software engineering jokes either from a predefined configuration or an external API. It includes a metrics server for monitoring using Prometheus and supports runtime configuration updates, centralized error handling, and graceful shutdown.

---

## Features

- **Fetch Jokes:** Periodically fetch jokes from:
  - Predefined configuration [Every 10 seconds].
  - External API [Every 15 seconds].
- **Configuration Updates:** Monitors configuration changes at runtime and applies them dynamically.
- **Metrics Server:** Prometheus-compatible metrics for monitoring application performance.
- **Centralized Error Handling:** Unified error management across the application.
- **Graceful Shutdown:** Ensures clean termination of running processes, including metrics server.
- **Docker Support:** Fully containerized for deployment on Docker.

---

## Project Structure

```plaintext
.
├── cmd
│   └── main.go                # Entry point of the application
├── config
│   └── config.json            # Application configuration file
├── internal
│   ├── business
│   │   └── timer_service.go   # Timer-based joke fetching logic
│   ├── infra
│   │   ├── config.go          # Configuration handling and validation
│   │   ├── metrics.go         # Metrics server management
│   │   ├── logger.go          # Logging infrastructure
│   │   ├── error.go           # Centralized error handling
│   └── models
│       └── config.go          # Configuration structure definition
├── Dockerfile                 # Dockerfile for building the service
├── docker-compose.yml         # Docker Compose configuration
└── README.md                  # Project documentation
```

---

## Prerequisites

Before running this service, ensure you have the following installed:

- **Go:** Version 1.20+ installed.
- **Docker & Docker Compose:** Installed and running.
- **Prometheus:** For metrics scraping (optional).

---

## Installation

### Clone the Repository

```bash
git clone https://github.com/ankitkmrpatel/go-joke-w-service.git
cd go-joke-w-service
```

## Configuration

### `config/config.json`

The application uses a configuration file located at (`config/config.json`) to manage the following:

- Predefined jokes list.
- Path to the log file for storing fetched jokes.
- Metrics server address for Prometheus.

Below is an example of the configuration structure:

```json
{
  "jokes": [
    "Why do programmers prefer dark mode? Because light attracts bugs.",
    "How do you comfort a JavaScript bug? You console it."
  ],
  "log_file_path": "/log/file.log",
  "metrics_server": "localhost:9090"
}
```

- **jokes**: List of predefined jokes.
- **log_file_path**: Path to the log file.
- **metrics_server**: Address for the metrics server.Configuration

---

## Running Locally

### Build the Project

```bash
go mod tidy
go build -o go-joke-service ./cmd/main.go
```

### Run the Service

```bash
./go-joke-service
```

---

## Running with Docker

### Build the Docker Image

```bash
docker build -t go-joke-service .
```

## Run the Service

```bash
docker-compose up
```

This starts the service and metrics server in Docker. The logs can be accessed using:

```bash
docker logs -f go-joke-service
```

---

## Features Overview

### Jokes Fetching

- **From Configuration:** Fetches jokes every 10 seconds from the predefined list in config.json.
- **From API:** Fetches jokes every 15 seconds from an external jokes API.

### Metrics

Metrics are exposed at `http://localhost:9090/metrics` (default).

- `jokes_from_config_total`: Number of jokes fetched from the configuration.
- `jokes_from_api_total`: Number of jokes fetched from the API.
- `config_reload_total`: Number of configuration reloads.

### Configuration Reload

The `WatchConfig` function monitors `config/config.json` for changes. Any updates are applied at runtime without restarting the application.

---

## Development and Testing

For development, the service includes:

- Unit tests for business logic and infrastructure components.
- Mock interfaces for dependency injection and testing.

You can run the unit tests to validate functionality and test the configuration watcher by modifying the `config.json` file while the service is running.

### Run Tests

Run unit tests to ensure application correctness:

```bash
go test ./...
```

### Test Configuration Watcher

Modify `config/config.json` while the service is running to test dynamic configuration updates.

---

## Deployment

This service can be deployed using Docker. Prometheus can also run in Docker to monitor the service metrics.

Deployment steps include:

- Building and pushing the Docker image to a registry.
- Using Docker Compose for service orchestration.

### Steps for Deployment in Docker

1. **Build and Push Docker Image:**

   ```bash
   docker build -t go-joke-service .
   docker tag go-joke-service <dockerhub_username>/go-joke-service
   docker push <dockerhub_username>/go-joke-service
   ```

2. **Run Using Docker Compose:**
   Update docker-compose.yml if necessary and deploy:
   ```bash
   docker-compose up -d
   ```

---

## Graceful Shutdown

The service supports graceful shutdown by:

- Metrics server stops listening for requests.
- Ensuring timers complete pending operations.
- Configuration watcher stops monitoring.

Simply send a termination signal (`Ctrl+C` or `docker-compose down`), and the service will shut down cleanly.

---

## Improvements and Future Enhancements

1. **Centralized Logging Aggregation:**
   Integrate with logging platforms like ELK Stack or Grafana Loki.

2. **Enhanced Error Reporting:**
   Add support for reporting errors to external monitoring tools like Sentry or Datadog.

3. **Horizontal Scalability:**
   Add support for distributed processing using RabbitMQ or Kafka.

4. **API Jokes Configuration:**
   Allow the jokes API URL to be configured dynamically.

---

## Acknowledgments

- **Golang Community**
- **Prometheus** for reliable metrics monitoring
- **Logrus** for simple and powerful logging

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
