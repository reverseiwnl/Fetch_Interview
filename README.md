# Fetch API Coding Assessment with Go

## Project Overview
This project implements a simple REST API using **Fiber**, a fast web framework for Go. The API allows you to:
- Add transactions with points.
- Spend points across payers.
- Check the current balance of points by payer.

The repository includes:
- The API router code (`routes.go`).
- The API handler code (`handlers.go`).
- The Sort interface (`transaction_sort.go`).
- A summary of the project (`summary.txt`).
- A test suite (`testing_test.go`) to validate the API's behavior using **Resty** and Go's built-in testing framework.

---

## Prerequisites
Ensure you have the following installed:
1. **Go**: Version 1.19 or later.
   - [Install Go](https://golang.org/doc/install)
2. **Dependencies**: The project uses:
   - [Fiber](https://github.com/gofiber/fiber): Web framework for building REST APIs.
   - [Resty](https://github.com/go-resty/resty): HTTP client for testing.
   - [Testify](https://github.com/stretchr/testify): For test assertions.

Install the required Go packages by running:
```bash
go get github.com/gofiber/fiber/v2
go get github.com/go-resty/resty/v2
go get github.com/stretchr/testify
```

---

## Running the API

### Step 1: Compile the Code
To compile the Go program, run:
```bash
go build -o app
```

This will generate an executable file named `app` in the current directory.

### Step 2: Run the API
To start the API server, execute:
```bash
./app
```

---

## Testing the API

### Step 1: Run the Tests
To run the test suite, use:
```bash
go test -v
```

This command will:
1. Start the API server in a goroutine.
2. Use **Resty** to send requests to the endpoints.
3. Validate the responses against expected results.

### Expected Output
On successful execution, you will see output similar to:
```plaintext
=== RUN   TestSolution
--- PASS: TestSolution (1.23s)
PASS
ok      example/fiber-api-tests   1.234s
```

---
