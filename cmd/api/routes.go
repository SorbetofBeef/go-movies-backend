package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes returns a variable, "router", which holds http methods
func (app *application) routes() http.Handler {
	router := httprouter.New()
	// GET
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:genre_id", app.getAllMoviesByGenre)
	// POST
	router.HandlerFunc(http.MethodPost, "/v1/admin/editmovie", app.editMovie)
	return app.enableCORS(router)
}
