package payload

import (
	"time"

	"stone/transaction-challenge/domain"
)

type InsertTransactionPayload struct {
	AuthType          string             `json:"auth_type"`
	TransationPayload TransactionPayload `json:"transaction"`
}

type TransactionPayload struct {
	Amount               int32                 `json:"amount"`
	CardHash             string                `json:"card_hash"`
	CardId               string                `json:"card_id"`
	CardHolderName       string                `json:"card_holder_name"`
	CardExpirationDate   string                `json:"card_expiration_date"`
	CardNumber           string                `json:"card_number"`
	CardCvv              string                `json:"card_cvv"`
	PaymentMethod        string                `json:"payment_method"`
	PostbackUrl          string                `json:"postback_url"`
	Installments         string                `json:"installments"`
	Capture              string                `json:"capture"`
	BoletoExpirationDate string                `json:"boleto_expiration_date"`
	SoftDescriptor       string                `json:"soft_descriptor"`
	Customer             domain.Customer       `json:"customer"`
	Billing              domain.Billing        `json:"billing"`
	Shipping             domain.Shipping       `json:"shipping"`
	Items                []domain.Item         `json:"items"`
	Metadata             interface{}           `json:"metadata"`
	SplitRules           []interface{}         `json:"split_rules"`
	BoletoFine           domain.BoletoFine     `json:"boleto_fine"`
	BoletoInterest       domain.BoletoInterest `json:"boleto_interest"`
	BoletoRules          []string              `json:"boleto_rules"`
	Session              string                `json:"session"`
	LocalTime            time.Time             `json:"local_time"`
}

func (p TransactionPayload) ToEntity() domain.Transaction {
	return domain.Transaction{
		Amount:               p.Amount,
		CardHash:             p.CardHash,
		CardId:               p.CardId,
		CardHolderName:       p.CardHolderName,
		CardExpirationDate:   p.CardExpirationDate,
		CardNumber:           p.CardNumber,
		CardCvv:              p.CardCvv,
		PaymentMethod:        p.PaymentMethod,
		PostbackUrl:          p.PostbackUrl,
		Installments:         p.Installments,
		BoletoExpirationDate: p.BoletoExpirationDate,
		SoftDescriptor:       p.SoftDescriptor,
		Customer:             p.Customer,
		Billing:              p.Billing,
		Shipping:             p.Shipping,
		Items:                p.Items,
		Metadata:             p.Metadata,
		SplitRules:           p.SplitRules,
		BoletoFine:           p.BoletoFine,
		BoletoInterest:       p.BoletoInterest,
		BoletoRules:          p.BoletoRules,
		Session:              p.Session,
		LocalTime:            p.LocalTime,
	}
}
