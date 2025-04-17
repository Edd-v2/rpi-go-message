# rpi-go-message
This repo will provide a scalable, efficiency, secure, reliable backend server for a messaging application




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
