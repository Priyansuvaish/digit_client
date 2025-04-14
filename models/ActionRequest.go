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

// ActionBuilder constructs Action instances
type ActionBuilder struct {
	action *Action
}

// NewActionBuilder creates a new ActionBuilder
func NewActionBuilder() *ActionBuilder {
	return &ActionBuilder{action: &Action{}}
}

// Builder methods
func (b *ActionBuilder) WithID(id int) *ActionBuilder {
	b.action.ID = id
	return b
}

func (b *ActionBuilder) WithName(name string) *ActionBuilder {
	b.action.Name = name
	return b
}

func (b *ActionBuilder) WithURL(url string) *ActionBuilder {
	b.action.URL = url
	return b
}

func (b *ActionBuilder) WithDisplayName(displayName string) *ActionBuilder {
	b.action.DisplayName = displayName
	return b
}

func (b *ActionBuilder) WithOrderNumber(orderNumber int) *ActionBuilder {
	b.action.OrderNumber = orderNumber
	return b
}

func (b *ActionBuilder) WithQueryParams(queryParams string) *ActionBuilder {
	b.action.QueryParams = queryParams
	return b
}

func (b *ActionBuilder) WithParentModule(parentModule string) *ActionBuilder {
	b.action.ParentModule = parentModule
	return b
}

func (b *ActionBuilder) WithEnabled(enabled bool) *ActionBuilder {
	b.action.Enabled = enabled
	return b
}

func (b *ActionBuilder) WithServiceCode(serviceCode string) *ActionBuilder {
	b.action.ServiceCode = serviceCode
	return b
}

func (b *ActionBuilder) WithTenantID(tenantID string) *ActionBuilder {
	b.action.TenantID = tenantID
	return b
}

func (b *ActionBuilder) WithCreatedDate(createdDate time.Time) *ActionBuilder {
	b.action.CreatedDate = createdDate
	return b
}

func (b *ActionBuilder) WithCreatedBy(createdBy int) *ActionBuilder {
	b.action.CreatedBy = createdBy
	return b
}

func (b *ActionBuilder) WithLastModifiedDate(lastModifiedDate time.Time) *ActionBuilder {
	b.action.LastModifiedDate = lastModifiedDate
	return b
}

func (b *ActionBuilder) WithLastModifiedBy(lastModifiedBy int) *ActionBuilder {
	b.action.LastModifiedBy = lastModifiedBy
	return b
}

func (b *ActionBuilder) WithPath(path string) *ActionBuilder {
	b.action.Path = path
	return b
}

func (b *ActionBuilder) WithNavigationURL(navigationURL string) *ActionBuilder {
	b.action.NavigationURL = navigationURL
	return b
}

func (b *ActionBuilder) WithLeftIcon(leftIcon string) *ActionBuilder {
	b.action.LeftIcon = leftIcon
	return b
}

func (b *ActionBuilder) WithRightIcon(rightIcon string) *ActionBuilder {
	b.action.RightIcon = rightIcon
	return b
}

// Build validates and returns the Action
func (b *ActionBuilder) Build() (*Action, error) {
	if len(b.action.Name) > 100 {
		return nil, errors.New("name exceeds 100 characters")
	}
	if len(b.action.URL) > 100 {
		return nil, errors.New("url exceeds 100 characters")
	}
	if len(b.action.DisplayName) > 100 {
		return nil, errors.New("displayName exceeds 100 characters")
	}
	if len(b.action.QueryParams) > 100 {
		return nil, errors.New("queryParams exceeds 100 characters")
	}
	if len(b.action.ParentModule) > 50 {
		return nil, errors.New("parentModule exceeds 50 characters")
	}
	if len(b.action.ServiceCode) > 50 {
		return nil, errors.New("serviceCode exceeds 50 characters")
	}
	if len(b.action.TenantID) > 50 {
		return nil, errors.New("tenantID exceeds 50 characters")
	}
	return b.action, nil
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

// ActionSearchCriteriaBuilder constructs search criteria
type ActionSearchCriteriaBuilder struct {
	criteria *ActionSearchCriteria
}

func NewActionSearchCriteriaBuilder() *ActionSearchCriteriaBuilder {
	return &ActionSearchCriteriaBuilder{
		criteria: &ActionSearchCriteria{
			RoleCodes:  []string{},
			FeatureIDs: []int{},
		},
	}
}

func (b *ActionSearchCriteriaBuilder) WithRoleCodes(roleCodes []string) *ActionSearchCriteriaBuilder {
	b.criteria.RoleCodes = roleCodes
	return b
}

func (b *ActionSearchCriteriaBuilder) WithFeatureIDs(featureIDs []int) *ActionSearchCriteriaBuilder {
	b.criteria.FeatureIDs = featureIDs
	return b
}

func (b *ActionSearchCriteriaBuilder) WithTenantID(tenantID string) *ActionSearchCriteriaBuilder {
	b.criteria.TenantID = tenantID
	return b
}

func (b *ActionSearchCriteriaBuilder) Build() *ActionSearchCriteria {
	return b.criteria
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

// ActionRequestBuilder constructs ActionRequest instances
type ActionRequestBuilder struct {
	request *ActionRequest
}

func NewActionRequestBuilder() *ActionRequestBuilder {
	return &ActionRequestBuilder{
		request: &ActionRequest{
			RoleCodes:  []string{},
			FeatureIDs: []int{},
			Actions:    []*Action{},
		},
	}
}

func (b *ActionRequestBuilder) WithRequestInfo(info *RequestInfo) *ActionRequestBuilder {
	b.request.RequestInfo = info
	return b
}

func (b *ActionRequestBuilder) WithRoleCodes(roleCodes []string) *ActionRequestBuilder {
	b.request.RoleCodes = roleCodes
	return b
}

func (b *ActionRequestBuilder) WithFeatureIDs(featureIDs []int) *ActionRequestBuilder {
	b.request.FeatureIDs = featureIDs
	return b
}

func (b *ActionRequestBuilder) WithTenantID(tenantID string) *ActionRequestBuilder {
	b.request.TenantID = tenantID
	return b
}

func (b *ActionRequestBuilder) WithEnabled(enabled bool) *ActionRequestBuilder {
	b.request.Enabled = &enabled
	return b
}

func (b *ActionRequestBuilder) WithActions(actions []*Action) *ActionRequestBuilder {
	b.request.Actions = actions
	return b
}

func (b *ActionRequestBuilder) WithActionMaster(master string) *ActionRequestBuilder {
	b.request.ActionMaster = master
	return b
}

func (b *ActionRequestBuilder) WithNavigationURL(url string) *ActionRequestBuilder {
	b.request.NavigationURL = url
	return b
}

func (b *ActionRequestBuilder) WithLeftIcon(icon string) *ActionRequestBuilder {
	b.request.LeftIcon = icon
	return b
}

func (b *ActionRequestBuilder) WithRightIcon(icon string) *ActionRequestBuilder {
	b.request.RightIcon = icon
	return b
}

func (b *ActionRequestBuilder) Build() *ActionRequest {
	return b.request
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
