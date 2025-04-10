package models

// CitizenUser represents a citizen user in the system
type CitizenUser struct {
	ID                   string
	UserName             string
	Password             string
	Salutation           string
	Name                 string
	Gender               string
	MobileNumber         string
	EmailID              string
	AltContactNumber     string
	Pan                  string
	AadhaarNumber        string
	PermanentAddress     string
	PermanentCity        string
	PermanentPinCode     string
	CorrespondenceAddress string
	CorrespondenceCity   string
	CorrespondencePinCode string
	Active               bool
	Dob                  string
	PwdExpiryDate        string
	Locale               string
	Type                 string
	Signature            string
	AccountLocked        bool
	Roles                []Role
	FatherOrHusbandName  string
	BloodGroup           string
	IdentificationMark   string
	Photo                string
	CreatedBy            string
	LastModifiedBy       string
	TenantID             string
	OtpReference         string
}

// Role represents a user role
type Role struct {
	ID       string
	Name     string
	Code     string
	TenantID string
}

// RoleBuilder creates a new Role instance
func RoleBuilder() *Role {
	return &Role{}
}
// Builder methods for Role
func (r *Role) WithID(id string) *Role {
	r.ID = id
	return r
}

func (r *Role) WithName(name string) *Role {
	r.Name = name
	return r
}

func (r *Role) WithCode(code string) *Role {
	r.Code = code
	return r
}

func (r *Role) WithTenantID(tenantID string) *Role {
	r.TenantID = tenantID
	return r
}

// Build validates and returns the Role instance
func (r *Role) Build() (*Role) {
	return r
}

// ToMap converts the Role to a map
func (r *Role) ToMap() map[string]interface{} {
	result := make(map[string]interface{}, 4) // Pre-allocate map with expected size
	if r.ID != "" {
		result["id"] = r.ID
	}
	if r.Name != "" {
		result["name"] = r.Name
	}
	if r.Code != "" {
		result["code"] = r.Code
	}
	if r.TenantID != "" {
		result["tenantId"] = r.TenantID
	}
	return result
}

// CreateCitizenUser creates a new CitizenUser with default values
func CreateCitizenUser() *CitizenUser {
	return &CitizenUser{
		Active:       true,
		AccountLocked: false,
		Locale:      "en_IN",
		Type:        "CITIZEN",
	}
}

// Builder methods for CitizenUser
func (u *CitizenUser) WithID(id string) *CitizenUser {
	u.ID = id
	return u
}

func (u *CitizenUser) WithUserName(userName string) *CitizenUser {
	u.UserName = userName
	return u
}

func (u *CitizenUser) WithPassword(password string) *CitizenUser {
	u.Password = password
	return u
}

func (u *CitizenUser) WithSalutation(salutation string) *CitizenUser {
	u.Salutation = salutation
	return u
}

func (u *CitizenUser) WithName(name string) *CitizenUser {
	u.Name = name
	return u
}

func (u *CitizenUser) WithGender(gender string) *CitizenUser {
	u.Gender = gender
	return u
}

func (u *CitizenUser) WithMobileNumber(mobileNumber string) *CitizenUser {
	u.MobileNumber = mobileNumber
	return u
}

func (u *CitizenUser) WithEmailID(emailID string) *CitizenUser {
	u.EmailID = emailID
	return u
}

func (u *CitizenUser) WithAltContactNumber(altContactNumber string) *CitizenUser {
	u.AltContactNumber = altContactNumber
	return u
}

func (u *CitizenUser) WithPan(pan string) *CitizenUser {
	u.Pan = pan
	return u
}

func (u *CitizenUser) WithAadhaarNumber(aadhaarNumber string) *CitizenUser {
	u.AadhaarNumber = aadhaarNumber
	return u
}

func (u *CitizenUser) WithPermanentAddress(address string) *CitizenUser {
	u.PermanentAddress = address
	return u
}

func (u *CitizenUser) WithPermanentCity(city string) *CitizenUser {
	u.PermanentCity = city
	return u
}

func (u *CitizenUser) WithPermanentPinCode(pinCode string) *CitizenUser {
	u.PermanentPinCode = pinCode
	return u
}

func (u *CitizenUser) WithCorrespondenceAddress(address string) *CitizenUser {
	u.CorrespondenceAddress = address
	return u
}

func (u *CitizenUser) WithCorrespondenceCity(city string) *CitizenUser {
	u.CorrespondenceCity = city
	return u
}

func (u *CitizenUser) WithCorrespondencePinCode(pinCode string) *CitizenUser {
	u.CorrespondencePinCode = pinCode
	return u
}

func (u *CitizenUser) WithActive(active bool) *CitizenUser {
	u.Active = active
	return u
}

func (u *CitizenUser) WithDob(dob string) *CitizenUser {
	u.Dob = dob
	return u
}

func (u *CitizenUser) WithPwdExpiryDate(expiryDate string) *CitizenUser {
	u.PwdExpiryDate = expiryDate
	return u
}

func (u *CitizenUser) WithLocale(locale string) *CitizenUser {
	u.Locale = locale
	return u
}

func (u *CitizenUser) WithType(type_ string) *CitizenUser {
	u.Type = type_
	return u
}

func (u *CitizenUser) WithSignature(signature string) *CitizenUser {
	u.Signature = signature
	return u
}

func (u *CitizenUser) WithAccountLocked(locked bool) *CitizenUser {
	u.AccountLocked = locked
	return u
}

func (u *CitizenUser) WithRoles(roles []Role) *CitizenUser {
	u.Roles = roles
	return u
}

func (u *CitizenUser) WithFatherOrHusbandName(name string) *CitizenUser {
	u.FatherOrHusbandName = name
	return u
}

func (u *CitizenUser) WithBloodGroup(bloodGroup string) *CitizenUser {
	u.BloodGroup = bloodGroup
	return u
}

func (u *CitizenUser) WithIdentificationMark(mark string) *CitizenUser {
	u.IdentificationMark = mark
	return u
}

func (u *CitizenUser) WithPhoto(photo string) *CitizenUser {
	u.Photo = photo
	return u
}

func (u *CitizenUser) WithCreatedBy(createdBy string) *CitizenUser {
	u.CreatedBy = createdBy
	return u
}

func (u *CitizenUser) WithLastModifiedBy(lastModifiedBy string) *CitizenUser {
	u.LastModifiedBy = lastModifiedBy
	return u
}

func (u *CitizenUser) WithTenantID(tenantID string) *CitizenUser {
	u.TenantID = tenantID
	return u
}

func (u *CitizenUser) WithOtpReference(otpReference string) *CitizenUser {
	u.OtpReference = otpReference
	return u
}

// ToMap converts the CitizenUser to a map
func (u *CitizenUser) ToMap() map[string]interface{} {
	roles := make([]map[string]interface{}, len(u.Roles))
	for i, role := range u.Roles {
		roles[i] = role.ToMap()
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
		"otpReference":        u.OtpReference,
	}
} 