package models

import "time"

type Multiplicator struct {
	Multiplicator  float64   `json:"multiplicator"`
	ExpirationTime time.Time `json:"expirationTime"`
}
