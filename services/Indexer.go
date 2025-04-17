package services

import (
	"fmt"

	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

type IndexerService struct {
	apiClient *client.APIClient
	baseURL   string
}

func NewIndexerService(apiClient *client.APIClient) *IndexerService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &IndexerService{
		apiClient: apiClient,
		baseURL:   "index-operations",
	}
}
func (s *IndexerService) Legacy_index(index *models.LegacyIndexRequest, requestinfo *models.RequestInfo) (interface{}, error) {
	// Build the URL for the legacy index endpoint
	url := s.baseURL + "/_legacyindex"

	// If requestInfo is provided, update the actionRequest's requestInfo
	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}
	index.RequestInfo = requestinfo
	// Convert the request to a map
	requestMap := index.ToMap()
	// Make the API call using the client
	response, err := s.apiClient.Post(url, requestMap, nil, nil, nil, nil, true)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *IndexerService) Index_to_topic(topic string, index_data map[string]interface{}, requestinfo *models.RequestInfo) (interface{}, error) {
	// Build the URL for the index to topic endpoint
	endpoint := fmt.Sprintf("%s/%s/_index", s.baseURL, topic)

	// If requestInfo is provided, update the actionRequest's requestInfo
	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"requestInfo": requestinfo.ToMap(),
		"indexJson":   index_data,
	}

	// Make the API call using the client
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)

}

func (s *IndexerService) Reindex_data(request *models.ReindexRequest, requestinfo *models.RequestInfo) (interface{}, error) {

	// Build the URL for the reindex endpoint
	url := s.baseURL + "/_reindex"

	// If requestInfo is provided, update the actionRequest's requestInfo
	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}
	request.RequestInfo = requestinfo
	// Convert the request to a map
	requestMap := request.ToMap()
	// Make the API call using the client
	return s.apiClient.Post(url, requestMap, nil, nil, nil, nil, true)

}
