package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

type Customer struct {
	gorm.Model
	CustomerId string
	FullName   string
}

type Accounts struct {
	gorm.Model
	AccountNumber int
	CustomerId    string
	Balance       int
}

type Transactions struct {
	gorm.Model
	TransactionsId   int
	TransactionsType string
	CustomerId       string
	Amount           int
	Timestamp        string
}

var db *gorm.DB
var err error

var (
	customer = []Customer{
		{CustomerId: "1", FullName: "Gerza"},
		{CustomerId: "2", FullName: "Gerza1"},
		{CustomerId: "3", FullName: "Gerza2"},
	}
	accounts = []Accounts{
		{AccountNumber: 2001, CustomerId: "2", Balance: 200.0},
		{AccountNumber: 2002, CustomerId: "3", Balance: 300.0},
		{AccountNumber: 2003, CustomerId: "4", Balance: 400.0},
	}
	transaction = []Transactions{
		{TransactionsId: 1111, TransactionsType: "deposit", CustomerId: "1", Amount: 100.0, Timestamp: "2006-01-02T15:04:05-0700"},
		{TransactionsId: 2222, TransactionsType: "deposit", CustomerId: "2", Amount: 200.0, Timestamp: "2006-01-02T15:04:05-0700"},
		{TransactionsId: 3333, TransactionsType: "withdraw", CustomerId: "3", Amount: 300.0, Timestamp: "2006-01-02T15:04:05-0700"},
		{TransactionsId: 4444, TransactionsType: "withdraw", CustomerId: "4", Amount: 400.0, Timestamp: "2006-01-02T15:04:05-0700"},
	}
)

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=example")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Transactions{})
	for index := range customer {
		db.Create(&customer[index])
	}

	for index := range accounts {
		db.Create(&accounts[index])
	}
	for index := range transaction {
		db.Create(&transaction[index])
	}

	router.HandleFunc("/customer", GetCustomer).Methods("GET")
	router.HandleFunc("/account/{customerId}", GetAccount).Methods("GET")
	router.HandleFunc("/transaction/{transactionsType}", GetTransaction).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8081", handler))
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	var customer []Customer
	db.Find(&customer)
	json.NewEncoder(w).Encode(&customer)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var accounts Accounts
	db.First(&accounts, params["customerId"])
	json.NewEncoder(w).Encode(&accounts)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transactions Transactions
	db.Where("transactions_type = ?", params["transactionsType"]).First(&transactions)
	json.NewEncoder(w).Encode(&transactions)
}
