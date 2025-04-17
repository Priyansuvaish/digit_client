package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// IDRequestService is a service for managing ID requests.
type IDRequestService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewIDRequestService creates a new IDRequestService.
func NewIDRequestService(apiClient *client.APIClient) *IDRequestService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &IDRequestService{
		apiClient: apiClient,
		baseURL:   "/egov-idgen/id",
	}
}

func (s *IDRequestService) Generate_id(id *models.IDRequest, requestInfo *models.RequestInfo) (interface{}, error) {
	if requestInfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	//create the payload
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"IDRequest":   id.ToMap(),
	}
	//make the request
	return s.apiClient.Post(s.baseURL+"/_generate", payload, nil, nil, nil, nil, true)
}
