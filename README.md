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

## Initializing Go Modules
Run the following command to ensure all dependencies are properly installed:
```bash
go mod tidy
```
---

## Starting the server
To start the API server, run:
```bash
go run main.go
```
- The server will start at `http://localhost:8000`.

## Using the API
You can interact with the API using tools like Postman, curl, or any HTTP client. Below are examples of how to use the API endpoints.

### 1. Add Transactions
- Endpoint: POST /add
- Request Body: { "payer": "DANNON", "points": 300, "timestamp": "2022-10-31T10:00:00Z" }
- Example Command (Using curl): `curl -X POST http://localhost:8000/add -H "Content-Type: application/json" -d '{"payer":"DANNON","points":300,"timestamp":"2022-10-31T10:00:00Z"}'`
### 2. Spend Points

- Endpoint: POST /spend
- Request Body: { "points": 5000 }
- Example Command (Using curl): `curl -X POST http://localhost:8000/spend -H "Content-Type: application/json" -d '{"points":5000}'`
### 3. Get Balance

- Endpoint: GET /balance
Example Command (Using curl): `curl -X GET http://localhost:8000/balance`
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

## Troubleshooting
### Common Issues and Fixes
1. `go: command not found`
- Ensure Go is properly installed and added to your system's PATH. Run go version to verify.

2. `module not found` or `cannot import errors`
- Run `go mod tidy` to clean up dependencies and resolve any missing imports.

3. Port Already in Use
- If the server fails to start because port `8000` is already in use, try running the server on a different port by modifying `app.Listen(":8000")` in `main.go`.

4. Tests Fail
- Ensure you have the latest dependencies by running `go mod tidy`.
---