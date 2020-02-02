package services

import (
	"time"

	. "../models"
	. "../repository"
)

type Service struct{}

var repository = Repository{}

func (s Service) FindPricingCalculated(city string, distance float64, minutes float64, userId string) (float64, error) {
	pricing, err := repository.Find(city)

	var multi Multiplicator

	if err != nil {
		return 0.0, err
	}

	multi, found := GetCache(city + userId)

	if !found {
		newMulti := &Multiplicator{Multiplicator: 1.0, ExpirationTime: time.Now().Add(5 * time.Minute)}

		SetCache(city+userId, *newMulti)

		total := pricing.Calc(distance, minutes, 1.0)

		return total, nil
	} else {
		multi.Multiplicator = pricing.GetMultiplicator(multi.Multiplicator, multi.ExpirationTime)

		SetCache(city+userId, multi)

		total := pricing.Calc(distance, minutes, multi.Multiplicator)

		return total, nil
	}
}
