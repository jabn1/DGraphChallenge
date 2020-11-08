package data

//These structs are used to insert the data into the DGraph database

type Entity struct {
	UID string `json:"uid"`
}

type Date struct {
	UID   string `json:"uid"`
	Value string `json:"date.value"`
}

type Product struct {
	UID   string `json:"uid"`
	ID    string `json:"product.id"`
	Name  string `json:"product.name"`
	Price int    `json:"product.price"`
	Date  Entity `json:"product.date"`
}

type Buyer struct {
	UID          string   `json:"uid"`
	ID           string   `json:"buyer.id"`
	Name         string   `json:"buyer.name"`
	Age          int      `json:"buyer.age"`
	Date         Entity   `json:"buyer.date"`
	Transactions []Entity `json:"buyer.transactions"`
}

type Transaction struct {
	UID      string   `json:"uid"`
	ID       string   `json:"transaction.id"`
	Buyer    Entity   `json:"transaction.buyer"`
	IP       string   `json:"transaction.ip"`
	Device   string   `json:"transaction.device"`
	Products []Entity `json:"transaction.products"`
	Date     Entity   `json:"transaction.date"`
}

type BusinessDay struct {
	Date         Date          `json:"date"`
	Products     []Product     `json:"products"`
	Buyers       []Buyer       `json:"buyers"`
	Transactions []Transaction `json:"transactions"`
}
