# Redis Clone in Go

A simple, Redis-inspired in-memory data store written in Go. This project emulates core Redis functionalities, including basic data storage, command handling, and TTL-based caching. The aim is to provide a fast, lightweight key-value store with similar command support to Redis.

## Features

- **Key-Value Storage**: Supports basic commands like `SET`, `GET`, and `DEL`.
- **TTL (Time-To-Live)**: Allows setting expiration times for keys, similar to Redis’s `EXPIRE`.
- **RESP Protocol Support**: Implements the Redis Serialization Protocol (RESP) for compatibility with Redis clients.
- **In-Memory Storage**: High-speed data access by storing all data in memory.

## Project Structure

```
redis-clone/
├── cmd/                  # Application entry point
│   └── redis-clone/      
│       └── main.go       # Initializes server and configuration
├── internal/             
│   ├── server/           # Manages server and client connections
│   ├── data/             # Core data structures and in-memory storage
│   └── protocol/         # Handles the Redis protocol and command parsing
├── config/               # Configuration files
└── README.md             # Project documentation
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
   go build -o redix ./cmd/redix
   ```

3. Run the server:
   ```bash
   ./redix
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
127.0.0.1:6379> EXPIRE mykey 10
(integer) 1
```

### Configuration

Configuration options, like the port and memory limits, can be modified in the `config/config.yaml` file. Example:

```yaml
port: 6379
max_memory: 64MB
```
