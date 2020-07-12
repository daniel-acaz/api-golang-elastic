package domain

import "time"

type Property struct {
	ID                  int
	BedroomQuantity     int
	SquareMetter        int
	Price               float64
	Address             Address
	BuldingDate         time.Time
	ParkingLotsQuantity int
	BathroomQuantity    int
	HasFurniture        bool
}

type Address struct {
	Street       string
	Number       int
	Neighborhood string
	City         string
	State        string
	Country      string
}
