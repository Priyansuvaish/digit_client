package services

import (
	"fmt"

	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// AuthorizeService handles authorization-related operations
type AuthorizeService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewAuthorizeService creates a new authorize service instance
func NewAuthorizeService(apiClient *client.APIClient) *AuthorizeService {
	return &AuthorizeService{
		apiClient: apiClient,
		baseURL:   "access/v1",
	}
}

// AuthorizeAction authorizes an action based on roles and tenant permissions
func (s *AuthorizeService) AuthorizeAction(authorizationRequest *models.AuthorizationRequest, requestInfo *models.RequestInfo) (interface{}, error) {

	if requestInfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	fmt.Printf("requestinfo %+v\n", requestInfo.ToMap())
	payload := map[string]interface{}{
		"RequestInfo":          requestInfo.ToMap(),
		"AuthorizationRequest": authorizationRequest.ToMap(),
	}

	endpoint := s.baseURL + "/actions/_authorize"
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

// GetMDMSAction gets MDMS action details
func (s *AuthorizeService) GetMDMSAction(actionRequest *models.ActionRequest, requestInfo *models.RequestInfo) (interface{}, error) {
	// If requestInfo is provided, update the actionRequest's requestInfo
	if requestInfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	actionRequest.RequestInfo = requestInfo
	payload := actionRequest.ToMap()
	endpoint := s.baseURL + "/actions/mdms/_get"
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
