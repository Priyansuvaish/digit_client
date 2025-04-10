package models

// RequestInfo represents the request information for API calls
type RequestInfo struct {
	AuthToken string
	// Add other fields as needed
}

// ToMap converts the RequestInfo to a map
func (r *RequestInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"authToken": r.AuthToken,
		// Add other fields as needed
	}
} 