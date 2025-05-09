package services

import (
	"encoding/json"
	"fmt"

	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// MdmsV2Service is a service to interact with the MDMS V2 API
type MdmsV2Service struct {
	apiClient     *client.APIClient
	baseURL       string
	Mdms_base_url string
}

// NewMdmsV2Service creates a new instance of MdmsV2Service
func NewMdmsV2Service(apiClient *client.APIClient, Mdms_base_url string) *MdmsV2Service {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &MdmsV2Service{
		apiClient: apiClient,
		baseURL:   "mdms-v2/schema/v1",
		Mdms_base_url: func() string {
			if Mdms_base_url != "" {
				return Mdms_base_url
			}
			return "mdms-v2/v2"
		}(),
	}
}

// Create a new schema definition
func (s *MdmsV2Service) CreateSchema(body *models.SchemaDefinition, requestinfo *models.RequestInfo) (interface{}, error) {
	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}

	payload := map[string]interface{}{
		"RequestInfo":      requestinfo.ToMap(),
		"SchemaDefinition": body.ToMap(),
	}
	//endpoint := s.baseURL + "/_create"
	endpoint := s.baseURL + "/_create"
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

//Search for schema definitions based on criteria

func (s *MdmsV2Service) SearchSchema(body *models.SchemaDefCriteria, requestinfo *models.RequestInfo) (interface{}, error) {
	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}

	payload := map[string]interface{}{
		"RequestInfo":       requestinfo.ToMap(),
		"SchemaDefCriteria": body.ToMap(),
	}
	endpoint := s.baseURL + "/_search"
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

// Create new MDMS data for a specific schema

func (s *MdmsV2Service) CreateMDMS(schema_code string, body *models.Mdms, requestinfo *models.RequestInfo) (interface{}, error) {
	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}

	payload := map[string]interface{}{
		"RequestInfo": requestinfo.ToMap(),
		"Mdms":        body.ToMap(),
	}
	endpoint := s.Mdms_base_url + "/_create/" + schema_code
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

// Search MDMS records
func (s *MdmsV2Service) SearchMDMS(body *models.MdmsCriteriaV2, requestinfo *models.RequestInfo) (interface{}, error) {
	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}

	payload := map[string]interface{}{
		"RequestInfo":  requestinfo.ToMap(),
		"MdmsCriteria": body.ToMap(),
	}
	endpoint := s.Mdms_base_url + "/_search"
	jsonData, _ := json.Marshal(body)
	fmt.Println(string(jsonData))
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

//Update existing MDMS data

func (s *MdmsV2Service) UpdateMDMS(schema_code string, body *models.Mdms, requestinfo *models.RequestInfo) (interface{}, error) {
	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}

	payload := map[string]interface{}{
		"RequestInfo": requestinfo.ToMap(),
		"Mdms":        body.ToMap(),
	}
	endpoint := s.Mdms_base_url + "/_update/" + schema_code
	return s.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
