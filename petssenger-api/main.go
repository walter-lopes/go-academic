package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Pricing struct {
	Id             string  `json:id`
	City           string  `json:city`
	BaseFee        float64 `json:basefee`
	PricePerMinute float64 `json:PricePerMinute`
	ServiceFee     float64 `json:ServiceFee`
	PricePerKm     float64 `json:PricePerKm`
}

var Pricings []Pricing

func (pricing *Pricing) Calc(distance float64, minutes float64) float64 {
	total := pricing.BaseFee + ((pricing.PricePerMinute * minutes) + (pricing.PricePerKm*distance)*1.0) + pricing.ServiceFee
	return total
}

func findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	city := vars["city"]
	distance, _ := strconv.ParseFloat(vars["distance"], 64)
	minutes, _ := strconv.ParseFloat(vars["minutes"], 64)

	for _, pricing := range Pricings {
		if pricing.City == city {
			total := pricing.Calc(distance, minutes)
			json.NewEncoder(w).Encode(total)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/princing/city/{city}/distance/{distance}/minutes/{minutes}", findOne)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Pricings = []Pricing{
		Pricing{Id: "1", City: "São Paulo", PricePerMinute: 1.00, BaseFee: 3.50, ServiceFee: 0.75, PricePerKm: 0.5},
		Pricing{Id: "2", City: "Rio de Janeiro", PricePerMinute: 0.95, BaseFee: 3.00, ServiceFee: 1.40},
	}

	handleRequests()
}
