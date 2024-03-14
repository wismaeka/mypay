package external

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Transfer struct {
	MyID           string `json:"my_id"`
	SenderName     string `json:"sender_name"`
	SenderNumber   string `json:"sender_number"`
	SenderBank     string `json:"sender_bank"`
	ReceiverName   string `json:"receiver_name"`
	ReceiverNumber string `json:"receiver_number"`
	ReceiverBank   string `json:"receiver_bank"`
	Amount         string `json:"amount"`
	Description    string `json:"description"`
}

type BankResponse struct {
	// ID is the bank transaction ID
	ID             string `json:"id"`
	MyID           string `json:"my_id"`
	SenderName     string `json:"sender_name"`
	SenderNumber   string `json:"sender_number"`
	SenderBank     string `json:"sender_bank"`
	ReceiverName   string `json:"receiver_name"`
	ReceiverNumber string `json:"receiver_number"`
	ReceiverBank   string `json:"receiver_bank"`
	Amount         string `json:"amount"`
	Description    string `json:"description"`
	Status         bool   `json:"status"`
}

func (b BankCall) Transaction(ctx context.Context, txn Transfer) error {
	u, err := b.baseURL.Parse(fmt.Sprint("/mock/bank/transaction"))
	if err != nil {
		return fmt.Errorf("invalid url: %v", err)
	}

	reqBody, err := json.Marshal(txn)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("status code: %v", resp.StatusCode)
	}
	return nil

}
