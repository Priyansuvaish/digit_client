package models

import (
	"fmt"
)

// AuthorizationRequest represents an authorization request in the system
type AuthorizationRequest struct {
	Roles     []Role
	URI       string
	TenantIDs []string
}

// Validate checks if the AuthorizationRequest has valid required fields
func (a *AuthorizationRequest) Validate() error {
	if len(a.Roles) == 0 {
		return fmt.Errorf("at least one role is required")
	}
	if a.URI == "" {
		return fmt.Errorf("URI is required")
	}
	if len(a.TenantIDs) == 0 {
		return fmt.Errorf("at least one tenant ID is required")
	}
	return nil
}

// ToMap converts the AuthorizationRequest to a map
func (a *AuthorizationRequest) ToMap() map[string]interface{} {
	roles := make([]map[string]interface{}, len(a.Roles))
	for i, role := range a.Roles {
		roles[i] = role.ToMap()
	}

	return map[string]interface{}{
		"roles":     roles,
		"uri":       a.URI,
		"tenantIds": a.TenantIDs,
	}
}


// AuthorizationRequestBuilder creates a new AuthorizationRequestBuilder
func AuthorizationRequestBuilder() *AuthorizationRequest {
	return &AuthorizationRequestBuilder{
		roles:     make([]Role, 0),
		tenantIDs: make([]string, 0),
	}
}

// WithURI sets the URI for the authorization request
func (b *AuthorizationRequestBuilder) WithURI(uri string) *AuthorizationRequest {
	b.uri = uri
	return b
}

// AddRole adds a role to the authorization request
func (b *AuthorizationRequestBuilder) AddRole(role Role) *AuthorizationRequest {
	b.roles = append(b.roles, role)
	return b
}

// AddTenantID adds a tenant ID to the authorization request
func (b *AuthorizationRequestBuilder) AddTenantID(tenantID string) *AuthorizationRequest {
	b.tenantIDs = append(b.tenantIDs, tenantID)
	return b
}

// Build creates and validates a new AuthorizationRequest
func (b *AuthorizationRequestBuilder) Build() (*AuthorizationRequest, error) {
	request := &AuthorizationRequest{
		Roles:     b.roles,
		URI:       b.uri,
		TenantIDs: b.tenantIDs,
	}

	if err := request.Validate(); err != nil {
		return nil, err
	}

	return request, nil
} 