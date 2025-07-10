# Go Basic

A simple Go project for learning the basics of Go programming.

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Node.js (optional, for development with auto-reload)

### Running the Application

Everything starts at `main.go`. You have several options to run the application:

#### Option 1: Direct execution

```bash
go run main.go
```

#### Option 2: Build and run

```bash
go build main.go
./main
```

#### Option 3: Development with auto-reload

For development with automatic reloading on file changes:

```bash
nodemon --exec go run main.go --signal SIGTERM
```

## Project Structure

```
go-basic/
├── main.go          # Entry point of the application
└── README.MD         # This file
```

## Module Management

Go modules are used to manage dependencies in Go projects:

- `go.mod` - Contains module dependencies and metadata
- `go.sum` - Contains checksums for dependency verification

### Initialize a new module
```bash
go mod init github.com/your-username/go-basic
```

### Common module commands
```bash
# Add a dependency
go get package-name

# Update dependencies
go mod tidy

# Download dependencies
go mod download

# Verify dependencies
go mod verify
```

**Note:** Module names are typically structured like Git repository URLs for better organization and discoverability.

## Development Tips

### Useful Go commands
```bash
# Format your code
go fmt ./...

# Run tests
go test ./...

# Build for different platforms
GOOS=linux GOARCH=amd64 go build main.go
GOOS=windows GOARCH=amd64 go build main.go

# Get help
go help
```

### Recommended tools
- `gofmt` - Code formatting
- `go vet` - Static analysis
- `golint` - Code linting
- `godoc` - Documentation generation

## Learning Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Tour](https://tour.golang.org/)
