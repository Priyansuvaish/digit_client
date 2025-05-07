package models

import (
	"fmt"
)

type MasterDetail struct {
	Name   string   `json:"name"`
	Filter []string `json:"filter"`
}

func (md *MasterDetail) Validate() error {
	if len(md.Name) > 100 {
		return fmt.Errorf("name is too long, max length is 100 characters")
	}
	if len(md.Filter) > 500 {
		return fmt.Errorf("filter is too long, max length is 500 characters")
	}
	return nil
}

func (md *MasterDetail) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":   md.Name,
		"filter": md.Filter,
	}
}

type ModuleDetail struct {
	ModuleName    string         `json:"module_name"`
	MasterDetails []MasterDetail `json:"master_details"`
}

func (md *ModuleDetail) Validate() error {
	if len(md.ModuleName) > 100 {
		return fmt.Errorf("module_name is too long, max length is 100 characters")
	}
	return nil
}

func (md *ModuleDetail) ToMap() map[string]interface{} {
	result := make([]map[string]interface{}, len(md.MasterDetails))
	for i, masterDetail := range md.MasterDetails {
		result[i] = masterDetail.ToMap()
	}
	return map[string]interface{}{
		"moduleName":    md.ModuleName,
		"masterDetails": result,
	}
}

type MdmsCriteria struct {
	TenantId             string            `json:"tenantId"`
	ModuleDetails        []ModuleDetail    `json:"moduleDetails"`
	Ids                  []string          `json:"ids"`
	Uniqueidentifer      string            `json:"uniqueIdentifier"`
	SchemaCodeFilter1map map[string]string `json:"schemaCodeFilter1map"`
	IsActive             bool              `json:"isActive"`
}

func (md *MdmsCriteria) Validate() error {
	if len(md.TenantId) > 100 {
		return fmt.Errorf("tenantId is too long, max length is 100 characters")
	}

	if len(md.Uniqueidentifer) > 64 {
		return fmt.Errorf("uniqueIdentifier is too long, max length is 100 characters")
	}
	return nil
}

func (md *MdmsCriteria) ToMap() map[string]interface{} {
	payload := map[string]interface{}{
		"tenantId":      md.TenantId,
		"moduleDetails": make([]interface{}, 0, len(md.ModuleDetails)),
	}

	// Add module details (always include even if empty)
	for _, moduleDetail := range md.ModuleDetails {
		payload["moduleDetails"] = append(payload["moduleDetails"].([]interface{}), moduleDetail.ToMap())
	}

	// Conditionally add other fields ONLY if they have values
	if len(md.Ids) > 0 {
		payload["ids"] = md.Ids
	}
	if md.Uniqueidentifer != "" {
		payload["uniqueIdentifier"] = md.Uniqueidentifer
	}
	if len(md.SchemaCodeFilter1map) > 0 {
		payload["schemaCodeFilter1map"] = md.SchemaCodeFilter1map
	}
	if md.IsActive { // Only include if true
		payload["isActive"] = md.IsActive
	}

	return payload
}

func MasterDetailBuilder() *MasterDetail {
	return &MasterDetail{}
}

func ModuleDetailBuilder() *ModuleDetail {
	return &ModuleDetail{}
}

func MdmsCriteriaBuilder() *MdmsCriteria {
	return &MdmsCriteria{}
}
func (md *MasterDetail) WithName(name string) *MasterDetail {
	md.Name = name
	return md
}
func (md *MasterDetail) WithFilter(filter []string) *MasterDetail {
	md.Filter = filter
	return md
}
func (md *MasterDetail) Build() (*MasterDetail, error) {
	if err := md.Validate(); err != nil {
		return nil, err
	}
	return md, nil
}
func (md *ModuleDetail) WithModuleName(moduleName string) *ModuleDetail {
	md.ModuleName = moduleName
	return md
}
func (md *ModuleDetail) WithMasterDetails(masterDetails []MasterDetail) *ModuleDetail {
	md.MasterDetails = masterDetails
	return md
}
func (md *ModuleDetail) Build() (*ModuleDetail, error) {
	if err := md.Validate(); err != nil {
		return nil, err
	}
	return md, nil
}

func (md *MdmsCriteria) WithTenantId(tenantId string) *MdmsCriteria {
	md.TenantId = tenantId
	return md
}
func (md *MdmsCriteria) WithModuleDetails(moduleDetails []ModuleDetail) *MdmsCriteria {
	md.ModuleDetails = moduleDetails
	return md
}
func (md *MdmsCriteria) WithIds(ids []string) *MdmsCriteria {
	md.Ids = ids
	return md
}

func (md *MdmsCriteria) WithUniqueidentifer(uniqueidentifer string) *MdmsCriteria {
	md.Uniqueidentifer = uniqueidentifer
	return md
}
func (md *MdmsCriteria) WithSchemaCodeFilter1map(schemaCodeFilter1map map[string]string) *MdmsCriteria {
	md.SchemaCodeFilter1map = schemaCodeFilter1map
	return md
}
func (md *MdmsCriteria) WithIsActive(isActive bool) *MdmsCriteria {
	md.IsActive = isActive
	return md
}
func (md *MdmsCriteria) Build() (*MdmsCriteria, error) {
	if err := md.Validate(); err != nil {
		return nil, err
	}
	return md, nil
}
