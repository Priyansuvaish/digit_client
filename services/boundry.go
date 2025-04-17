package services

import (
	"fmt"

	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// BoundaryService handles operations related to boundaries
type BoundaryService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewBoundaryService creates a new BoundaryService instance
func NewBoundaryService(apiClient *client.APIClient) *BoundaryService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &BoundaryService{
		apiClient: apiClient,
		baseURL:   "boundarys/_search",
	}
}

// SearchBoundaries searches for boundaries based on the provided request

func (s *BoundaryService) BoundarySearch(request *models.BoundarySearchRequest) (interface{}, error) {
	param := request.ToMap()
	stringParam := make(map[string]string)
	for key, value := range param {
		if strValue, ok := value.(string); ok {
			stringParam[key] = strValue
		} else {
			return nil, fmt.Errorf("invalid parameter type for key %s: expected string", key)
		}
	}
	endpoint := s.baseURL
	return s.apiClient.Post(endpoint, nil, nil, nil, stringParam, nil, true)
}

func (s *BoundaryService) LocationSearchBoundary(request *models.LocationBoundarySearchRequest) (interface{}, error) {
	param := request.ToMap()
	stringParam := make(map[string]string)
	for key, value := range param {
		if strValue, ok := value.(string); ok {
			stringParam[key] = strValue
		} else {
			return nil, fmt.Errorf("invalid parameter type for key %s: expected string", key)
		}
	}
	endpoint := "/location/v11" + s.baseURL
	return s.apiClient.Post(endpoint, nil, nil, nil, stringParam, nil, true)
}
