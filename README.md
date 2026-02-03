<p align="center">
  <img src="https://img.shields.io/badge/Go-1.23+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version">
  <img src="https://img.shields.io/badge/Fiber-v2-00ACD7?style=for-the-badge&logo=go&logoColor=white" alt="Fiber">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/badge/Version-1.0.0-blue?style=for-the-badge" alt="Version">
</p>

<h1 align="center">🚀 GoFiber Creator</h1>

<p align="center">
  <strong>A powerful CLI tool to scaffold production-ready Go Fiber projects in seconds</strong>
</p>

<p align="center">
  Stop wasting time on boilerplate. Start building features.
</p>

---

## ✨ Features

- 🏗️ **Production-Ready Structure** - Battle-tested layered architecture
- ⚡ **Fiber v2** - Express-inspired, fastest HTTP framework for Go
- 🗄️ **GORM Integration** - Elegant ORM with MySQL support out of the box
- 🔴 **Redis Cache** - High-performance caching layer included
- 🔐 **JWT Authentication** - Secure token-based auth ready to use
- 📝 **Zap Logger** - Structured logging with rotation support
- 🎯 **Clean Architecture** - Controller → Service → Storage pattern
- 📚 **AI-Friendly Docs** - Database design docs for AI code generation

## 📦 Installation

### Via Homebrew (Recommended for macOS/Linux)

```bash
brew tap songqii/tap https://github.com/songqii/go_base_temp
brew install gofiber-creator
```

### Via Go Install

```bash
go install github.com/halfhuman88/gofiber-creator@latest
```

### Build from Source

```bash
git clone https://github.com/songqii/go_base_temp.git
cd go_base_temp
go build -o gofiber-creator .

# Create global symlink
sudo ln -sf $(pwd)/gofiber-creator /usr/local/bin/gofiber-creator
```

## 🚀 Quick Start

```bash
# Create a new project
gofiber-creator init -n myproject -m github.com/yourname/myproject

# Enter project directory
cd myproject

# Install dependencies
go mod tidy

# Configure your database & redis
vim dev.yaml

# Run the server
go run cmd/main.go -config_path=dev.yaml
```

Your API server is now running at `http://localhost:8080` 🎉

## 📁 Project Structure

```
myproject/
├── cmd/
│   └── main.go                 # Application entry point
├── conf/
│   └── conf.go                 # Configuration management (Viper)
├── log/
│   └── log.go                  # Structured logging (Zap)
├── docs/
│   └── base_database.md        # Database design for AI generation
├── pkg/
│   ├── cache/                  # Redis cache layer
│   │   ├── redis.go
│   │   └── keys.go
│   ├── controller/             # HTTP handlers
│   │   ├── api.go              # Route registration
│   │   └── v1/                 # API version 1
│   ├── error_code/             # Unified error handling
│   ├── request/                # Request DTOs
│   ├── response/               # Response DTOs
│   ├── service/                # Business logic layer
│   ├── storage/                # Data access layer (GORM)
│   └── utils/                  # Utility functions
├── dev.yaml                    # Development config
├── online.yaml                 # Production config
└── build.sh                    # Build script
```

## 🛠️ Tech Stack

| Component | Technology |
|-----------|------------|
| Web Framework | [Fiber v2](https://gofiber.io/) |
| ORM | [GORM](https://gorm.io/) |
| Cache | [go-redis](https://github.com/redis/go-redis) |
| Config | [Viper](https://github.com/spf13/viper) |
| Logger | [Zap](https://github.com/uber-go/zap) |
| Auth | [JWT](https://github.com/golang-jwt/jwt) |

## 📖 API Examples

The generated project includes a complete user module:

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/v1/login` | User login |
| POST | `/v1/get_user_info` | Get user profile |
| POST | `/v1/update_user_info` | Update user profile |
| GET | `/health` | Health check |

## 🤖 AI-Powered Development

The generated project includes `docs/base_database.md` where you can describe your database tables in natural language. Then ask AI to generate:

- GORM models
- CRUD operations
- Request/Response structs
- Controller handlers
- Service layer logic

## 📋 Commands

```bash
# Show help
gofiber-creator --help

# Show version
gofiber-creator -v

# Initialize new project
gofiber-creator init -n <project-name> -m <module-name>
```

## 🤝 Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

## 📄 License

This project is licensed under the MIT License.

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/songqii">songqii</a>
</p>
