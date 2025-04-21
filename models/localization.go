package models

import (
	"fmt"
)

type LocaleRequest struct {
	Locale   string   `json:"locale"`
	TenantID string   `json:"tenant_id"`
	Module   string   `json:"module"`
	Codes    []string `json:"codes"`
}

func (l *LocaleRequest) Validate() error {
	if l.Locale == "" {
		return fmt.Errorf("locale is required")
	}
	if l.TenantID == "" {
		return fmt.Errorf("tenant ID is required")
	}

	return nil
}

func (l *LocaleRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"locale":    l.Locale,
		"tenant_id": l.TenantID,
		"module":    l.Module,
		"codes":     l.Codes,
	}
}

type Message struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Locale  string `json:"locale"`
	Module  string `json:"module"`
}

func (m *Message) Validate() error {
	if m.Code == "" {
		return fmt.Errorf("code is required")
	}
	if m.Message == "" {
		return fmt.Errorf("message is required")
	}
	if m.Locale == "" {
		return fmt.Errorf("locale is required")
	}
	if m.Module == "" {
		return fmt.Errorf("module is required")
	}

	return nil
}

func (m *Message) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    m.Code,
		"message": m.Message,
		"locale":  m.Locale,
		"module":  m.Module,
	}
}

type CreateMessagesRequest struct {
	TenantID    string       `json:"tenant_id"`
	Messages    []Message    `json:"messages"`
	RequestInfo *RequestInfo `json:"request_info"`
}

func (c *CreateMessagesRequest) Validate() error {
	if c.TenantID == "" {
		return fmt.Errorf("tenant ID is required")
	}
	if len(c.Messages) < 1 {
		return fmt.Errorf("at least one message is required")
	}
	return nil
}

func (c *CreateMessagesRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":    c.TenantID,
		"messages":    c.Messages,
		"RequestInfo": c.RequestInfo.ToMap(),
	}
}

type UpdateMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (u *UpdateMessage) Validate() error {
	if u.Code == "" {
		return fmt.Errorf("code is required")
	}
	if u.Message == "" {
		return fmt.Errorf("message is required")
	}
	return nil
}

func (u *UpdateMessage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    u.Code,
		"message": u.Message,
	}
}

type UpdateMessagesRequest struct {
	TenantID    string          `json:"tenant_id"`
	Messages    []UpdateMessage `json:"messages"`
	RequestInfo *RequestInfo    `json:"request_info"`
	Locale      string          `json:"locale"`
	Module      string          `json:"module"`
}

func (u *UpdateMessagesRequest) Validate() error {
	if u.TenantID == "" || len(u.TenantID) > 256 {
		return fmt.Errorf("tenant ID is required")
	}
	if len(u.Messages) < 1 {
		return fmt.Errorf("at least one message is required")
	}
	if u.Locale == "" || len(u.Locale) > 255 {
		return fmt.Errorf("locale is required")
	}
	if u.Module == "" || len(u.Module) > 255 {
		return fmt.Errorf("module is required")
	}
	return nil
}
func (u *UpdateMessagesRequest) ToMap() map[string]interface{} {
	result := make([]map[string]interface{}, len(u.Messages))
	for i, msg := range u.Messages {
		result[i] = msg.ToMap()
	}

	return map[string]interface{}{
		"tenantId":    u.TenantID,
		"RequestInfo": u.RequestInfo.ToMap(),
		"locale":      u.Locale,
		"module":      u.Module,
		"messages":    result,
	}
}

type DeleteMessage struct {
	Code   string `json:"code"`
	Locale string `json:"locale"`
	Module string `json:"module"`
}

func (d *DeleteMessage) Validate() error {
	if d.Code == "" {
		return fmt.Errorf("code is required")
	}
	if d.Locale == "" {
		return fmt.Errorf("locale is required")
	}
	if d.Module == "" {
		return fmt.Errorf("module is required")
	}
	return nil
}
func (d *DeleteMessage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":   d.Code,
		"locale": d.Locale,
		"module": d.Module,
	}
}

type DeleteMessagesRequest struct {
	TenantID    string          `json:"tenant_id"`
	Messages    []DeleteMessage `json:"messages"`
	RequestInfo *RequestInfo    `json:"request_info"`
}

func (d *DeleteMessagesRequest) Validate() error {
	if d.TenantID == "" {
		return fmt.Errorf("tenant ID is required")
	}
	if len(d.Messages) < 1 {
		return fmt.Errorf("at least one message is required")
	}
	return nil
}
func (d *DeleteMessagesRequest) ToMap() map[string]interface{} {
	result := make([]map[string]interface{}, len(d.Messages))
	for i, msg := range d.Messages {
		result[i] = msg.ToMap()
	}

	return map[string]interface{}{
		"tenantId":    d.TenantID,
		"RequestInfo": d.RequestInfo.ToMap(),
		"messages":    result,
	}
}

func LocaleRequestBuilder() *LocaleRequest {
	return &LocaleRequest{}
}
func MessageBuilder() *Message {
	return &Message{}
}
func CreateMessagesRequestBuilder() *CreateMessagesRequest {
	return &CreateMessagesRequest{}
}
func UpdateMessageBuilder() *UpdateMessage {
	return &UpdateMessage{}
}
func UpdateMessagesRequestBuilder() *UpdateMessagesRequest {
	return &UpdateMessagesRequest{}
}
func DeleteMessageBuilder() *DeleteMessage {
	return &DeleteMessage{}
}
func DeleteMessagesRequestBuilder() *DeleteMessagesRequest {
	return &DeleteMessagesRequest{}
}

func (l *LocaleRequest) WithLocale(locale string) *LocaleRequest {
	l.Locale = locale
	return l
}
func (l *LocaleRequest) WithTenantID(tenantID string) *LocaleRequest {
	l.TenantID = tenantID
	return l
}
func (l *LocaleRequest) WithModule(module string) *LocaleRequest {
	l.Module = module
	return l
}
func (l *LocaleRequest) WithCodes(codes []string) *LocaleRequest {
	l.Codes = codes
	return l
}

func (l *LocaleRequest) Build() (*LocaleRequest, error) {
	err := l.Validate()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (m *Message) WithCode(code string) *Message {
	m.Code = code
	return m
}
func (m *Message) WithMessage(message string) *Message {
	m.Message = message
	return m
}
func (m *Message) WithLocale(locale string) *Message {
	m.Locale = locale
	return m
}
func (m *Message) WithModule(module string) *Message {
	m.Module = module
	return m
}
func (m *Message) Build() (*Message, error) {
	err := m.Validate()
	if err != nil {
		return nil, err
	}
	return m, nil
}
func (c *CreateMessagesRequest) WithTenantID(tenantID string) *CreateMessagesRequest {
	c.TenantID = tenantID
	return c
}
func (c *CreateMessagesRequest) WithMessages(messages []Message) *CreateMessagesRequest {
	c.Messages = messages
	return c
}
func (c *CreateMessagesRequest) WithRequestInfo(requestInfo *RequestInfo) *CreateMessagesRequest {
	c.RequestInfo = requestInfo
	return c
}
func (c *CreateMessagesRequest) Build() (*CreateMessagesRequest, error) {
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (u *UpdateMessage) WithCode(code string) *UpdateMessage {
	u.Code = code
	return u
}
func (u *UpdateMessage) WithMessage(message string) *UpdateMessage {
	u.Message = message
	return u
}
func (u *UpdateMessage) Build() (*UpdateMessage, error) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UpdateMessagesRequest) WithTenantID(tenantID string) *UpdateMessagesRequest {
	u.TenantID = tenantID
	return u
}
func (u *UpdateMessagesRequest) WithMessages(messages []UpdateMessage) *UpdateMessagesRequest {
	u.Messages = messages
	return u
}
func (u *UpdateMessagesRequest) WithRequestInfo(requestInfo *RequestInfo) *UpdateMessagesRequest {
	u.RequestInfo = requestInfo
	return u
}
func (u *UpdateMessagesRequest) WithLocale(locale string) *UpdateMessagesRequest {
	u.Locale = locale
	return u
}
func (u *UpdateMessagesRequest) WithModule(module string) *UpdateMessagesRequest {
	u.Module = module
	return u
}
func (u *UpdateMessagesRequest) Build() (*UpdateMessagesRequest, error) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (d *DeleteMessage) WithCode(code string) *DeleteMessage {
	d.Code = code
	return d
}
func (d *DeleteMessage) WithLocale(locale string) *DeleteMessage {

	d.Locale = locale
	return d
}
func (d *DeleteMessage) WithModule(module string) *DeleteMessage {
	d.Module = module
	return d
}
func (d *DeleteMessage) Build() (*DeleteMessage, error) {
	err := d.Validate()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DeleteMessagesRequest) WithTenantID(tenantID string) *DeleteMessagesRequest {
	d.TenantID = tenantID
	return d
}
func (d *DeleteMessagesRequest) WithMessages(messages []DeleteMessage) *DeleteMessagesRequest {
	d.Messages = messages
	return d
}
func (d *DeleteMessagesRequest) WithRequestInfo(requestInfo *RequestInfo) *DeleteMessagesRequest {
	d.RequestInfo = requestInfo
	return d
}
func (d *DeleteMessagesRequest) Build() (*DeleteMessagesRequest, error) {
	err := d.Validate()
	if err != nil {
		return nil, err
	}
	return d, nil
}
