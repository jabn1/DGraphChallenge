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
func QueryDateExists(timestamp int64) bool {
	client := newClient()
	if client == nil {
		return false
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
		log.Fatal(err)
	}

	var decode struct {
		Date []BusinessDay
	}

	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		log.Fatal(err)
	}
	if len(decode.Date) != 0 {
		return true
	}
	return false
}

//WriteBusinessDay performs a DGraph mutation to load the data containing buyers, products and transactions pertaining to a day
func WriteBusinessDay(timestamp int64) bool {
	client := newClient()
	if client == nil {
		return false
	}
	txn := client.NewTxn()
	defer txn.Discard(context.Background())

	dayData := GetDayData(timestamp)
	if dayData == nil {
		return false
	}
	dayDataJSON, err := json.Marshal(dayData)
	if err != nil {
		log.Println(err)
		return false
	}

	queryS := fmt.Sprintf(`{ v as var(func: eq(date.value,"%v")) }`, dayData.Date.Value)
	condS := "@if(eq(len(v),0))"

	dgMutation := api.Mutation{SetJson: dayDataJSON, Cond: condS}

	dgRequest := api.Request{Mutations: []*api.Mutation{&dgMutation}, Query: queryS}

	response, err := txn.Do(context.Background(), &dgRequest)
	if err != nil {
		log.Println(err)
		return false
	}

	if len(response.Uids) == 0 {
		log.Println(fmt.Sprintf("Error in WriteBusinessDay(%d), the requested date already exists", timestamp))
		return false
	}

	err = txn.Commit(context.Background())
	return true
}

func QueryBuyerList() {

}

func QueryBuyerData() {

}
