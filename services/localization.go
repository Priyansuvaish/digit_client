package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// LocalizationService is a service for managing localization data.
type LocalizationService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewLocalizationService creates a new LocalizationService.
func NewLocalizationService(apiClient *client.APIClient) *LocalizationService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &LocalizationService{
		apiClient: apiClient,
		baseURL:   "localization/messages/v1",
	}
}

// Create new localization messages
func (s *LocalizationService) Create_messages(body *models.CreateMessagesRequest, request *models.RequestInfo) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_create"
	if request == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		request = requestConfig.GetRequestInfo("", nil)
	}
	// Set the request info in the body
	body.RequestInfo = request
	// Make the POST request
	return s.apiClient.Post(endpoint, body, nil, nil, nil, nil, true)
}

//Search localized messages

func (s *LocalizationService) Search_messages(body *models.LocaleRequest) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_search"
	// Make the POST request
	// Convert map[string]interface{} to map[string]string
	convertedMap := make(map[string]string)
	for key, value := range body.ToMap() {
		if strValue, ok := value.(string); ok {
			convertedMap[key] = strValue
		}
	}
	return s.apiClient.Post(endpoint, nil, nil, nil, convertedMap, nil, true)
}

// Update existing localized messages
func (s *LocalizationService) Update_messages(body *models.UpdateMessagesRequest, request *models.RequestInfo) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_update"
	if request == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		request = requestConfig.GetRequestInfo("", nil)
	}
	// Set the request info in the body
	body.RequestInfo = request
	// Make the POST request
	return s.apiClient.Post(endpoint, body, nil, nil, nil, nil, true)
}

//Delete localized messages

func (s *LocalizationService) Delete_messages(body *models.DeleteMessagesRequest, request *models.RequestInfo) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_delete"
	if request == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		request = requestConfig.GetRequestInfo("", nil)
	}
	// Set the request info in the body
	body.RequestInfo = request
	// Make the POST request
	return s.apiClient.Post(endpoint, body, nil, nil, nil, nil, true)
}

// Create or update localized messages
func (s *LocalizationService) Upsert_messages(body *models.CreateMessagesRequest, request *models.RequestInfo) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_upsert"
	if request == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		request = requestConfig.GetRequestInfo("", nil)
	}
	// Set the request info in the body
	body.RequestInfo = request
	// Make the POST request
	return s.apiClient.Post(endpoint, body, nil, nil, nil, nil, true)
}
