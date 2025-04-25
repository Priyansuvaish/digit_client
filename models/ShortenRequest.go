package models

import (
	"fmt"
)

type ShortenRequest struct {
	Id        string `json:"id"`
	URL       string `json:"url"`
	ValidFrom int    `json:"validfrom"`
	ValidTill int    `json:"validtill"`
}

func (s *ShortenRequest) Validate() error {
	if s.URL == "" {
		return fmt.Errorf("url is required")
	} else if len(s.URL) > 2000 {
		return fmt.Errorf("url must not exceed 2000 characters")
	}
	return nil
}
func (s *ShortenRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":        s.Id,
		"url":       s.URL,
		"validFrom": s.ValidFrom,
		"validTill": s.ValidTill,
	}
}

func ShortenRequestBuilder() *ShortenRequest {
	return &ShortenRequest{}
}

func (s *ShortenRequest) WithId(code string) *ShortenRequest {
	s.Id = code
	return s
}

func (s *ShortenRequest) WithURL(code string) *ShortenRequest {
	s.URL = code
	return s
}

func (s *ShortenRequest) WithValidFrom(code int) *ShortenRequest {
	s.ValidFrom = code
	return s
}

func (s *ShortenRequest) WithValidtill(code int) *ShortenRequest {
	s.ValidTill = code
	return s
}

func (s *ShortenRequest) Build() (*ShortenRequest, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return s, err
}
