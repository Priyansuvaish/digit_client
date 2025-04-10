package main

import (
	"fmt"
	"log"
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
	"github.com/Priyansuvaish/digit_client/services"
)

func main() {
	// Create a new API client with DIGIT sandbox URL
	apiClient := client.NewAPIClient("https://sandbox.digit.org", "")
	
	// Create user service
	userService := services.NewUserService(apiClient)

	// Create request info
	requestInfo := &models.RequestInfo{
		AuthToken: "your-auth-token-here", // Replace with actual auth token
	}

	// Example: Create a new citizen user
	citizenUser := &models.CitizenUser{
		UserName:     "testuser123",
		Password:     "Test@123",
		Salutation:   "Mr",
		Name:         "Test User",
		Gender:       "Male",
		MobileNumber: "9876543210",
		EmailID:      "test@example.com",
		Active:       true,
		Locale:       "en_IN",
		Type:         "CITIZEN",
		TenantID:     "pb",
	}

	// Create citizen
	result, err := userService.CreateCitizen(citizenUser, requestInfo)
	if err != nil {
		log.Fatalf("Error creating citizen: %v", err)
	}
	fmt.Printf("Create Citizen Response: %v\n", result)

	// Example: Search users
	searchCriteria := &models.UserSearchModel{
		TenantID: "pb",
		// Add other search criteria as needed
	}

	searchResult, err := userService.SearchUsers(searchCriteria, requestInfo)
	if err != nil {
		log.Fatalf("Error searching users: %v", err)
	}
	fmt.Printf("Search Users Response: %v\n", searchResult)

	// Example: Get user details
	userDetails, err := userService.GetUserDetails("pb", requestInfo)
	if err != nil {
		log.Fatalf("Error getting user details: %v", err)
	}
	fmt.Printf("User Details Response: %v\n", userDetails)
} 