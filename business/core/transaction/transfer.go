package transaction

import (
	"context"
	"time"

	"mypayment/business/data/external"
	"mypayment/business/data/transaction"
)

type RawTransfer struct {
	SenderName         string
	SenderBankNumber   string
	SenderBank         string
	ReceiverName       string
	ReceiverBankNumber string
	ReceiverBank       string
	Amount             string
	Currency           string
	Status             string
	Description        string
}

type Transfer struct {
	ID                 string
	SenderName         string
	SenderBankNumber   string
	SenderBank         string
	ReceiverName       string
	ReceiverBankNumber string
	ReceiverBank       string
	Amount             string
	Currency           string
	Status             string
	Description        string
	Created            time.Time
	Updated            time.Time
}

type BankResponse struct {
	// ID is the bank transaction ID
	ID             string
	MyID           string
	SenderName     string
	SenderNumber   string
	SenderBank     string
	ReceiverName   string
	ReceiverNumber string
	ReceiverBank   string
	Amount         string
	Description    string
	Status         string
}

type StatusUpdate struct {
	ID      string
	Status  string
	Updated time.Time
}

func (c Core) InsertTransfer(ctx context.Context, rt RawTransfer) (Transfer, error) {
	tf, err := c.insertTransfer(ctx, rt)
	if err != nil {
		return Transfer{}, err
	}
	err = c.callBank(ctx, tf)
	if err != nil {
		return Transfer{}, err
	}
	return Transfer{
		ID:                 tf.ID,
		SenderName:         tf.SenderName,
		SenderBankNumber:   tf.SenderBankNumber,
		SenderBank:         tf.SenderBank,
		ReceiverName:       tf.ReceiverName,
		ReceiverBankNumber: tf.ReceiverBankNumber,
		ReceiverBank:       tf.ReceiverBank,
		Amount:             tf.Amount,
		Currency:           tf.Currency,
		Status:             "PENDING",
		Description:        tf.Description,
		Created:            tf.Created,
		Updated:            tf.Updated,
	}, nil
}

func (c Core) insertTransfer(ctx context.Context, rt RawTransfer) (Transfer, error) {
	tf, err := c.ts.InsertTransfer(ctx, transaction.RawTransfer{
		SenderName:         rt.SenderName,
		SenderBankNumber:   rt.SenderBankNumber,
		SenderBank:         rt.SenderBank,
		ReceiverName:       rt.ReceiverName,
		ReceiverBankNumber: rt.ReceiverBankNumber,
		ReceiverBank:       rt.ReceiverBank,
		Amount:             rt.Amount,
		Currency:           rt.Currency,
		Status:             "PENDING",
		Description:        rt.Description,
	})
	if err != nil {
		return Transfer{}, err
	}
	return Transfer{
		ID:                 tf.ID,
		SenderName:         tf.SenderName,
		SenderBankNumber:   tf.SenderBankNumber,
		SenderBank:         tf.SenderBank,
		ReceiverName:       tf.ReceiverName,
		ReceiverBankNumber: tf.ReceiverBankNumber,
		ReceiverBank:       tf.ReceiverBank,
		Amount:             tf.Amount,
		Currency:           tf.Currency,
		Status:             tf.Status,
		Description:        tf.Description,
		Created:            tf.Created,
		Updated:            tf.Updated,
	}, nil
}

func (c Core) callBank(ctx context.Context, tf Transfer) error {
	err := c.cb.Transaction(ctx, external.Transfer{
		MyID:           tf.ID,
		SenderName:     tf.SenderName,
		SenderNumber:   tf.SenderBankNumber,
		SenderBank:     tf.SenderBank,
		ReceiverName:   tf.ReceiverName,
		ReceiverNumber: tf.ReceiverBankNumber,
		ReceiverBank:   tf.ReceiverBank,
		Amount:         tf.Amount,
		Description:    tf.Description,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c Core) UpdateStatus(ctx context.Context, id, status string) (StatusUpdate, error) {
	st, err := c.ts.UpdateStatus(ctx, id, status)
	if err != nil {
		return StatusUpdate{}, err
	}
	return StatusUpdate{
		ID:      st.ID,
		Status:  st.Status,
		Updated: st.Updated,
	}, nil
}
