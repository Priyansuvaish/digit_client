package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// Encrypts encrypts the given data using the provided key and returns the encrypted data.
type WorkflowV2Service struct {
	apiClient *client.APIClient
	baseURL   string
	URL       string
}

// NewWorkflowV2Service creates a new WorkflowV2Service with the given API client and base URL.
func NewWorkflowV2Service(apiClient *client.APIClient, baseURL string) *WorkflowV2Service {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &WorkflowV2Service{
		apiClient: apiClient,
		baseURL:   "egov-workflow-v2/egov-wf/process",
		URL:       "egov-workflow-v2/egov-wf",
	}
}

func (w *WorkflowV2Service) Transition_process(process []models.ProcessInstance, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/_transition"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"ProcessInstances": func() []map[string]interface{} {
			var instances []map[string]interface{}
			for _, instance := range process {
				instances = append(instances, instance.ToMap())
			}
			return instances
		}(),
	}
	return w.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

func (w *WorkflowV2Service) Search_process(process *models.ProcessInstanceSearchCriteria, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/_search"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := process.ToMap()
	return w.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}

func (w *WorkflowV2Service) Count_process(process *models.ProcessInstanceSearchCriteria, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/_count"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := process.ToMap()
	return w.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}

func (w *WorkflowV2Service) Get_nearing_sla_count(process *models.ProcessInstanceSearchCriteria, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/_nearingslacount"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := process.ToMap()
	return w.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}
func (w *WorkflowV2Service) Get_status_count(process *models.ProcessInstanceSearchCriteria, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/_statuscount"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := process.ToMap()
	return w.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}

func (w *WorkflowV2Service) Auto_escalate(process string, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/auto/" + process + "/_escalate"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	return w.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
func (w *WorkflowV2Service) Search_escalations(process *models.ProcessInstanceSearchCriteria, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/escalate/_search"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := process.ToMap()
	return w.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}

func (w *WorkflowV2Service) Create_business_serviceV2(process []models.BusinessService, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/escalate/_search"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"BusinessServices": func() []map[string]interface{} {
			var instances []map[string]interface{}
			for _, instance := range process {
				instances = append(instances, instance.ToMap())
			}
			return instances
		}(),
	}
	return w.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}

func (w *WorkflowV2Service) Search_business_serviceV2(process *models.BusinessServiceSearchCriteria, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/businessservice/v2/_search"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := process.ToMap()
	return w.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}
func (w *WorkflowV2Service) Update_business_serviceV2(process []models.BusinessService, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/businessservice/v2/_update"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"BusinessServices": func() []map[string]interface{} {
			var instances []map[string]interface{}
			for _, instance := range process {
				instances = append(instances, instance.ToMap())
			}
			return instances
		}(),
	}
	return w.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
func (w *WorkflowV2Service) Create_business_service(process []models.BusinessService, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/businessservice/_create"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"BusinessServices": func() []map[string]interface{} {
			var instances []map[string]interface{}
			for _, instance := range process {
				instances = append(instances, instance.ToMap())
			}
			return instances
		}(),
	}
	return w.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
func (w *WorkflowV2Service) Search_business_service(process *models.BusinessServiceSearchCriteria, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/businessservice/_search"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
	}
	param := process.ToMap()
	return w.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}

func (w *WorkflowV2Service) Update_business_service(process []models.BusinessService, requestInfo *models.RequestInfo) (interface{}, error) {
	endpoint := w.baseURL + "/businessservice/_update"

	if requestInfo == nil {
		requestConfig := (&models.RequestConfig{}).GetInstance()
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestInfo.ToMap(),
		"BusinessServices": func() []map[string]interface{} {
			var instances []map[string]interface{}
			for _, instance := range process {
				instances = append(instances, instance.ToMap())
			}
			return instances
		}(),
	}
	return w.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)
}
