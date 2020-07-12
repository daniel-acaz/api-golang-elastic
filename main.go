package main

import "time"

type Property struct {
	ID                  int
	bedroomQuantity     string
	squareMetter        int
	price               float64
	address             Address
	buldingDate         time.Time
	parkingLotsQuantity int
	bathroomQuantity    int
	hasFurniture        bool
}

type Address struct {
	street       string
	number       int
	neighborhood string
	city         string
	state        string
	country      string
}

func main() {

}
