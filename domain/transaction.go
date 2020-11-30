package domain

import (
	"context"
	"time"
)

const (
	AUTH_TYPE_BASIC = "basic"
	AUTH_TYPE_BODY  = "body"
	AUTH_TYPE_URL   = "url"
)

type Transaction struct {
	ApiKey               string         `json:"api_key,omitempty"`
	Amount               int32          `json:"amount"`
	CardHash             string         `json:"card_hash,omitempty"`
	CardId               string         `json:"card_id,omitempty"`
	CardHolderName       string         `json:"card_holder_name,omitempty"`
	CardExpirationDate   string         `json:"card_expiration_date,omitempty"`
	CardNumber           string         `json:"card_number,omitempty"`
	CardCvv              string         `json:"card_cvv,omitempty"`
	PaymentMethod        string         `json:"payment_method,omitempty"`
	PostbackUrl          string         `json:"postback_url,omitempty"`
	Async                bool           `json:"async"`
	Installments         string         `json:"installments,omitempty"`
	Capture              string         `json:"capture,omitempty"`
	BoletoExpirationDate string         `json:"boleto_expiration_date,omitempty"`
	SoftDescriptor       string         `json:"soft_descriptor,omitempty"`
	Customer             Customer       `json:"customer"`
	Billing              Billing        `json:"billing"`
	Shipping             Shipping       `json:"shipping"`
	Items                []Item         `json:"items"`
	Metadata             interface{}    `json:"metadata,omitempty"`
	SplitRules           []interface{}  `json:"split_rules,omitempty"`
	BoletoFine           BoletoFine     `json:"boleto_fine,omitempty"`
	BoletoInterest       BoletoInterest `json:"boleto_interest,omitempty"`
	BoletoRules          []string       `json:"boleto_rules,omitempty"`
	ReferenceKey         string         `json:"reference_key,omitempty"`
	Session              string         `json:"session,omitempty"`
	LocalTime            time.Time      `json:"local_time"`
}

type Customer struct {
	ExternalId   string     `json:"external_id" validate:"required"`
	Name         string     `json:"name" validate:"required"`
	Email        string     `json:"email" validate:"required"`
	Country      string     `json:"country" validate:"required"`
	Type         string     `json:"type" validate:"required"`
	Documents    []Document `json:"documents" validate:"required"`
	PhoneNumbers []string   `json:"phone_numbers" validate:"required"`
}

type Billing struct {
	Name    string  `json:"name" validate:"required"`
	Address Address `json:"address" validate:"required"`
}

type Shipping struct {
	Name         string  `json:"name" validate:"required"`
	Fee          int32   `json:"fee" validate:"required"`
	DeliveryDate string  `json:"delivery_date"`
	Expedited    bool    `json:"expedited"`
	Address      Address `json:"address" validate:"required"`
}

type Item struct {
	ID        string `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	UnitPrice int32  `json:"unit_price" validate:"required"`
	Quantity  int32  `json:"quantity" validate:"required"`
	Tangible  bool   `json:"tangible" validate:"required"`
	Category  string `json:"category"`
	Venue     string `json:"venue"`
	Date      string `json:"date"`
}

type Address struct {
	Street        string `json:"street" validate:"required"`
	StreetNumber  string `json:"street_number" validate:"required"`
	ZipCode       string `json:"zipcode" validate:"required"`
	Country       string `json:"country" validate:"required"`
	State         string `json:"state" validate:"required"`
	City          string `json:"city" validate:"required"`
	Neighborhood  string `json:"neighborhood"`
	Complementary string `json:"complementary,omitempty"`
}

type Document struct {
	Type   string `json:"type" validate:"required"`
	Number string `json:"number" validate:"required"`
}

type BoletoFine struct {
	Days       int32  `json:"days"`
	Amount     int32  `json:"amount"`
	Percentage string `json:"percentage"`
}

type BoletoInterest struct {
	Days       int32  `json:"days"`
	Amount     int32  `json:"amount"`
	Percentage string `json:"percentage"`
}

type TransactionUsecase interface {
	Insert(string, *Transaction) (map[string]interface{}, error)
}

type TransactionRepository interface {
	Insert(context.Context, string, *Transaction) (map[string]interface{}, error)
}
