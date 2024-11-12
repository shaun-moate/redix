# Redix

A simple, Redis-inspired in-memory data store written in Go. This project emulates core Redis functionalities, including basic data storage, command handling, and TTL-based caching. The aim is to provide a fast, lightweight key-value store with similar command support to Redis.

## Features

- **Key-Value Storage**: Supports basic commands like `SET`, `GET`, `HSET` and `HGET`.
- **RESP Protocol Support**: Implements the Redis Serialization Protocol (RESP) for compatibility with Redis clients (`redis-cli`).
- **In-Memory Storage**: High-speed data access by storing all data in memory.

## Project Structure

```
redis-clone/
├── cmd/                  
    └── redix/      
        ├── main.go        # entry point
        ├── aof.go         # data persistence
        ├── handler.go     # commands
        ├── parser.go      # parser the input from redis-cli
        ├── serializer.go  # serialization
        └── writer.go      # write the response to redis-cli
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.21 or higher

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/shaun-moate/redix.git
   cd redix
   ```

2. Build the project:
   ```bash
   go build -o build/redix ./cmd/redix
   ```

3. Run the server:
   ```bash
   ./build/redix
   ```

The server will start listening on the default port (e.g., `6379`).

### Usage

You can interact with the server using a Redis client or telnet. Here’s a quick example:

```bash
$ redis-cli -p 6379
127.0.0.1:6379> SET mykey "Hello, World!"
OK
127.0.0.1:6379> GET mykey
"Hello, World!"
```
