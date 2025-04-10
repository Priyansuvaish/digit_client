package models

// UserSearchModel represents the search criteria for users
type UserSearchModel struct {
	TenantID     string
	ID           []string
	UserName     string
	Name         string
	MobileNumber string
	AadhaarNumber string
	Pan          string
	EmailID      string
	Active       *bool
	AccountLocked *bool
	Roles        []string
	Type         string
	PageSize     int
	PageNumber   int
	SortBy       string
	SortOrder    string
}

// ToMap converts the UserSearchModel to a map
func (s *UserSearchModel) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	if s.TenantID != "" {
		result["tenantId"] = s.TenantID
	}
	if len(s.ID) > 0 {
		result["id"] = s.ID
	}
	if s.UserName != "" {
		result["userName"] = s.UserName
	}
	if s.Name != "" {
		result["name"] = s.Name
	}
	if s.MobileNumber != "" {
		result["mobileNumber"] = s.MobileNumber
	}
	if s.AadhaarNumber != "" {
		result["aadhaarNumber"] = s.AadhaarNumber
	}
	if s.Pan != "" {
		result["pan"] = s.Pan
	}
	if s.EmailID != "" {
		result["emailId"] = s.EmailID
	}
	if s.Active != nil {
		result["active"] = *s.Active
	}
	if s.AccountLocked != nil {
		result["accountLocked"] = *s.AccountLocked
	}
	if len(s.Roles) > 0 {
		result["roles"] = s.Roles
	}
	if s.Type != "" {
		result["type"] = s.Type
	}
	if s.PageSize > 0 {
		result["pageSize"] = s.PageSize
	}
	if s.PageNumber > 0 {
		result["pageNumber"] = s.PageNumber
	}
	if s.SortBy != "" {
		result["sortBy"] = s.SortBy
	}
	if s.SortOrder != "" {
		result["sortOrder"] = s.SortOrder
	}

	return result
} 