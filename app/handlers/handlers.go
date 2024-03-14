package handlers

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"mypayment/app/handlers/transaction"
	"mypayment/app/handlers/validate"
	tfCore "mypayment/business/core/transaction"
	valData "mypayment/business/data/external"
	tfData "mypayment/business/data/transaction"
	"mypayment/foundation/web"
)

type Config struct {
	Client  *http.Client
	BankURL *url.URL
	DB      *pgxpool.Pool
}

func Init(c Config) http.Handler {
	router := mux.NewRouter()

	app := web.NewApp(router)
	routes(app, c)

	return app
}

func routes(app *web.App, c Config) {
	validateHandler := validate.NewHandler(valData.New(c.Client, c.BankURL))
	transactionHandler := transaction.NewHandler(tfCore.New(
		tfData.New(c.DB),
		valData.New(c.Client, c.BankURL),
	))

	app.Handle(http.MethodGet, "/validation/accounts", validateHandler.Account)
	app.Handle(http.MethodPost, "/transfers", transactionHandler.Transfer)
	app.Handle(http.MethodPost, "/transfers/status", transactionHandler.UpdateStatus)
}
