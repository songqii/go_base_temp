package templates

import "fmt"

func GetAllTemplates(moduleName, projectName string) map[string]string {
	return map[string]string{
		"go.mod":                    GoMod(moduleName),
		"dev.yaml":                  DevYaml(projectName),
		"online.yaml":               OnlineYaml(projectName),
		".gitignore":                GitIgnore(),
		"build.sh":                  BuildSh(projectName),
		"cmd/main.go":               MainGo(moduleName),
		"conf/conf.go":              ConfGo(moduleName),
		"log/log.go":                LogGo(moduleName),
		"pkg/cache/redis.go":        RedisGo(moduleName),
		"pkg/cache/keys.go":         KeysGo(),
		"pkg/storage/db.go":         DbGo(moduleName),
		"pkg/storage/t_user.go":     UserModelGo(),
		"pkg/error_code/error.go":   ErrorGo(moduleName),
		"pkg/request/request.go":    RequestGo(),
		"pkg/response/response.go":  ResponseGo(),
		"pkg/controller/api.go":     ApiGo(moduleName),
		"pkg/controller/v1/user.go": UserControllerGo(moduleName),
		"pkg/service/service.go":    ServiceGo(moduleName),
		"pkg/service/user.go":       UserServiceGo(moduleName),
		"pkg/utils/common.go":       CommonGo(),
		"pkg/utils/jwt.go":          JwtGo(moduleName),
		"docs/base_database.md":     BaseDatabaseMd(),
	}
}

func GoMod(moduleName string) string {
	return fmt.Sprintf(`module %s

go 1.23

require (
	github.com/gofiber/fiber/v2 v2.52.8
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/redis/go-redis/v9 v9.8.0
	github.com/spf13/viper v1.20.1
	go.uber.org/zap v1.27.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.7
)
`, moduleName)
}

func DevYaml(projectName string) string {
	return fmt.Sprintf(`Debug: true
Mode: "debug"
Env: "dev"

MysqlAddr: root:123456@tcp(127.0.0.1:3306)/%s?parseTime=true&loc=Local&charset=utf8mb4

Socket:
  Port: 8080

Redis:
  Addr: 127.0.0.1:6379
  Pwd: ""
  Db: 0

Log:
  LogLevel: Debug
  LogPath: ./logs/
  LogName: %s.dev

JWT:
  Secret: your-secret-key-change-me
  Expire: 168  # hours
`, projectName, projectName)
}

func OnlineYaml(projectName string) string {
	return fmt.Sprintf(`Debug: false
Mode: "release"
Env: "online"

MysqlAddr: root:password@tcp(127.0.0.1:3306)/%s?parseTime=true&loc=Local&charset=utf8mb4

Socket:
  Port: 8080

Redis:
  Addr: 127.0.0.1:6379
  Pwd: ""
  Db: 0

Log:
  LogLevel: Info
  LogPath: /var/log/%s/
  LogName: %s.online

JWT:
  Secret: your-production-secret-key
  Expire: 168
`, projectName, projectName, projectName)
}

func GitIgnore() string {
	return `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Output
/bin/
/logs/
*.log

# IDE
.idea/
.vscode/
*.swp
*.swo

# Config (keep templates)
online.yaml
!online.yaml.example

# OS
.DS_Store
Thumbs.db

# Go
vendor/
go.sum
`
}

func BuildSh(projectName string) string {
	return fmt.Sprintf(`#!/bin/bash
set -e

APP_NAME="%s"
OUTPUT_DIR="./bin"

echo "🔨 Building ${APP_NAME}..."

mkdir -p ${OUTPUT_DIR}

# Build for current platform
go build -o ${OUTPUT_DIR}/${APP_NAME} ./cmd/main.go

echo "✅ Build complete: ${OUTPUT_DIR}/${APP_NAME}"
`, projectName)
}

func BaseDatabaseMd() string {
	return `# Database Design Document

This document describes the database tables and fields for this project.
AI can generate corresponding GORM models, storage layer code, and CRUD operations based on this document.

## How to Use

1. Describe your table requirements in natural language below
2. Ask AI to generate code based on this document
3. AI will create files in ` + "`pkg/storage/`" + ` directory

## Table Definitions

### Example: Product Table

**Table Name:** t_product

**Description:** Store product information for e-commerce

**Fields:**
- id: Primary key, auto increment
- product_id: Unique product identifier (string, indexed)
- name: Product name (string, max 256 chars)
- description: Product description (text)
- price: Product price (decimal, 10,2)
- stock: Stock quantity (integer, default 0)
- category_id: Category ID (string, indexed)
- status: Product status (integer, 1=active, 0=inactive, default 1)
- create_time: Creation timestamp (auto)
- update_time: Update timestamp (auto)

**Indexes:**
- Unique index on product_id
- Index on category_id
- Index on status

**Relations:**
- Belongs to Category (category_id -> t_category.category_id)

---

## Your Table Definitions

<!-- Add your table definitions below -->

### Table 1: [Your Table Name]

**Table Name:** t_xxx

**Description:** [What this table stores]

**Fields:**
- id: Primary key, auto increment
- [field_name]: [type], [constraints], [description]

**Indexes:**
- [index description]

**Relations:**
- [relation description]

---

## Code Generation Prompt

When you need to generate code, tell AI:

> Based on docs/base_database.md, generate GORM model and CRUD operations for [table_name] table.
> Create file: pkg/storage/t_[table_name].go

AI will generate:
1. GORM struct with proper tags
2. TableName() method
3. Basic CRUD functions (Create, GetByID, Update, Delete, List)
`
}
