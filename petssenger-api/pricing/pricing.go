package pricing

import (
	"fmt"
	"math"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Pricing struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	City           string        `bson:"city"  json:"city"`
	BaseFee        float64       `bson:"basefee" json:"basefee"`
	PricePerMinute float64       `bson:"pricePerMinute" json:"pricePerMinute"`
	ServiceFee     float64       `bson:"serviceFee" json:"serviceFee"`
	PricePerKm     float64       `bson:"pricePerKm" json:"pricePerKm"`
}

func (pricing *Pricing) Calc(distance float64, minutes float64, multiplicator float64) float64 {
	fmt.Println(multiplicator)
	total := pricing.BaseFee + ((pricing.PricePerMinute * minutes) + (pricing.PricePerKm*distance)*multiplicator) + pricing.ServiceFee
	return math.Floor(total*100) / 100
}

func (pricing *Pricing) GetMultiplicator(multiplicator float64, expiredTime time.Time) float64 {
	now := time.Now()

	if expiredTime.After(now) {
		return multiplicator + 0.1
	}

	return multiplicator - 0.1
}

var Pricings []Pricing
