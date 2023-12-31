package main

import (
	"fmt"
	"log"
	"time"

	//"html/template"
	"net/http"

	"github.com/AlvinTsou/bookings/pkg/config"
	"github.com/AlvinTsou/bookings/pkg/handlers"
	"github.com/AlvinTsou/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

// const cant be changed
const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main application function
func main() {
	// using 2nd templates loading with getring the template cache from config.go
	var app config.AppConfig

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // 24 hours, must be useing time package
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // change to true when https

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplatesCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	// 2nd solution

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	//_ = http.ListenAndServe(portNumber, nil)

	// another way to create a server without using http.HandleFunc
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
