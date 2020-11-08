package main

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/go-chi/chi"
	restaurant "restaurant.com"
)

func main() {
	//_ = restaurant.Entity{UID: "as"} //REMOVE for debugging

	//fmt.Println(restaurant.WriteBusinessDay(1603843200))

	bl := restaurant.QueryBuyerList(10, 20)
	for _, b := range *bl {
		fmt.Println("Id: " + b.ID + " - " + "Name: " + b.Name + " - " + "Age: " + strconv.Itoa(b.Age))
	}

	// buffer := new(bytes.Buffer)
	// encoder := json.NewEncoder(buffer)
	// encoder.SetEscapeHTML(false)
	// encoder.SetIndent("", "  ")

	// encoder.Encode(restaurant.GetDayData(1603843200))

	// _ = ioutil.WriteFile("day.json", buffer.Bytes(), 0644)

}

//route declarations
func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/load?date={timestamp}", loadData) //POST /load?date=16050000
		r.Get("/buyers", getBuyers)                //GET /buyers
		r.Get("/buyer?id={id}", getBuyer)          //GET /buyer?id=234e2f
	})
	return r
}

func loadData(w http.ResponseWriter, r *http.Request) {

}

func getBuyers(w http.ResponseWriter, r *http.Request) {

}

func getBuyer(w http.ResponseWriter, r *http.Request) {

}
