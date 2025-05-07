// package example

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	// "github.com/Priyansuvaish/digit_client/client"
// 	"github.com/Priyansuvaish/digit_client/config"
// 	"github.com/Priyansuvaish/digit_client/models"
// 	"github.com/Priyansuvaish/digit_client/services"
// 	"github.com/google/uuid"
// )

// func AuthorizeExample() {
// 	// Create a new API client with DIGIT sandbox URL
// 	// apiClient := client.NewAPIClient("https://sandbox.digit.org", "")
// 	config.GetGlobalConfig().Initialize(
// 		"https://sandbox.digit.org",
// 		"your-auth-token",
// 	)
// 	// Initialize RequestConfig properly
// 	requestConfig := (&models.RequestConfig{}).GetInstance()
// 	requestConfig.Initialize(
// 		"digit",                  // apiID
// 		"1.0",                    // version
// 		"your-auth-token-here",   // authToken
// 		map[string]interface{}{}, // userInfo
// 		"device123",              // did
// 		"key123",                 // key
// 		uuid.New().String(),      // msgID
// 		"requester123",           // requesterID
// 		uuid.New().String(),      // correlationID
// 		"authorize",              // action
// 		time.Now().UnixMilli(),   // ts
// 	)

// 	// Create authorize service
// 	authorizeService := services.NewAuthorizeService(nil)

// 	// Build Role first
// 	role := models.RoleBuilder().WithCode("ADMIN").Build()

// 	// Create authorization request
// 	authRequest, err := models.AuthorizationRequestBuilder().
// 		WithURI("/api/v1/users").
// 		AddRole(*role). // Dereference the pointer
// 		AddTenantID("pb").
// 		Build()

// 	if err != nil {
// 		log.Fatalf("Error building authorization request: %v", err)
// 	}

// 	action, err := models.ActionBuilder().
// 		WithName("view").
// 		WithURL("/api/v1/users").
// 		WithDisplayName("View Users").
// 		Build()
// 	if err != nil {
// 		log.Fatalf("Error building ActionBuilder: %v", err)
// 	}
// 	actionrequest, err := models.ActionRequestBuilder().
// 		WithActions([]*models.Action{action}). // Dereference the pointer
// 		Build()

// 	if err != nil {
// 		log.Fatalf("Error building actionrequest: %v", err)
// 	}

// 	// Get authorized actions (pass nil to use default RequestInfo)
// 	authorizedActions, err := authorizeService.AuthorizeAction(authRequest, nil)
// 	if err != nil {
// 		log.Fatalf("Error authorizing action: %v", err)
// 	}
// 	// GetMDMSAction gets MDMS action details
// 	getmdmsaction, err := authorizeService.GetMDMSAction(actionrequest, nil)
// 	if err != nil {
// 		log.Fatalf("Error getting MDMS action: %v", err)
// 	}
// 	// Print the MDMS action details
// 	fmt.Printf("MDMS action details: %+v\n", getmdmsaction)
// 	// Print the authorized actions

// 	fmt.Printf("Authorized actions: %+v\n", authorizedActions)
// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "github.com/Priyansuvaish/digit_client/client"
	"github.com/Priyansuvaish/digit_client/config"
	"github.com/Priyansuvaish/digit_client/models"
	"github.com/Priyansuvaish/digit_client/services"
)

func main() {
	//for the auth token
	// token, err := digit_init.Authenticate()
	// // configre the url and auth token globally

	// if err != nil {
	// 	log.Fatalf("Failed to authenticate: %v", err)
	// }
	// fmt.Println("Received token:", token)

	// // Step 2: Update or append AUTH_TOKEN in .env file
	// envFile := ".env"

	// // Read existing lines
	// lines, err := readLines(envFile)
	// if err != nil {
	// 	// If file doesn't exist, create new
	// 	if os.IsNotExist(err) {
	// 		lines = []string{}
	// 	} else {
	// 		log.Fatalf("Failed to read .env file: %v", err)
	// 	}
	// }

	// // Update or add AUTH_TOKEN line
	// found := false
	// for i, line := range lines {
	// 	if strings.HasPrefix(line, "AUTH_TOKEN=") {
	// 		lines[i] = fmt.Sprintf("AUTH_TOKEN=%s", token)
	// 		found = true
	// 		break
	// 	}
	// }
	// if !found {
	// 	lines = append(lines, fmt.Sprintf("AUTH_TOKEN=%s", token))
	// }

	// // Write back to .env file
	// err = writeLines(lines, envFile)
	// if err != nil {
	// 	log.Fatalf("Failed to write to .env file: %v", err)
	// }

	// fmt.Println(".env file updated successfully")

	config.GetGlobalConfig().Initialize(
		"https://unified-dev.digit.org",
		"",
	)

	// Initialize RequestConfig properly
	requestConfig := (&models.RequestConfig{}).GetInstance()
	requestConfig.Initialize(
		"hcm",                                  // apiID
		".01",                                  // version
		"6aed1556-7385-4a93-a490-8c7ba937f304", // authToken
		map[string]interface{}{},               // userInfo
		"1",                                    // did
		"1",                                    // key
		"",                                     // msgID
		"",                                     // requesterID
		"",                                     // correlationID
		"_get",                                 // action
		1709096352589,                          // ts
	)

	masterdeatils, err := models.MasterDetailBuilder().WithName("ServiceConfiguration").Build()
	moduledetails := models.ModuleDetailBuilder().WithModuleName("Studio").WithMasterDetails([]models.MasterDetail{*masterdeatils})
	MdmsCriteria := models.MdmsCriteriaBuilder().WithTenantId("dev").WithModuleDetails([]models.ModuleDetail{*moduledetails})
	mdmsserice := services.NewMDMSService(nil)
	// Call CreateMDMS method
	createdMDMS, err := mdmsserice.Search(MdmsCriteria, nil)
	if err != nil {
		log.Printf("API call failed: %v", err)
	}
	fmt.Println("MDMS created:", createdMDMS.(map[string]interface{}))

	// // Optionally do something with createdMDMS
	// mdmsRes, ok := createdMDMS.(map[string]interface{})["MdmsRes"].(map[string]interface{})
	// if !ok {
	// 	log.Println("Failed to assert MdmsRes as map[string]interface{}")
	// 	return
	// }

	// fmt.Println("MDMS created:", mdmsRes["PropertyTax"].(map[string]interface{})["code"])

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return w.Flush()
}
