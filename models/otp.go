package models

import (
	"fmt"
)

type Otp struct {
	Otp                  string `json:"otp"`
	Uuid                 string `json:"uuid"`
	Identity             string `json:"identity"`
	TenantID             string `json:"tenantid"`
	ValidationSuccessful bool   `json:"validationsuccessful"`
}

func (o *Otp) Validate() error {
	if len(o.Otp) > 128 {
		return fmt.Errorf("otp must be at most 128 characters")

	}
	if len(o.Uuid) > 36 {
		return fmt.Errorf("uuid must be at most 36 characters")
	}

	if len(o.Identity) > 100 {
		return fmt.Errorf("identity must be at most 100 characters")
	}

	if len(o.TenantID) > 256 {
		return fmt.Errorf("tenant_id must be at most 256 characters")
	}

	return nil
}

func (o *Otp) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"otp":                    o.Otp,
		"uuid":                   o.Uuid,
		"identity":               o.Identity,
		"tenantId":               o.TenantID,
		"isValidationSuccessful": o.ValidationSuccessful,
	}
}

type UserOtp struct {
	MobileNumber string `json:"mobilenumber"`
	TenantID     string `json:"tenantid"`
	Type         string `json:"type"`
	UserType     string `json:"usertype"`
}

func (u *UserOtp) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"mobileNumber": u.MobileNumber,
		"tenantId":     u.TenantID,
		"type":         u.Type,
		"userType":     u.UserType,
	}
}

func OtpBuilder() *Otp {
	return &Otp{}
}

func UserOtpBuilder() *UserOtp {
	return &UserOtp{}
}

func (o *Otp) WithOtp(otp string) *Otp {
	o.Otp = otp
	return o
}

func (o *Otp) WithUuid(uuid string) *Otp {
	o.Uuid = uuid
	return o
}
func (o *Otp) WithIdentify(identify string) *Otp {
	o.Identity = identify
	return o
}
func (o *Otp) WithTenantID(tenant_id string) *Otp {
	o.TenantID = tenant_id
	return o
}
func (o *Otp) WithValidationSuccessful(validation_successful bool) *Otp {
	o.ValidationSuccessful = validation_successful
	return o
}
func (o *Otp) Build() (*Otp, error) {
	err := o.Validate()
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (u *UserOtp) WithMobileNumber(mobile string) *UserOtp {
	u.MobileNumber = mobile
	return u
}

func (u *UserOtp) WithTenantID(tenant_id string) *UserOtp {
	u.TenantID = tenant_id
	return u
}

func (u *UserOtp) WithType(ty string) *UserOtp {
	u.Type = ty
	return u
}

func (u *UserOtp) WithUserType(usertype string) *UserOtp {
	u.UserType = usertype
	return u
}

func (u *UserOtp) Build() (*UserOtp, error) {
	return u, nil
}
