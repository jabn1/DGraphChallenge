package data

//this group of functions is responsible for performing queries and mutations in the DGraph database

import (
	"context"
	"flag"
	"fmt"
	"log"
    
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
    
	"google.golang.org/grpc"
)

func SendQuery(query string) {
    var dgraph = flag.String("d", "127.0.0.1:9080", "Dgraph Alpha address")
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
    
	resp, err := dg.NewTxn().Query(context.Background(), query)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)
}


func WriteBusinessDay() {

}

func ReadBuyerList() {

}

func ReadBuyerData() {

}
