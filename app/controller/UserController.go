package controller

import (
	"TaxiStop/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"sync"
)

type UserController struct{}

var once sync.Once
var singletonUserControllerInstance *UserController

func GetInstance() *UserController {
	once.Do(func() {
		singletonUserControllerInstance = &UserController{}
	})
	return singletonUserControllerInstance
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Validation

	if user, err = user.Create(user); err != nil {
		http.Error(w, "Could Not Register User!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(user)

}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var user model.User

	users, err := user.GetAll()

	if err != nil {
		http.Error(w, "Could Not Get Users!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(users)

}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = user.FindBy(uint(userId))

	if err != nil {
		http.Error(w, "Could Not Get User!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(user)

}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = user.DeleteBy(uint(userId)); err != nil {
		http.Error(w, "Could Not Delete User!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w)

}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = uint(userId)
	if err = user.Update(user); err != nil {
		http.Error(w, "Could Not Update User!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w)

}
