package models

import (
	"errors"
	"time"
)

// Action represents a system action
type Action struct {
	ID               int
	Name             string
	URL              string
	DisplayName      string
	OrderNumber      int
	QueryParams      string
	ParentModule     string
	Enabled          bool
	ServiceCode      string
	TenantID         string
	CreatedDate      time.Time
	CreatedBy        int
	LastModifiedDate time.Time
	LastModifiedBy   int
	Path             string
	NavigationURL    string
	LeftIcon         string
	RightIcon        string
}

// // ActionBuilder constructs Action instances
// type ActionBuilder struct {
// 	action *Action
// }

// NewActionBuilder creates a new ActionBuilder
func ActionBuilder() *Action {
	return &Action{}
}

// Builder methods
func (b *Action) WithID(id int) *Action {
	b.ID = id
	return b
}

func (b *Action) WithName(name string) *Action {
	b.Name = name
	return b
}

func (b *Action) WithURL(url string) *Action {
	b.URL = url
	return b
}

func (b *Action) WithDisplayName(displayName string) *Action {
	b.DisplayName = displayName
	return b
}

func (b *Action) WithOrderNumber(orderNumber int) *Action {
	b.OrderNumber = orderNumber
	return b
}

func (b *Action) WithQueryParams(queryParams string) *Action {
	b.QueryParams = queryParams
	return b
}

func (b *Action) WithParentModule(parentModule string) *Action {
	b.ParentModule = parentModule
	return b
}

func (b *Action) WithEnabled(enabled bool) *Action {
	b.Enabled = enabled
	return b
}

func (b *Action) WithServiceCode(serviceCode string) *Action {
	b.ServiceCode = serviceCode
	return b
}

func (b *Action) WithTenantID(tenantID string) *Action {
	b.TenantID = tenantID
	return b
}

func (b *Action) WithCreatedDate(createdDate time.Time) *Action {
	b.CreatedDate = createdDate
	return b
}

func (b *Action) WithCreatedBy(createdBy int) *Action {
	b.CreatedBy = createdBy
	return b
}

func (b *Action) WithLastModifiedDate(lastModifiedDate time.Time) *Action {
	b.LastModifiedDate = lastModifiedDate
	return b
}

func (b *Action) WithLastModifiedBy(lastModifiedBy int) *Action {
	b.LastModifiedBy = lastModifiedBy
	return b
}

func (b *Action) WithPath(path string) *Action {
	b.Path = path
	return b
}

func (b *Action) WithNavigationURL(navigationURL string) *Action {
	b.NavigationURL = navigationURL
	return b
}

func (b *Action) WithLeftIcon(leftIcon string) *Action {
	b.LeftIcon = leftIcon
	return b
}

func (b *Action) WithRightIcon(rightIcon string) *Action {
	b.RightIcon = rightIcon
	return b
}

// Build validates and returns the Action
func (b *Action) Build() (*Action, error) {
	if len(b.Name) > 100 {
		return nil, errors.New("name exceeds 100 characters")
	}
	if len(b.URL) > 100 {
		return nil, errors.New("url exceeds 100 characters")
	}
	if len(b.DisplayName) > 100 {
		return nil, errors.New("displayName exceeds 100 characters")
	}
	if len(b.QueryParams) > 100 {
		return nil, errors.New("queryParams exceeds 100 characters")
	}
	if len(b.ParentModule) > 50 {
		return nil, errors.New("parentModule exceeds 50 characters")
	}
	if len(b.ServiceCode) > 50 {
		return nil, errors.New("serviceCode exceeds 50 characters")
	}
	if len(b.TenantID) > 50 {
		return nil, errors.New("tenantID exceeds 50 characters")
	}
	return b, nil
}

// ToMap converts Action to a map
func (a *Action) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               a.ID,
		"name":             a.Name,
		"url":              a.URL,
		"displayName":      a.DisplayName,
		"orderNumber":      a.OrderNumber,
		"queryParams":      a.QueryParams,
		"parentModule":     a.ParentModule,
		"enabled":          a.Enabled,
		"serviceCode":      a.ServiceCode,
		"tenantId":         a.TenantID,
		"createdDate":      a.CreatedDate.Format(time.RFC3339),
		"createdBy":        a.CreatedBy,
		"lastModifiedDate": a.LastModifiedDate.Format(time.RFC3339),
		"lastModifiedBy":   a.LastModifiedBy,
		"path":             a.Path,
		"navigationURL":    a.NavigationURL,
		"leftIcon":         a.LeftIcon,
		"rightIcon":        a.RightIcon,
	}
}

// ActionSearchCriteria represents search parameters
type ActionSearchCriteria struct {
	RoleCodes  []string
	FeatureIDs []int
	TenantID   string
}

func ActionSearchCriteriaBuilder() *ActionSearchCriteria {
	return &ActionSearchCriteria{RoleCodes: []string{}, FeatureIDs: []int{}}
}

func (b *ActionSearchCriteria) WithRoleCodes(roleCodes []string) *ActionSearchCriteria {
	b.RoleCodes = roleCodes
	return b
}

func (b *ActionSearchCriteria) WithFeatureIDs(featureIDs []int) *ActionSearchCriteria {
	b.FeatureIDs = featureIDs
	return b
}

func (b *ActionSearchCriteria) WithTenantID(tenantID string) *ActionSearchCriteria {
	b.TenantID = tenantID
	return b
}

func (b *ActionSearchCriteria) Build() *ActionSearchCriteria {
	return b
}

// ActionRequest represents an action API request
type ActionRequest struct {
	RequestInfo   *RequestInfo
	RoleCodes     []string
	FeatureIDs    []int
	TenantID      string
	Enabled       *bool
	Actions       []*Action
	ActionMaster  string
	NavigationURL string
	LeftIcon      string
	RightIcon     string
}

func ActionRequestBuilder() *ActionRequest {
	return &ActionRequest{RoleCodes: []string{}, FeatureIDs: []int{}, Actions: []*Action{}}
}

func (b *ActionRequest) WithRequestInfo(info *RequestInfo) *ActionRequest {
	b.RequestInfo = info
	return b
}

func (b *ActionRequest) WithRoleCodes(roleCodes []string) *ActionRequest {
	b.RoleCodes = roleCodes
	return b
}

func (b *ActionRequest) WithFeatureIDs(featureIDs []int) *ActionRequest {
	b.FeatureIDs = featureIDs
	return b
}

func (b *ActionRequest) WithTenantID(tenantID string) *ActionRequest {
	b.TenantID = tenantID
	return b
}

func (b *ActionRequest) WithEnabled(enabled bool) *ActionRequest {
	b.Enabled = &enabled
	return b
}

func (b *ActionRequest) WithActions(actions []*Action) *ActionRequest {
	b.Actions = actions
	return b
}

func (b *ActionRequest) WithActionMaster(master string) *ActionRequest {
	b.ActionMaster = master
	return b
}

func (b *ActionRequest) WithNavigationURL(url string) *ActionRequest {
	b.NavigationURL = url
	return b
}

func (b *ActionRequest) WithLeftIcon(icon string) *ActionRequest {
	b.LeftIcon = icon
	return b
}

func (b *ActionRequest) WithRightIcon(icon string) *ActionRequest {
	b.RightIcon = icon
	return b
}

func (b *ActionRequest) Build() (*ActionRequest, error) {
	return b, nil
}

// ToMap converts ActionRequest to a map
func (r *ActionRequest) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	result["RequestInfo"] = r.RequestInfo.ToMap()
	if len(r.RoleCodes) > 0 {
		result["roleCodes"] = r.RoleCodes
	}
	if len(r.FeatureIDs) > 0 {
		result["featureIds"] = r.FeatureIDs
	}
	if r.TenantID != "" {
		result["tenantId"] = r.TenantID
	}
	if r.Enabled != nil {
		result["enabled"] = *r.Enabled
	}
	if len(r.Actions) > 0 {
		actions := make([]map[string]interface{}, len(r.Actions))
		for i, a := range r.Actions {
			actions[i] = a.ToMap()
		}
		result["actions"] = actions
	}
	if r.ActionMaster != "" {
		result["actionMaster"] = r.ActionMaster
	}
	if r.NavigationURL != "" {
		result["navigationURL"] = r.NavigationURL
	}
	if r.LeftIcon != "" {
		result["leftIcon"] = r.LeftIcon
	}
	if r.RightIcon != "" {
		result["rightIcon"] = r.RightIcon
	}

	return result
}
