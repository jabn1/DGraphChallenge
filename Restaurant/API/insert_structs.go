package data

//These structs are used to insert the data into the DGraph database


type Entity struct {
	UID string `json:"uid"`
}

type Product struct {
    UID string `json:"uid"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Buyer struct {
    UID string `json:"uid"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Transaction struct {
	ID       string   `json:"id"`
	Buyer    Entity   `json:"buyer"`
	IP       string   `json:"ip"`
	Device   string   `json:"device"`
	Products []Entity `json:"products"`
}

type BusinessDay struct {
	Year            int           `json:"year"`
	Month           int           `json:"month"`
	Day             int           `json:"day"`
	DayProducts     []Product     `json:"dayproducts"`
	DayBuyers       []Buyer       `json:"daybuyers"`
	DayTransactions []Transaction `json:"daytransactions"`
}
