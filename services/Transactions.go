package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// Encrypts encrypts the given data using the provided key and returns the encrypted data.
type TransactionService struct {
	apiClient  *client.APIClient
	baseURL    string
	gatewayURL string
}

// NewTransactionService creates a new TransactionService with the given API client and base URL.
func NewTransactionService(apiClient *client.APIClient, baseURL string) *TransactionService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &TransactionService{
		apiClient:  apiClient,
		baseURL:    "transaction/v1",
		gatewayURL: "gateway/v1",
	}
}

func (t *TransactionService) Create_transaction(transaction *models.Transaction, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := t.baseURL + "/_create"
	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"Transaction": transaction.ToMap(),
	}
	return t.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

func (t *TransactionService) Search_transaction(transaction *models.TransactionCriteria, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := t.baseURL + "/_search"
	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := transaction.ToMap()
	return t.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}

func (t *TransactionService) Update_transaction(transaction map[string]string, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := t.baseURL + "/_update"
	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := make(map[string]interface{})
	for key, value := range transaction {
		param[key] = value
	}
	return t.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}

func (t *TransactionService) Search_gateway_transaction() (interface{}, error) {
	endpoint := t.gatewayURL + "/_search"
	return t.apiClient.Post(endpoint, nil, nil, nil, nil, nil, true)
}
