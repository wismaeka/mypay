package transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// RawTransfer represents pre insertion data of Transfer.
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

type StatusUpdate struct {
	ID      string
	Status  string
	Updated time.Time
}

func (s Store) InsertTransfer(ctx context.Context, tf RawTransfer) (Transfer, error) {
	rows, err := s.db.Query(ctx, `INSERT INTO transfers (sender_name, sender_number, sender_bank, receiver_name, receiver_number, receiver_bank, amount, currency, status, description) 
VALUES ($1, $2, $3, $4, $5,$6,$7,$8,$9, $10) RETURNING 
    id, sender_name, sender_number, sender_bank, receiver_name, receiver_number, receiver_bank, amount, currency, status, description, created, updated`,
		tf.SenderName,
		tf.SenderBankNumber,
		tf.SenderBank,
		tf.ReceiverName,
		tf.ReceiverBankNumber,
		tf.ReceiverBank,
		tf.Amount,
		tf.Currency,
		tf.Status,
		tf.Description)

	if err != nil {
		return Transfer{}, fmt.Errorf("inserting: %v", err)
	}

	defer rows.Close()

	t, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[Transfer])
	if err != nil {
		return Transfer{}, fmt.Errorf("collecting: %v", err)
	}
	return *t, nil
}

func (s Store) UpdateStatus(ctx context.Context, id, status string) (StatusUpdate, error) {
	rows, err := s.db.Query(ctx, `UPDATE transfers SET status = $1 WHERE id = $2 RETURNING id, status, updated`, status, id)
	if err != nil {
		return StatusUpdate{}, fmt.Errorf("updating status: %v", err)
	}
	defer rows.Close()

	st, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[StatusUpdate])
	if err != nil {
		return StatusUpdate{}, fmt.Errorf("collecting: %v", err)
	}
	return *st, nil
}
