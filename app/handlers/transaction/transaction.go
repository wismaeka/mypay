// Package transaction provides the service for transaction management.
package transaction

import (
	"context"
	"time"

	txCore "mypayment/business/core/transaction"
)

type store interface {
	InsertTransfer(ctx context.Context, tf txCore.RawTransfer) (txCore.Transfer, error)
	UpdateStatus(ctx context.Context, id, status string) (txCore.StatusUpdate, error)
}

type Handler struct {
	s store
}

func NewHandler(s store) Handler {
	return Handler{
		s: s,
	}
}

type transferReq struct {
	SenderName         string `json:"sender_name"`
	SenderBankNumber   string `json:"sender_bank_number"`
	SenderBank         string `json:"sender_bank"`
	ReceiverName       string `json:"receiver_name"`
	ReceiverBankNumber string `json:"receiver_bank_number"`
	ReceiverBank       string `json:"receiver_bank"`
	Amount             string `json:"amount"`
	Currency           string `json:"currency"`
	Description        string `json:"description"`
}

type transferResponse struct {
	ID                 string    `json:"id"`
	SenderName         string    `json:"sender_name"`
	SenderBankNumber   string    `json:"sender_bank_number"`
	SenderBank         string    `json:"sender_bank"`
	ReceiverName       string    `json:"receiver_name"`
	ReceiverBankNumber string    `json:"receiver_bank_number"`
	ReceiverBank       string    `json:"receiver_bank"`
	Amount             string    `json:"amount"`
	Currency           string    `json:"currency"`
	Status             string    `json:"status"`
	Description        string    `json:"description"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
}
