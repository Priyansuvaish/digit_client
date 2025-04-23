package digit_init

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// BaseURL for the API
const BaseURL = "https://sandbox.digit.org"

// RequestInfo struct for common request info
type Requestinfo struct {
	ApiID              string      `json:"apiId"`
	AuthToken          interface{} `json:"authToken"`
	UserInfo           interface{} `json:"userInfo"`
	MsgID              string      `json:"msgId"`
	PlainAccessRequest interface{} `json:"plainAccessRequest"`
}

// --- Register User ---

type RegisterTenantPayload struct {
	Tenant      map[string]string `json:"tenant"`
	RequestInfo Requestinfo       `json:"RequestInfo"`
}

// --- Send OTP ---
type OtpPayload struct {
	Otp         map[string]string `json:"otp"`
	RequestInfo Requestinfo       `json:"RequestInfo"`
}

func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func registerUser(email, tenantID string) (interface{}, error) {
	url := fmt.Sprintf("%s/tenant-management/tenant/_create", BaseURL)
	payload := RegisterTenantPayload{
		Tenant: map[string]string{
			"name":  tenantID,
			"email": email,
		},
		RequestInfo: Requestinfo{
			ApiID:              "Rainmaker",
			AuthToken:          nil,
			UserInfo:           nil,
			MsgID:              "registration",
			PlainAccessRequest: map[string]interface{}{},
		},
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)
	return result, nil
}

func sendOTP(email, tenantID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/user-otp/v1/_send", BaseURL)
	params := url.Values{}
	params.Add("tenantId", tenantID)
	fullURL := endpoint + "?" + params.Encode()
	payload := OtpPayload{
		Otp: map[string]string{
			"userName": email,
			"type":     "login",
			"tenantId": tenantID,
			"userType": "EMPLOYEE",
		},
		RequestInfo: Requestinfo{
			ApiID:              "Rainmaker",
			AuthToken:          nil,
			UserInfo:           nil,
			MsgID:              "otp_request",
			PlainAccessRequest: map[string]interface{}{},
		},
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)
	return result, nil
}

func validateOTP(email, otp, tenantID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("%s/user/oauth/token", BaseURL)
	data := url.Values{}
	data.Set("username", email)
	data.Set("password", otp)
	data.Set("tenantId", tenantID)
	data.Set("userType", "EMPLOYEE")
	data.Set("scope", "read")
	data.Set("grant_type", "password")

	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("authorization", "Basic ZWdvdi11c2VyLWNsaWVudDo=")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("origin", "https://sandbox.digit.org")
	req.Header.Set("referer", fmt.Sprintf("https://sandbox.digit.org/sandbox-ui/%s/employee/user/login/otp", tenantID))
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func handleRegistration() (string, error) {
	email := readInput("Enter your email: ")
	tenantID := readInput("Enter your tenant ID: ")

	_, err := registerUser(email, tenantID)
	if err != nil {
		return "", err
	}

	_, err = sendOTP(email, tenantID)
	if err != nil {
		return "", err
	}

	otp := readInput("Enter the OTP: ")
	token, err := validateOTP(email, otp, tenantID)
	if err != nil {
		return "", fmt.Errorf("OTP validation failed: %v", err)
	}
	tok := token["access_token"].(string)
	return tok, nil
}

func handleLogin() (string, error) {
	email := readInput("Enter your email: ")
	tenantID := readInput("Enter your tenant ID: ")
	sendOTP(email, tenantID)
	otp := readInput("Enter the OTP: ")
	token, err := validateOTP(email, otp, tenantID)
	if err != nil {
		return "", fmt.Errorf("OTP validation failed: %v", err)
	}
	tok := token["access_token"].(string)
	// fmt.Println("Access Token:", tok)
	return tok, nil
}

func Authenticate() (string, error) {
	fmt.Println("=== DIGIT Authentication ===")
	fmt.Println("Before using the package, you need to authenticate.")
	for {
		choice := readInput("Do you want to (1) Register or (2) Login? ")
		if choice == "1" {
			token, err := handleRegistration()
			if err != nil {
				return "", err
			}
			fmt.Printf("Access token: %s\n", token)
			return token, nil
		} else if choice == "2" {
			token, err := handleLogin()
			if err != nil {
				return "", err
			}
			fmt.Printf("Access token: %s\n", token)

			return token, nil
		}
		fmt.Println("Invalid choice. Try again.")

	}
}
