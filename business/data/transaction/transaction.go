// Package transaction contains everything needed to manage transaction.
package transaction

import database "mypayment/business/sys"

// Store contains everything needed to manage transaction.
type Store struct {
	db database.PgxDB
}

// New returns new Store.
func New(db database.PgxDB) Store {
	return Store{
		db: db,
	}
}
