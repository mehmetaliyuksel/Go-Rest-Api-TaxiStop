package controller

import (
	"TaxiStop/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CarController struct {
}

var singletonCarControllerInstance *CarController

func GetCarControllerInstance() *CarController {
	once.Do(func() {
		singletonCarControllerInstance = &CarController{}
	})
	return singletonCarControllerInstance
}

func (cc *CarController) RegisterCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car

	err := json.NewDecoder(r.Body).Decode(&car)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Validation

	if car, err = car.Create(car); err != nil {
		http.Error(w, "Could Not Register Car!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(car)

}

func (cc *CarController) GetCars(w http.ResponseWriter, r *http.Request) {
	var car model.Car

	cars, err := car.GetAll()

	if err != nil {
		http.Error(w, "Could Not Get Cars!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(cars)

}

func (cc *CarController) GetCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car

	carId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car, err = car.FindBy(uint(carId))

	if err != nil {
		http.Error(w, "Could Not Get Car!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(car)

}

func (cc *CarController) DeleteCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car

	carId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = car.DeleteBy(uint(carId)); err != nil {
		http.Error(w, "Could Not Delete Car!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w)

}

func (cc *CarController) UpdateCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car

	err := json.NewDecoder(r.Body).Decode(&car)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	carId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car.ID = uint(carId)
	if err = car.Update(car); err != nil {
		http.Error(w, "Could Not Update Car!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w)

}
