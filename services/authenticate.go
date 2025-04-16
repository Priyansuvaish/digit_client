package services

import (
	"fmt"

	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

type AuthenticateService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewAuthenticateService creates a new authenticate service instance
func NewAuthenticateService(apiClient *client.APIClient) *AuthenticateService {
	return &AuthenticateService{
		apiClient: apiClient,
		baseURL:   "user",
	}
}

// Get authentication token using password grant type
func (s *AuthenticateService) GetTokenUsingPasswordGrantType(authenticationrequest *models.AuthenticationRequest) (interface{}, error) {
	data := authenticationrequest.ToMap()
	convertedData := make(map[string]string)
	for key, value := range data {
		if strValue, ok := value.(string); ok {
			convertedData[key] = strValue
		} else {
			return nil, fmt.Errorf("invalid data type for key %s, expected string", key)
		}
	}
	headers := map[string]string{
		"Authorization": "Basic ZWdvdi11c2VyLWNsaWVudDo=",
		"Content-Type":  "application/x-www-form-urlencoded",
	}
	endpoint := s.baseURL + "/oauth/token"
	return s.apiClient.Post(endpoint, nil, convertedData, headers, nil, nil, true)
}

// Update password without requiring login
func (s *AuthenticateService) UpdatePasswordNoLogin(requestInfo *models.RequestInfo) (interface{}, error) {
	if requestInfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := requestInfo.ToMap()
	headers := map[string]string{
		"Authorization": "Basic ZWdvdi11c2VyLWNsaWVudDo=",
		"Content-Type":  "application/x-www-form-urlencoded",
	}
	endpoint := s.baseURL + "/password/nologin/_update"
	return s.apiClient.Post(endpoint, payload, nil, headers, nil, nil, true)
}

//Logout the user using their access token

func (s *AuthenticateService) Logout(requestInfo *models.RequestInfo) (interface{}, error) {
	if requestInfo == nil {
		// Get the singleton instance of RequestConfig
		requestConfig := (&models.RequestConfig{}).GetInstance()
		// Call GetRequestInfo on the instance
		requestInfo = requestConfig.GetRequestInfo("", nil)
	}
	payload := requestInfo.ToMap()
	param := map[string]string{
		"access_token": requestInfo.AuthToken,
	}
	endpoint := s.baseURL + "/_logout"
	return s.apiClient.Post(endpoint, payload, nil, nil, param, nil, true)
}
