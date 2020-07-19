package route

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	model "github.com/daniel-acaz/api-golang-elastic/domain"
	repository "github.com/daniel-acaz/api-golang-elastic/repository"
	"github.com/gorilla/mux"
)

var models = []model.Property{
	{
		ID:              1,
		BedroomQuantity: 1,
		SquareMetter:    100,
		Price:           10.0,
		Address: model.Address{
			Street:       "Rua Oliver",
			Number:       386,
			Neighborhood: "Andrada",
			City:         "São Paulo",
			State:        "São Paulo",
			Country:      "Brasil",
		},
		BuldingDate:         time.Date(1990, time.June, 12, 0, 0, 0, 0, time.UTC),
		ParkingLotsQuantity: 2,
		BathroomQuantity:    2,
		HasFurniture:        true,
	},
	{
		ID:              2,
		BedroomQuantity: 3,
		SquareMetter:    500,
		Price:           1000.0,
		Address: model.Address{
			Street:       "Rua Alexandre Magno",
			Number:       200,
			Neighborhood: "Vila Abrão",
			City:         "São Paulo",
			State:        "São Paulo",
			Country:      "Brasil",
		},
		BuldingDate:         time.Date(2013, time.March, 3, 0, 0, 0, 0, time.UTC),
		ParkingLotsQuantity: 2,
		BathroomQuantity:    2,
		HasFurniture:        false,
	},
}

func GetAllProperty(w http.ResponseWriter, r *http.Request) {

	properties := repository.FindAll()

	response, err := json.Marshal(properties)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func GetPropertyById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, property := range models {

		if property.ID == id {

			response, err := json.Marshal(property)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(response)

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func CreateProperty(w http.ResponseWriter, r *http.Request) {

	var property model.Property
	err := json.NewDecoder(r.Body).Decode(&property)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	property.ID = len(models) + 1

	models = append(models, property)

	response, err := json.Marshal(models)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func UpdateProperty(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for index, updateProperty := range models {

		if updateProperty.ID == id {

			var property model.Property
			err := json.NewDecoder(r.Body).Decode(&property)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updateProperty = updateObject(index, property)

			response, err := json.Marshal(updateProperty)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(response)

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

}

func updateObject(index int, property model.Property) model.Property {
	models[index].Address = property.Address
	models[index].BathroomQuantity = property.BathroomQuantity
	models[index].BedroomQuantity = property.BedroomQuantity
	models[index].BuldingDate = property.BuldingDate
	models[index].HasFurniture = property.HasFurniture
	models[index].ParkingLotsQuantity = property.ParkingLotsQuantity
	models[index].Price = property.Price
	models[index].SquareMetter = property.SquareMetter
	return models[index]
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for index, removeProperty := range models {

		if removeProperty.ID == id {

			models = append(models[:index], models[index+1:]...)

			response, err := json.Marshal(models)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(response)

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

}
