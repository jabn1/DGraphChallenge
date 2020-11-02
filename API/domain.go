package domain

import(
	"time"
)

type transaction struct {
	ID         string   `json:"id"`
	BuyerID    string   `json:"buyerId"`
	IP         string   `json:"ip"`
	Device     string   `json:"device"`
	ProductIds []string `json:"productIds"`
	Date time.Time `json:"date"`
}

type product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Date time.Time `json:"date"`
}

type buyer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age int `json:"age"`
	Date time.Time `json:"date"`
}