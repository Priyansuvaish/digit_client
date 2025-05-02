package example

import (
	"fmt"
	"log"

	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
	"github.com/Priyansuvaish/digit_client/services"
)

func Userexample() {
	// Create a new API client with DIGIT sandbox URL
	apiClient := client.NewAPIClient("https://sandbox.digit.org", "")

	// Create user service
	userService := services.NewUserService(apiClient)

	// Create request info using builder pattern
	requestInfo := models.RequestInfoBuilder().
		WithAuthToken("your-auth-token-here"). // Replace with actual auth token
		WithAPIID("digit").
		WithVersion("1.0").
		WithAction("create").
		WithUserInfo(map[string]interface{}{
			"id":   "1",
			"name": "System",
		})

	// Example: Create a new citizen user using builder pattern
	citizenUser := models.CreateCitizenUser().
		WithUserName("testuser123").
		WithPassword("Test@123").
		WithSalutation("Mr").
		WithName("Test User").
		WithGender("Male").
		WithMobileNumber("9876543210").
		WithEmailID("test@example.com").
		WithActive(true).
		WithLocale("en_IN").
		WithType("CITIZEN").
		WithTenantID("pb").
		WithRoles([]models.Role{
			*models.RoleBuilder().
				WithCode("CITIZEN").
				WithName("Citizen").
				WithTenantID("pb"),
		})

	// Create citizen
	result, err := userService.CreateCitizen(citizenUser, requestInfo)
	if err != nil {
		log.Fatalf("Error creating citizen: %v", err)
	}
	fmt.Printf("Create Citizen Response: %v\n", result)

	// Example: Search users using builder pattern
	searchCriteria := models.UserSearchBuilder().
		WithTenantID("pb").
		WithUserName("testuser123").
		WithActive(true).
		WithPageSize(10).
		WithPageNumber(0)

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

	// Example: Update user profile
	updatedProfile := models.CreateCitizenUser().
		WithUserName("testuser123").
		WithName("Updated Test User").
		WithEmailID("updated@example.com").
		WithMobileNumber("9876543210").
		WithTenantID("pb")

	updateResult, err := userService.UpdateProfile(updatedProfile, requestInfo)
	if err != nil {
		log.Fatalf("Error updating profile: %v", err)
	}
	fmt.Printf("Update Profile Response: %v\n", updateResult)
}
