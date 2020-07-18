package domain

import "time"

type Property struct {
	ID                  int       `json:"id,omitempty"`
	BedroomQuantity     int       `json:"bedroom_quantity,omitempty"`
	SquareMetter        int       `json:"square_metter,omitempty"`
	Price               float64   `json:"price,omitempty"`
	Address             Address   `json:"address,omitempty"`
	BuldingDate         time.Time `json:"building_date,omitempty"`
	ParkingLotsQuantity int       `json:"parking_lots_quantity,omitempty"`
	BathroomQuantity    int       `json:"bathroom_quantity,omitempty"`
	HasFurniture        bool      `json:"has_furniture,omitempty"`
}

type Address struct {
	Street       string `json:"street,omitempty"`
	Number       int    `json:"number,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Country      string `json:"country,omitempty"`
}
