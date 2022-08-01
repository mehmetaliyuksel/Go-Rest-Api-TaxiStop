package app

import (
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

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware)

	userController := controller.GetUserControllerInstance()
	taxiStopController := controller.GetTaxiStopControllerInstance()
	driverController := controller.GetDriverControllerInstance()
	carController := controller.GetCarControllerInstance()

	// User Endpoints TODO: Add Login Endpoint
	router.HandleFunc("/createUser", userController.RegisterUser).Methods("POST")
	router.HandleFunc("/getUsers", userController.GetUsers).Methods("GET")
	router.HandleFunc("/getUser/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/updateUser/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteUser/{id}", userController.DeleteUser).Methods("DELETE")

	// TaxiStop Endpoints TODO: Revision
	router.HandleFunc("/createTaxiStop", taxiStopController.RegisterTaxiStop).Methods("POST")
	router.HandleFunc("/getTaxiStops", taxiStopController.GetTaxiStops).Methods("GET")
	router.HandleFunc("/getTaxiStop/{id}", taxiStopController.GetTaxiStop).Methods("GET")
	router.HandleFunc("/updateTaxiStop/{id}", taxiStopController.UpdateTaxiStop).Methods("PUT")
	router.HandleFunc("/deleteTaxiStop/{id}", taxiStopController.DeleteTaxiStop).Methods("DELETE")

	// Driver Endpoints TODO: Revision
	router.HandleFunc("/createDriver", driverController.RegisterDriver).Methods("POST")
	router.HandleFunc("/getDrivers", driverController.GetDrivers).Methods("GET")
	router.HandleFunc("/getDriver/{id}", driverController.GetDriver).Methods("GET")
	router.HandleFunc("/updateDriver/{id}", driverController.UpdateDriver).Methods("PUT")
	router.HandleFunc("/deleteDriver/{id}", driverController.DeleteDriver).Methods("DELETE")

	// Car Endpoints TODO: Revision
	router.HandleFunc("/createCar", carController.RegisterCar).Methods("POST")
	router.HandleFunc("/getCars", carController.GetCars).Methods("GET")
	router.HandleFunc("/getCar/{id}", carController.GetCar).Methods("GET")
	router.HandleFunc("/updateCar/{id}", carController.UpdateCar).Methods("PUT")
	router.HandleFunc("/deleteCar/{id}", carController.DeleteCar).Methods("DELETE")

	// TODO: Improve logging
	fmt.Println("Server is Running!")

	// TODO: Configure Go Routines
	log.Fatal(http.ListenAndServe(":8000", router))
}

func middleware(next http.Handler) http.Handler {
	// TODO: Revision for improving
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
