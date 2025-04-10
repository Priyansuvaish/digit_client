package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// APIClient represents the main client for making API requests
type APIClient struct {
	baseURL    string
	authToken  string
	httpClient *http.Client
}

// NewAPIClient creates a new API client instance
func NewAPIClient(baseURL, authToken string) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		authToken:  authToken,
		httpClient: &http.Client{},
	}
}

// Get makes a GET request to the specified endpoint
func (c *APIClient) Get(endpoint string, params map[string]string, requireAuth bool) (interface{}, error) {
	// Create URL with query parameters
	u, err := url.Parse(c.baseURL + "/" + endpoint)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	// Create request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Add authorization header if required
	if requireAuth && c.authToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.authToken)
	}

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle redirects
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		return location, nil
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Try to parse as JSON
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		// If not JSON, return as string
		return string(body), nil
	}

	return result, nil
}

// Post makes a POST request to the specified endpoint
func (c *APIClient) Post(endpoint string, jsonData interface{}, params map[string]string, requireAuth bool) (interface{}, error) {
	// Create URL with query parameters
	u, err := url.Parse(c.baseURL + "/" + endpoint)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	// Marshal JSON data
	var body []byte
	if jsonData != nil {
		body, err = json.Marshal(jsonData)
		if err != nil {
			return nil, err
		}
	}

	// Create request
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Add headers
	if jsonData != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if requireAuth && c.authToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.authToken)
	}

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Try to parse as JSON
	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		// If not JSON, return as string
		return string(respBody), nil
	}

	return result, nil
} 