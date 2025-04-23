package models

import (
	"errors"
	"fmt"
)

// OrderEnum is the enum for sorting order
type OrderEnum string

const (
	ASC  OrderEnum = "asc"
	DESC OrderEnum = "desc"
)

type DataTypeEnum string

const (
	STRING          DataTypeEnum = "String"
	NUMBER          DataTypeEnum = "Number"
	TEXT            DataTypeEnum = "Text"
	DATETIME        DataTypeEnum = "Datetime"
	SINGLEVALUELIST DataTypeEnum = "SingleValueList"
	MULTIVALUELIST  DataTypeEnum = "MultiValueList"
	FILE            DataTypeEnum = "File"
)

type AttributeValue struct {
	AttributeCode     string                 `json:"attributeCode"`
	Value             interface{}            `json:"value"`
	ID                string                 `json:"id,omitempty"`
	ReferenceID       string                 `json:"referenceId,omitempty"`
	AuditDetails      *AuditDetails          `json:"auditDetails,omitempty"`
	AdditionalDetails map[string]interface{} `json:"additionalDetails,omitempty"`
}

func (a *AttributeValue) Validate() error {
	if a.AttributeCode == "" {
		return errors.New("attribute_code is required")
	}
	if a.ReferenceID != "" {
		l := len(a.ReferenceID)
		if l < 2 || l > 64 {
			return errors.New("reference_id must be between 2-64 characters")
		}
	}
	return nil
}

func (a *AttributeValue) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"attributeCode":     a.AttributeCode,
		"value":             a.Value,
		"id":                a.ID,
		"referenceId":       a.ReferenceID,
		"auditDetails":      a.AuditDetails.ToMap(),
		"additionalDetails": a.AdditionalDetails,
	}
}

type Service struct {
	TenantID          string                 `json:"tenantId"`
	ServiceDefID      string                 `json:"serviceDefId"`
	AccountID         string                 `json:"accountId"`
	Attributes        []AttributeValue       `json:"attributes"`
	ID                *string                `json:"id,omitempty"`
	ReferenceID       *string                `json:"referenceId,omitempty"`
	ClientID          *string                `json:"clientId,omitempty"`
	AuditDetails      *AuditDetails          `json:"auditDetails,omitempty"`
	AdditionalDetails map[string]interface{} `json:"additionalDetails,omitempty"`
}

func (s *Service) Validate() error {
	if len(s.TenantID) < 2 || len(s.TenantID) > 64 {
		return errors.New("tenant_id must be between 2-64 characters")
	}
	if len(s.ServiceDefID) < 2 || len(s.ServiceDefID) > 64 {
		return errors.New("service_def_id must be between 2-64 characters")
	}
	if s.ReferenceID != nil {
		l := len(*s.ReferenceID)
		if l < 2 || l > 64 {
			return errors.New("reference_id must be between 2-64 characters")
		}
	}
	if s.ClientID != nil && len(*s.ClientID) > 64 {
		return errors.New("client_id must be at most 64 characters")
	}
	if len(s.Attributes) == 0 {
		return errors.New("attributes cannot be empty")
	}
	return nil
}

func (s *Service) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":     s.TenantID,
		"serviceDefId": s.ServiceDefID,
		"accountId":    s.AccountID,
		"attributes": func() []map[string]interface{} {
			var attrs []map[string]interface{}
			for _, attr := range s.Attributes {
				attrs = append(attrs, attr.ToMap())
			}
			return attrs
		}(),
		"id":                s.ID,
		"referenceId":       s.ReferenceID,
		"clientId":          s.ClientID,
		"auditDetails":      s.AuditDetails.ToMap(),
		"additionalDetails": s.AdditionalDetails,
	}
}

type ServiceCriteria struct {
	TenantID        string   `json:"tenant_id"`
	Ids             []string `json:"ids"`
	Service_def_ids []string `json:"serviceDefId"`
	ReferenceID     []string `json:"referenceId"`
	AccountID       string   `json:"accountId"`
	Client_id       string   `json:"client_id"`
}

func (s *ServiceCriteria) Validate() error {
	if len(s.TenantID) < 2 || len(s.TenantID) > 64 {
		return fmt.Errorf("tenant_id must be between 2-64 characters")
	}
	return nil
}

func (s *ServiceCriteria) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":      s.TenantID,
		"ids":           s.Ids,
		"serviceDefIds": s.Service_def_ids,
		"referenceIds":  s.ReferenceID,
		"accountId":     s.AccountID,
		"clientId":      s.Client_id,
	}
}

type Pagination struct {
	Limit      int       `json:"limit"`
	Offset     int       `json:"offset"`
	Totalcount int       `json:"totalCount"`
	SortBy     string    `json:"sortBy"`
	Order      OrderEnum `json:"order"`
}

func (p *Pagination) Validate() error {
	if p.Limit < 1 {
		return fmt.Errorf("limit must be at least 1")
	}
	if p.Offset < 0 {
		return fmt.Errorf("offset cannot be negative")
	}
	return nil
}

func (p *Pagination) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"limit":      p.Limit,
		"offset":     p.Offset,
		"totalCount": p.Totalcount,
		"sortBy":     p.SortBy,
		"order":      p.Order,
	}
}

type AttributeDefinition struct {
	Code               string                 `json:"code"`
	Data_type          DataTypeEnum           `json:"datatype"`
	Id                 string                 `json:"id"`
	Reference_id       string                 `json:"referenceId"`
	Tenant_id          string                 `json:"tenantId"`
	Values             []string               `json:"values"`
	Is_active          bool                   `json:"isactive"`
	Required           bool                   `json:"required"`
	Regex              string                 `json:"regex"`
	Order              string                 `json:"order"`
	Audit_details      *AuditDetails          `json:"auditDetails"`
	Additional_details map[string]interface{} `json:"additionalDetails"`
}

func (a *AttributeDefinition) Validate() error {
	for _, field := range []string{a.Id, a.Reference_id, a.Tenant_id, a.Regex} {
		if len(field) < 2 || len(field) > 64 {
			return fmt.Errorf("Field must be between 2-64 characters")
		}
	}
	if len(a.Code) < 2 || len(a.Code) > 64 {
		return fmt.Errorf("code must be 2-64 characters")
	}
	return nil
}

func (a *AttributeDefinition) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":        a.Code,
		"dataType":    a.Data_type,
		"id":          a.Id,
		"referenceId": a.Reference_id,
		"tenantId":    a.Tenant_id,
		"values":      a.Values,
		"isActive":    a.Is_active,
		"required":    a.Required,
		"regex":       a.Regex,
		"order":       a.Order,
		"auditDetails": func() map[string]interface{} {
			if a.Audit_details != nil {
				return a.Audit_details.ToMap()
			}
			return nil
		}(),
		"additionalDetails": a.Additional_details,
	}
}

type ServiceDefinition struct {
	Code              string                 `json:"code"`
	TenantID          string                 `json:"tenantId"`
	AttributeS        []AttributeDefinition  `json:"attributes"`
	Id                string                 `json:"id"`
	IsActive          bool                   `json:"isActive"`
	ClientID          string                 `json:"clientId"`
	AuditDetails      *AuditDetails          `json:"auditdetails"`
	AdditionalDetails map[string]interface{} `json:"additionalDetails"`
}

func (s *ServiceDefinition) Validate() error {
	for _, field := range []string{s.Id, s.TenantID, s.Code, s.ClientID} {
		if len(field) < 2 || len(field) > 64 {
			return fmt.Errorf("Field must be between 2-64 characters")
		}
	}
	if s.AttributeS == nil {
		return fmt.Errorf("attributes cannot be empty")
	}
	return nil
}

func (s *ServiceDefinition) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":       s.Id,
		"tenantId": s.TenantID,
		"code":     s.Code,
		"isActive": s.IsActive,
		"attributes": func() []map[string]interface{} {
			var attrs []map[string]interface{}
			for _, attr := range s.AttributeS {
				attrs = append(attrs, attr.ToMap())
			}
			return attrs
		}(),
		"clientId": s.ClientID,
		"auditDetails": func() map[string]interface{} {
			if s.AuditDetails != nil {
				return s.AuditDetails.ToMap()
			}
			return nil
		}(),
		"additionalDetails": s.AdditionalDetails,
	}
}

type ServiceDefinitionCriteria struct {
	TenantID string   `json:"tenantId"`
	Ids      []string `json:"ids"`
	Codes    []string `json:"codes"`
	ClientID string   `json:"clientId"`
}

func (s *ServiceDefinitionCriteria) Validate() error {
	if len(s.TenantID) < 2 || len(s.TenantID) > 64 {
		return fmt.Errorf("tenant_id must be 2-64 characters")
	}
	return nil
}

func (s *ServiceDefinitionCriteria) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId": s.TenantID,
		"ids":      s.Ids,
		"code":     s.Codes,
		"clientId": s.ClientID,
	}
}

func AttributeValueBuilder() *AttributeValue {
	return &AttributeValue{}
}
func ServiceBuilder() *Service {
	return &Service{}
}
func ServiceCriteriaBuilder() *ServiceCriteria {
	return &ServiceCriteria{}
}
func PaginationBuilder() *Pagination {
	return &Pagination{}
}
func AttributeDefinitionBuilder() *AttributeDefinition {
	return &AttributeDefinition{}
}
func ServiceDefinitionBuilder() *ServiceDefinition {
	return &ServiceDefinition{}
}
func ServiceDefinitionCriteriaBuilder() *ServiceDefinitionCriteria {
	return &ServiceDefinitionCriteria{}
}
func (a *AttributeValue) WithAttributesCode(code string) *AttributeValue {
	a.AttributeCode = code
	return a
}
func (a *AttributeValue) WithValue(value any) *AttributeValue {
	a.Value = value
	return a
}
func (a *AttributeValue) WithId(code string) *AttributeValue {
	a.AttributeCode = code
	return a
}
func (a *AttributeValue) WithReferenceId(code string) *AttributeValue {
	a.ReferenceID = code
	return a
}
func (a *AttributeValue) WithAuditDetail(code *AuditDetails) *AttributeValue {
	a.AuditDetails = code
	return a
}
func (a *AttributeValue) WithAddtionalDetails(code map[string]interface{}) *AttributeValue {
	a.AdditionalDetails = code
	return a
}
