package routers

import (
	"github.com/gorilla/mux"
	"github.com/gurdeep/crud/controllers"
)

func GetMuxRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
	return router
}
