package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	. "./config"
	. "./repository"
	. "./services"
	"github.com/gorilla/mux"
)

var service = Service{}

var config = Config{}

var repository = Repository{}

func findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	city := vars["city"]
	distance, _ := strconv.ParseFloat(vars["distance"], 64)
	minutes, _ := strconv.ParseFloat(vars["minutes"], 64)
	userId, _ := vars["userId"]

	total, err := service.FindPricingCalculated(city, distance, minutes, userId)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error")
	}

	json.NewEncoder(w).Encode(total)

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/pricing/city/{city}/distance/{distance}/minutes/{minutes}/userId/{userId}", findOne)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	server := config.Server
	database := config.Database

	repository.Connect(server, database)
}

func main() {
	handleRequests()
}
