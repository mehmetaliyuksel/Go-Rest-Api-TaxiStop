package controller

import (
	"TaxiStop/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DriverController struct {
}

var singletonDriverControllerInstance *DriverController

func GetDriverControllerInstance() *DriverController {
	once.Do(func() {
		singletonDriverControllerInstance = &DriverController{}
	})
	return singletonDriverControllerInstance
}

func (tsc *DriverController) RegisterDriver(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver

	err := json.NewDecoder(r.Body).Decode(&driver)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Validation

	if driver, err = driver.Create(driver); err != nil {
		http.Error(w, "Could Not Register Driver!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(driver)

}

func (tsc *DriverController) GetDrivers(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver

	drivers, err := driver.GetAll()

	if err != nil {
		http.Error(w, "Could Not Get Drivers!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(drivers)

}

func (tsc *DriverController) GetDriver(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver

	driverId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	driver, err = driver.FindBy(uint(driverId))

	if err != nil {
		http.Error(w, "Could Not Get Driver!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(driver)

}

func (tsc *DriverController) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver

	driverId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = driver.DeleteBy(uint(driverId)); err != nil {
		http.Error(w, "Could Not Delete Driver!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w)

}

func (tsc *DriverController) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver

	err := json.NewDecoder(r.Body).Decode(&driver)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	driverId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	driver.ID = uint(driverId)
	if err = driver.Update(driver); err != nil {
		http.Error(w, "Could Not Update Driver!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w)

}
