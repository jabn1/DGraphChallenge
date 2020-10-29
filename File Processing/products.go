package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("products.csv")

	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var prodStrings []string

	for scanner.Scan() {
		prodStrings = append(prodStrings, scanner.Text())
	}

	file.Close()

	var products []product

	for _, prodString := range prodStrings {
		temp := strings.Replace(prodString, "\"", "", -1)
		prodValues := strings.Split(temp, "'")
		prodPrice, err := strconv.Atoi(prodValues[len(prodValues)-1])
		check(err)
		prodName := strings.Join(prodValues[1:len(prodValues)-1], "'")
		prod := product{ID: prodValues[0], Name: prodName, Price: prodPrice}
		products = append(products, prod)
	}

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	encoder.Encode(products)

	err = ioutil.WriteFile("products.json", buffer.Bytes(), 0644)
	check(err)

}
