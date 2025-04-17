package main

import (
	"fmt"
	"log"
	"time"

	// "github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/config"
	"github.com/Priyansuvaish/digit_client/models"
	"github.com/Priyansuvaish/digit_client/services"
	"github.com/google/uuid"
)

func main() {
	// Initialize the configuration
	config.GetGlobalConfig().Initialize(
		"https://sandbox.digit.org",
		"0e9b955f-5e25-4809-b680-97ef37ccf53f",
	)

	requestConfig := (&models.RequestConfig{}).GetInstance()
	requestConfig.Initialize(
		"digit",                                // apiID
		"1.0",                                  // version
		"0e9b955f-5e25-4809-b680-97ef37ccf53f", // authToken
		map[string]interface{}{},               // userInfo
		"device123",                            // did
		"key123",                               // key
		uuid.New().String(),                    // msgID
		"requester123",                         // requesterID
		uuid.New().String(),                    // correlationID
		"authorize",                            // action
		time.Now().UnixMilli(),                 // ts
	)
	IDRequestService := services.NewIDRequestService(nil)

	// roles := []models.Role{
	// 	*models.RoleBuilder().WithCode("EMPLOYEE").WithName("EMPLOYEE").WithTenantID("LMN").Build(),
	// 	*models.RoleBuilder().WithCode("SYSTEM").WithName("System user").WithTenantID("LMN").Build(),
	// }

	//create the idrequest
	idRequest, err := models.IdRequestBuilder().
		SetIdName("test_id_name").
		SetTenantID("test_tenant_id").
		SetFormat("test_format").Build()

	if err != nil {
		log.Fatalf("Error building IDRequest: %v", err)
	}

	result, err := IDRequestService.Generate_id(idRequest, nil)
	if err != nil {
		log.Fatalf("Error generating ID: %v", err)
	}
	fmt.Printf("Generated ID: %v\n", result)
}
