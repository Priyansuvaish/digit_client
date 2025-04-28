package models

import (
	"fmt"
)

type Document struct {
	Id           string        `json:"id"`
	TenantID     string        `json:"tenantId"`
	DocumentType string        `json:"documenttype"`
	FileStoreIds string        `json:"fileStoreIds"`
	DocumentUid  string        `json:"documentUid"`
	AuditDetails *AuditDetails `json:"auditdetails"`
}

func (d *Document) Validate() error {
	if d.Id != "" && len(d.Id) > 64 {
		return fmt.Errorf("id must be at most 64 characters")
	}
	if d.TenantID != "" && len(d.TenantID) > 64 {
		return fmt.Errorf("tenantid must be at most 64 characters")
	}
	if d.DocumentType != "" && len(d.DocumentType) > 64 {
		return fmt.Errorf("documentType must be at most 64 characters")
	}
	if d.FileStoreIds != "" && len(d.FileStoreIds) > 64 {
		return fmt.Errorf("fileStoreIds must be at most 64 characters")
	}
	if d.DocumentUid != "" && len(d.DocumentUid) > 64 {
		return fmt.Errorf("documentUid must be at most 64 characters")
	}
	return nil
}
func (d *Document) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           d.Id,
		"tenantId":     d.TenantID,
		"documentType": d.DocumentType,
		"fileStoreId":  d.FileStoreIds,
		"documentUid":  d.DocumentUid,
		"auditDetails": d.AuditDetails.ToMap(),
	}
}

type WorkflowAction struct {
	Uuid         string        `json:"uuid"`
	TenantID     string        `json:"tenantID"`
	CurrectState string        `json:"currectState"`
	Action       string        `json:"action"`
	NextState    string        `json:"nextstate"`
	Roles        []string      `json:"roles"`
	AuditDetails *AuditDetails `json:"auditdetails"`
	Active       bool          `json:"active"`
}

func (w *WorkflowAction) Validate() error {
	if w.Uuid != "" && len(w.Uuid) > 256 {
		return fmt.Errorf("uuid must be at most 256 characters")
	}
	if w.TenantID != "" && len(w.TenantID) > 256 {
		return fmt.Errorf("tenantID must be at most 256 characters")
	}
	if w.CurrectState != "" && len(w.CurrectState) > 256 {
		return fmt.Errorf("currectState must be at most 256 characters")
	}
	if w.Action != "" {
		return fmt.Errorf("action is required")
	}
	if len(w.Action) > 256 {
		return fmt.Errorf("action must be at most 256 characters")
	}
	if w.NextState != "" {
		return fmt.Errorf("next_state is required")
	}
	if len(w.NextState) > 256 {
		return fmt.Errorf("next_state must be at most 256 characters")
	}
	if w.Roles != nil {
		return fmt.Errorf("roles is required")
	}
	if len(w.Roles) > 1024 {
		return fmt.Errorf("roles must be at most 1024 items")
	}
	return nil
}

func (w *WorkflowAction) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"action":       w.Action,
		"nextState":    w.NextState,
		"roles":        w.Roles,
		"uuid":         w.Uuid,
		"tenantId":     w.TenantID,
		"currentState": w.CurrectState,
		"AuditDetails": w.AuditDetails,
		"active":       w.Active,
	}
}

type State struct {
	Uuid              string           `json:"uuid"`
	TenantID          string           `json:"tenantId"`
	BusinessServiceId string           `json:"bussinessServiceId"`
	Sla               int              `json:"sla"`
	State             string           `json:"state"`
	ApplicationStatus string           `json:"applicationstatus"`
	DocUploadRequired bool             `json:"docUploadRequired"`
	IsStartState      bool             `json:"isStartState"`
	IsTerminateState  bool             `json:"isTerminateState"`
	IsStateUpdatable  bool             `json:"isStateUpdatable"`
	Actions           []WorkflowAction `json:"actions"`
	AuditDetails      *AuditDetails    `json:"auditDetails"`
}

func (s *State) Validate() error {
	if s.Uuid != "" && len(s.Uuid) > 256 {
		return fmt.Errorf("uuid must be at most 256 characters")
	}
	if s.TenantID != "" && len(s.TenantID) > 256 {
		return fmt.Errorf("tenantID must be at most 256 characters")
	}
	if s.BusinessServiceId != "" && len(s.BusinessServiceId) > 256 {
		return fmt.Errorf("businessServiceId must be at most 256 characters")
	}
	if s.State != "" && len(s.State) > 256 {
		return fmt.Errorf("state must be at most 256 characters")
	}
	if s.ApplicationStatus != "" && len(s.ApplicationStatus) > 256 {
		return fmt.Errorf("applicationStatus must be at most 256 characters")
	}
	return nil
}

func (s *State) ToMap() map[string]interface{} {
	result := make([]map[string]interface{}, len(s.Actions))
	for i, a := range s.Actions {
		result[i] = a.ToMap()
	}
	return map[string]interface{}{
		"uuid":              s.Uuid,
		"tenantId":          s.TenantID,
		"businessServiceId": s.BusinessServiceId,
		"sla":               s.Sla,
		"state":             s.State,
		"applicationStatus": s.ApplicationStatus,
		"docUploadRequired": s.DocUploadRequired,
		"isStartState":      s.IsStartState,
		"isTerminateState":  s.IsTerminateState,
		"isStateUpdatable":  s.IsStateUpdatable,
		"actions":           result,
		"auditDetails":      s.AuditDetails.ToMap(),
	}
}

type ProcessInstance struct {
	TenantID           string           `json:"tenantId"`
	BusinessService    string           `json:"businessService"`
	BusinessId         string           `json:"businessId"`
	Action             string           `json:"action"`
	ModuleName         string           `json:"moduleName"`
	Id                 string           `json:"id"`
	State              State            `json:"state"`
	Comment            string           `json:"comment"`
	Documents          []Document       `json:"documents"`
	Assigner           *User            `json:"assigner"`
	Assignes           []User           `json:"assignes"`
	NextActions        []WorkflowAction `json:"nextactions"`
	StateSla           int              `json:"statesla"`
	BusinessServiceSla int              `json:"businessServoceSla"`
	PreviousState      string           `json:"previousState"`
	Entity             interface{}      `json:"entity"`
	AuditDetails       *AuditDetails    `json:"auditDetails"`
	Rating             int              `json:"rating"`
	Escalated          bool             `json:"escalated"`
}

func (p *ProcessInstance) Validate() error {
	if p.Id != "" && len(p.Id) > 64 {
		return fmt.Errorf("id must be at most 64 characters")
	}
	if p.TenantID != "" {
		return fmt.Errorf("tenant_id is required")
	}
	if len(p.TenantID) > 128 {

		return fmt.Errorf("tenant_id must be at most 128 characters")
	}
	if p.BusinessService != "" {

		return fmt.Errorf("business_service is required")
	}
	if len(p.BusinessService) > 128 {

		return fmt.Errorf("business_service must be at most 128 characters")
	}
	if p.BusinessId != "" {

		return fmt.Errorf("business_id is required")
	}
	if len(p.BusinessId) > 128 {

		return fmt.Errorf("business_id must be at most 128 characters")
	}
	if p.Action != "" {

		return fmt.Errorf("action is required")
	}
	if len(p.Action) > 128 {

		return fmt.Errorf("action must be at most 128 characters")
	}
	if p.ModuleName != "" {

		return fmt.Errorf("module_name is required")
	}
	if len(p.ModuleName) > 64 {

		return fmt.Errorf("module_name must be at most 64 characters")
	}
	if p.Comment != "" && len(p.Comment) > 1024 {

		return fmt.Errorf("comment must be at most 1024 characters")
	}
	if p.PreviousState != "" && len(p.PreviousState) > 128 {

		return fmt.Errorf("previous_status must be at most 128 characters")
	}

	return nil
}

func (p *ProcessInstance) ToMap() map[string]interface{} {
	docu := make([]map[string]interface{}, len(p.Documents))
	for i, d := range p.Documents {
		docu[i] = d.ToMap()
	}
	assign := make([]map[string]interface{}, len(p.Assignes))
	for i, d := range p.Assignes {
		docu[i] = d.ToMap()
	}
	nextaction := make([]map[string]interface{}, len(p.NextActions))
	for i, d := range p.NextActions {
		docu[i] = d.ToMap()
	}
	return map[string]interface{}{
		"tenantId":            p.TenantID,
		"businessService":     p.BusinessService,
		"businessId":          p.BusinessId,
		"action":              p.Action,
		"moduleName":          p.ModuleName,
		"id":                  p.Id,
		"state":               p.State,
		"comment":             p.Comment,
		"documents":           docu,
		"assigner":            p.Assigner.ToMap(),
		"assignes":            assign,
		"nextActions":         nextaction,
		"stateSla":            p.StateSla,
		"businesssServiceSla": p.BusinessServiceSla,
		"previousStatus":      p.PreviousState,
		"entity":              p.Entity,
		"auditDetails":        p.AuditDetails.ToMap(),
		"rating":              p.Rating,
		"escalated":           p.Escalated,
	}
}

type ProcessInstanceSearchCriteria struct {
	TenantID                     string   `json:"tenantId"`
	Status                       []string `json:"status"`
	BusinessIds                  []string `json:"businessIds"`
	Assignee                     string   `json:"assignee"`
	Ids                          []string `json:"ids"`
	History                      bool     `json:"history"`
	FromDate                     int      `json:"fromdate"`
	Todate                       int      `json:"todate"`
	Offset                       int      `json:"offset"`
	Limit                        int      `json:"limit"`
	BusinessService              string   `json:"businessService"`
	ModuleName                   string   `json:"modulename"`
	IsNearingSlaCount            bool     `json:"isNearingSlaCount"`
	TenantSpecificStatus         []string `json:"tenantSpecificStatus"`
	MultipleAssignee             []string `json:"multipleassignee"`
	StatesToIgnore               []string `json:"statesToIgnore"`
	IsescalatedCount             bool     `json:"isescalatedcount"`
	IsAssignedToMeCount          bool     `json:"isassignedtomecount"`
	StatusesIrrespectiveOfTenant []string `json:"statusesIrrespectiveSlaLimit"`
	SlotPercentageSlaLimit       int      `json:"slotPercem=ntageSlaLimit"`
}

func (p *ProcessInstanceSearchCriteria) Validate() error {
	if p.TenantID != "" {
		return fmt.Errorf("tenant_id is required")
	}
	if p.BusinessIds != nil {
		for _, b := range p.BusinessIds {
			if len(b) < 4 {
				return fmt.Errorf("business_id must be at least 4 characters")
			}
		}
	}
	return nil
}

func (p *ProcessInstanceSearchCriteria) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":                     p.TenantID,
		"status":                       p.Status,
		"businessIds":                  p.BusinessIds,
		"assignee":                     p.Assignee,
		"ids":                          p.Ids,
		"history":                      p.History,
		"fromDate":                     p.FromDate,
		"toDate":                       p.Todate,
		"offset":                       p.Offset,
		"limit":                        p.Limit,
		"businessService":              p.BusinessService,
		"moduleName":                   p.ModuleName,
		"isNearingSlaCount":            p.IsNearingSlaCount,
		"tenantSpecifiStatus":          p.TenantSpecificStatus,
		"multipleAssignees":            p.MultipleAssignee,
		"statesToIgnore":               p.StatesToIgnore,
		"isEscalatedCount":             p.IsescalatedCount,
		"isAssignedToMeCount":          p.IsAssignedToMeCount,
		"statusesIrrespectiveOfTenant": p.StatusesIrrespectiveOfTenant,
		"slotPercentageSlaLimit":       p.SlotPercentageSlaLimit,
	}
}
func DocumentBuilder() *Document {
	return &Document{}
}
func WorkflowActionBuilder() *WorkflowAction {
	return &WorkflowAction{}
}
func StateBuilder() *State {
	return &State{}
}
func ProcessInstanceBuilder() *ProcessInstance {
	return &ProcessInstance{}
}
func ProcessInstanceSearchCriteriaBuilder() *ProcessInstanceSearchCriteria {
	return &ProcessInstanceSearchCriteria{}
}

func (d *Document) With_id(id string) *Document {
	d.Id = id
	return d
}

func (d *Document) With_tenant_id(tenant_id string) *Document {
	d.TenantID = tenant_id
	return d
}
func (d *Document) With_document_type(document_type string) *Document {
	d.DocumentType = document_type
	return d
}
func (d *Document) With_file_store_id(file_store_id string) *Document {
	d.FileStoreIds = file_store_id
	return d
}
func (d *Document) With_document_uid(document_uid string) *Document {
	d.DocumentUid = document_uid
	return d
}
func (d *Document) With_audit_details(audit_details *AuditDetails) *Document {
	d.AuditDetails = audit_details
	return d
}
func (d *Document) Build() (*Document, error) {
	err := d.Validate()
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (w *WorkflowAction) With_uuid(uuid string) *WorkflowAction {
	w.Uuid = uuid
	return w
}
func (w *WorkflowAction) With_tenant_id(tenant_id string) *WorkflowAction {
	w.TenantID = tenant_id
	return w
}
func (w *WorkflowAction) With_current_state(current_state string) *WorkflowAction {
	w.CurrectState = current_state
	return w
}
func (w *WorkflowAction) With_action(action string) *WorkflowAction {
	w.Action = action
	return w
}
func (w *WorkflowAction) With_next_state(next_state string) *WorkflowAction {
	w.NextState = next_state
	return w
}
func (w *WorkflowAction) With_roles(roles []string) *WorkflowAction {
	w.Roles = roles
	return w
}
func (w *WorkflowAction) With_audit_details(audit_details *AuditDetails) *WorkflowAction {
	w.AuditDetails = audit_details
	return w
}
func (w *WorkflowAction) With_active(active bool) *WorkflowAction {
	w.Active = active
	return w
}
func (w *WorkflowAction) Build() (*WorkflowAction, error) {
	err := w.Validate()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s *State) With_uuid(uuid string) *State {
	s.Uuid = uuid
	return s
}
func (s *State) With_tenant_id(tenant_id string) *State {
	s.TenantID = tenant_id
	return s
}
func (s *State) With_business_service_id(business_service_id string) *State {
	s.BusinessServiceId = business_service_id
	return s
}
func (s *State) With_sla(sla int) *State {
	s.Sla = sla
	return s
}
func (s *State) With_state(state string) *State {
	s.State = state
	return s
}
func (s *State) With_application_status(application_status string) *State {
	s.ApplicationStatus = application_status
	return s
}
func (s *State) With_doc_upload_required(doc_upload_required bool) *State {
	s.DocUploadRequired = doc_upload_required
	return s
}
func (s *State) With_is_start_state(is_start_state bool) *State {
	s.IsStartState = is_start_state
	return s
}
func (s *State) With_is_terminate_state(is_terminate_state bool) *State {
	s.IsTerminateState = is_terminate_state
	return s
}
func (s *State) With_is_state_updatable(is_state_updatable bool) *State {
	s.IsStateUpdatable = is_state_updatable
	return s
}
func (s *State) With_actions(actions []WorkflowAction) *State {
	s.Actions = actions
	return s
}
func (s *State) With_audit_details(audit_details *AuditDetails) *State {
	s.AuditDetails = audit_details
	return s
}
func (s *State) Build() (*State, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (p *ProcessInstance) With_tenant_id(tenant_id string) *ProcessInstance {
	p.TenantID = tenant_id
	return p
}
func (p *ProcessInstance) With_business_service(business_service string) *ProcessInstance {
	p.BusinessService = business_service
	return p
}
func (p *ProcessInstance) With_business_id(business_id string) *ProcessInstance {
	p.BusinessId = business_id
	return p
}
func (p *ProcessInstance) With_action(action string) *ProcessInstance {
	p.Action = action
	return p
}
func (p *ProcessInstance) With_module_name(module_name string) *ProcessInstance {
	p.ModuleName = module_name
	return p
}
func (p *ProcessInstance) With_id(id string) *ProcessInstance {
	p.Id = id
	return p
}
func (p *ProcessInstance) With_state(state State) *ProcessInstance {
	p.State = state
	return p
}
func (p *ProcessInstance) With_comment(comment string) *ProcessInstance {
	p.Comment = comment
	return p
}
func (p *ProcessInstance) With_documents(documents []Document) *ProcessInstance {
	p.Documents = documents
	return p
}
func (p *ProcessInstance) With_assigner(assigner *User) *ProcessInstance {
	p.Assigner = assigner
	return p
}
func (p *ProcessInstance) With_assignes(assignes []User) *ProcessInstance {
	p.Assignes = assignes
	return p
}
func (p *ProcessInstance) With_next_actions(next_actions []WorkflowAction) *ProcessInstance {
	p.NextActions = next_actions
	return p
}
func (p *ProcessInstance) With_state_sla(state_sla int) *ProcessInstance {
	p.StateSla = state_sla
	return p
}
func (p *ProcessInstance) With_business_service_sla(business_service_sla int) *ProcessInstance {
	p.BusinessServiceSla = business_service_sla
	return p
}
func (p *ProcessInstance) With_previous_status(previous_status string) *ProcessInstance {
	p.PreviousState = previous_status
	return p
}
func (p *ProcessInstance) With_entity(entity interface{}) *ProcessInstance {
	p.Entity = entity
	return p
}
func (p *ProcessInstance) With_audit_details(audit_details *AuditDetails) *ProcessInstance {
	p.AuditDetails = audit_details
	return p
}
func (p *ProcessInstance) With_rating(rating int) *ProcessInstance {
	p.Rating = rating
	return p
}
func (p *ProcessInstance) With_escalated(escalated bool) *ProcessInstance {
	p.Escalated = escalated
	return p
}
func (p *ProcessInstance) Build() (*ProcessInstance, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}
func (p *ProcessInstanceSearchCriteria) With_tenant_id(tenant_id string) *ProcessInstanceSearchCriteria {
	p.TenantID = tenant_id
	return p
}
func (p *ProcessInstanceSearchCriteria) With_status(status []string) *ProcessInstanceSearchCriteria {
	p.Status = status
	return p
}
func (p *ProcessInstanceSearchCriteria) With_business_ids(business_ids []string) *ProcessInstanceSearchCriteria {
	p.BusinessIds = business_ids
	return p
}
func (p *ProcessInstanceSearchCriteria) With_assignee(assignee string) *ProcessInstanceSearchCriteria {
	p.Assignee = assignee
	return p
}
func (p *ProcessInstanceSearchCriteria) With_ids(ids []string) *ProcessInstanceSearchCriteria {
	p.Ids = ids
	return p
}
func (p *ProcessInstanceSearchCriteria) With_history(history bool) *ProcessInstanceSearchCriteria {
	p.History = history
	return p
}
func (p *ProcessInstanceSearchCriteria) With_from_date(from_date int) *ProcessInstanceSearchCriteria {
	p.FromDate = from_date
	return p
}
func (p *ProcessInstanceSearchCriteria) With_to_date(to_date int) *ProcessInstanceSearchCriteria {
	p.Todate = to_date
	return p
}
func (p *ProcessInstanceSearchCriteria) With_offset(offset int) *ProcessInstanceSearchCriteria {
	p.Offset = offset
	return p
}
func (p *ProcessInstanceSearchCriteria) With_limit(limit int) *ProcessInstanceSearchCriteria {
	p.Limit = limit
	return p
}
func (p *ProcessInstanceSearchCriteria) With_business_service(business_service string) *ProcessInstanceSearchCriteria {
	p.BusinessService = business_service
	return p
}
func (p *ProcessInstanceSearchCriteria) With_module_name(module_name string) *ProcessInstanceSearchCriteria {
	p.ModuleName = module_name
	return p
}
func (p *ProcessInstanceSearchCriteria) With_is_nearing_sla_count(is_nearing_sla_count bool) *ProcessInstanceSearchCriteria {
	p.IsNearingSlaCount = is_nearing_sla_count
	return p
}
func (p *ProcessInstanceSearchCriteria) With_tenant_specific_status(tenant_specific_status []string) *ProcessInstanceSearchCriteria {
	p.TenantSpecificStatus = tenant_specific_status
	return p
}
func (p *ProcessInstanceSearchCriteria) With_multiple_assignees(multiple_assignees []string) *ProcessInstanceSearchCriteria {
	p.MultipleAssignee = multiple_assignees
	return p
}
func (p *ProcessInstanceSearchCriteria) With_states_to_ignore(states_to_ignore []string) *ProcessInstanceSearchCriteria {
	p.StatesToIgnore = states_to_ignore
	return p
}
func (p *ProcessInstanceSearchCriteria) With_is_escalated_count(is_escalated_count bool) *ProcessInstanceSearchCriteria {
	p.IsescalatedCount = is_escalated_count
	return p
}
func (p *ProcessInstanceSearchCriteria) With_is_assigned_to_me_count(is_assigned_to_me_count bool) *ProcessInstanceSearchCriteria {
	p.IsAssignedToMeCount = is_assigned_to_me_count
	return p
}
func (p *ProcessInstanceSearchCriteria) With_statuses_irrespective_of_tenant(statuses_irrespective_of_tenant []string) *ProcessInstanceSearchCriteria {
	p.StatusesIrrespectiveOfTenant = statuses_irrespective_of_tenant
	return p
}
func (p *ProcessInstanceSearchCriteria) With_slot_percentage_sla_limit(slot_percentage_sla_limit int) *ProcessInstanceSearchCriteria {
	p.SlotPercentageSlaLimit = slot_percentage_sla_limit
	return p
}
func (p *ProcessInstanceSearchCriteria) Build() (*ProcessInstanceSearchCriteria, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}
