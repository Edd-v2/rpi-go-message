# rpi-go-message

`rpi-go-message` is a scalable, efficient, secure, and reliable backend server built with Go for a messaging application similar to WhatsApp or Telegram. It is designed to run on a Raspberry Pi and uses MongoDB for storage, Gin for HTTP handling, and WebSocket for real-time messaging.

## Features

- User registration and authentication with JWT
- Group and private chat support
- MongoDB for message and user storage
- Secure password hashing (bcrypt)
- Clean modular architecture: handler, service, repository layers
- Real-time communication with WebSocket (WIP)
- Prometheus-ready metrics and health endpoints
- Built-in logger using logrus (file + stdout)
- Configurable via `configuration.yaml`

## Architecture Overview


```
rpi-go-message/
│
├── src/
│   ├── cmd/                  # Main app entrypoint(s)
│   │   └── main.go           # Bootstraps everything
│   ├── config/               # Loads env/configuration
│   │   └── config.go
│   └── internal/             # All internal modules
│       ├── api/              # HTTP route handlers
│       ├── db/               # MongoDB init & logic
│       ├── message/          # Message-related business logic
│       ├── model/            # Structs: User, Chat, Message
│       └── middleware/       # Logging, auth, recovery
│
├── configuration.yaml        # Config file (envs, secrets)
├── docker-compose.yaml       # DB, app, Prometheus setup
├── Dockerfile                # Containerize app
└── README.md
```



## Technologies

- [Go (Golang)](https://golang.org)
- [MongoDB](https://www.mongodb.com)
- [Gin](https://github.com/gin-gonic/gin)
- [Logrus](https://github.com/sirupsen/logrus)
- [Viper](https://github.com/spf13/viper)
- [JWT](https://github.com/golang-jwt/jwt)
- WebSocket (coming soon)
- Prometheus (coming soon)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/Edd-v2/rpi-go-message.git

2. Adjust configuration.yaml as needed.

3. Run the project:
    
```
    go run ./src/cmd
```