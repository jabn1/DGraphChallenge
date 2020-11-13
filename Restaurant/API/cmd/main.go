package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"net/http"

	"github.com/go-chi/chi"
	restaurant "restaurant.com"
)

func main() {
	port := "5000"
	r := registerRoutes()
	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))

	//_ = restaurant.Entity{UID: "as"} //REMOVE for debugging

	//fmt.Println(restaurant.WriteBusinessDay(1603843200))

	// bl := restaurant.QueryBuyerList(10, 20)
	// for _, b := range *bl {
	// 	fmt.Println("Id: " + b.ID + " - " + "Name: " + b.Name + " - " + "Age: " + strconv.Itoa(b.Age))
	// }

	// buffer := new(bytes.Buffer)
	// encoder := json.NewEncoder(buffer)
	// encoder.SetEscapeHTML(false)
	// encoder.SetIndent("", "  ")

	// encoder.Encode(restaurant.QueryBuyerData("733fb35a"))
	// fmt.Println(string(buffer.Bytes()))
	// _ = ioutil.WriteFile("day.json", buffer.Bytes(), 0644)

}

//route declarations
func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/load", loadData)   //POST /load or /load?date=16050000
		r.Get("/buyers", getBuyers) //GET /buyers?first=50offset=100
		r.Get("/buyer", getBuyer)   //GET /buyer?id=234e2f
	})
	return r
}

func loadData(w http.ResponseWriter, r *http.Request) {
	timestampstring := r.URL.Query().Get("date")
	var timestamp int64
	if timestampstring == "" {
		timestamp = time.Now().Unix()
	} else {
		var err error
		timestamp, err = strconv.ParseInt(timestampstring, 10, 64)
		if err != nil {
			http.Error(w, "Invalid query parameters", 400)
			return
		}

		status := restaurant.WriteBusinessDay(timestamp)
		if status == nil {
			http.Error(w, "Error while processing request", 500)
			return
		}
		json.NewEncoder(w).Encode(status)

	}

}

func getBuyers(w http.ResponseWriter, r *http.Request) {
	firstString := r.URL.Query().Get("first")
	offsetString := r.URL.Query().Get("offset")

	if firstString == "" || offsetString == "" {
		http.Error(w, "Invalid query parameters", 400)
		return
	}
	first, err := strconv.Atoi(firstString)
	if err != nil {
		http.Error(w, "Invalid query parameters", 400)
		return
	}
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		http.Error(w, "Invalid query parameters", 400)
		return
	}

	buyers := restaurant.QueryBuyerList(first, offset)

	if buyers == nil {
		http.Error(w, "Error while processing request", 500)
		return
	}
	json.NewEncoder(w).Encode(buyers)

}

func getBuyer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Invalid query parameters", 400)
		return
	}
	buyerdata := restaurant.QueryBuyerData(id)
	if buyerdata == nil {
		http.Error(w, "Error while processing request", 500)
		return
	}
	json.NewEncoder(w).Encode(buyerdata)
}
