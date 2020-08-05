package route

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/daniel-acaz/api-golang-elastic/domain"
	repository "github.com/daniel-acaz/api-golang-elastic/repository"
	"github.com/gorilla/mux"
)

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

	property := repository.FindById(id)

	response, err := json.Marshal(property)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateProperty(w http.ResponseWriter, r *http.Request) {

	var property model.Property
	err := json.NewDecoder(r.Body).Decode(&property)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := repository.GetMaxId() + 1

	property.ID = id

	response, err := json.Marshal(repository.Save(property))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func UpdateProperty(w http.ResponseWriter, r *http.Request) {

	var property model.Property
	err := json.NewDecoder(r.Body).Decode(&property)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	property.ID = id

	response, err := json.Marshal(repository.Save(property))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repository.Delete(id)

	w.WriteHeader(http.StatusNotFound)

}
