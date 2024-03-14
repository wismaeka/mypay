// Package database provides database access
package database

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PgxDB is something that provides DB querying access.
type PgxDB interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...any) pgx.Row
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

// New provides a new db pool.
func New(host, dbName, user, password string, port int) (*pgxpool.Pool, error) {
	u := assembleURL(host, dbName, user, password, port)

	pool, err := pgxpool.New(context.Background(), u)
	if err != nil {
		return nil, fmt.Errorf("new db pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("pinging db: %v", err)
	}

	return pool, err
}

func assembleURL(host, name, user, password string, port int) string {
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, password),
		Host:     fmt.Sprintf("%v:%v", host, port),
		Path:     name,
		RawQuery: q.Encode(),
	}

	return u.String()
}
