package services

import (
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
	endpoint := s.baseURL
	return s.apiClient.Post(endpoint, nil, nil, nil, param, nil, true)
}

func (s *BoundaryService) LocationSearchBoundary(request *models.LocationBoundarySearchRequest) (interface{}, error) {
	param := request.ToMap()
	endpoint := "/location/v11" + s.baseURL
	return s.apiClient.Post(endpoint, nil, nil, nil, param, nil, true)
}
