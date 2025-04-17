package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// Encrypts encrypts the given data using the provided key and returns the encrypted data.
type EncryptsService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewEncryptsService creates a new EncryptsService with the given API client and base URL.
func NewEncryptsService(apiClient *client.APIClient, baseURL string) *EncryptsService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &EncryptsService{
		apiClient: apiClient,
		baseURL:   "egov-enc-service/crypto/v1",
	}
}

func (s *EncryptsService) EncryptData(data []*models.EncReqObject) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_encrypt"

	//create payload for the request
	// Convert EncReqObject to map for each item in the data slice
	payload := make([]map[string]interface{}, len(data))
	for i, item := range data {
		// Convert EncReqObject to map
		payload[i] = item.ToMap()
	}

	// Make the POST request to the API
	return s.apiClient.Post(endpoint, data, nil, nil, nil, nil, true)
}

func (s *EncryptsService) DecryptData(data map[string]string) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_decrypt"

	// Make the POST request to the API
	return s.apiClient.Post(endpoint, data, nil, nil, nil, nil, true)
}

func (s *EncryptsService) Create_Digital_Signature(data *models.SignRequest) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_sign"
	payload := data.ToMap()

	// Make the POST request to the API
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

func (s *EncryptsService) Verify_Signature(data *models.VerifyRequest) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_verify"
	payload := data.ToMap()

	// Make the POST request to the API
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

func (s *EncryptsService) Rotate_all_keys(data *models.RotateKeyRequest) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_rotateallkeys"
	payload := data.ToMap()

	// Make the Post request to the API
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

func (s *EncryptsService) Rotate_single_key(data *models.RotateKeyRequest) (interface{}, error) {
	// Define the endpoint URL
	endpoint := s.baseURL + "/_rotatekey"
	payload := data.ToMap()

	// Make the Post request to the API
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
