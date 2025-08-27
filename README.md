# GO Echo Server

Echo server with Go.

## Usage

### Running the Server

```bash
go run main.go
```

Or with custom port:

```bash
go run main.go -p 8080
```

## Connect to the server

### Using the built-in client

```bash
go run main.go -m client
```

### Using external tools

```bash
nc localhost 9000
```

or

```bash
telnet localhost 9000
```

Replace `9000` with your custom port if using `-p` flag.
