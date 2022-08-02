package controller

import (
	"TaxiStop/app/auth"
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

func GetUserControllerInstance() *UserController {
	once.Do(func() {
		singletonUserControllerInstance = &UserController{}
	})
	return singletonUserControllerInstance
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var tokenString string

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Validation
	if err = user.HashPassword(user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user, err = user.Create(user); err != nil {
		http.Error(w, "Could Not Register User!", http.StatusInternalServerError)
		return
	}

	if tokenString, err = auth.GenerateJWT(user.Email, user.Username, user.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{}

	resp["user"] = user
	resp["JWT"] = tokenString

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(resp)

}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var tokenString string

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user, err = user.FindByEmail(user.Email); err != nil {
		http.Error(w, "User does not exists!", http.StatusBadRequest)
		return
	}

	if credentialError := user.CheckPassword(user.Password); credentialError != nil {
		http.Error(w, "Password is incorrect!", http.StatusUnauthorized)
		return
	}

	if tokenString, err = auth.GenerateJWT(user.Email, user.Username, user.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(tokenString)
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

	user, err = user.FindById(uint(userId))

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
