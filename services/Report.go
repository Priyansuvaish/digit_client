package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

type ReportService struct {
	apiClient *client.APIClient
	baseURL   string
}

func NewReportService(apiClient *client.APIClient) *ReportService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &ReportService{
		apiClient: apiClient,
		baseURL:   "report",
	}
}
func (r *ReportService) Create_v1_metadata(module_name string, version string, request *models.MetadataRequest, requestinfo *models.RequestInfo) (interface{}, error) {
	endpoint := r.baseURL + "/" + module_name + "/" + version + "/metadata/_get"

	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}

	request.RequestInfo = requestinfo

	return r.apiClient.Post(endpoint, request.ToMap(), nil, nil, nil, nil, true)

}

func (r *ReportService) Get_report_data_v1(module_name string, version string, request *models.MetadataRequest, requestinfo *models.RequestInfo) (interface{}, error) {
	endpoint := r.baseURL + "/" + module_name + "/" + version + "/_get"

	if requestinfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestinfo = requestConfig.GetRequestInfo("", nil)
	}

	request.RequestInfo = requestinfo

	return r.apiClient.Post(endpoint, request.ToMap(), nil, nil, nil, nil, true)

}
