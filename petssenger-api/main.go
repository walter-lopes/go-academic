package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	. "./pricing"
	"github.com/gorilla/mux"
)

var repository = Repository{}

func findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	city := vars["city"]
	distance, _ := strconv.ParseFloat(vars["distance"], 64)
	minutes, _ := strconv.ParseFloat(vars["minutes"], 64)

	pricing, err := repository.Find(city)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid city")
		return
	}

	total := pricing.Calc(distance, minutes)
	json.NewEncoder(w).Encode(total)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/pricing/city/{city}/distance/{distance}/minutes/{minutes}", findOne)

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
	repository.Connect()
}

func main() {
	handleRequests()
}
