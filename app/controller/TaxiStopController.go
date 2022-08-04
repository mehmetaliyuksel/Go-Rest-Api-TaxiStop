package controller

import (
	"TaxiStop/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TaxiStopController struct {
}

var singletonTaxiStopControllerInstance *TaxiStopController

func GetTaxiStopControllerInstance() *TaxiStopController {
	once.Do(func() {
		singletonTaxiStopControllerInstance = &TaxiStopController{}
	})
	return singletonTaxiStopControllerInstance
}

func (tsc *TaxiStopController) RegisterTaxiStop(w http.ResponseWriter, r *http.Request) {
	var taxiStop model.TaxiStop
	var user model.User

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&taxiStop)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if model.IsExist(taxiStop) {
		http.Error(w, "TaxiStop Already Exists!", http.StatusBadRequest)
		return
	}

	// TODO: Validation

	user, _ = user.FindById(uint(userId))
	taxiStop.User = append(taxiStop.User, user)

	if taxiStop, err = taxiStop.Create(taxiStop); err != nil {
		http.Error(w, "Could Not Register TaxiStop!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(taxiStop)

}

func (tsc *TaxiStopController) GetTaxiStops(w http.ResponseWriter, r *http.Request) {
	var taxiStop model.TaxiStop

	taxiStops, err := taxiStop.GetAll()

	if err != nil {
		http.Error(w, "Could Not Get TaxiStops!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(taxiStops)

}

func (tsc *TaxiStopController) GetTaxiStop(w http.ResponseWriter, r *http.Request) {
	var taxiStop model.TaxiStop

	taxiStopId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taxiStop, err = taxiStop.FindBy(uint(taxiStopId))

	if err != nil {
		http.Error(w, "Could Not Get TaxiStop!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(taxiStop)

}

func (tsc *TaxiStopController) GetTaxiStopUsers(w http.ResponseWriter, r *http.Request) {
	var taxiStop model.TaxiStop
	var users []model.User

	taxiStopId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taxiStop.ID = uint(taxiStopId)
	users, err = taxiStop.FindAssociatedUsers(taxiStop)

	if err != nil {
		http.Error(w, "Could Not Get Users of the TaxiStop!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(users)
}

func (tsc *TaxiStopController) DeleteTaxiStop(w http.ResponseWriter, r *http.Request) {
	var taxiStop model.TaxiStop

	taxiStopId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = taxiStop.DeleteBy(uint(taxiStopId)); err != nil {
		http.Error(w, "Could Not Delete TaxiStop!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w)

}

func (tsc *TaxiStopController) UpdateTaxiStop(w http.ResponseWriter, r *http.Request) {
	var taxiStop model.TaxiStop

	err := json.NewDecoder(r.Body).Decode(&taxiStop)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taxiStopId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taxiStop.ID = uint(taxiStopId)
	if err = taxiStop.Update(taxiStop); err != nil {
		http.Error(w, "Could Not Update TaxiStop!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w)

}
