package models

import (
	"fmt"
)

type BusinessService struct {
	TenantID           string        `json:"TenantID"`
	BusinessService    string        `json:"BusinessService"`
	Business           string        `json:"Business"`
	States             []string      `json:"States"`
	Uuid               string        `json:"Uuid"`
	GetUri             string        `json:"GetUri"`
	PostUri            string        `json:"PostUri"`
	BusinessServiceSla int           `json:"BusinessServiceSla"`
	AuditDetails       *AuditDetails `json:"AuditDetails"`
}

func (b *BusinessService) Validate() error {
	if b.TenantID == "" {
		return fmt.Errorf("tenant_id is required")
	}
	if len(b.TenantID) > 256 {
		return fmt.Errorf("tenant_id must be at most 256 characters")
	}
	if b.Uuid != "" && len(b.Uuid) > 256 {
		return fmt.Errorf("uuid must be at most 256 characters")
	}
	if b.BusinessService != "" && len(b.BusinessService) > 256 {
		return fmt.Errorf("business_service must be at most 64 characters")
	}
	if b.Business != "" && len(b.Business) > 256 {
		return fmt.Errorf("business must be at most 256 characters")
	}
	if b.GetUri != "" && len(b.GetUri) > 1024 {
		return fmt.Errorf("get_uri must be at most 1024 characters")
	}
	if b.PostUri != "" && len(b.PostUri) > 1024 {
		return fmt.Errorf("post_uri must be at most 1024 characters")
	}
	if b.States == nil {
		return fmt.Errorf("states is required")
	}
	return nil
}

func (b *BusinessService) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":           b.TenantID,
		"businessService":    b.BusinessService,
		"business":           b.Business,
		"states":             b.States,
		"uuid":               b.Uuid,
		"getUri":             b.GetUri,
		"postUri":            b.PostUri,
		"businessServiceSla": b.BusinessServiceSla,
		"auditDetails":       b.AuditDetails.ToMap(),
	}
}

func BusinessServiceBuilder() *BusinessService {
	return &BusinessService{}
}

func (b *BusinessService) WithTenantID(tenantID string) *BusinessService {
	b.TenantID = tenantID
	return b
}

func (b *BusinessService) WithBusinessService(businessService string) *BusinessService {
	b.BusinessService = businessService
	return b
}

func (b *BusinessService) WithBusiness(business string) *BusinessService {
	b.Business = business
	return b
}

func (b *BusinessService) WithStates(states []string) *BusinessService {
	b.States = states
	return b
}

func (b *BusinessService) WithUuid(uuid string) *BusinessService {
	b.Uuid = uuid
	return b
}

func (b *BusinessService) WithGetUri(getUri string) *BusinessService {
	b.GetUri = getUri
	return b
}

func (b *BusinessService) WithPostUri(postUri string) *BusinessService {
	b.PostUri = postUri
	return b
}

func (b *BusinessService) WithBusinessServiceSla(businessServiceSla int) *BusinessService {
	b.BusinessServiceSla = businessServiceSla
	return b
}

func (b *BusinessService) WithAuditDetails(auditDetails *AuditDetails) *BusinessService {
	b.AuditDetails = auditDetails
	return b
}

func (b *BusinessService) Build() *BusinessService {
	businessService := &BusinessService{
		TenantID:           b.TenantID,
		BusinessService:    b.BusinessService,
		Business:           b.Business,
		States:             b.States,
		Uuid:               b.Uuid,
		GetUri:             b.GetUri,
		PostUri:            b.PostUri,
		BusinessServiceSla: b.BusinessServiceSla,
		AuditDetails:       b.AuditDetails,
	}
	if err := businessService.Validate(); err != nil {
		panic(fmt.Sprintf("BusinessService validation failed: %v", err))
	}
	// If validation is successful, return the constructed object
	return businessService
}

type BusinessServiceSearchCriteria struct {
	TenantID         string   `json:"tenantId"`
	BusinessServices []string `json:"businessServices"`
	StateUUIDs       []string `json:"stateUuids"`
	ActionUUIDs      []string `json:"actionUuids"`
}

func (b *BusinessServiceSearchCriteria) Validate() error {
	if b.TenantID == "" {
		return fmt.Errorf("tenant_id is required")
	}
	return nil
}

func (b *BusinessServiceSearchCriteria) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":         b.TenantID,
		"businessServices": b.BusinessServices,
		"stateUuids":       b.StateUUIDs,
		"actionUuids":      b.ActionUUIDs,
	}
}

func BusinessServiceSearchCriteriaBuilder() *BusinessServiceSearchCriteria {
	return &BusinessServiceSearchCriteria{}
}

func (b *BusinessServiceSearchCriteria) WithTenantID(tenantID string) *BusinessServiceSearchCriteria {
	b.TenantID = tenantID
	return b
}

func (b *BusinessServiceSearchCriteria) WithBusinessServices(businessServices []string) *BusinessServiceSearchCriteria {
	b.BusinessServices = businessServices
	return b
}

func (b *BusinessServiceSearchCriteria) WithStateUUIDs(stateUUIDs []string) *BusinessServiceSearchCriteria {
	b.StateUUIDs = stateUUIDs
	return b
}

func (b *BusinessServiceSearchCriteria) WithActionUUIDs(actionUUIDs []string) *BusinessServiceSearchCriteria {
	b.ActionUUIDs = actionUUIDs
	return b
}

func (b *BusinessServiceSearchCriteria) Build() *BusinessServiceSearchCriteria {
	businessServiceSearchCriteria := &BusinessServiceSearchCriteria{
		TenantID:         b.TenantID,
		BusinessServices: b.BusinessServices,
		StateUUIDs:       b.StateUUIDs,
		ActionUUIDs:      b.ActionUUIDs,
	}
	if err := businessServiceSearchCriteria.Validate(); err != nil {
		panic(fmt.Sprintf("BusinessServiceSearchCriteria validation failed: %v", err))
	}
	// If validation is successful, return the constructed object
	return businessServiceSearchCriteria
}
