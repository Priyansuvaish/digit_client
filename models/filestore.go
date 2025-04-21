package models

import (
	"fmt"
	"mime/multipart"
)

// FileStore represents a file store with a name and path.
type FileUploadRequest struct {
	Files       []*multipart.FileHeader `json:"files"`
	TenantID    string                  `json:"tenant_id"`
	Module      string                  `json:"module"`
	Tag         string                  `json:"tag"`
	RequestInfo *RequestInfo            `json:"request_info"`
}

func (f *FileUploadRequest) Validate() error {
	if len(f.Files) == 0 {
		return fmt.Errorf("files cannot be empty")
	}
	if f.TenantID == "" {
		return fmt.Errorf("tenant_id cannot be empty")
	}
	if f.Module == "" {
		return fmt.Errorf("module cannot be empty")
	}
	return nil
}

func (r *FileUploadRequest) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"tenantId": r.TenantID,
		"module":   r.Module,
	}
	if r.Tag != "" {
		result["tag"] = r.Tag
	}
	if r.RequestInfo != nil {
		result["requestInfo"] = r.RequestInfo
	}
	// Optionally include filenames for reference
	if len(r.Files) > 0 {
		filenames := make([]string, len(r.Files))
		for i, file := range r.Files {
			filenames[i] = file.Filename
		}
		result["files"] = filenames
	}
	return result
}

type FileRetrieveByIdRequest struct {
	TenantID    string `json:"tenant_id"`
	FileStoreId string `json:"file_store_id"`
}

func (f *FileRetrieveByIdRequest) Validate() error {
	if f.FileStoreId == "" {
		return fmt.Errorf("file_store_id cannot be empty")
	}
	if f.TenantID == "" {
		return fmt.Errorf("tenant_id cannot be empty")
	}
	return nil
}
func (f *FileRetrieveByIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":    f.TenantID,
		"fileStoreId": f.FileStoreId,
	}
}

type FileRetrieveByUrlRequest struct {
	TenantID     string   `json:"tenant_id"`
	FileStoreIds []string `json:"file_store_ids"`
}

func (f *FileRetrieveByUrlRequest) Validate() error {
	if len(f.FileStoreIds) == 0 {
		return fmt.Errorf("file_store_ids cannot be empty")
	}
	if f.TenantID == "" {
		return fmt.Errorf("tenant_id cannot be empty")
	}
	return nil
}
func (f *FileRetrieveByUrlRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":     f.TenantID,
		"fileStoreIds": f.FileStoreIds,
	}
}

type FileRetrieveByTagRequest struct {
	TenantID string `json:"tenant_id"`
	Tag      string `json:"tag"`
}

func (f *FileRetrieveByTagRequest) Validate() error {
	if f.Tag == "" {
		return fmt.Errorf("tag cannot be empty")
	}
	if f.TenantID == "" {
		return fmt.Errorf("tenant_id cannot be empty")
	}
	return nil
}
func (f *FileRetrieveByTagRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId": f.TenantID,
		"tag":      f.Tag,
	}
}

func FileUploadRequestBuilder() *FileUploadRequest {
	return &FileUploadRequest{}
}
func FileRetrieveByIdRequestBuilder() *FileRetrieveByIdRequest {
	return &FileRetrieveByIdRequest{}
}
func FileRetrieveByUrlRequestBuilder() *FileRetrieveByUrlRequest {
	return &FileRetrieveByUrlRequest{}
}
func FileRetrieveByTagRequestBuilder() *FileRetrieveByTagRequest {
	return &FileRetrieveByTagRequest{}
}

func (f *FileUploadRequest) WithFiles(files []*multipart.FileHeader) *FileUploadRequest {
	f.Files = files
	return f
}

func (f *FileUploadRequest) WithTenantID(tenantID string) *FileUploadRequest {
	f.TenantID = tenantID
	return f
}
func (f *FileUploadRequest) WithModule(module string) *FileUploadRequest {
	f.Module = module
	return f
}
func (f *FileUploadRequest) WithTag(tag string) *FileUploadRequest {
	f.Tag = tag
	return f
}
func (f *FileUploadRequest) WithRequestInfo(requestInfo *RequestInfo) *FileUploadRequest {
	f.RequestInfo = requestInfo
	return f
}
func (f *FileUploadRequest) Build() *FileUploadRequest {
	if err := f.Validate(); err != nil {
		panic(fmt.Sprintf("FileUploadRequest validation failed: %v", err))
	}
	// Perform any additional setup or processing if needed
	return f
}

func (f *FileRetrieveByIdRequest) WithTenantID(tenantID string) *FileRetrieveByIdRequest {
	f.TenantID = tenantID
	return f
}
func (f *FileRetrieveByIdRequest) WithFileStoreId(fileStoreId string) *FileRetrieveByIdRequest {
	f.FileStoreId = fileStoreId
	return f
}
func (f *FileRetrieveByIdRequest) Build() *FileRetrieveByIdRequest {
	if err := f.Validate(); err != nil {
		panic(fmt.Sprintf("FileRetrieveByIdRequest validation failed: %v", err))
	}
	return f
}
func (f *FileRetrieveByUrlRequest) WithTenantID(tenantID string) *FileRetrieveByUrlRequest {
	f.TenantID = tenantID
	return f
}
func (f *FileRetrieveByUrlRequest) WithFileStoreIds(fileStoreIds []string) *FileRetrieveByUrlRequest {
	f.FileStoreIds = fileStoreIds
	return f
}
func (f *FileRetrieveByUrlRequest) Build() *FileRetrieveByUrlRequest {
	if err := f.Validate(); err != nil {
		panic(fmt.Sprintf("FileRetrieveByUrlRequest validation failed: %v", err))
	}
	return f
}

func (f *FileRetrieveByTagRequest) WithTenantID(tenantID string) *FileRetrieveByTagRequest {
	f.TenantID = tenantID
	return f
}
func (f *FileRetrieveByTagRequest) WithTag(tag string) *FileRetrieveByTagRequest {
	f.Tag = tag
	return f
}
func (f *FileRetrieveByTagRequest) Build() *FileRetrieveByTagRequest {
	if err := f.Validate(); err != nil {
		panic(fmt.Sprintf("FileRetrieveByTagRequest validation failed: %v", err))
	}
	return f
}
