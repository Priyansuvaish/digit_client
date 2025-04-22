package models

import (
	"fmt"
)

type MetadataRequest struct {
	RequestInfo *RequestInfo `json:"requestinfo"`
	TenantID    string       `json:"teantid"`
	ReportName  string       `json:"reportname"`
}

func (m *MetadataRequest) Validate() error {
	if m.TenantID == "" {
		return fmt.Errorf("tenant_id is required")
	}
	if m.ReportName == "" {
		return fmt.Errorf("report_name is required")
	}
	return nil

}

func (m *MetadataRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"requestinfo": m.RequestInfo,
		"tenantId":    m.TenantID,
		"reportName":  m.ReportName,
	}
}

type SearchParam struct {
	Input any `json:"input"`
}

func (s *SearchParam) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"input": s.Input,
	}
}

type ReportRequest struct {
	SearchParam []SearchParam `json:"searchParams"`
}

func (r *ReportRequest) ToMap() map[string]interface{} {
	result := make([]map[string]interface{}, len(r.SearchParam))
	for i, r := range r.SearchParam {
		result[i] = r.ToMap()
	}
	return map[string]interface{}{
		"searchParams": result,
	}
}

func MetadataRequestBuilder() *MetadataRequest {
	return &MetadataRequest{}
}
func SearchParamBuilder() *SearchParam {
	return &SearchParam{}
}
func ReportRequestBuilder() *ReportRequest {
	return &ReportRequest{}
}

func (m *MetadataRequest) WithRequestInfo(RequestInfo *RequestInfo) *MetadataRequest {
	m.RequestInfo = RequestInfo
	return m
}
func (m *MetadataRequest) WithTenantID(teantid string) *MetadataRequest {
	m.TenantID = teantid
	return m
}
func (m *MetadataRequest) WithReportName(reportname string) *MetadataRequest {
	m.ReportName = reportname
	return m
}
func (m *MetadataRequest) Build() (*MetadataRequest, error) {
	err := m.Validate()
	if err != nil {
		return nil, err
	}
	return m, nil
}
func (s *SearchParam) WithInput(input any) *SearchParam {
	s.Input = input
	return s
}
func (s *SearchParam) Build() *SearchParam {
	return s
}

func (r *ReportRequest) WithSearchParam(searchParam []SearchParam) *ReportRequest {
	r.SearchParam = searchParam
	return r
}

func (r *ReportRequest) Build() *ReportRequest {
	return r
}
