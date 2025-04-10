# Digit Client Go

A Go client library for interacting with the Digit API.

## Installation

```bash
go get github.com/priya/digit_client_go
```

## Usage

```go
import "github.com/priya/digit_client_go/client"

// Create a new client
client := client.NewAPIClient("https://api.example.com", "your-auth-token")

// Make a GET request
result, err := client.Get("endpoint", map[string]string{"param": "value"}, true)
if err != nil {
    // Handle error
}

// Make a POST request
data := map[string]interface{}{
    "key": "value",
}
result, err = client.Post("endpoint", data, nil, true)
if err != nil {
    // Handle error
}
```

## Features

- Simple HTTP client for making API requests
- Support for GET and POST methods
- Automatic JSON handling
- Authorization token support
- Query parameter support

## License

MIT License 