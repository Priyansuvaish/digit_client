package models

import (
	"time"
	"github.com/google/uuid"
)

// RequestInfo represents the request information for API calls
type RequestInfo struct {
	APIID         string
	Ver           string
	Ts            int64
	Action        string
	Did           string
	Key           string
	MsgID         string
	RequesterID   string
	AuthToken     string
	UserInfo      map[string]interface{}
	CorrelationID string
}

// RequestInfoBuilder creates a new RequestInfo instance
func RequestInfoBuilder() *RequestInfo {
	return &RequestInfo{
		MsgID:         uuid.New().String(),
		CorrelationID: uuid.New().String(),
		Ts:            time.Now().UnixMilli(),
	}
}

// Builder methods for RequestInfo
func (r *RequestInfo) WithAPIID(apiID string) *RequestInfo {
	r.APIID = apiID
	return r
}

func (r *RequestInfo) WithVersion(ver string) *RequestInfo {
	r.Ver = ver
	return r
}

func (r *RequestInfo) WithTimestamp(ts int64) *RequestInfo {
	r.Ts = ts
	return r
}

func (r *RequestInfo) WithAction(action string) *RequestInfo {
	r.Action = action
	return r
}

func (r *RequestInfo) WithDid(did string) *RequestInfo {
	r.Did = did
	return r
}

func (r *RequestInfo) WithKey(key string) *RequestInfo {
	r.Key = key
	return r
}

func (r *RequestInfo) WithMsgID(msgID string) *RequestInfo {
	r.MsgID = msgID
	return r
}

func (r *RequestInfo) WithRequesterID(requesterID string) *RequestInfo {
	r.RequesterID = requesterID
	return r
}

func (r *RequestInfo) WithAuthToken(authToken string) *RequestInfo {
	r.AuthToken = authToken
	return r
}

func (r *RequestInfo) WithUserInfo(userInfo map[string]interface{}) *RequestInfo {
	r.UserInfo = userInfo
	return r
}

func (r *RequestInfo) WithCorrelationID(correlationID string) *RequestInfo {
	r.CorrelationID = correlationID
	return r
}

// WithAuthToken creates a new RequestInfo instance with a temporary auth token
func (r *RequestInfo) WithTempAuthToken(tempAuthToken string) *RequestInfo {
	return RequestInfoBuilder().
		WithAPIID(r.APIID).
		WithVersion(r.Ver).
		WithTimestamp(time.Now().UnixMilli()).
		WithAction(r.Action).
		WithDid(r.Did).
		WithKey(r.Key).
		WithMsgID(uuid.New().String()).
		WithRequesterID(r.RequesterID).
		WithAuthToken(tempAuthToken).
		WithUserInfo(r.UserInfo).
		WithCorrelationID(uuid.New().String())
}

// ToMap converts the RequestInfo to a map
func (r *RequestInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"apiId":        r.APIID,
		"ver":          r.Ver,
		"ts":           r.Ts,
		"action":       r.Action,
		"did":          r.Did,
		"key":          r.Key,
		"msgId":        r.MsgID,
		"requesterId":  r.RequesterID,
		"authToken":    r.AuthToken,
		"userInfo":     r.UserInfo,
		"correlationId": r.CorrelationID,
	}
}

// RequestConfig manages the default request configuration
type RequestConfig struct {
	defaultRequestInfo *RequestInfo
}

var requestConfigInstance *RequestConfig

// GetInstance returns the singleton instance of RequestConfig
func (rc *RequestConfig) GetInstance() *RequestConfig {
	if requestConfigInstance == nil {
		requestConfigInstance = &RequestConfig{}
	}
	return requestConfigInstance
}

// Initialize sets up the default request configuration
func (rc *RequestConfig) Initialize(
	apiID string,
	version string,
	authToken string,
	userInfo map[string]interface{},
	did string,
	key string,
	msgID string,
	requesterID string,
	correlationID string,
	action string,
	ts int64,
) {
	rc.defaultRequestInfo = RequestInfoBuilder().
		WithAPIID(apiID).
		WithVersion(version).
		WithTimestamp(ts).
		WithAction(action).
		WithAuthToken(authToken).
		WithUserInfo(userInfo).
		WithDid(did).
		WithKey(key).
		WithMsgID(msgID).
		WithRequesterID(requesterID).
		WithCorrelationID(correlationID)
}

// UpdateAPIID updates the API ID in the existing configuration
func (rc *RequestConfig) UpdateAPIID(apiID string) {
	if rc.defaultRequestInfo == nil {
		panic("RequestConfig not initialized. Call RequestConfig.Initialize() first.")
	}

	rc.defaultRequestInfo = rc.defaultRequestInfo.WithAPIID(apiID).
		WithTimestamp(time.Now().UnixMilli()).
		WithMsgID(uuid.New().String()).
		WithCorrelationID(uuid.New().String())
}

// UpdateVersion updates the version in the existing configuration
func (rc *RequestConfig) UpdateVersion(version string) {
	if rc.defaultRequestInfo == nil {
		panic("RequestConfig not initialized. Call RequestConfig.Initialize() first.")
	}

	rc.defaultRequestInfo = rc.defaultRequestInfo.WithVersion(version).
		WithTimestamp(time.Now().UnixMilli()).
		WithMsgID(uuid.New().String()).
		WithCorrelationID(uuid.New().String())
}

// UpdateAuthToken updates the auth token in the existing configuration
func (rc *RequestConfig) UpdateAuthToken(authToken string) {
	if rc.defaultRequestInfo == nil {
		panic("RequestConfig not initialized. Call RequestConfig.Initialize() first.")
	}

	rc.defaultRequestInfo = rc.defaultRequestInfo.WithAuthToken(authToken).
		WithTimestamp(time.Now().UnixMilli()).
		WithMsgID(uuid.New().String()).
		WithCorrelationID(uuid.New().String())
}

// UpdateUserInfo updates the user info in the existing configuration
func (rc *RequestConfig) UpdateUserInfo(userInfo map[string]interface{}) {
	if rc.defaultRequestInfo == nil {
		panic("RequestConfig not initialized. Call RequestConfig.Initialize() first.")
	}

	rc.defaultRequestInfo = rc.defaultRequestInfo.WithUserInfo(userInfo).
		WithTimestamp(time.Now().UnixMilli()).
		WithMsgID(uuid.New().String()).
		WithCorrelationID(uuid.New().String())
}

// UpdateDid updates the device ID in the existing configuration
func (rc *RequestConfig) UpdateDid(did string) {
	if rc.defaultRequestInfo == nil {
		panic("RequestConfig not initialized. Call RequestConfig.Initialize() first.")
	}

	rc.defaultRequestInfo = rc.defaultRequestInfo.WithDid(did).
		WithTimestamp(time.Now().UnixMilli()).
		WithMsgID(uuid.New().String()).
		WithCorrelationID(uuid.New().String())
}

// UpdateKey updates the key in the existing configuration
func (rc *RequestConfig) UpdateKey(key string) {
	if rc.defaultRequestInfo == nil {
		panic("RequestConfig not initialized. Call RequestConfig.Initialize() first.")
	}

	rc.defaultRequestInfo = rc.defaultRequestInfo.WithKey(key).
		WithTimestamp(time.Now().UnixMilli()).
		WithMsgID(uuid.New().String()).
		WithCorrelationID(uuid.New().String())
}

// UpdateRequesterID updates the requester ID in the existing configuration
func (rc *RequestConfig) UpdateRequesterID(requesterID string) {
	if rc.defaultRequestInfo == nil {
		panic("RequestConfig not initialized. Call RequestConfig.Initialize() first.")
	}

	rc.defaultRequestInfo = rc.defaultRequestInfo.WithRequesterID(requesterID).
		WithTimestamp(time.Now().UnixMilli()).
		WithMsgID(uuid.New().String()).
		WithCorrelationID(uuid.New().String())
}

// GetRequestInfo returns a new RequestInfo instance with default values and specified overrides
func (rc *RequestConfig) GetRequestInfo(tempAuthToken string, overrides map[string]interface{}) *RequestInfo {
	if rc.defaultRequestInfo == nil {
		panic("RequestConfig not initialized. Call RequestConfig.Initialize() first.")
	}

	requestInfo := rc.defaultRequestInfo.WithTimestamp(time.Now().UnixMilli())

	if tempAuthToken != "" {
		requestInfo = requestInfo.WithAuthToken(tempAuthToken)
	}

	// Apply overrides
	for key, value := range overrides {
		switch key {
		case "apiId":
			requestInfo = requestInfo.WithAPIID(value.(string))
		case "ver":
			requestInfo = requestInfo.WithVersion(value.(string))
		case "action":
			requestInfo = requestInfo.WithAction(value.(string))
		case "did":
			requestInfo = requestInfo.WithDid(value.(string))
		case "key":
			requestInfo = requestInfo.WithKey(value.(string))
		case "msgId":
			requestInfo = requestInfo.WithMsgID(value.(string))
		case "requesterId":
			requestInfo = requestInfo.WithRequesterID(value.(string))
		case "userInfo":
			requestInfo = requestInfo.WithUserInfo(value.(map[string]interface{}))
		case "correlationId":
			requestInfo = requestInfo.WithCorrelationID(value.(string))
		}
	}

	return requestInfo
} 