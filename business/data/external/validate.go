package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Account struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	BankNumber string `json:"bank_number"`
	IsValid    bool   `json:"is_valid"`
}

// Validate returns account information for a given account number and name.
func (b BankCall) Validate(ctx context.Context, number, name string) (Account, error) {

	params := url.Values{}
	params.Set("name", name)
	params.Set("bank_number", number)

	u, err := b.baseURL.Parse(fmt.Sprintf("/mock/bank/validate?%v", params.Encode()))
	if err != nil {
		return Account{}, fmt.Errorf("invalid url: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return Account{}, fmt.Errorf("new request: %v", err)
	}
	resp, err := b.client.Do(req)
	if err != nil {
		return Account{}, fmt.Errorf("do request: %v", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return Account{}, ErrNotFound
	default:
		return Account{}, fmt.Errorf("status code: %v", resp.StatusCode)
	}

	var result []Account
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Account{}, fmt.Errorf("decoding: %v", err)
	}

	return result[0], nil
}
