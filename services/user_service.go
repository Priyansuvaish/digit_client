package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// UserService handles user-related operations
type UserService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewUserService creates a new user service instance
func NewUserService(apiClient *client.APIClient) *UserService {
	return &UserService{
		apiClient: apiClient,
		baseURL:   "user",
	}
}

// CreateCitizen creates a new citizen user
func (s *UserService) CreateCitizen(citizenUser *models.CitizenUser, requestInfo *models.RequestInfo) (interface{}, error) {
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"user":        citizenUser.ToMap(),
	}

	endpoint := s.baseURL + "/citizen/_create"
	return s.apiClient.Post(endpoint, payload, nil, true)
}

// GetUserDetails retrieves user details
func (s *UserService) GetUserDetails(tenantID string, requestInfo *models.RequestInfo) (interface{}, error) {
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}

	params := map[string]string{
		"tenantId": tenantID,
	}

	// Add auth token from RequestInfo if available
	if requestInfo.AuthToken != "" {
		params["access_token"] = requestInfo.AuthToken
	}

	endpoint := s.baseURL + "/_details"
	return s.apiClient.Post(endpoint, payload, params, true)
}

// UpdateProfile updates a user's profile
func (s *UserService) UpdateProfile(userProfile *models.CitizenUser, requestInfo *models.RequestInfo) (interface{}, error) {
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"user":        userProfile.ToMap(),
	}

	endpoint := s.baseURL + "/profile/_update"
	return s.apiClient.Post(endpoint, payload, nil, true)
}

// SearchUsers searches for users based on criteria
func (s *UserService) SearchUsers(searchCriteria *models.UserSearchModel, requestInfo *models.RequestInfo) (interface{}, error) {
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}

	// Add search criteria to payload
	for k, v := range searchCriteria.ToMap() {
		payload[k] = v
	}

	endpoint := s.baseURL + "/_search"
	return s.apiClient.Post(endpoint, payload, nil, true)
}

// CreateUserNoValidate creates a user without validation
func (s *UserService) CreateUserNoValidate(citizenUser *models.CitizenUser, requestInfo *models.RequestInfo) (interface{}, error) {
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"user":        citizenUser.ToMap(),
	}

	endpoint := s.baseURL + "/users/_createnovalidate"
	return s.apiClient.Post(endpoint, payload, nil, true)
}

// UpdateUserNoValidate updates a user without validation
func (s *UserService) UpdateUserNoValidate(userProfile *models.CitizenUser, requestInfo *models.RequestInfo) (interface{}, error) {
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"user":        userProfile.ToMap(),
	}

	endpoint := s.baseURL + "/users/_updatenovalidate"
	return s.apiClient.Post(endpoint, payload, nil, true)
} 