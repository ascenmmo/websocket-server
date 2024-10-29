# High-Load WebSocket Service for Game Clients

## Русская версия

Для русскоязычных пользователей документацию можно найти [здесь](https://github.com/ascenmmo/websocket-server/blob/master/RU_README.md).

## Description
This project is a high-load **WebSocket service** tailored for connecting game clients. It is designed for free deployment on servers, making it a robust solution for game developers.

## Key Features

- **High Performance**: Optimized for managing a large volume of simultaneous WebSocket connections.
- **Free Deployment**: Easily deployable on any server without additional costs.
- **Docker Support**: Available as both a binary and a Docker container, ensuring compatibility across platforms.
- **Flexibility and Scalability**: Configurable to meet specific needs and easily scalable.

## Installation

### Installation via Docker

1. Make sure Docker is installed. If it’s not installed, follow the Docker installation instructions.
2. Run the command:
   ```bash
   docker compose up -d --force-recreate --build && docker image prune -f
	```

## Configuration

The project uses the **env/env.go** package to define key configuration parameters. All services interacting with this WebSocket server must use the same token to ensure secure connections and authentication.

### Configuration Parameters

```go
package env

var (
   ServerAddress       = "0.0.0.0" // Server IP address
   TCPPort             = "8082"    // Port for TCP connections
   WebsocketPort       = "4240"    // Port for WebSocket connections
   TokenKey            = "_remember_token_must_be_32_bytes" // Unique token for authentication
   MaxRequestPerSecond = 50        // Max requests per second
)
```

* ServerAddress: Specifies the IP address where the server will operate.
* TCPPort: The port on which the server will listen for TCP connections.
* WebsocketPort: The port used for handling WebSocket connections.
* TokenKey: A unique token that must be the same across all services interacting with this Websocket server. This ensures the security and integrity of connections.
* MaxRequestPerSecond: A limit on the maximum number of requests the server can handle per second.



##  Importance of a Single Token
### Using a single token allows:

* **Security Assurance:** All services check the token before establishing a connection, helping to prevent unauthorized access.
* **Simplified Authentication:** A single token for all services simplifies the authentication and access management process.
* **Ease of Maintenance:** If the token needs to be changed, it can be done in one place, and all services will be updated simultaneously.

Make sure that all your services are configured to use this token to ensure the correct operation and security of the system.



## Troubleshooting

If issues arise when starting the service, check the following:

- Ensure that ports 8082 and 4240 are not in use by other applications.
- Confirm that the configuration parameters in env/env.go are set correctly.
- If using Docker, verify it is running and configured properly.





## Теги

`WebSocket`, `игровой сервер`, `высоконагруженный`, `бесплатное развертывание`, `Docker`, `кроссплатформенный`, `игровая разработка`, `сеть`, `многопользовательская игра`, `сервис для игр`, `настройка сервера`, `аутентификация`, `токены`, `Golang`, `open-source`

## Tags

`WebSocket`, `game server`, `high-performance`, `free deployment`, `Docker`, `cross-platform`, `game development`, `network`, `multiplayer game`, `game service`, `server setup`, `authentication`, `tokens`, `Golang`, `open-source`
