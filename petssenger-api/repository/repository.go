package repository

import (
	"fmt"
	"log"

	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const SERVER = "mongodb://admin:admin123@ds253368.mlab.com:53368/live-dotnet"

const DBNAME = "live-dotnet"

const DOCNAME = "Pricing"

var db *mgo.Database

func (r *Repository) Connect() {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	db = session.DB(DBNAME)
}

func (r Repository) Find(city string) (Pricing, error) {

	var pricing Pricing

	err := db.C(DOCNAME).Find(bson.M{"city": city}).One(&pricing)

	if err != nil {
		fmt.Println(err)
	}

	return pricing, err
}
