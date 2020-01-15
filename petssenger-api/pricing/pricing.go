package pricing

import "gopkg.in/mgo.v2/bson"

type Pricing struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	City           string        `bson:"city"  json:"city"`
	BaseFee        float64       `bson:"basefee" json:"basefee"`
	PricePerMinute float64       `bson:"pricePerMinute" json:"pricePerMinute"`
	ServiceFee     float64       `bson:"serviceFee" json:"serviceFee"`
	PricePerKm     float64       `bson:"pricePerKm" json:"pricePerKm"`
}

func (pricing *Pricing) Calc(distance float64, minutes float64) float64 {
	total := pricing.BaseFee + ((pricing.PricePerMinute * minutes) + (pricing.PricePerKm*distance)*1.0) + pricing.ServiceFee
	return total
}

var Pricings []Pricing
