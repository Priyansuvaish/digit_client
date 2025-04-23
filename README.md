# Digit Client Go

A Go client library for interacting with the Digit API.

## Installation

```bash
go get github.com/Priyansuvaish/digit_client
```

## Usage

```go
import ("fmt"
	"log"
	"time"

	// "github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/config"
	"github.com/Priyansuvaish/digit_client/models"
	"github.com/Priyansuvaish/digit_client/services"
	"github.com/google/uuid")

func main() {
    // configre the url and auth token globally
    config.GetGlobalConfig().Initialize(
		"https://sandbox.digit.org",
		"your-auth-token",
	)

    //if you want to create a new API client with DIGIT sandbox URL and you have to pass in the service instanciation
	
    // apiClient := client.NewAPIClient("https://sandbox.digit.org", "")

    // Initialize RequestConfig properly
	requestConfig := (&models.RequestConfig{}).GetInstance()
	requestConfig.Initialize(
		"digit",                  // apiID
		"1.0",                    // version
		"your-auth-token-here",   // authToken
		map[string]interface{}{}, // userInfo
		"device123",              // did
		"key123",                 // key
		uuid.New().String(),      // msgID
		"requester123",           // requesterID
		uuid.New().String(),      // correlationID
		"authorize",              // action
		time.Now().UnixMilli(),   // ts
	)

	// Create authorize service
	authorizeService := services.NewAuthorizeService(nil)

    // if not set globally created a new API client

	//authorizeService := services.NewAuthorizeService(apiClient)
    

	// Build Role first
	role := models.RoleBuilder().WithCode("ADMIN").Build()

	// Create authorization request
	authRequest, err := models.AuthorizationRequestBuilder().
		WithURI("/api/v1/users").
		AddRole(*role). // Dereference the pointer
		AddTenantID("pb").
		Build()

	if err != nil {
		log.Fatalf("Error building authorization request: %v", err)
	}

	action, err := models.ActionBuilder().
		WithName("view").
		WithURL("/api/v1/users").
		WithDisplayName("View Users").
		Build()
	if err != nil {
		log.Fatalf("Error building ActionBuilder: %v", err)
	}
	actionrequest, err := models.ActionRequestBuilder().
		WithActions([]*models.Action{action}). // Dereference the pointer
		Build()

	if err != nil {
		log.Fatalf("Error building actionrequest: %v", err)
	}

	// Get authorized actions (pass nil to use default RequestInfo)
	authorizedActions, err := authorizeService.AuthorizeAction(authRequest, nil)
	if err != nil {
		log.Fatalf("Error authorizing action: %v", err)
	}
	// GetMDMSAction gets MDMS action details
	getmdmsaction, err := authorizeService.GetMDMSAction(actionrequest, nil)
	if err != nil {
		log.Fatalf("Error getting MDMS action: %v", err)
	}
	// Print the MDMS action details
	fmt.Printf("MDMS action details: %+v\n", getmdmsaction)
	// Print the authorized actions

	fmt.Printf("Authorized actions: %+v\n", authorizedActions)
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