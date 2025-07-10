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
