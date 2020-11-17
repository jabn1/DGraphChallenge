package domain_model

//these are the DTOs for the API functions

type Status struct {
	Success bool `json:"success"`
}

type BuyerDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type BuyerHistoryDTO struct {
	Buyer               BuyerDTO         `json:"buyer,omitempty"`
	Transactions        []TransactionDTO `json:"transactions,omitempty"`
	OtherBuyers         []BuyerDTO       `json:"otherbuyers,omitempty"`
	RecommendedProducts []ProductDTO     `json:"recommendedproducts,omitempty"`
}

type TransactionDTO struct {
	ID       string       `json:"id"`
	IP       string       `json:"ip"`
	Device   string       `json:"device"`
	Date     string       `json:"date"`
	Products []ProductDTO `json:"products"`
}

type ProductDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type DateExistsDTO struct {
	Exists bool `json:"exists"`
}
