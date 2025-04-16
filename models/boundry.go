package models

import (
	"fmt"
)

type BoundarySearchRequest struct {
	TenantID      string   `json:"tenant_id"`
	BoundryIds    []string `json:"boundary_ids"`
	BoundaryType  string   `json:"boundary_type"`
	BoundaryNum   []string `json:"boundary_name"`
	HierarchyType string   `json:"HierarchyType"`
	Codes         []string `json:"codes"`
}

func (b *BoundarySearchRequest) Validate() error {
	if b.TenantID == "" {
		return fmt.Errorf("tenant_id is required")
	}
	if len(b.TenantID) > 256 {
		return fmt.Errorf("tenant_id must be at most 256 characters")
	}
	if b.BoundaryType != "" && len(b.BoundaryType) > 64 {
		return fmt.Errorf("boundary_type is required")
	}
	if b.HierarchyType != "" && len(b.HierarchyType) > 128 {
		return fmt.Errorf("HierarchyType is required")
	}
	return nil
}

func (b *BoundarySearchRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":      b.TenantID,
		"boundaryIds":   b.BoundryIds,
		"boundaryType":  b.BoundaryType,
		"boundaryNum":   b.BoundaryNum,
		"HierarchyType": b.HierarchyType,
		"codes":         b.Codes,
	}
}

func BoundarySearchRequestBuilder() *BoundarySearchRequest {
	return &BoundarySearchRequest{}
}

func (b *BoundarySearchRequest) WithTenantID(tenantID string) *BoundarySearchRequest {
	b.TenantID = tenantID
	return b
}

func (b *BoundarySearchRequest) WithBoundaryIds(boundaryIds []string) *BoundarySearchRequest {
	b.BoundryIds = boundaryIds
	return b
}

func (b *BoundarySearchRequest) WithBoundaryType(boundaryType string) *BoundarySearchRequest {

	b.BoundaryType = boundaryType
	return b
}

func (b *BoundarySearchRequest) WithBoundaryNum(boundaryNum []string) *BoundarySearchRequest {
	b.BoundaryNum = boundaryNum
	return b
}

func (b *BoundarySearchRequest) WithHierarchyType(hierarchyType string) *BoundarySearchRequest {
	b.HierarchyType = hierarchyType
	return b
}

func (b *BoundarySearchRequest) WithCodes(codes []string) *BoundarySearchRequest {
	b.Codes = codes
	return b
}

func (b *BoundarySearchRequest) Build() *BoundarySearchRequest {
	boundarySearchRequest := &BoundarySearchRequest{
		TenantID:      b.TenantID,
		BoundryIds:    b.BoundryIds,
		BoundaryType:  b.BoundaryType,
		BoundaryNum:   b.BoundaryNum,
		HierarchyType: b.HierarchyType,
		Codes:         b.Codes,
	}
	if err := boundarySearchRequest.Validate(); err != nil {
		fmt.Printf("Error validating BoundarySearchRequest: %v\n", err)
		return nil
	}
	// If validation is successful, return the constructed object
	return boundarySearchRequest
}

type LocationBoundarySearchRequest struct {
	TenantID          string       `json:"tenant_id"`
	RequestInfo       *RequestInfo `json:"RequestInfo"`
	HierarchyTypeCode string       `json:"hierarchyTypeCode"`
	Codes             []string     `json:"codes"`
	BoundaryType      string       `json:"boundaryType"`
}

func (l *LocationBoundarySearchRequest) Validate() error {
	if l.TenantID == "" {
		return fmt.Errorf("tenant_id is required")
	}
	if len(l.TenantID) > 256 {
		return fmt.Errorf("tenant_id must be at most 256 characters")
	}
	if l.HierarchyTypeCode != "" && len(l.HierarchyTypeCode) > 128 {
		return fmt.Errorf("hierarchyTypeCode must be at most 128 characters")
	}
	if l.BoundaryType != "" && len(l.BoundaryType) > 64 {
		return fmt.Errorf("boundaryType must be at most 64 characters")
	}
	return nil
}

func (l *LocationBoundarySearchRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":          l.TenantID,
		"RequestInfo":       l.RequestInfo.ToMap(),
		"hierarchyTypeCode": l.HierarchyTypeCode,
		"codes":             l.Codes,
		"boundaryType":      l.BoundaryType,
	}
}

func LocationBoundarySearchRequestBuilder() *LocationBoundarySearchRequest {
	return &LocationBoundarySearchRequest{}
}

func (l *LocationBoundarySearchRequest) WithTenantID(tenantID string) *LocationBoundarySearchRequest {
	l.TenantID = tenantID
	return l
}

func (l *LocationBoundarySearchRequest) WithRequestInfo(requestInfo *RequestInfo) *LocationBoundarySearchRequest {
	l.RequestInfo = requestInfo
	return l
}

func (l *LocationBoundarySearchRequest) WithHierarchyTypeCode(hierarchyTypeCode string) *LocationBoundarySearchRequest {
	l.HierarchyTypeCode = hierarchyTypeCode
	return l
}

func (l *LocationBoundarySearchRequest) WithCodes(codes []string) *LocationBoundarySearchRequest {
	l.Codes = codes
	return l
}

func (l *LocationBoundarySearchRequest) WithBoundaryType(boundaryType string) *LocationBoundarySearchRequest {
	l.BoundaryType = boundaryType
	return l
}

func (l *LocationBoundarySearchRequest) Build() *LocationBoundarySearchRequest {
	lb := &LocationBoundarySearchRequest{
		TenantID:          l.TenantID,
		RequestInfo:       l.RequestInfo,
		HierarchyTypeCode: l.HierarchyTypeCode,
		Codes:             l.Codes,
		BoundaryType:      l.BoundaryType,
	}
	if err := lb.Validate(); err != nil {
		fmt.Printf("Error validating LocationBoundarySearchRequest: %v\n", err)
		return nil
	}
	return lb
}
