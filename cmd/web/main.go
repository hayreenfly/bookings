package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/hayreenfly/bookings/pkg/config"
	"github.com/hayreenfly/bookings/pkg/handlers"
	"github.com/hayreenfly/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main(){

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // production should be true

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)


	render.NewTemplates(&app)


	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
