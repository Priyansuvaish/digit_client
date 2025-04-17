package models

import (
	"fmt"
)

type PaginationDetails struct {
	OffsetKey      string `json:"offset"`
	SizeKey        string `json:"size"`
	MaxPageSize    int    `json:"maxPageSize"`
	StartingOffset int    `json:"startingOffset"`
	MaxRecords     int    `json:"maxRecords"`
}

func (p *PaginationDetails) Validate() error {
	if p.OffsetKey == "" {
		return fmt.Errorf("offset key cannot be empty")
	}
	if p.SizeKey == "" {
		return fmt.Errorf("size key cannot be empty")
	}
	if p.MaxPageSize <= 0 {
		return fmt.Errorf("max page size must be greater than 0")
	}
	return nil
}

func (P *PaginationDetails) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"offset":         P.OffsetKey,
		"sizeKey":        P.SizeKey,
		"maxPageSize":    P.MaxPageSize,
		"startingOffset": P.StartingOffset,
		"maxRecords":     P.MaxRecords,
	}
}

type APIDetails struct {
	Uri                       string                 `json:"uri"`
	Paginationdeatils         *PaginationDetails     `json:"paginationDetails"`
	ResponseJsonPath          string                 `json:"responseJsonPath"`
	Request                   map[string]interface{} `json:"request"`
	Tenant_id_for_open_search string                 `json:"tenant_id_for_open_search"`
	Custom_query_param        string                 `json:"custom_query_param"`
}

func (a *APIDetails) Validate() error {
	if a.Uri == "" {
		return fmt.Errorf("uri cannot be empty")
	}
	if a.ResponseJsonPath == "" {
		return fmt.Errorf("response json path cannot be empty")
	}
	return nil
}

func (a *APIDetails) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uri":                   a.Uri,
		"paginationDetails":     a.Paginationdeatils.ToMap(),
		"responseJsonPath":      a.ResponseJsonPath,
		"request":               a.Request,
		"tenantIdForOpenSearch": a.Tenant_id_for_open_search,
		"customQueryParam":      a.Custom_query_param,
	}
}

type ReindexRequest struct {
	RequestInfo  *RequestInfo `json:"requestInfo"`
	Index        string       `json:"index"`
	Type         string       `json:"type"`
	ReindexTopic string       `json:"reindexTopic"`
	TenantID     string       `json:"tenantId"`
	BatchSize    int          `json:"batchSize"`
	JobId        string       `json:"jobId"`
	StartTime    int64        `json:"startTime"`
	TotalRecords int          `json:"totalRecords"`
}

func (r *ReindexRequest) Validate() error {

	if r.Index == "" {
		return fmt.Errorf("index cannot be empty")
	}
	if r.Type == "" {
		return fmt.Errorf("type cannot be empty")
	}
	if r.ReindexTopic == "" {
		return fmt.Errorf("reindex topic cannot be empty")
	}
	if r.TenantID == "" {
		return fmt.Errorf("tenant ID cannot be empty")
	}
	return nil
}

func (r *ReindexRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"RequestInfo":  r.RequestInfo.ToMap(),
		"index":        r.Index,
		"type":         r.Type,
		"reindexTopic": r.ReindexTopic,
		"tenantId":     r.TenantID,
		"batchSize":    r.BatchSize,
		"jobId":        r.JobId,
		"startTime":    r.StartTime,
		"totalRecords": r.TotalRecords,
	}
}

type LegacyIndexRequest struct {
	RequestInfo      *RequestInfo `json:"requestInfo"`
	APIDetails       *APIDetails  `json:"apiDetails"`
	LegacyIndexTopic string       `json:"legacyIndexTopic"`
	TenantID         string       `json:"tenantId"`
	JobId            string       `json:"jobId"`
	StartTime        int64        `json:"startTime"`
	TotalRecords     int          `json:"totalRecords"`
}

func (l *LegacyIndexRequest) Validate() error {
	if l.LegacyIndexTopic == "" {
		return fmt.Errorf("legacy index topic cannot be empty")
	}
	if l.TenantID == "" {
		return fmt.Errorf("tenant ID cannot be empty")
	}
	return nil
}

func (l *LegacyIndexRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"RequestInfo":      l.RequestInfo.ToMap(),
		"apiDetails":       l.APIDetails.ToMap(),
		"legacyIndexTopic": l.LegacyIndexTopic,
		"tenantId":         l.TenantID,
		"jobId":            l.JobId,
		"startTime":        l.StartTime,
		"totalRecords":     l.TotalRecords,
	}
}

func PaginationDetailsBuilder() *PaginationDetails {
	return &PaginationDetails{}
}
func APIDetailsBuilder() *APIDetails {
	return &APIDetails{}
}

func ReindexRequestBuilder() *ReindexRequest {
	return &ReindexRequest{}
}
func LegacyIndexRequestBuilder() *LegacyIndexRequest {
	return &LegacyIndexRequest{}
}

func (p *PaginationDetails) WithOffsetKey(offsetKey string) *PaginationDetails {
	p.OffsetKey = offsetKey
	return p
}
func (p *PaginationDetails) WithSizeKey(sizeKey string) *PaginationDetails {
	p.SizeKey = sizeKey
	return p
}
func (p *PaginationDetails) WithMaxPageSize(maxPageSize int) *PaginationDetails {
	p.MaxPageSize = maxPageSize
	return p
}
func (p *PaginationDetails) WithStartingOffset(startingOffset int) *PaginationDetails {
	p.StartingOffset = startingOffset
	return p
}
func (p *PaginationDetails) WithMaxRecords(maxRecords int) *PaginationDetails {
	p.MaxRecords = maxRecords
	return p
}
func (P *PaginationDetails) Build() (*PaginationDetails, error) {
	if err := P.Validate(); err != nil {
		return nil, err
	}
	return P, nil
}
func (a *APIDetails) WithUri(uri string) *APIDetails {
	a.Uri = uri
	return a
}
func (a *APIDetails) WithPaginationDetails(paginationDetails *PaginationDetails) *APIDetails {
	a.Paginationdeatils = paginationDetails
	return a
}
func (a *APIDetails) WithResponseJsonPath(responseJsonPath string) *APIDetails {
	a.ResponseJsonPath = responseJsonPath
	return a
}
func (a *APIDetails) WithRequest(request map[string]interface{}) *APIDetails {
	a.Request = request
	return a
}
func (a *APIDetails) WithTenantIdForOpenSearch(tenantIdForOpenSearch string) *APIDetails {
	a.Tenant_id_for_open_search = tenantIdForOpenSearch
	return a
}
func (a *APIDetails) WithCustomQueryParam(customQueryParam string) *APIDetails {
	a.Custom_query_param = customQueryParam
	return a
}
func (a *APIDetails) Build() (*APIDetails, error) {
	if err := a.Validate(); err != nil {
		return nil, err
	}
	return a, nil
}
func (r *ReindexRequest) WithRequestInfo(requestInfo *RequestInfo) *ReindexRequest {
	r.RequestInfo = requestInfo
	return r
}
func (r *ReindexRequest) WithIndex(index string) *ReindexRequest {
	r.Index = index
	return r
}
func (r *ReindexRequest) WithType(type_ string) *ReindexRequest {
	r.Type = type_
	return r
}
func (r *ReindexRequest) WithReindexTopic(reindexTopic string) *ReindexRequest {
	r.ReindexTopic = reindexTopic
	return r
}
func (r *ReindexRequest) WithTenantId(tenantId string) *ReindexRequest {
	r.TenantID = tenantId
	return r
}
func (r *ReindexRequest) WithBatchSize(batchSize int) *ReindexRequest {
	r.BatchSize = batchSize
	return r
}
func (r *ReindexRequest) WithJobId(jobId string) *ReindexRequest {
	r.JobId = jobId
	return r
}
func (r *ReindexRequest) WithStartTime(startTime int64) *ReindexRequest {
	r.StartTime = startTime
	return r
}
func (r *ReindexRequest) WithTotalRecords(totalRecords int) *ReindexRequest {
	r.TotalRecords = totalRecords
	return r
}
func (r *ReindexRequest) Build() (*ReindexRequest, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return r, nil
}

func (l *LegacyIndexRequest) WithRequestInfo(requestInfo *RequestInfo) *LegacyIndexRequest {
	l.RequestInfo = requestInfo
	return l
}
func (l *LegacyIndexRequest) WithAPIDetails(apiDetails *APIDetails) *LegacyIndexRequest {
	l.APIDetails = apiDetails
	return l
}
func (l *LegacyIndexRequest) WithLegacyIndexTopic(legacyIndexTopic string) *LegacyIndexRequest {
	l.LegacyIndexTopic = legacyIndexTopic
	return l
}
func (l *LegacyIndexRequest) WithTenantId(tenantId string) *LegacyIndexRequest {
	l.TenantID = tenantId
	return l
}
func (l *LegacyIndexRequest) WithJobId(jobId string) *LegacyIndexRequest {
	l.JobId = jobId
	return l
}
func (l *LegacyIndexRequest) WithStartTime(startTime int64) *LegacyIndexRequest {
	l.StartTime = startTime
	return l
}
func (l *LegacyIndexRequest) WithTotalRecords(totalRecords int) *LegacyIndexRequest {
	l.TotalRecords = totalRecords
	return l
}
func (l *LegacyIndexRequest) Build() (*LegacyIndexRequest, error) {
	if err := l.Validate(); err != nil {
		return nil, err
	}
	return l, nil
}
