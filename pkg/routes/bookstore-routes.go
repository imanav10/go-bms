package routes


import {
	"github.com/gorilla/mux"
	"github.com/imanav10/go-bms/pkg/controllers"
}


var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
}