package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

type ServiceRequestService struct {
	apiClient *client.APIClient
	baseURL   string
}

func NewServiceRequestService(apiClient *client.APIClient) *ServiceRequestService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &ServiceRequestService{
		apiClient: apiClient,
		baseURL:   "service-request/service",
	}
}

func (s *ServiceRequestService) Create_service_definition(definition *models.ServiceDefinition, request_info *models.RequestInfo) (interface{}, error) {
	endpoint := s.baseURL + "/definition/v1/_create"
	if request_info == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		request_info = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo":       request_info.ToMap(),
		"serviceDefinition": definition.ToMap(),
	}
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
func (s *ServiceRequestService) Search_service_definition(definition *models.ServiceDefinitionCriteria, page *models.Pagination, request_info *models.RequestInfo) (interface{}, error) {
	endpoint := s.baseURL + "/definition/v1/_search"
	if request_info == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		request_info = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo":       request_info.ToMap(),
		"serviceDefinition": definition.ToMap(),
		"pagination":        page.ToMap(),
	}
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
func (s *ServiceRequestService) Create_service(definition *models.Service, request_info *models.RequestInfo) (interface{}, error) {
	endpoint := s.baseURL + "/v1/_create"
	if request_info == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		request_info = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": request_info.ToMap(),
		"service":     definition.ToMap(),
	}
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
func (s *ServiceRequestService) Search_service(definition *models.ServiceCriteria, page *models.Pagination, request_info *models.RequestInfo) (interface{}, error) {
	endpoint := s.baseURL + "/v1/_search"
	if request_info == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		request_info = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo":     request_info.ToMap(),
		"ServiceCriteria": definition.ToMap(),
		"Pagination":      page.ToMap(),
	}
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
