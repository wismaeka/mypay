package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"mypayment/app/handlers"
	database "mypayment/business/sys"
)

type configuration struct {
	Port int `envconfig:"PORT" default:"3002"`

	DBHost     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	DBPort     int    `envconfig:"POSTGRES_PORT" default:"5451"`
	DBName     string `envconfig:"POSTGRES_DATABASE" default:"mypayment"`
	DBUser     string `envconfig:"POSTGRES_USER" default:"user"`
	DBPassword string `envconfig:"POSTGRES_PASSWORD" default:"password"`

	BankURL string `envconfig:"BANK_URL" default:"https://65f19893034bdbecc7631ed1.mockapi.io"`
}

func main() {
	_ = godotenv.Load()
	var conf configuration

	if err := envconfig.Process("", &conf); err != nil {
		log.Println("reading env:", err)
		os.Exit(1)
	}

	db, err := database.New(conf.DBHost, conf.DBName, conf.DBUser, conf.DBPassword, conf.DBPort)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	defer func() {
		fmt.Println("closing DB")
		db.Close()
	}()

	bankURL, err := url.Parse(conf.BankURL)
	if err != nil {
		fmt.Println("Parsing URL error")
		os.Exit(2)
	}

	server := initServer(conf, db, bankURL)

	fmt.Println("Server is running on port: ", conf.Port)

	server.ListenAndServe()
}

func initServer(
	conf configuration,
	db *pgxpool.Pool,
	bankURL *url.URL,
) *http.Server {
	mux := handlers.Init(handlers.Config{
		Client:  http.DefaultClient,
		BankURL: bankURL,
		DB:      db,
	})

	return &http.Server{
		Addr:        fmt.Sprintf(":%v", conf.Port),
		IdleTimeout: 181 * time.Second,
		Handler:     mux,
	}
}
