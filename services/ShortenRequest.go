package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

type ShortenRequestService struct {
	apiClient *client.APIClient
	baseURL   string
}

func NewShortenRequestService(apiClient *client.APIClient) *ShortenRequestService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &ShortenRequestService{
		apiClient: apiClient,
		baseURL:   "",
	}
}
func (s *ShortenRequestService) Redirect_url(id string) (interface{}, error) {
	endpoint := s.baseURL + "/eus/" + id
	return s.apiClient.Post(endpoint, nil, nil, nil, nil, nil, false)
}
func (s *ShortenRequestService) Shorten_url(sh *models.ShortenRequest, header map[string]string) (interface{}, error) {
	endpoint := s.baseURL + "/eus/shortener"
	payload := sh.ToMap()
	return s.apiClient.Post(endpoint, payload, nil, header, nil, nil, false)
}
