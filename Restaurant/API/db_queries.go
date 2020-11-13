package data

//this group of functions is responsible for performing queries and mutations in the DGraph database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"

	"google.golang.org/grpc"
)

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

//QueryDateExists is used to check if a day is already loaded into the database
func QueryDateExists(timestamp int64) *DateExistsDTO {
	client := newClient()
	if client == nil {
		return nil
	}
	txn := client.NewTxn()
	defer txn.Discard(context.Background())

	query := `
	{
		day(func: eq(date.value,"%v")) {
			date
		}
	}
	`
	query = fmt.Sprintf(query, UnixToDate(timestamp))

	resp, err := txn.Query(context.Background(), query)
	if err != nil {
		log.Println(err)
		return nil
	}

	var decode struct {
		Date []BusinessDay
	}

	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		log.Println(err)
		return nil
	}
	if len(decode.Date) != 0 {
		return &DateExistsDTO{Exists: true}
	}
	return &DateExistsDTO{Exists: false}
}

//WriteBusinessDay performs a DGraph mutation to load the data containing buyers, products and transactions pertaining to a day
func WriteBusinessDay(timestamp int64) *Status {
	client := newClient()
	if client == nil {
		return nil
	}
	txn := client.NewTxn()
	defer txn.Discard(context.Background())

	dayData := GetDayData(timestamp)
	if dayData == nil {
		return nil
	}
	dayDataJSON, err := json.Marshal(dayData)
	if err != nil {
		log.Println(err)
		return nil
	}

	queryS := fmt.Sprintf(`{ v as var(func: eq(date.value,"%v")) }`, dayData.Date.Value)
	condS := "@if(eq(len(v),0))"

	dgMutation := api.Mutation{SetJson: dayDataJSON, Cond: condS}

	dgRequest := api.Request{Mutations: []*api.Mutation{&dgMutation}, Query: queryS}

	response, err := txn.Do(context.Background(), &dgRequest)
	if err != nil {
		log.Println(err)
		return nil
	}

	if len(response.Uids) == 0 {
		log.Println(fmt.Sprintf("Error in WriteBusinessDay(%d), the requested date already exists", timestamp))
		return &Status{Success: false}
	}

	err = txn.Commit(context.Background())
	return &Status{Success: true}
}

//QueryBuyerList retunrs a list of all buyers with pagination parameters
func QueryBuyerList(first int, offset int) *[]BuyerDTO {
	client := newClient()
	if client == nil {
		return nil
	}
	txn := client.NewTxn()
	defer txn.Discard(context.Background())

	query := `
	{
		buyers(func: has(buyer.id),first: %d,offset: %d) {   
			buyer.id 
			buyer.name 
			buyer.age
		}
	}
	`
	query = fmt.Sprintf(query, first, offset)

	resp, err := txn.Query(context.Background(), query)
	if err != nil {
		log.Println(err)
		return nil
	}

	var decode struct {
		Buyers []struct {
			ID   string `json:"buyer.id"`
			Name string `json:"buyer.name"`
			Age  int    `json:"buyer.age"`
		} `json:"buyers"`
	}

	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		log.Println(err)
		return nil
	}

	var buyers []BuyerDTO
	for _, buyer := range decode.Buyers {
		buyers = append(buyers, BuyerDTO{ID: buyer.ID, Name: buyer.Name, Age: buyer.Age})
	}
	return &buyers
}

//QueryBuyerData returns the buyer history of the buyer with the corresponding entered id
func QueryBuyerData(ID string) *BuyerHistoryDTO {
	client := newClient()
	if client == nil {
		return nil
	}
	txn := client.NewTxn()
	defer txn.Discard(context.Background())

	query := `{
		buyer(func: eq(buyer.id,"%v")) {   
		  ID AS buyer.id 
		  buyer.name 
		  buyer.age
		  buyer.transactions{
			transaction.id
			IP AS transaction.ip
			transaction.device
			transaction.products{
				product.id
				product.name
				product.price
			}
			transaction.date{
				date.value
			}
					  
			}
		}
		otherbuyers(func: has(buyer.id)) {
			buyer.id @filter(not eq(buyer.id,val(ID)))
			buyer.name
			buyer.age
			buyer.transactions @filter(eq(transaction.ip,val(IP)))
		}
		  recproducts(func: has(transaction.id),first: 5){
			transaction.buyer @filter(not eq(buyer.id,val(ID)))
		  transaction.products (first:1){
			product.id
			product.name
			product.price
		  }
		  }
	  }`
	query = fmt.Sprintf(query, ID)

	var decode struct {
		Buyer []struct {
			ID           string `json:"buyer.id"`
			Name         string `json:"buyer.name"`
			Age          int    `json:"buyer.age"`
			Transactions []struct {
				ID       string `json:"transaction.id"`
				IP       string `json:"transaction.ip"`
				Device   string `json:"transaction.device"`
				Products []struct {
					ID    string `json:"product.id"`
					Name  string `json:"product.name"`
					Price int    `json:"product.price"`
				} `json:"transaction.products"`
				Date struct {
					Value string `json:"date.value"`
				} `json:"transaction.date"`
			} `json:"buyer.transactions"`
		} `json:"buyer"`
		OtherBUyers []struct {
			ID   string `json:"buyer.id"`
			Name string `json:"buyer.name"`
			Age  int    `json:"buyer.age"`
		} `json:"otherbuyers"`
		RecProducts []struct {
			Products []struct {
				ID    string `json:"product.id"`
				Name  string `json:"product.name"`
				Price int    `json:"product.price"`
			} `json:"transaction.products"`
		} `json:"recproducts"`
	}

	resp, err := txn.Query(context.Background(), query)
	if err != nil {
		log.Println(err)
		return nil
	}

	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		log.Println(err)
		return nil
	}

	//mapping transactions
	var transactions []TransactionDTO
	for _, respbuyer := range decode.Buyer {
		for _, resptrans := range respbuyer.Transactions {
			var products []ProductDTO
			for _, respprod := range resptrans.Products {
				products = append(products, ProductDTO{
					ID:    respprod.ID,
					Name:  respprod.Name,
					Price: respprod.Price,
				})
			}
			transactions = append(transactions, TransactionDTO{
				ID:       resptrans.ID,
				IP:       resptrans.IP,
				Device:   resptrans.Device,
				Date:     resptrans.Date.Value,
				Products: products,
			})
		}
	}

	//mapping other buyers
	var otherbuyers []BuyerDTO
	for _, respotherbuyer := range decode.OtherBUyers {
		otherbuyers = append(otherbuyers, BuyerDTO{
			ID:   respotherbuyer.ID,
			Name: respotherbuyer.Name,
			Age:  respotherbuyer.Age,
		})
	}

	var recproducts []ProductDTO
	for _, resprecproducts := range decode.RecProducts {
		for _, product := range resprecproducts.Products {
			recproducts = append(recproducts, ProductDTO{
				ID:    product.ID,
				Name:  product.Name,
				Price: product.Price,
			})
		}
	}

	buyerHistory := BuyerHistoryDTO{
		Buyer: BuyerDTO{
			ID:   decode.Buyer[0].ID,
			Name: decode.Buyer[0].Name,
			Age:  decode.Buyer[0].Age,
		},
		Transactions:        transactions,
		OtherBuyers:         otherbuyers,
		RecommendedProducts: recproducts,
	}

	return &buyerHistory
}
