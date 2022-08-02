package app

import (
	"TaxiStop/app/auth"
	"TaxiStop/app/controller"
	"TaxiStop/app/model"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func RunApp() {
	app := new(App)
	app.init()
	app.run()
}

func (a *App) init() {

	model.CreateTables()

}

func (a *App) run() {

	insecureRouter := mux.NewRouter().StrictSlash(true)
	insecureRouter.Use(inSecureMiddleware)

	secureRouter := insecureRouter.PathPrefix("").Subrouter()
	secureRouter.Use(commonMiddleware)

	userController := controller.GetUserControllerInstance()
	taxiStopController := controller.GetTaxiStopControllerInstance()
	driverController := controller.GetDriverControllerInstance()
	carController := controller.GetCarControllerInstance()

	// User Endpoints TODO: Add Login Endpoint
	insecureRouter.HandleFunc("/createUser", userController.RegisterUser).Methods("POST")
	secureRouter.HandleFunc("/login", userController.Login).Methods("POST")
	secureRouter.HandleFunc("/getUsers", userController.GetUsers).Methods("GET")
	secureRouter.HandleFunc("/getUser/{id}", userController.GetUser).Methods("GET")
	secureRouter.HandleFunc("/updateUser/{id}", userController.UpdateUser).Methods("PUT")
	secureRouter.HandleFunc("/deleteUser/{id}", userController.DeleteUser).Methods("DELETE")

	// TaxiStop Endpoints TODO: Revision
	secureRouter.HandleFunc("/createTaxiStop", taxiStopController.RegisterTaxiStop).Methods("POST")
	secureRouter.HandleFunc("/getTaxiStops", taxiStopController.GetTaxiStops).Methods("GET")
	secureRouter.HandleFunc("/getTaxiStop/{id}", taxiStopController.GetTaxiStop).Methods("GET")
	secureRouter.HandleFunc("/updateTaxiStop/{id}", taxiStopController.UpdateTaxiStop).Methods("PUT")
	secureRouter.HandleFunc("/deleteTaxiStop/{id}", taxiStopController.DeleteTaxiStop).Methods("DELETE")

	// Driver Endpoints TODO: Revision
	secureRouter.HandleFunc("/createDriver", driverController.RegisterDriver).Methods("POST")
	secureRouter.HandleFunc("/getDrivers", driverController.GetDrivers).Methods("GET")
	secureRouter.HandleFunc("/getDriver/{id}", driverController.GetDriver).Methods("GET")
	secureRouter.HandleFunc("/updateDriver/{id}", driverController.UpdateDriver).Methods("PUT")
	secureRouter.HandleFunc("/deleteDriver/{id}", driverController.DeleteDriver).Methods("DELETE")

	// Car Endpoints TODO: Revision
	secureRouter.HandleFunc("/createCar", carController.RegisterCar).Methods("POST")
	secureRouter.HandleFunc("/getCars", carController.GetCars).Methods("GET")
	secureRouter.HandleFunc("/getCar/{id}", carController.GetCar).Methods("GET")
	secureRouter.HandleFunc("/updateCar/{id}", carController.UpdateCar).Methods("PUT")
	secureRouter.HandleFunc("/deleteCar/{id}", carController.DeleteCar).Methods("DELETE")

	// TODO: Improve logging
	fmt.Println("Server is Running!")

	// TODO: Configure Go Routines
	log.Fatal(http.ListenAndServe(":8000", insecureRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	// TODO: Revision for improving
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Missing Token! Unauthorized!", http.StatusUnauthorized)
			return
		}

		if err := auth.ValidateToken(tokenString); err != nil {
			http.Error(w, "Invalid Token! Unauthorized!", http.StatusUnauthorized)
			return
		}

		//w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func inSecureMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
