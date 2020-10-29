package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
)

type transaction struct {
	ID         string   `json:"id"`
	BuyerID    string   `json:"buyerId"`
	IP         string   `json:"ip"`
	Device     string   `json:"device"`
	ProductIds []string `json:"productIds"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func removeEmpties(input []string) []string {
	var output []string
	for _, element := range input {
		if element != "" {
			output = append(output, element)
		}
	}
	return output
}

func main() {
	file, err := ioutil.ReadFile("transactions.txt")
	check(err)
	transactionStrings := strings.Split(string(file), "#")
	transactionStrings = removeEmpties(transactionStrings)

	var transactions []transaction

	for _, transactionString := range transactionStrings {
		ts := strings.Replace(transactionString, "$$", "", -1)
		values := strings.Split(ts, "$")
		prodString := values[4][1 : len(values[4])-1]
		productIds := strings.Split(prodString, ",")

		tran := transaction{ID: values[0], BuyerID: values[1], IP: values[2], Device: values[3], ProductIds: productIds}
		transactions = append(transactions, tran)
	}

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	encoder.Encode(transactions)

	err = ioutil.WriteFile("transactions.json", buffer.Bytes(), 0644)
	check(err)

}
