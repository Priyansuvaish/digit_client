package models

import (
	"fmt"
	"strconv"
	"strings"
)

type EncReqObject struct {
	TenantID string `json:"tenant_id"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

func (e *EncReqObject) Validate() error {
	if e.TenantID == "" {
		return fmt.Errorf("tenant ID is required")
	}
	if e.Type == "" {
		return fmt.Errorf("type is required")
	}
	if e.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

func (e *EncReqObject) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId": e.TenantID,
		"type":     e.Type,
		"value":    e.Value,
	}
}

type SignRequest struct {
	TenantID string `json:"tenant_id"`
	Value    string `json:"value"`
}

func (s *SignRequest) Validate() error {
	if s.TenantID == "" {
		return fmt.Errorf("tenant ID is required")
	}
	if s.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

func (s *SignRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId": s.TenantID,
		"value":    s.Value,
	}
}

type Signature struct {
	KeyID          int    `json:"key_id"`
	SignatureValue string `json:"value"`
}

// NewSignature parses the input string and returns a Signature struct
func NewSignature(signatureValue string) (*Signature, error) {
	parts := strings.Split(signatureValue, "|")
	if len(parts) != 2 {
		return nil, fmt.Errorf("%s: Invalid Signature", signatureValue)
	}
	keyID, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("%s: Invalid Signature", signatureValue)
	}
	return &Signature{
		KeyID:          keyID,
		SignatureValue: parts[1],
	}, nil
}

// String returns the string representation of the Signature
func (s *Signature) String() string {
	return fmt.Sprintf("%d|%s", s.KeyID, s.SignatureValue)
}

// ToMap returns a map representation of the Signature
func (s *Signature) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"keyId":          s.KeyID,
		"signatureValue": s.SignatureValue,
	}
}

type VerifyRequest struct {
	Value     string     `json:"value"`
	Signature *Signature `json:"signature"`
}

func (v *VerifyRequest) Validate() error {
	if v.Value == "" {
		return fmt.Errorf("value is required")
	}
	if v.Signature == nil {
		return fmt.Errorf("signature is required")
	}
	return nil
}

func (v *VerifyRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"value":     v.Value,
		"signature": v.Signature.ToMap(),
	}
}

type RotateKeyRequest struct {
	TenantID string `json:"tenant_id"`
}

func (r *RotateKeyRequest) Validate() error {
	if r.TenantID == "" {
		return fmt.Errorf("tenant ID is required")
	}
	return nil
}

func (r *RotateKeyRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId": r.TenantID,
	}
}

func EncReqObjectBuilder() *EncReqObject {
	return &EncReqObject{}
}
func SignRequestBuilder() *SignRequest {
	return &SignRequest{}
}
func VerifyRequestBuilder() *VerifyRequest {
	return &VerifyRequest{}
}
func RotateKeyRequestBuilder() *RotateKeyRequest {
	return &RotateKeyRequest{}
}

func (e *EncReqObject) SetTenantID(tenantID string) *EncReqObject {
	e.TenantID = tenantID
	return e
}
func (e *EncReqObject) SetType(typeValue string) *EncReqObject {
	e.Type = typeValue
	return e
}
func (e *EncReqObject) SetValue(value string) *EncReqObject {
	e.Value = value
	return e
}
func (e *EncReqObject) Build() (*EncReqObject, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *SignRequest) SetTenantID(tenantID string) *SignRequest {
	s.TenantID = tenantID
	return s
}
func (s *SignRequest) SetValue(value string) *SignRequest {
	s.Value = value
	return s
}
func (s *SignRequest) Build() (*SignRequest, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}

func (v *VerifyRequest) SetValue(value string) *VerifyRequest {
	v.Value = value
	return v
}

func (v *VerifyRequest) SetSignature(signature *Signature) *VerifyRequest {
	v.Signature = signature
	return v
}

func (v *VerifyRequest) Build() (*VerifyRequest, error) {
	if err := v.Validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func (r *RotateKeyRequest) SetTenantID(tenantID string) *RotateKeyRequest {
	r.TenantID = tenantID
	return r
}
func (r *RotateKeyRequest) Build() (*RotateKeyRequest, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return r, nil
}
