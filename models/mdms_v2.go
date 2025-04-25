package models

import (
	"fmt"
)

type AuditDetails struct {
	CreatedBy        string `json:"createdBy"`
	CreatedTime      string `json:"createdTime"`
	LastModifiedBy   string `json:"lastModifiedBy"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

func (a *AuditDetails) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"createdBy":        a.CreatedBy,
		"createdTime":      a.CreatedTime,
		"lastModifiedBy":   a.LastModifiedBy,
		"lastModifiedTime": a.LastModifiedTime,
	}
}

func (a *AuditDetails) AuditDetailsBuilder() *AuditDetails {
	return &AuditDetails{}
}

func (a *AuditDetails) WithCreatedBy(createdBy string) *AuditDetails {
	a.CreatedBy = createdBy
	return a
}

func (a *AuditDetails) WithCreatedTime(createdTime string) *AuditDetails {
	a.CreatedTime = createdTime
	return a
}

func (a *AuditDetails) WithLastModifiedBy(lastModifiedBy string) *AuditDetails {
	a.LastModifiedBy = lastModifiedBy
	return a
}

func (a *AuditDetails) WithLastModifiedTime(lastModifiedTime string) *AuditDetails {
	a.LastModifiedTime = lastModifiedTime
	return a
}

func (a *AuditDetails) Build() *AuditDetails {
	return a
}

type SchemaDefinition struct {
	TenantID     string        `json:"tenantId"`
	Code         string        `json:"code"`
	Defination   string        `json:"definition"`
	ID           string        `json:"id"`
	Discription  string        `json:"description"`
	ISactive     bool          `json:"is_active"`
	AuditDetails *AuditDetails `json:"auditDetails"`
}

func (s *SchemaDefinition) Validate() error {
	if len(s.ID) < 2 || len(s.ID) > 128 {
		return fmt.Errorf("id must be between 2 and 128 characters")
	}
	if len(s.TenantID) < 2 || len(s.TenantID) > 128 {
		return fmt.Errorf("tenant_id must be between 2 and 128 characters")
	}
	if s.Code != "" && len(s.Code) < 2 || len(s.Code) > 128 {
		return fmt.Errorf("code must be between 2 and 128 characters")
	}
	if s.Discription != "" && len(s.Discription) < 2 || len(s.Discription) > 512 {
		return fmt.Errorf("description must be between 2 and 512 characters")
	}
	return nil
}

func (s *SchemaDefinition) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":     s.TenantID,
		"code":         s.Code,
		"definition":   s.Defination,
		"description":  s.Discription,
		"id":           s.ID,
		"isActive":     s.ISactive,
		"auditDetails": s.AuditDetails.ToMap(),
	}
}

func SchemaDefinitionBuilder() *SchemaDefinition {
	return &SchemaDefinition{}
}

func (s *SchemaDefinition) WithTenantID(tenantID string) *SchemaDefinition {
	s.TenantID = tenantID
	return s
}

func (s *SchemaDefinition) WithCode(code string) *SchemaDefinition {
	s.Code = code
	return s
}

func (s *SchemaDefinition) WithDefination(defination string) *SchemaDefinition {
	s.Defination = defination
	return s
}

func (s *SchemaDefinition) WithID(id string) *SchemaDefinition {
	s.ID = id
	return s
}

func (s *SchemaDefinition) WithDiscription(discription string) *SchemaDefinition {
	s.Discription = discription
	return s
}

func (s *SchemaDefinition) WithISactive(isActive bool) *SchemaDefinition {
	s.ISactive = isActive
	return s
}

func (s *SchemaDefinition) WithAuditDetails(auditDetails *AuditDetails) *SchemaDefinition {

	s.AuditDetails = auditDetails
	return s
}

func (s *SchemaDefinition) Build() *SchemaDefinition {
	sd := &SchemaDefinition{
		TenantID:     s.TenantID,
		Code:         s.Code,
		Defination:   s.Defination,
		ID:           s.ID,
		Discription:  s.Discription,
		ISactive:     s.ISactive,
		AuditDetails: s.AuditDetails,
	}
	if err := sd.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return nil
	}
	// If validation is successful, return the constructed object
	return s
}

type SchemaDefCriteria struct {
	TenantID string   `json:"tenantId"`
	Codes    []string `json:"codes"`
	Offeset  int      `json:"offset"`
	Limit    int      `json:"limit"`
}

func (s *SchemaDefCriteria) Validate() error {
	if len(s.TenantID) < 2 || len(s.TenantID) > 128 {
		return fmt.Errorf("tenant_id must be between 2 and 128 characters")
	}
	return nil
}

func (s *SchemaDefCriteria) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId": s.TenantID,
		"codes":    s.Codes,
		"offset":   s.Offeset,
		"limit":    s.Limit,
	}
}
func SchemaDefCriteriaBuilder() *SchemaDefCriteria {
	return &SchemaDefCriteria{}
}

func (s *SchemaDefCriteria) WithTenantID(tenantID string) *SchemaDefCriteria {
	s.TenantID = tenantID
	return s
}

func (s *SchemaDefCriteria) WithCodes(codes []string) *SchemaDefCriteria {

	s.Codes = codes
	return s
}

func (s *SchemaDefCriteria) WithOffeset(offeset int) *SchemaDefCriteria {
	s.Offeset = offeset
	return s
}

func (s *SchemaDefCriteria) WithLimit(limit int) *SchemaDefCriteria {
	s.Limit = limit
	return s
}

func (s *SchemaDefCriteria) Build() *SchemaDefCriteria {
	schemaDefCriteria := &SchemaDefCriteria{
		TenantID: s.TenantID,
		Codes:    s.Codes,
		Offeset:  s.Offeset,
		Limit:    s.Limit,
	}
	if err := schemaDefCriteria.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return nil
	}
	// If validation is successful, return the constructed object
	return s
}

type Mdms struct {
	TenantID         string         `json:"tenantId"`
	SchemeCode       string         `json:"schemeCode"`
	Data             map[string]any `json:"data"`
	ID               string         `json:"id"`
	UniqueIdentifier string         `json:"uniqueIdentifier"`
	IsActive         bool           `json:"isActive"`
	AuditDetails     *AuditDetails  `json:"auditDetails"`
}

func (m *Mdms) Validate() error {
	if len(m.ID) < 2 || len(m.ID) > 64 {
		return fmt.Errorf("id must be between 2 and 64 characters")
	}
	if len(m.TenantID) < 2 || len(m.TenantID) > 128 {
		return fmt.Errorf("tenant_id must be between 2 and 128 characters")
	}
	if m.SchemeCode != "" && len(m.SchemeCode) < 2 || len(m.SchemeCode) > 128 {
		return fmt.Errorf("scheme_code must be between 2 and 128 characters")
	}
	if m.UniqueIdentifier != "" && len(m.UniqueIdentifier) < 1 || len(m.UniqueIdentifier) > 128 {
		return fmt.Errorf("unique_identifier must be between 2 and 128 characters")
	}
	return nil
}

func (m *Mdms) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":         m.TenantID,
		"schemeCode":       m.SchemeCode,
		"data":             m.Data,
		"id":               m.ID,
		"uniqueIdentifier": m.UniqueIdentifier,
		"isActive":         m.IsActive,
		"auditDetails":     m.AuditDetails.ToMap(),
	}
}

func MdmsBuilder() *Mdms {
	return &Mdms{}
}

func (m *Mdms) WithTenantID(tenantID string) *Mdms {
	m.TenantID = tenantID
	return m
}

func (m *Mdms) WithSchemeCode(schemeCode string) *Mdms {

	m.SchemeCode = schemeCode
	return m
}

func (m *Mdms) WithData(data map[string]any) *Mdms {
	m.Data = data
	return m
}

func (m *Mdms) WithID(id string) *Mdms {

	m.ID = id
	return m
}

func (m *Mdms) WithUniqueIdentifier(uniqueIdentifier string) *Mdms {

	m.UniqueIdentifier = uniqueIdentifier
	return m
}

func (m *Mdms) WithIsActive(isActive bool) *Mdms {

	m.IsActive = isActive
	return m
}

func (m *Mdms) WithAuditDetails(auditDetails *AuditDetails) *Mdms {

	m.AuditDetails = auditDetails
	return m
}

func (m *Mdms) Build() *Mdms {
	err := m.Validate()
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return nil
	}
	// If validation is successful, return the constructed object
	return m
}

type MdmsCriteriaV2 struct {
	TenantID                           string         `json:"tenantId"`
	IDS                                []string       `json:"ids"`
	UniqueIdentifiers                  []string       `json:"uniqueIdentifiers"`
	SchemaCode                         string         `json:"schemeCode"`
	FilterMap                          map[string]any `json:"filterMap"`
	IsActive                           bool           `json:"isActive"`
	SchemaCodeFilterMap                map[string]any `json:"schemaCodeFilterMap"`
	UniqueIdentifierForRefVerification []string       `json:"uniqueIdentifierForRefVerification"`
	Offeset                            int            `json:"offset"`
	Limit                              int            `json:"limit"`
}

func (m *MdmsCriteriaV2) Validate() error {
	if len(m.TenantID) < 1 || len(m.TenantID) > 100 {
		return fmt.Errorf("tenant_id must be between 2 and 100 characters")
	}
	if len(m.UniqueIdentifiers) > 0 {
		for _, uid := range m.UniqueIdentifiers {
			if len(uid) < 1 || len(uid) > 64 {
				return fmt.Errorf("unique_identifiers must be between 1 and 64 characters")
			}
		}
	}
	return nil
}

func (m *MdmsCriteriaV2) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":                            m.TenantID,
		"ids":                                 m.IDS,
		"uniqueIdentifiers":                   m.UniqueIdentifiers,
		"schemeCode":                          m.SchemaCode,
		"filters":                             m.FilterMap,
		"isActive":                            m.IsActive,
		"schemaCodeFilterMap":                 m.SchemaCodeFilterMap,
		"uniqueIdentifiersForRefVerification": m.UniqueIdentifierForRefVerification,
		"offset":                              m.Offeset,
		"limit":                               m.Limit,
	}
}

func MdmsCriteriaV2Builder() *MdmsCriteriaV2 {
	return &MdmsCriteriaV2{}
}

func (m *MdmsCriteriaV2) WithTenantID(tenantID string) *MdmsCriteriaV2 {
	m.TenantID = tenantID
	return m
}

func (m *MdmsCriteriaV2) WithIDS(ids []string) *MdmsCriteriaV2 {
	m.IDS = ids
	return m
}

func (m *MdmsCriteriaV2) WithUniqueIdentifiers(uniqueIdentifiers []string) *MdmsCriteriaV2 {
	m.UniqueIdentifiers = uniqueIdentifiers
	return m
}

func (m *MdmsCriteriaV2) WithSchemaCode(schemaCode string) *MdmsCriteriaV2 {
	m.SchemaCode = schemaCode
	return m
}

func (m *MdmsCriteriaV2) WithFilterMap(filterMap map[string]any) *MdmsCriteriaV2 {
	m.FilterMap = filterMap
	return m
}

func (m *MdmsCriteriaV2) WithIsActive(isActive bool) *MdmsCriteriaV2 {
	m.IsActive = isActive
	return m
}

func (m *MdmsCriteriaV2) WithSchemaCodeFilterMap(schemaCodeFilterMap map[string]any) *MdmsCriteriaV2 {
	m.SchemaCodeFilterMap = schemaCodeFilterMap
	return m
}

func (m *MdmsCriteriaV2) WithUniqueIdentifierForRefVerification(uniqueIdentifierForRefVerification []string) *MdmsCriteriaV2 {

	m.UniqueIdentifierForRefVerification = uniqueIdentifierForRefVerification
	return m
}

func (m *MdmsCriteriaV2) WithOffeset(offeset int) *MdmsCriteriaV2 {
	m.Offeset = offeset
	return m
}

func (m *MdmsCriteriaV2) WithLimit(limit int) *MdmsCriteriaV2 {
	m.Limit = limit
	return m
}

func (m *MdmsCriteriaV2) Build() *MdmsCriteriaV2 {
	mdmsCriteriaV2 := &MdmsCriteriaV2{
		TenantID:                           m.TenantID,
		IDS:                                m.IDS,
		UniqueIdentifiers:                  m.UniqueIdentifiers,
		SchemaCode:                         m.SchemaCode,
		FilterMap:                          m.FilterMap,
		IsActive:                           m.IsActive,
		SchemaCodeFilterMap:                m.SchemaCodeFilterMap,
		UniqueIdentifierForRefVerification: m.UniqueIdentifierForRefVerification,
		Offeset:                            m.Offeset,
		Limit:                              m.Limit,
	}
	if err := mdmsCriteriaV2.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return nil
	}
	return m
}
