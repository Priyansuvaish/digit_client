package models

import (
	"fmt"
)

type IDRequest struct {
	Idname   string `json:"idname"`
	TenantID string `json:"tenant_id"`
	Format   string `json:"format"`
	Count    int    `json:"count"`
}

func (idRequest *IDRequest) Validate() error {
	if idRequest.Idname == "" || len(idRequest.Idname) > 200 {
		return fmt.Errorf("idname must be between 1 and 200 characters")
	}
	if idRequest.TenantID == "" || len(idRequest.TenantID) > 200 {
		return fmt.Errorf("tenant_id must be between 1 and 200 characters")
	}
	if len(idRequest.Format) > 200 {
		return fmt.Errorf("format must be between 1 and 200 characters")
	}
	return nil
}

func (idRequest *IDRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"idName":   idRequest.Idname,
		"tenantId": idRequest.TenantID,
		"format":   idRequest.Format,
		"count":    idRequest.Count,
	}
}

func IdRequestBuilder() *IDRequest {
	return &IDRequest{}
}

func (idRequest *IDRequest) SetIdName(idname string) *IDRequest {
	idRequest.Idname = idname
	return idRequest
}
func (idRequest *IDRequest) SetTenantID(tenantID string) *IDRequest {
	idRequest.TenantID = tenantID
	return idRequest
}
func (idRequest *IDRequest) SetFormat(format string) *IDRequest {
	idRequest.Format = format
	return idRequest
}
func (idRequest *IDRequest) SetCount(count int) *IDRequest {
	idRequest.Count = count
	return idRequest
}
func (idRequest *IDRequest) Build() (*IDRequest, error) {
	if err := idRequest.Validate(); err != nil {
		return nil, err
	}
	return idRequest, nil
}
