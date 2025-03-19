# TMDB CLI Tool

TMDB CLI Tool is a simple command-line application that allows you to fetch information about movies from [The Movie Database (TMDB)](https://www.themoviedb.org/). The application fetches data in categories such as popular, top-rated, now playing, and upcoming movies, caches the results in Redis, and displays them in the terminal.

## Features

- Fetch movie information using the TMDB API.
- Support for various movie types: `playing` (now playing), `popular` (popular), `top` (top-rated), and `upcoming` (upcoming).
- Cache query results in Redis (TTL — 10 minutes) for faster subsequent requests.
- Configuration through environment variables (with support for `.env` file).

## Requirements

- Go 1.18 or higher.
- Redis — for caching (can be installed locally or run in a Docker container).
- TMDB API key — get it from TMDB (free).

## Installation and Running

### Clone the repository:

```bash
git clone https://github.com/che1nov/tmdb-cli.git
cd tmdb-cli
```

### Install dependencies:

```bash
go mod tidy
```

### Configuration:

Create a `.env` file in the root of the project and specify the environment variables (default values are used if variables are missing):

```env
TMDB_API_KEY=your_tmdb_api_key
PORT=8080
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

If the `.env` file is missing, the configuration will be loaded from system variables or default values will be used.

### Start the Redis server:

If Redis is not installed locally, you can run a container using Docker:

```bash
docker run --name redis -p 6379:6379 -d redis
```

### Run the application:

Using `go run`:

```bash
go run ./cmd/main.go --type popular
```

Using `Makefile`:

```bash
make run
```

## Usage

The application accepts the `--type` argument, which defines the type of movies to fetch. Usage examples:

```bash
tmdb-cli --type popular
tmdb-cli --type playing
tmdb-cli --type top
tmdb-cli --type upcoming
```

The application will first attempt to fetch data from the Redis cache; if the data is missing or expired, a request will be made to the TMDB API, and the result will be cached for 10 minutes.

## Docker

The `Makefile` includes targets for working with Docker:

### Build the image:

```bash
make docker-build
```

### Start containers using `docker-compose`:

```bash
make docker-up
```

### Stop containers:

```bash
make docker-down
```

### Remove Docker image:

```bash
make docker-clean
```

Note: To fully utilize Docker, you may need to create your own `Dockerfile` and `docker-compose.yml`.

## Logging

The application uses the `slog` package for logging, outputting messages to the console at INFO level and above. Logging is configured in the `internal/config/logger.go` file.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

https://roadmap.sh/projects/tmdb-cli