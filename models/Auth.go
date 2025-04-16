package models

type AuthenticationRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	TenantID  string `json:"tenant_id"`
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	Usertype  string `json:"Usertype"`
}

func (a *AuthenticationRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"username":   a.Username,
		"password":   a.Password,
		"tenantId":   a.TenantID,
		"grant_type": a.GrantType,
		"scope":      a.Scope,
		"UserType":   a.Usertype,
	}
}
func AuthenticationRequestBuilder() *AuthenticationRequest {
	return &AuthenticationRequest{}
}

func (b *AuthenticationRequest) WithUsername(username string) *AuthenticationRequest {
	b.Username = username
	return b
}

func (b *AuthenticationRequest) WithPassword(password string) *AuthenticationRequest {
	b.Password = password
	return b
}

func (b *AuthenticationRequest) WithTenantID(tenantID string) *AuthenticationRequest {
	b.TenantID = tenantID
	return b
}

func (b *AuthenticationRequest) WithGrantType(grantType string) *AuthenticationRequest {
	b.GrantType = grantType
	return b
}

func (b *AuthenticationRequest) WithScope(scope string) *AuthenticationRequest {
	b.Scope = scope
	return b
}

func (b *AuthenticationRequest) WithUserType(userType string) *AuthenticationRequest {
	b.Usertype = userType
	return b
}

func (b *AuthenticationRequest) Build() *AuthenticationRequest {
	return b
}
