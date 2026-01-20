# pingenemy

`pingenemy` is a **small CLI application written in Go**, using **only the Go standard library**.  
It is a **personal project** created to practice Go fundamentals, concurrency, context handling, and HTTP.

The tool periodically checks whether HTTP services are responding as expected, working as a simple HTTP-based “ping” for microservices.

---

## Purpose

This project is **for learning purposes only**.  
It focuses on:

- writing idiomatic Go code
- using goroutines and `sync.WaitGroup`
- handling cancellation with `context`
- building a simple CLI without external dependencies
- testing HTTP code with `httptest`

---

## Platform support

- Currently tested **only on macOS**
- Other operating systems (Linux / Windows) were not tested yet and may require small adjustments

---

## Features

- Periodic HTTP checks
- Multiple endpoints support
- Graceful shutdown with `Ctrl + C`
- Request timeout
- Response time measurement
- Expected HTTP status validation
- Concurrent execution

---

## Project structure

```
.
├── main.go
├── internal/
│   ├── httpclient/
│   └── job/
├── Makefile
├── go.mod
└── README.md
```

---

## Running

### Requirements
- Go installed (version compatible with `go.mod`)
- macOS (currently tested platform)

### Run directly
```bash
go run main.go
```

---

## Using Makefile

The project includes a simple `Makefile` to standardize common tasks.

### Build
```bash
make build
```

### Run
```bash
make run
```

### Test
```bash
make test
```

---

## Notes

- This is a personal learning project
- Not intended for production use
- Code favors clarity over abstraction

---

## License

Personal project. No license defined.
