package data

//this group of functions is responsible for consuming the source data API
//processing the source data and returning it in a usable form

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type rawTransaction struct {
	ID         string   `json:"id"`
	BuyerID    string   `json:"buyerId"`
	IP         string   `json:"ip"`
	Device     string   `json:"device"`
	ProductIds []string `json:"productIds"`
}

var baseAPIURL = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com"

//GetDayData returns the struct with the necessary structure to load the data into the database
func GetDayData(timestamp int64) *BusinessDay {
	var businessDay BusinessDay = BusinessDay{}
	products := getDataProducts(timestamp)
	buyers := getDataBuyers(timestamp)
	rawTran := getDataTransactions(timestamp)
	if products == nil || buyers == nil || rawTran == nil {
		return nil
	}

	businessDay.DayProducts = *products
	businessDay.DayBuyers = *buyers

	dayTime := time.Unix(timestamp, 0)
	businessDay.Year = dayTime.Year()
	businessDay.Month = int(dayTime.Month())
	businessDay.Day = dayTime.Day()

	buyersMap := make(map[string]Entity)   //key: id, value: uid
	productsMap := make(map[string]Entity) //key: id, value: uid

	for i := 0; i < len(*buyers); i++ {
		uid := "buyer" + string(i)
		(*buyers)[i].UID = uid
		buyersMap[(*buyers)[i].ID] = Entity{UID: uid}
	}

	for i := 0; i < len(*products); i++ {
		uid := "product" + string(i)
		(*products)[i].UID = uid
		productsMap[(*products)[i].ID] = Entity{UID: uid}
	}

	return &businessDay
}

func getDataProducts(timestamp int64) *[]Product {
	response, err := http.Get(baseAPIURL + "/products?date=" + string(timestamp))
	if err != nil {
		log.Println(err)
		return nil
	}

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanLines)
	var prodStrings []string

	for scanner.Scan() {
		prodStrings = append(prodStrings, scanner.Text())
	}

	var products []Product

	for _, prodString := range prodStrings {
		temp := strings.Replace(prodString, "\"", "", -1)
		prodValues := strings.Split(temp, "'")
		prodPrice, err := strconv.Atoi(prodValues[len(prodValues)-1])
		if err != nil {
			log.Println(err)
			return nil
		}
		prodName := strings.Join(prodValues[1:len(prodValues)-1], "'")
		prod := Product{ID: prodValues[0], Name: prodName, Price: prodPrice}
		products = append(products, prod)
	}

	return &products
}

func getDataBuyers(timestamp int64) *[]Buyer {
	response, err := http.Get(baseAPIURL + "/buyers?date=" + string(timestamp))
	if err != nil {
		log.Println(err)
		return nil
	}
	buyersBytes, _ := ioutil.ReadAll(response.Body)
	var buyers []Buyer
	json.Unmarshal(buyersBytes, &buyers)
	return &buyers
}

func getDataTransactions(timestamp int64) *[]rawTransaction {
	response, err := http.Get(baseAPIURL + "/transactions?date=" + string(timestamp))
	if err != nil {
		log.Println(err)
		return nil
	}

	transactionsBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	transactionsBytes = replaceInvalid(transactionsBytes)
	transactionStrings := strings.Split(string(transactionsBytes), "#")
	transactionStrings = removeEmpties(transactionStrings)

	var transactions []rawTransaction

	for _, transactionString := range transactionStrings {
		ts := strings.Replace(transactionString, "$$", "", -1)
		values := strings.Split(ts, "$")
		prodString := values[4][1 : len(values[4])-1]
		productIds := strings.Split(prodString, ",")

		tran := rawTransaction{ID: values[0], BuyerID: values[1], IP: values[2], Device: values[3], ProductIds: productIds}
		transactions = append(transactions, tran)
	}
	return &transactions
}

//for transaction processing
func replaceInvalid(input []byte) []byte {
	var output []byte
	for _, b := range input {
		if b == 0 {
			output = append(output, byte('$'))
		} else {
			output = append(output, b)
		}
	}
	return output
}

//for transaction processing
func removeEmpties(input []string) []string {
	var output []string
	for _, element := range input {
		if element != "" {
			output = append(output, element)
		}
	}
	return output
}
