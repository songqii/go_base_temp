# gofiber-creator

A CLI tool to scaffold Go Fiber projects with GORM + Redis.

## Installation

```bash
go install github.com/halfhuman88/gofiber-creator@latest
```

Or build locally:

```bash
go build -o gofiber-creator .

# Create symlink to global
sudo ln -sf $(pwd)/gofiber-creator /usr/local/bin/gofiber-creator
```

## Usage

```bash
# Create new project
gofiber-creator init -n myproject -m github.com/yourname/myproject

# Enter project
cd myproject
go mod tidy

# Edit config
vim dev.yaml

# Run
go run cmd/main.go -config_path=dev.yaml
```

## Generated Project Structure

```
myproject/
├── cmd/
│   └── main.go              # Entry point
├── conf/
│   └── conf.go              # Config management (Viper)
├── log/
│   └── log.go               # Logger (Zap)
├── pkg/
│   ├── cache/               # Redis cache
│   ├── controller/          # Controller layer
│   │   ├── api.go           # Route registration
│   │   └── v1/              # API v1
│   ├── error_code/          # Unified error codes
│   ├── request/             # Request structs
│   ├── response/            # Response structs
│   ├── service/             # Business logic layer
│   ├── storage/             # Data access layer (GORM)
│   └── utils/               # Utility functions
├── dev.yaml                 # Development config
├── online.yaml              # Production config
├── build.sh                 # Build script
└── go.mod
```

## Tech Stack

- Web Framework: [Fiber v2](https://gofiber.io/)
- ORM: [GORM](https://gorm.io/)
- Cache: [go-redis](https://github.com/redis/go-redis)
- Config: [Viper](https://github.com/spf13/viper)
- Logger: [Zap](https://github.com/uber-go/zap)
- Auth: [JWT](https://github.com/golang-jwt/jwt)

## API Examples

The generated project includes user module examples:

- `POST /v1/login` - Login
- `POST /v1/get_user_info` - Get user info
- `POST /v1/update_user_info` - Update user info
