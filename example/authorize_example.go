package main

import (
	"fmt"
	"log"
	"time"
	"github.com/google/uuid"
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
	"github.com/Priyansuvaish/digit_client/services"
)

func main() {
	// Create a new API client with DIGIT sandbox URL
	apiClient := client.NewAPIClient("https://sandbox.digit.org", "")
	
	// Initialize RequestConfig properly
	requestConfig := (&models.RequestConfig{}).GetInstance()
	requestConfig.Initialize(
		"digit",                // apiID
		"1.0",                  // version
		"your-auth-token-here", // authToken
		map[string]interface{}{}, // userInfo
		"device123",            // did
		"key123",               // key
		uuid.New().String(),    // msgID
		"requester123",         // requesterID
		uuid.New().String(),    // correlationID
		"authorize",            // action
		time.Now().UnixMilli(), // ts
	)

	// Create authorize service
	authorizeService := services.NewAuthorizeService(apiClient)
	
	// Build Role first
	role:= models.RoleBuilder().WithCode("ADMIN").Build()

	// Create authorization request
	authRequest, err := models.AuthorizationRequestBuilder().
		WithURI("/api/v1/users").
		AddRole(*role). // Dereference the pointer
		AddTenantID("pb").
		Build()
		
	if err != nil {
		log.Fatalf("Error building authorization request: %v", err)
	}

	// Get authorized actions (pass nil to use default RequestInfo)
	authorizedActions, err := authorizeService.AuthorizeAction(authRequest, nil)
	if err != nil {
		log.Fatalf("Error authorizing action: %v", err)
	}
	
	fmt.Printf("Authorized actions: %+v\n", authorizedActions)
}
