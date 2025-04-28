package models

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type TxnStatusEnum string

const (
	SUCCESS TxnStatusEnum = "SUCCESS"
	FAILURE TxnStatusEnum = "FAILURE"
	PENDING TxnStatusEnum = "PENDING"
)

type TaxAndPayment struct {
	TaxAmount  interface{} `json:"taxamount"`
	AmountPaid interface{} `json:"amountpaid"`
	BillId     string      `json:"billid"`
}

func (t *TaxAndPayment) Validate() error {
	switch v := t.TaxAmount.(type) {
	case decimal.Decimal:
		if v.LessThan(decimal.Zero) {
			return fmt.Errorf("tax_amount cannot be negative")
		}
	case float64:
		if v < 0 {
			return fmt.Errorf("tax_amount cannot be negative")
		}
	case int:
		if v < 0 {
			return fmt.Errorf("tax_amount cannot be negative")
		}
	default:
		return fmt.Errorf("tax_amount has an unsupported type")
	}
	switch v := t.AmountPaid.(type) {
	case decimal.Decimal:
		if v.LessThan(decimal.Zero) {
			return fmt.Errorf("amount_paid cannot be negative")
		}
	case float64:
		if v < 0 {
			return fmt.Errorf("amount_paid cannot be negative")
		}
	case int:
		if v < 0 {
			return fmt.Errorf("amount_paid cannot be negative")
		}
	default:
		return fmt.Errorf("amount_paid has an unsupported type")
	}
	if t.BillId == "" {
		return fmt.Errorf("bill_id is required")
	}
	return nil
}

func (t *TaxAndPayment) ToMap() map[string]interface{} {
	var taxAmount interface{}
	var amountPaid interface{}

	// Convert TaxAmount to string if it's decimal.Decimal
	switch v := t.TaxAmount.(type) {
	case decimal.Decimal:
		taxAmount = v.String()
	case *decimal.Decimal:
		taxAmount = v.String()
	default:
		taxAmount = t.TaxAmount
	}

	// Convert AmountPaid to string if it's decimal.Decimal
	switch v := t.AmountPaid.(type) {
	case decimal.Decimal:
		amountPaid = v.String()
	case *decimal.Decimal:
		amountPaid = v.String()
	default:
		amountPaid = t.AmountPaid
	}

	return map[string]interface{}{
		"taxAmount":  taxAmount,
		"amountPaid": amountPaid,
		"billId":     t.BillId,
	}
}

type Transaction struct {
	TenantID           string                 `json:"tenantId"`
	TaxAmount          string                 `json:"taxamount"`
	BillId             string                 `json:"billId"`
	ConsumerCode       string                 `json:"consumercode"`
	TaxAndPayments     []TaxAndPayment        `json:"taxandpayments"`
	ProductInfo        string                 `json:"productinfo"`
	Gateway            string                 `json:"gateway"`
	CallbackUrl        string                 `json:"callbackurl"`
	User               *User                  `json:"user"`
	Module             string                 `json:"module"`
	TxnId              string                 `json:"txnId"`
	Redirect_url       string                 `json:"redirect_url"`
	TxnStatus          TxnStatusEnum          `json:"txtstatus"`
	TxtStatusMessage   string                 `json:"txtstatusmessage"`
	GatewayTxnId       string                 `json:"gatewaytxnid"`
	GatewayPaymentMode string                 `json:"gatewaypaymentmode"`
	GatewayStatusCode  string                 `json:"gatestatuscode"`
	GatewayStatusMsg   string                 `json:"gatewaystatusmsg"`
	Receipt            string                 `json:"receipt"`
	AuditDetails       *AuditDetails          `json:"auditdetails"`
	AdditionalDetails  map[string]interface{} `json:"additionaldetails"`
	ResponseJson       map[string]interface{} `json:"responsejson"`
	AdditionalFields   map[string]interface{} `json:"additionalfields"`
}

func (t *Transaction) Validate() error {
	if len(t.TenantID) < 2 || len(t.TenantID) > 50 {
		return fmt.Errorf("tenant_id must be between 2-50 characters")
	}
	if t.TaxAmount != "" {
		return fmt.Errorf("txn_amount is required")
	}
	if t.BillId != "" {
		return fmt.Errorf("bill_id is required")
	}
	if t.ConsumerCode != "" || len(t.ConsumerCode) > 128 {
		return fmt.Errorf("consumer_code must be 1-128 characters")
	}
	if t.ProductInfo != "" || len(t.ProductInfo) > 512 {
		return fmt.Errorf("product_info must be 1-512 characters")
	}
	if t.Gateway != "" {
		return fmt.Errorf("gateway is required")
	}
	if t.CallbackUrl != "" {
		return fmt.Errorf("callback_url is required")
	}
	if t.TaxAndPayments != nil {
		return fmt.Errorf("tax_and_payments cannot be empty")
	}
	return nil
}

func (t *Transaction) ToMap() map[string]interface{} {
	taxandpay := make([]map[string]interface{}, len(t.TaxAndPayments))
	for _, f := range t.TaxAndPayments {
		taxandpay = append(taxandpay, f.ToMap())
	}
	return map[string]interface{}{
		"tenantId":           t.TenantID,
		"txnAmount":          t.TaxAmount,
		"billId":             t.BillId,
		"consumerCode":       t.ConsumerCode,
		"taxAndPayments":     taxandpay,
		"productInfo":        t.ProductInfo,
		"gateway":            t.Gateway,
		"callbackUrl":        t.CallbackUrl,
		"user":               t.User,
		"module":             t.Module,
		"txnId":              t.TxnId,
		"redirectUrl":        t.Redirect_url,
		"txnStatus":          t.TxnStatus,
		"txnStatusMsg":       t.TxtStatusMessage,
		"gatewayTxnId":       t.GatewayTxnId,
		"gatewayPaymentMode": t.GatewayPaymentMode,
		"gatewayStatusCode":  t.GatewayStatusCode,
		"gatewayStatusMsg":   t.GatewayStatusMsg,
		"receipt":            t.Receipt,
		"auditDetails":       t.AuditDetails,
		"additionalDetails":  t.AdditionalDetails,
	}
}

type TransactionCriteria struct {
	TenantID     string        `json:"tenantId"`
	TxnId        string        `json:"txnId"`
	BillId       string        `json:"billId"`
	UserUuid     string        `json:"useruuid"`
	Receipt      string        `json:"receipt"`
	ConsumerCode string        `json:"consumercode"`
	CreatedTime  int           `json:"createdtime"`
	TxnStatus    TxnStatusEnum `json:"txnstatus"`
	Limit        int           `json:"limit"`
	Offset       int           `json:"offset"`
}

func (t *TransactionCriteria) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"tenantId":     t.TenantID,
		"txnId":        t.TxnId,
		"billId":       t.BillId,
		"userUuid":     t.UserUuid,
		"receipt":      t.Receipt,
		"consumerCode": t.ConsumerCode,
		"createdTime":  t.CreatedTime,
		"txnStatus":    t.TxnStatus,
		"limit":        t.Limit,
		"offset":       t.Offset,
	}
}

func TaxAndPaymentBuilder() *TaxAndPayment {
	return &TaxAndPayment{}
}
func TransactionBuilder() *Transaction {
	return &Transaction{}
}
func TransactionCriteriaBuilder() *TransactionCriteria {
	return &TransactionCriteria{}
}
func (t *TaxAndPayment) WithTaxAmount(amt interface{}) *TaxAndPayment {
	t.TaxAmount = amt
	return t
}
func (t *TaxAndPayment) WithAmountPaid(amt interface{}) *TaxAndPayment {
	t.AmountPaid = amt
	return t
}
func (t *TaxAndPayment) WithBillId(id string) *TaxAndPayment {
	t.BillId = id
	return t
}

func (t *TaxAndPayment) Build() (*TaxAndPayment, error) {
	err := t.Validate()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Transaction) WithTenantId(code string) *Transaction {
	t.TenantID = code
	return t
}
func (t *Transaction) WithTxnAmount(code string) *Transaction {
	t.TaxAmount = code
	return t
}
func (t *Transaction) WithBillId(code string) *Transaction {
	t.BillId = code
	return t
}
func (t *Transaction) WithConsumerCode(code string) *Transaction {
	t.ConsumerCode = code
	return t
}
func (t *Transaction) WithTaxAndPayments(code []TaxAndPayment) *Transaction {
	t.TaxAndPayments = code
	return t
}
func (t *Transaction) WithProductInfo(code string) *Transaction {
	t.ProductInfo = code
	return t
}
func (t *Transaction) WithGateway(code string) *Transaction {
	t.Gateway = code
	return t
}
func (t *Transaction) WithCallBack(code string) *Transaction {
	t.CallbackUrl = code
	return t
}
func (t *Transaction) WithUser(code *User) *Transaction {
	t.User = code
	return t
}
func (t *Transaction) WithModule(code string) *Transaction {
	t.Module = code
	return t
}
func (t *Transaction) WithTxnId(code string) *Transaction {
	t.TxnId = code
	return t
}
func (t *Transaction) WithRedirectUrl(code string) *Transaction {
	t.Redirect_url = code
	return t
}
func (t *Transaction) WithTxnStatus(code TxnStatusEnum) *Transaction {
	t.TxnStatus = code
	return t
}
func (t *Transaction) WithTxnStatusMsg(code string) *Transaction {
	t.TxtStatusMessage = code
	return t
}
func (t *Transaction) WithGatewayTxnId(code string) *Transaction {
	t.GatewayTxnId = code
	return t
}
func (t *Transaction) WithGatewayPaymentMode(code string) *Transaction {
	t.GatewayPaymentMode = code
	return t
}
func (t *Transaction) WithGatewayStatusCode(code string) *Transaction {
	t.GatewayStatusCode = code
	return t
}
func (t *Transaction) WithGatewayStatusMsg(code string) *Transaction {
	t.GatewayStatusMsg = code
	return t
}
func (t *Transaction) WithReceipt(code string) *Transaction {
	t.Receipt = code
	return t
}
func (t *Transaction) WithAuditDetails(code *AuditDetails) *Transaction {
	t.AuditDetails = code
	return t
}
func (t *Transaction) WithAdditionalDetails(code map[string]interface{}) *Transaction {
	t.AdditionalDetails = code
	return t
}
func (t *Transaction) WithResponseJson(code map[string]interface{}) *Transaction {
	t.ResponseJson = code
	return t
}
func (t *Transaction) WithAdditionalFields(code map[string]interface{}) *Transaction {
	t.AdditionalFields = code
	return t
}
func (t *Transaction) Build() (*Transaction, error) {
	err := t.Validate()
	if err != nil {
		return nil, err
	}
	return t, nil
}
func (t *TransactionCriteria) WithTenantId(code string) *TransactionCriteria {
	t.TenantID = code
	return t
}
func (t *TransactionCriteria) WithTxnId(code string) *TransactionCriteria {
	t.TxnId = code
	return t
}
func (t *TransactionCriteria) WithBillId(code string) *TransactionCriteria {
	t.BillId = code
	return t
}
func (t *TransactionCriteria) WithUserUuid(code string) *TransactionCriteria {
	t.UserUuid = code
	return t
}
func (t *TransactionCriteria) WithReceipt(code string) *TransactionCriteria {
	t.Receipt = code
	return t
}
func (t *TransactionCriteria) WithConsumerCode(code string) *TransactionCriteria {
	t.ConsumerCode = code
	return t
}
func (t *TransactionCriteria) WithCreateTime(code int) *TransactionCriteria {
	t.CreatedTime = code
	return t
}
func (t *TransactionCriteria) WithTxnStatus(code TxnStatusEnum) *TransactionCriteria {
	t.TxnStatus = code
	return t
}
func (t *TransactionCriteria) WithLimit(code int) *TransactionCriteria {
	t.Limit = code
	return t
}
func (t *TransactionCriteria) WithOffset(code int) *TransactionCriteria {
	t.Offset = code
	return t
}
func (t *TransactionCriteria) Build() (*TransactionCriteria, error) {
	return t, nil
}
