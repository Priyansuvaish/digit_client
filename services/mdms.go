package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

type MDMSService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewMDMSService creates a new MDMSService instance

func NewMDMSService(apiClient *client.APIClient) *MDMSService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &MDMSService{
		apiClient: apiClient,
		baseURL:   "egov-mdms-service/v1",
	}
}

// Search for MDMS data based on criteria
func (s *MDMSService) Search(criteria *models.MdmsCriteria, requestinfo *models.RequestInfo) (interface{}, error) {
	// Define the endpoint and parameters
	endpoint := s.baseURL + "/_search"
	if requestinfo == nil {
		requestconfig := (&models.RequestConfig{}).GetInstance()
		requestinfo = requestconfig.GetRequestInfo("", nil)
	}

	payload := map[string]interface{}{
		"MdmsCriteria": criteria.ToMap(),
		"RequestInfo":  requestinfo.ToMap(),
	}

	// Make the Post request using the API client
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

//Get specific MDMS data based on module and master name

func (s *MDMSService) Get(moduleName string, masterName string, tenant_id string, filter []string, mdmsCriteria *models.MdmsCriteria, requestinfo *models.RequestInfo) (interface{}, error) {
	// Define the endpoint and parameters
	endpoint := s.baseURL + "/_get"
	if requestinfo == nil {
		requestconfig := (&models.RequestConfig{}).GetInstance()
		requestinfo = requestconfig.GetRequestInfo("", nil)
	}

	param := map[string]interface{}{
		"moduleName": moduleName,
		"masterName": masterName,
		"tenantId":   tenant_id,
	}
	if filter != nil {
		param["filter"] = filter
	}

	payload := map[string]interface{}{
		"MdmsCriteria": mdmsCriteria.ToMap(),
		"RequestInfo":  requestinfo.ToMap(),
	}

	// Make the Post request using the API client
	return s.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}
