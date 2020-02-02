package repository

import (
	"fmt"
	"log"

	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const DOCNAME = "Pricing"

var db *mgo.Database

func (r *Repository) Connect(server string, database string) {
	session, err := mgo.Dial(server)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	db = session.DB(database)
}

func (r Repository) Find(city string) (Pricing, error) {

	var pricing Pricing

	err := db.C(DOCNAME).Find(bson.M{"city": city}).One(&pricing)

	if err != nil {
		fmt.Println(err)
	}

	return pricing, err
}
