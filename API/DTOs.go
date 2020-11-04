package main

type Status struct {
    Success bool `json:"success"`
}

type BuyersDTO struct {
    Buyers []BuyerDTO `json:"buyers"`
}

type BuyerDTO struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Age string `json:"age"`
}

type BuyerHistoryDTO {
    Buyer BuyerDTO `json:"buyer"`
    Transactions []TransactionDTO `json:"transactions"`
    OtherBuyers []BuyerDTO `json:"otherbuyers"`
    RecommendedProducts []ProductDTO `json:"recommendedproducts"`
}

type TransactionDTO {
    ID string `json:"id"`
    IP string `json:"ip"`
    Device string `json:"device"`
    Year int `json:"year"`
    Month int `json:"month"`
    Day int `json:"day"`
    Products []ProductDTO `json:"products"`
}

type ProductDTO {
    ID string `json:"id"`
    Name string `json:"name"`
    Price int `json:"price"`
}


