package client

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
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
func (c *APIClient) Post(
	endpoint string,
	jsonData interface{},
	formData map[string]string,
	additionalHeaders map[string]string,
	params map[string]string,
	files map[string]io.Reader,
	requireAuth bool,
) (interface{}, error) {
	// Build URL with query parameters
	u, err := url.Parse(c.baseURL + "/" + endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	var body io.Reader
	headers := make(http.Header)

	// Handle files (multipart/form-data)
	if len(files) > 0 {
		var b bytes.Buffer
		writer := multipart.NewWriter(&b)
		// Add files
		for fieldName, fileReader := range files {
			part, err := writer.CreateFormFile(fieldName, fieldName)
			if err != nil {
				return nil, err
			}
			if _, err := io.Copy(part, fileReader); err != nil {
				return nil, err
			}
		}
		// Add form fields
		for k, v := range formData {
			_ = writer.WriteField(k, v)
		}
		writer.Close()
		body = &b
		headers.Set("Content-Type", writer.FormDataContentType())
	} else if jsonData != nil {
		// JSON body
		b, err := json.Marshal(jsonData)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
		headers.Set("Content-Type", "application/json")
	} else if len(formData) > 0 {
		// x-www-form-urlencoded
		data := url.Values{}
		for k, v := range formData {
			data.Set(k, v)
		}
		body = strings.NewReader(data.Encode())
		headers.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// Add authentication header if required
	if requireAuth && c.authToken != "" {
		headers.Set("Authorization", "Bearer "+c.authToken)
	}
	// Add any additional headers
	for k, v := range additionalHeaders {
		headers.Set(k, v)
	}

	// Create request
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header = headers

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

	// Try to parse as JSON, fallback to string
	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return string(respBody), nil
	}
	return result, nil
}
