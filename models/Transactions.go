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
