package main

import (
	"net/http"

	"github.com/AlvinTsou/bookings/pkg/config"
	"github.com/AlvinTsou/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	/* mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	*/
	//mux.Use(WriteToConsole)
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Use(NoSurf)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// Get the static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
