package app

import (
	"TaxiStop/app/controller"
	"TaxiStop/app/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	// DB     *gorm.DB ?
	//Queue  *amqp.Channel
}

func RunApp() {
	app := new(App)
	app.init()
	app.run()
}

func (a *App) init() {

	//database.InitDBConnection()
	model.CreateTables()

}

func (a *App) run() {

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware)

	userController := controller.GetInstance()

	router.HandleFunc("/createUser", userController.RegisterUser).Methods("POST")
	router.HandleFunc("/getUsers", userController.GetUsers).Methods("GET")
	router.HandleFunc("/getUser/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/updateUser/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteUser/{id}", userController.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
