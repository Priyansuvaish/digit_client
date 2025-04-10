package models

// UserSearchModel represents the search criteria for user search
type UserSearchModel struct {
	TenantID     string
	ID           []int
	UUID         []string
	UserName     string
	Name         string
	MobileNumber string
	AadhaarNumber string
	Pan          string
	EmailID      string
	FuzzyLogic   bool
	Active       bool
	PageSize     int
	PageNumber   int
	Sort         []string
	UserType     string
	RoleCodes    []string
}

// UserSearchBuilder creates a new UserSearchModel instance
func UserSearchBuilder() *UserSearchModel {
	return &UserSearchModel{
		PageNumber: 0,
		Sort:      []string{"name"},
	}
}

// Builder methods for UserSearchModel
func (u *UserSearchModel) WithTenantID(tenantID string) *UserSearchModel {
	u.TenantID = tenantID
	return u
}

func (u *UserSearchModel) WithID(id []int) *UserSearchModel {
	u.ID = id
	return u
}

func (u *UserSearchModel) WithUUID(uuid []string) *UserSearchModel {
	u.UUID = uuid
	return u
}

func (u *UserSearchModel) WithUserName(userName string) *UserSearchModel {
	u.UserName = userName
	return u
}

func (u *UserSearchModel) WithName(name string) *UserSearchModel {
	u.Name = name
	return u
}

func (u *UserSearchModel) WithMobileNumber(mobileNumber string) *UserSearchModel {
	u.MobileNumber = mobileNumber
	return u
}

func (u *UserSearchModel) WithAadhaarNumber(aadhaarNumber string) *UserSearchModel {
	u.AadhaarNumber = aadhaarNumber
	return u
}

func (u *UserSearchModel) WithPan(pan string) *UserSearchModel {
	u.Pan = pan
	return u
}

func (u *UserSearchModel) WithEmailID(emailID string) *UserSearchModel {
	u.EmailID = emailID
	return u
}

func (u *UserSearchModel) WithFuzzyLogic(fuzzyLogic bool) *UserSearchModel {
	u.FuzzyLogic = fuzzyLogic
	return u
}

func (u *UserSearchModel) WithActive(active bool) *UserSearchModel {
	u.Active = active
	return u
}

func (u *UserSearchModel) WithPageSize(pageSize int) *UserSearchModel {
	u.PageSize = pageSize
	return u
}

func (u *UserSearchModel) WithPageNumber(pageNumber int) *UserSearchModel {
	u.PageNumber = pageNumber
	return u
}

func (u *UserSearchModel) WithSort(sort []string) *UserSearchModel {
	u.Sort = sort
	return u
}

func (u *UserSearchModel) WithUserType(userType string) *UserSearchModel {
	u.UserType = userType
	return u
}

func (u *UserSearchModel) WithRoleCodes(roleCodes []string) *UserSearchModel {
	u.RoleCodes = roleCodes
	return u
}

// ToMap converts the UserSearchModel to a map
func (u *UserSearchModel) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	// Add fields to map only if they have values
	if u.TenantID != "" {
		result["tenantId"] = u.TenantID
	}
	if len(u.ID) > 0 {
		result["id"] = u.ID
	}
	if len(u.UUID) > 0 {
		result["uuid"] = u.UUID
	}
	if u.UserName != "" {
		result["userName"] = u.UserName
	}
	if u.Name != "" {
		result["name"] = u.Name
	}
	if u.MobileNumber != "" {
		result["mobileNumber"] = u.MobileNumber
	}
	if u.AadhaarNumber != "" {
		result["aadhaarNumber"] = u.AadhaarNumber
	}
	if u.Pan != "" {
		result["pan"] = u.Pan
	}
	if u.EmailID != "" {
		result["emailId"] = u.EmailID
	}
	if u.UserType != "" {
		result["userType"] = u.UserType
	}
	if len(u.RoleCodes) > 0 {
		result["roleCodes"] = u.RoleCodes
	}

	// Add numeric fields if they have non-zero values
	if u.PageSize > 0 {
		result["pageSize"] = u.PageSize
	}
	if u.PageNumber > 0 {
		result["pageNumber"] = u.PageNumber
	}

	// Add boolean fields if they are set
	if u.FuzzyLogic {
		result["fuzzyLogic"] = u.FuzzyLogic
	}
	if u.Active {
		result["active"] = u.Active
	}

	// Add sort field if it's different from default
	if len(u.Sort) > 0 && (len(u.Sort) != 1 || u.Sort[0] != "name") {
		result["sort"] = u.Sort
	}

	return result
} 