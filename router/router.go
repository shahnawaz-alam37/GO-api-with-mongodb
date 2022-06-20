package router

import (
	"github.com/gorilla/mux"
	"github.com/shahnawaz-alam37/newrepo/controller"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("api/movies",controller.Getmyallmovies).Methods("GET")
	router.HandleFunc("api/movie",controller.CreateMovie).Methods("POST")
	router.HandleFunc("api/movie/{id}",controller.Markaswatched).Methods("PUT")
	router.HandleFunc("api/movie/{id}",controller.Deletemymovie).Methods("DELETE")
	router.HandleFunc("api/movie-delete-all",controller.Deletemyallmovie).Methods("DELETE")
	return router
}