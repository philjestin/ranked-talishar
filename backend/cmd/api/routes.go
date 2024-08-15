package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/heroes", app.showAllHeroesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/heroes/:id", app.showHeroHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	// router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	// router.HandlerFunc(http.MethodPut, "/v1/users/password", app.updateUserPasswordHandler)

	return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
