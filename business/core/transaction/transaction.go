// Package transaction provides the use-case of transaction.
package transaction

import (
	"context"

	"mypayment/business/data/external"
	tfData "mypayment/business/data/transaction"
)

type transferStore interface {
	InsertTransfer(ctx context.Context, tf tfData.RawTransfer) (tfData.Transfer, error)
	UpdateStatus(ctx context.Context, id, status string) (tfData.StatusUpdate, error)
}

type callBank interface {
	Transaction(ctx context.Context, txn external.Transfer) error
}

type Core struct {
	ts transferStore
	cb callBank
}

// New returns a new Core.
func New(ts transferStore, cb callBank) Core {
	return Core{
		ts: ts,
		cb: cb,
	}
}
