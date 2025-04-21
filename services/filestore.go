package services

import (
	"github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/models"
)

// FilestoreService is a service for interacting with the filestore API.
type FilestoreService struct {
	apiClient *client.APIClient
	baseURL   string
}

// NewFilestoreService creates a new FilestoreService.
func NewFilestoreService(apiClient *client.APIClient) *FilestoreService {
	if apiClient == nil {
		apiClient = client.DefaultAPIClient()
	}
	return &FilestoreService{
		apiClient: apiClient,
		baseURL:   "filestore/v1/files",
	}
}

// UploadFile uploads a file to the filestore.
func (s *FilestoreService) UploadFile(file *models.FileUploadRequest) (interface{}, error) {
	data := make(map[string]string)
	for key, value := range file.ToMap() {
		if strValue, ok := value.(string); ok {
			data[key] = strValue
		}
	}
	// Add files to the request
	// filesMap := make(map[string]string)
	// if file.Files != nil && len(file.Files) > 0 {
	// 	for _, f := range file.Files {
	// 		fileContent, err := f.Open()
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		defer fileContent.Close()
	// 		contentBytes, err := io.ReadAll(fileContent)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		filesMap[f.Filename] = string(contentBytes) // Convert file content to string
	// 	}
	// }
	endpoint := s.baseURL
	// Make the API call to upload the file
	return s.apiClient.Post(
		endpoint,
		nil,
		data,
		nil,
		nil, // Pass filesMap as map[string]string
		nil,
		true,
	)
}

func (s *FilestoreService) Get_file_by_id(fileID *models.FileRetrieveByIdRequest) (interface{}, error) {
	param := fileID.ToMap()
	stringParam := make(map[string]string)
	for key, value := range param {
		if strValue, ok := value.(string); ok {
			stringParam[key] = strValue
		}
	}
	endpoint := s.baseURL + "/id"
	// Make the API call to retrieve the file by ID
	return s.apiClient.Get(
		endpoint,
		stringParam,
		true,
		true,
	)
}

func (s *FilestoreService) Get_file_by_metadata(fileID *models.FileRetrieveByIdRequest) (interface{}, error) {
	param := fileID.ToMap()
	stringParam := make(map[string]string)
	for key, value := range param {
		if strValue, ok := value.(string); ok {
			stringParam[key] = strValue
		}
	}
	endpoint := s.baseURL + "/metadata"
	// Make the API call to retrieve the file by tag
	return s.apiClient.Get(
		endpoint,
		stringParam,
		true,
		true,
	)
}

func (s *FilestoreService) Get_file_by_Url(fileID *models.FileRetrieveByUrlRequest) (interface{}, error) {
	param := fileID.ToMap()
	stringParam := make(map[string]string)
	for key, value := range param {
		if strValue, ok := value.(string); ok {
			stringParam[key] = strValue
		}
	}
	endpoint := s.baseURL + "/url"
	// Make the API call to retrieve the file by URL
	return s.apiClient.Get(
		endpoint,
		stringParam,
		true,
		true,
	)
}

func (s *FilestoreService) Get_file_by_tag(fileID *models.FileRetrieveByTagRequest) (interface{}, error) {
	param := fileID.ToMap()
	stringParam := make(map[string]string)
	for key, value := range param {
		if strValue, ok := value.(string); ok {
			stringParam[key] = strValue
		}
	}
	endpoint := s.baseURL + "/tag"
	// Make the API call to retrieve the file by tag
	return s.apiClient.Get(
		endpoint,
		stringParam,
		true,
		true,
	)
}
