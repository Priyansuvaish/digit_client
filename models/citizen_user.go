package models

// CitizenUser represents a citizen user in the system
type CitizenUser struct {
	ID           string
	UserName     string
	Password     string
	Salutation   string
	Name         string
	Gender       string
	MobileNumber string
	EmailID      string
	AltContactNumber string
	Pan          string
	AadhaarNumber string
	PermanentAddress string
	PermanentCity string
	PermanentPinCode string
	CorrespondenceAddress string
	CorrespondenceCity string
	CorrespondencePinCode string
	Active       bool
	Dob          string
	PwdExpiryDate string
	Locale       string
	Type         string
	Signature    string
	AccountLocked bool
	Roles        []Role
	FatherOrHusbandName string
	BloodGroup   string
	IdentificationMark string
	Photo        string
	CreatedBy    string
	LastModifiedBy string
	TenantID     string
}

// Role represents a user role
type Role struct {
	ID          string
	Name        string
	Code        string
	TenantID    string
}

// ToMap converts the CitizenUser to a map
func (u *CitizenUser) ToMap() map[string]interface{} {
	roles := make([]map[string]interface{}, len(u.Roles))
	for i, role := range u.Roles {
		roles[i] = map[string]interface{}{
			"id":       role.ID,
			"name":     role.Name,
			"code":     role.Code,
			"tenantId": role.TenantID,
		}
	}

	return map[string]interface{}{
		"id":                   u.ID,
		"userName":            u.UserName,
		"password":            u.Password,
		"salutation":          u.Salutation,
		"name":                u.Name,
		"gender":              u.Gender,
		"mobileNumber":        u.MobileNumber,
		"emailId":             u.EmailID,
		"altContactNumber":    u.AltContactNumber,
		"pan":                 u.Pan,
		"aadhaarNumber":       u.AadhaarNumber,
		"permanentAddress":    u.PermanentAddress,
		"permanentCity":       u.PermanentCity,
		"permanentPinCode":    u.PermanentPinCode,
		"correspondenceAddress": u.CorrespondenceAddress,
		"correspondenceCity":   u.CorrespondenceCity,
		"correspondencePinCode": u.CorrespondencePinCode,
		"active":              u.Active,
		"dob":                 u.Dob,
		"pwdExpiryDate":       u.PwdExpiryDate,
		"locale":              u.Locale,
		"type":                u.Type,
		"signature":           u.Signature,
		"accountLocked":       u.AccountLocked,
		"roles":               roles,
		"fatherOrHusbandName": u.FatherOrHusbandName,
		"bloodGroup":          u.BloodGroup,
		"identificationMark":  u.IdentificationMark,
		"photo":               u.Photo,
		"createdBy":           u.CreatedBy,
		"lastModifiedBy":      u.LastModifiedBy,
		"tenantId":            u.TenantID,
	}
} 