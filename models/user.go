package models

type User struct {
	Id           int    `json:"id"`
	Uuid         string `json:"uuid"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobilenumber"`
	EmailId      string `json:"emailid"`
	Type         string `json:"type"`
	Roles        []Role `json:"roles"`
	TenantID     string `json:"tenantId"`
}

func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           u.Id,
		"uuid":         u.Uuid,
		"userName":     u.Username,
		"name":         u.Name,
		"mobileNumber": u.MobileNumber,
		"emailId":      u.EmailId,
		"type":         u.Type,
		"roles":        u.Roles,
		"tenantId":     u.TenantID,
	}
}

func UserBuilder() *User {
	return &User{}
}
func (u *User) WithID(code int) *User {
	u.Id = code
	return u
}

func (u *User) WithUuid(code string) *User {
	u.Uuid = code
	return u
}
func (u *User) WithUserName(code string) *User {
	u.Username = code
	return u
}
func (u *User) WithName(code string) *User {
	u.Name = code
	return u
}
func (u *User) WithMobilenumber(code string) *User {
	u.MobileNumber = code
	return u
}
func (u *User) WithEmailId(code string) *User {
	u.EmailId = code
	return u
}
func (u *User) WithType(code string) *User {
	u.Type = code
	return u
}
func (u *User) WithRoles(code []Role) *User {
	u.Roles = code
	return u
}
func (u *User) WithTenantID(code string) *User {
	u.TenantID = code
	return u
}

func (u *User) Build() (*User, error) {
	return u, nil
}
