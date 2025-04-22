package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

type EgovOtpService struct {
	apiClient     *client.APIClient
	user_otp_base string
	otp_base      string
}

func NewEgovOtpService(apiClient *client.APIClient) *EgovOtpService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &EgovOtpService{
		apiClient:     apiClient,
		user_otp_base: "user-otp/v1",
		otp_base:      "otp/v1",
	}
}

func (e *EgovOtpService) Validate_otp(otp *models.Otp, requestinfo *models.RequestInfo) (interface{}, error) {
	endpoint := e.otp_base + "/_validate"

	if requestinfo == nil {
		requestconfig := (&models.RequestConfig{}).GetInstance()
		requestinfo = requestconfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestinfo.ToMap(),
		"otp":         otp.ToMap(),
	}
	return e.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)

}

func (e *EgovOtpService) Create_otp(otp *models.Otp, requestinfo *models.RequestInfo) (interface{}, error) {
	endpoint := e.otp_base + "/_create"

	if requestinfo == nil {
		requestconfig := (&models.RequestConfig{}).GetInstance()
		requestinfo = requestconfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestinfo.ToMap(),
		"otp":         otp.ToMap(),
	}
	return e.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)

}
func (e *EgovOtpService) Search_otp(otp *models.Otp, requestinfo *models.RequestInfo) (interface{}, error) {
	endpoint := e.otp_base + "/_search"

	if requestinfo == nil {
		requestconfig := (&models.RequestConfig{}).GetInstance()
		requestinfo = requestconfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestinfo.ToMap(),
		"otp":         otp.ToMap(),
	}
	return e.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)

}

type UserOtpService struct {
	apiClient *client.APIClient
	otp_base  string
}

func NewUserOtpService(apiClient *client.APIClient) *EgovOtpService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &EgovOtpService{
		apiClient:     apiClient,
		user_otp_base: "user-otp/v1",
	}
}

func (e *UserOtpService) User_send_otp(otp *models.UserOtp, requestinfo *models.RequestInfo) (interface{}, error) {
	endpoint := e.otp_base + "/_send"

	if requestinfo == nil {
		requestconfig := (&models.RequestConfig{}).GetInstance()
		requestinfo = requestconfig.GetRequestInfo("", nil)
	}
	payload := map[string]interface{}{
		"RequestInfo": requestinfo.ToMap(),
		"otp":         otp.ToMap(),
	}
	return e.apiClient.Post(endpoint, payload, nil, nil, nil, nil, true)

}
