package finance

import (
	"encoding/json"

	"github.com/hgiasac/hasura-utils/v2/types"
)

type PaymentMethod string

const (
	PaymentMethodCash   PaymentMethod = "cash"
	PaymentMethodCod    PaymentMethod = "cod"
	PaymentMethodBank   PaymentMethod = "bank"
	PaymentMethodWallet PaymentMethod = "wallet"
	PaymentMethodCard   PaymentMethod = "card"
	PaymentMethodPaypal PaymentMethod = "paypal"
)

// PaymentStatus represents payment status enum
type PaymentStatus string

const (
	PaymentPending  PaymentStatus = "pending"
	PaymentOnHold   PaymentStatus = "on_hold"
	PaymentRefunded PaymentStatus = "refunded"
	PaymentSuccess  PaymentStatus = "success"
	PaymentFailure  PaymentStatus = "failure"
)

// PaymentCardInput represents the payment card information
type PaymentCardInput struct {
	CardNumber     string     `json:"card_number"`
	CardholderName string     `json:"cardholder_name"`
	ExpiryDate     types.Date `json:"expiry_date"`
	CVV            string     `json:"cvv"`
}

// CreatePaymentInput represents the input of the CreatePayment interface
type CreatePaymentInput struct {
	RequestID      string            `json:"-"`
	PaymentMethod  string            `json:"payment_method"`
	PaymentGateway string            `json:"payment_gateway"`
	Amount         int64             `json:"amount"`
	CurrencyCode   string            `json:"currency_code"`
	Nonce          string            `json:"nonce"`
	OrderID        string            `json:"order_id"`
	CustomerID     string            `json:"customer_id"`
	CallbackURL    string            `json:"callback_url"`
	CancelURL      string            `json:"cancel_url"`
	PaymentCard    *PaymentCardInput `json:"payment_card"`
	AdditionalInfo json.RawMessage   `json:"additional_info"`
}

// CreatePaymentOutput represents the output of the CreatePayment interface
type CreatePaymentOutput struct {
	ID          string        `graphql:"id" json:"id"`
	Status      PaymentStatus `graphql:"status" json:"status"`
	Message     string        `graphql:"message" json:"message,omitempty"`
	CheckoutURL string        `graphql:"checkout_url" json:"checkout_url,omitempty"`
	PaymentInfo any           `graphql:"payment_info" scalar:"true" json:"payment_info,omitempty"`
	Response    any           `graphql:"response" scalar:"true" json:"response"`
}

// RefundPaymentInput represents the refund payment input data
type RefundPaymentInput struct {
	RequestID      string `json:"-"`
	PaymentGateway string `json:"payment_gateway"`
	OrderID        string `json:"order_id"`
	TransactionID  string `json:"transaction_id"`
	Amount         int64  `json:"amount"`
	CurrencyCode   string `json:"currency_code"`
	CustomerID     string `json:"customer_id"`
	OrderNumber    string `json:"order_number"`
}

// RefundPaymentOutput represents the refund output data
type RefundPaymentOutput struct {
	Status   PaymentStatus `graphql:"status" json:"status"`
	Message  string        `graphql:"message" json:"message,omitempty"`
	Response any           `graphql:"response" scalar:"true" json:"response"`
}
